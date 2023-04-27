package repository

import (
	"github.com/F7icK/pc-accessories-store/internal/clients/repository/storage"
	"github.com/F7icK/pc-accessories-store/internal/types"
	"gorm.io/gorm"
)

type Storage interface {
	Begin() *gorm.DB

	GetProduct(productID string) (*types.ProductResp, error)
	AddProduct(product *types.Product) (*types.Product, error)
	UpdateProduct(product *types.Product) (*types.Product, error)
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
