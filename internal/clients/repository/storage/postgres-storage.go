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
	if err := p.db.Debug().Table("products p").
		Joins("JOIN categories c on c.id = p.category_id").
		Select(`p.id, p.name, p.price, p.category_id, c.name as category,
			(SELECT name FROM categories WHERE id = c.parent_id) AS parent_category,
			p.created_at, p.updated_at, p.deleted_at`).
		Where("p.id = ?", productID).Take(product).Error; err != nil {
		return nil, err
	}

	if err := p.db.Debug().Table("product_properties p1").
		Joins("JOIN properties p2 on p2.id = p1.property_id").
		Select("p2.name, p1.value").
		Where("p1.product_id = ?", productID).Find(&product.Properties).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *StoragePostgres) AddProduct(tx *gorm.DB, product *types.Product) (*types.Product, error) {
	if err := tx.Debug().Table("products").Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *StoragePostgres) GetCategory(categoryID string) (*types.Category, error) {
	category := new(types.Category)

	if err := p.db.Debug().Table("categories").Where("id = ?", categoryID).Take(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (p *StoragePostgres) AddProperty(tx *gorm.DB, property *types.Property) (*types.Property, error) {
	oldProperty := new(types.Property)
	if err := tx.Debug().Table("properties").Where("name = ?", property.Name).Take(oldProperty).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if oldProperty != nil {
		return oldProperty, nil
	}

	if err := tx.Debug().Table("properties").Create(property).Error; err != nil {
		return nil, err
	}

	return property, nil
}

func (p *StoragePostgres) AddProductProperty(tx *gorm.DB, productProperty *types.ProductProperty) error {
	if err := tx.Debug().Table("product_properties").Create(productProperty).Error; err != nil {
		return err
	}
	return nil
}

func (p *StoragePostgres) UpdateProduct(product *types.Product) (*types.Product, error) {
	if err := p.db.Debug().Table("products").Where("id = ?", product.ID).Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *StoragePostgres) DeleteProduct(product *types.Product) (*types.Product, error) {
	if err := p.db.Debug().Table("products").Where("id = ?", product.ID).Delete(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *StoragePostgres) AddCategory(category *types.Category) (*types.Category, error) {
	if err := p.db.Debug().Table("categories").Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
