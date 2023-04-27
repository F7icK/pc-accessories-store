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
		return nil, echo.ErrBadRequest
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

	_, err := s.db.GetCategory(newProduct.CategoryID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "need valid category id")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	tx := s.db.Begin()

	product, err := s.db.AddProduct(tx, newProduct)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	if len(productProperty) != 0 {
		for _, prop := range productProperty {
			property, err := s.db.AddProperty(tx, &types.Property{
				Name: prop.Name,
			})
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

	tx.Commit()

	productResp, err := s.db.GetProduct(product.ID)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	return productResp, nil
}

func (s *StorageService) NewCategory(newCategory *types.Category) (*types.Category, error) {
	// тут валидация, проверка дублирования и т.д. так как по условию не было, не буду тратить время на это
	if newCategory.Name == "" {
		return nil, echo.ErrBadRequest
	}

	category, err := s.db.AddCategory(newCategory)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return category, nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
