package storage

import (
	"github.com/F7icK/pc-accessories-store/internal/types"
	"gorm.io/gorm"
)

type StoragePostgres struct {
	db *gorm.DB
}

func NewStoragePostgres(db *gorm.DB) *StoragePostgres {
	return &StoragePostgres{db: db}
}

func (p *StoragePostgres) Begin() *gorm.DB {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return tx
}

func (p *StoragePostgres) GetProduct(productID string) (*types.ProductResp, error) {
	product := new(types.ProductResp)
	if err := p.db.Debug().
		Select(`p.id, p.name, p.price, p.category_id, c.name as category,
			(SELECT name FROM categories WHERE id = c.parent_id) AS parent_category,
			p.created_at, p.updated_at, p.deleted_at`).
		Table("products p").
		Joins("JOIN categories c on c.id = p.category_id").
		Where("p.id = ?", productID).Take(product).Error; err != nil {
		return nil, err
	}

	if err := p.db.Debug().
		Select("p2.id AS property_id, p2.name, p1.value").
		Table("product_properties p1").
		Joins("JOIN properties p2 on p2.id = p1.property_id").
		Where("p1.product_id = ? AND p1.deleted_at IS NULL", productID).Find(&product.Properties).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *StoragePostgres) GetProductByName(nameProduct string) (*types.Product, error) {
	product := new(types.Product)

	if err := p.db.Debug().
		Table("products").
		Where("name = ?", nameProduct).
		Take(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *StoragePostgres) GetProductByNameWithRemote(nameProduct string) (*types.Product, error) {
	product := new(types.Product)

	if err := p.db.Debug().Unscoped().
		Table("products").
		Where("name = ?", nameProduct).
		Take(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *StoragePostgres) AddProduct(tx *gorm.DB, product *types.Product) (*types.Product, error) {
	if err := tx.Debug().
		Table("products").
		Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *StoragePostgres) GetCategory(categoryID string) (*types.Category, error) {
	category := new(types.Category)

	if err := p.db.Debug().
		Table("categories").
		Where("id = ?", categoryID).
		Take(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (p *StoragePostgres) AddProperty(tx *gorm.DB, property *types.Property) (*types.Property, error) {
	oldProperty := new(types.Property)

	if err := tx.Debug().Table("properties").Where("name = ?", property.Name).Take(oldProperty).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err = tx.Debug().Table("properties").Create(property).Error; err != nil {
				return nil, err
			}

			return property, nil
		}

		return nil, err
	}

	return oldProperty, nil
}

func (p *StoragePostgres) AddProductProperty(tx *gorm.DB, productProperty *types.ProductProperty) error {
	if tx.Debug().Table("product_properties").Where("product_id = ? AND property_id = ?", productProperty.ProductID, productProperty.PropertyID).Updates(productProperty).RowsAffected == 0 {
		if err := tx.Debug().Table("product_properties").Create(productProperty).Error; err != nil {
			return err
		}
	}

	return nil
}

func (p *StoragePostgres) UpdateProduct(tx *gorm.DB, product *types.Product) (*types.Product, error) {
	if err := tx.Debug().
		Table("products").
		Where("id = ?", product.ID).
		Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *StoragePostgres) DeleteOldProductProperties(tx *gorm.DB, productID string, productProperties []string) error {
	if err := tx.Debug().
		Table("product_properties").
		Where("product_id = ? AND property_id NOT IN ?", productID, productProperties).
		Delete(&types.ProductProperty{}).Error; err != nil {
		return err
	}
	return nil
}

func (p *StoragePostgres) DeleteProductPropertiesByProductID(tx *gorm.DB, productID string) error {
	if err := tx.Debug().Table("product_properties").
		Where("product_id = ?", productID).
		Delete(&types.ProductProperty{}).Error; err != nil {
		return err
	}
	return nil
}

func (p *StoragePostgres) DeleteProduct(tx *gorm.DB, product *types.Product) error {
	if err := tx.Debug().Unscoped().
		Table("products").
		Where("id = ?", product.ID).
		Delete(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *StoragePostgres) AddCategory(category *types.Category) (*types.Category, error) {
	if p.db.Debug().Table("categories").Where("name = ? AND parent_id = ?", category.Name, category.ParentID).Take(category).Limit(1).RowsAffected == 0 {
		if err := p.db.Debug().Table("categories").Create(category).Error; err != nil {
			return nil, err
		}
	}

	return category, nil
}

func (p *StoragePostgres) GetProductPropertiesByProductID(productID string) ([]types.ProductProperty, error) {
	ProductProperties := make([]types.ProductProperty, 0)

	if err := p.db.Debug().
		Table("product_properties").
		Where("product_id = ?", productID).
		Find(&ProductProperties).Error; err != nil {
		return nil, err
	}

	return ProductProperties, nil
}

func (p *StoragePostgres) GetCategories() ([]types.CategoriesResp, error) {
	product := make([]types.CategoriesResp, 0)

	if err := p.db.Debug().
		Select("c2.id, c2.name, c1.name AS parent_name").
		Table("categories c1").
		Joins("RIGHT JOIN categories c2 on c2.parent_id = c1.id").
		Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *StoragePostgres) GetProducts(filter *types.ReqFilterProducts) ([]types.ProductResp, error) {
	products := make([]types.ProductResp, 0)

	tx := p.db.Debug().
		Select("p.id, p.name, p.price, p.category_id, c.name as category, (?) AS parent_category, p.created_at, p.updated_at, p.deleted_at",
			p.db.Debug().
				Select("name").
				Table("categories").
				Where("id = c.parent_id")).
		Table("products p").
		Joins("JOIN categories c on c.id = p.category_id")

	if filter.PropertyID != "" {
		tx = tx.Joins("JOIN product_properties pp on p.id = pp.product_id").
			Group("p.id, p.name, p.price, p.category_id, c.name, parent_category, p.created_at, p.updated_at, p.deleted_at").
			Where("pp.property_id = ? AND pp.value = ?", filter.PropertyID, filter.PropertyVal)
	}

	if filter.CategoryID != "" {
		tx = tx.Where("p.category_id = ?", filter.CategoryID)
	}

	if err := tx.Find(&products).Error; err != nil {
		return nil, err
	}

	for i := range products {
		if err := p.db.Debug().
			Select("p2.id AS property_id, p2.name, p1.value").
			Table("product_properties p1").
			Joins("JOIN properties p2 on p2.id = p1.property_id").
			Where("p1.product_id = ? AND p1.deleted_at IS NULL", products[i].ID).
			Find(&products[i].Properties).Error; err != nil {
			return nil, err
		}
	}

	return products, nil
}
