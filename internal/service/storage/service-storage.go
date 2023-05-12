package storage

import (
	"net/http"

	"github.com/F7icK/pc-accessories-store/internal/clients/repository"
	"github.com/F7icK/pc-accessories-store/internal/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StorageService struct {
	db repository.Storage
}

func NewStorageService(db repository.Storage) *StorageService {
	return &StorageService{
		db: db,
	}
}

func (s *StorageService) GetProduct(productID string) (*types.ProductResp, error) {
	if !IsValidUUID(productID) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "no such product")
	}

	product, err := s.db.GetProduct(productID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, echo.ErrInternalServerError
	}

	if product == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "no such product")
	}

	return product, nil
}

func (s *StorageService) NewProduct(newProduct *types.Product, productProperty []types.ProductPropertyResp) (*types.ProductResp, error) {
	if !IsValidUUID(newProduct.CategoryID) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "select a category")
	}

	if _, err := s.db.GetCategory(newProduct.CategoryID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "need valid category id")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	tx := s.db.Begin()

	oldProduct, err := s.db.GetProductByNameWithRemoteTx(tx, newProduct.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	if oldProduct != nil {
		if !oldProduct.DeletedAt.Valid {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "this product is already in the database")
		}

	}

	product, err := s.db.AddProduct(tx, newProduct)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	if len(productProperty) != 0 {
		for _, prop := range productProperty {
			property, err := s.db.AddProperty(tx, &types.Property{Name: prop.Name})
			if err != nil {
				return nil, echo.ErrInternalServerError
			}

			if err = s.db.AddProductProperty(tx, &types.ProductProperty{
				ProductID:  product.ID,
				PropertyID: property.ID,
				Value:      prop.Value,
			}); err != nil {
				return nil, echo.ErrInternalServerError
			}
		}
	}

	if err = tx.Commit().Error; err != nil {
		return nil, echo.ErrInternalServerError
	}

	productResp, err := s.db.GetProduct(product.ID)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	return productResp, nil
}

func (s *StorageService) NewCategory(newCategory *types.Category) (*types.Category, error) {
	if newCategory.Name == "" {
		return nil, echo.ErrBadRequest
	}

	category, err := s.db.AddCategory(newCategory)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return category, nil
}

func (s *StorageService) UpdateProduct(product *types.Product, productProperty []types.ProductPropertyResp) (*types.ProductResp, error) {
	if !IsValidUUID(product.ID) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "no such product")
	}

	tx := s.db.Begin()

	if _, err := s.db.GetProductTx(tx, product.ID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "no such product")
		}
		return nil, echo.ErrInternalServerError
	}

	doubleProduct, err := s.db.GetProductByNameTx(tx, product.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, echo.ErrInternalServerError
	}

	if doubleProduct != nil {
		if doubleProduct.Name != product.Name {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "a product with the same name already exists")
		}
	}

	if _, err := s.db.UpdateProduct(tx, product); err != nil {
		return nil, echo.ErrInternalServerError
	}

	arrPropertyID := make([]string, 0)
	for _, prop := range productProperty {
		property, err := s.db.AddProperty(tx, &types.Property{Name: prop.Name})
		if err != nil {
			return nil, echo.ErrInternalServerError
		}

		if err = s.db.AddProductProperty(tx, &types.ProductProperty{
			ProductID:  product.ID,
			PropertyID: property.ID,
			Value:      prop.Value,
		}); err != nil {
			return nil, echo.ErrInternalServerError
		}

		arrPropertyID = append(arrPropertyID, property.ID)
	}

	if err = s.db.DeleteOldProductProperties(tx, product.ID, arrPropertyID); err != nil {
		return nil, echo.ErrInternalServerError
	}

	if err = tx.Commit().Error; err != nil {
		return nil, echo.ErrInternalServerError
	}

	productResp, err := s.db.GetProduct(product.ID)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	return productResp, nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func (s *StorageService) DeleteProduct(productID string) error {
	if !IsValidUUID(productID) {
		return echo.NewHTTPError(http.StatusBadRequest, "no such product")
	}

	tx := s.db.Begin()

	if err := s.db.DeleteProduct(tx, &types.Product{ID: productID}); err != nil {
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return echo.ErrInternalServerError
	}

	return nil
}

func (s *StorageService) GetCategories() ([]types.CategoriesResp, error) {
	categories, err := s.db.GetCategories()
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	return categories, nil
}

func (s *StorageService) GetProducts(filter *types.ReqFilterProducts) ([]types.ProductResp, error) {
	if !IsValidUUID(filter.PropertyID) && filter.PropertyID != "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "bad filter, specify id property")
	}

	if filter.PropertyID != "" && filter.PropertyVal == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "bad filter, specify the value of the property")
	}

	if !IsValidUUID(filter.CategoryID) && filter.CategoryID != "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "bad filter, specify id category")
	}

	products, err := s.db.GetProducts(filter)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	return products, nil
}
