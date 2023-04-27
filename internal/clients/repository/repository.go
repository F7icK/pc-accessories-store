package repository

import (
	"github.com/F7icK/pc-accessories-store/internal/clients/repository/storage"
	"github.com/F7icK/pc-accessories-store/internal/types"
	"gorm.io/gorm"
)

type Storage interface {
	Begin() *gorm.DB

	GetProduct(productID string) (*types.ProductResp, error)
	AddProduct(tx *gorm.DB, product *types.Product) (*types.Product, error)
	GetCategory(categoryID string) (*types.Category, error)
	GetProductByName(nameProduct string) (*types.Product, error)
	AddProperty(tx *gorm.DB, property *types.Property) (*types.Property, error)
	AddProductProperty(tx *gorm.DB, property *types.ProductProperty) error
	UpdateProduct(tx *gorm.DB, product *types.Product) (*types.Product, error)
	GetProductPropertiesByProductID(productID string) ([]types.ProductProperty, error)
	DeleteOldProductProperties(tx *gorm.DB, productID string, productProperties []string) error
	DeleteProduct(product *types.Product) (*types.Product, error)
	AddCategory(category *types.Category) (*types.Category, error)
}

type Repository struct {
	Storage
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Storage: storage.NewStoragePostgres(db),
	}
}
