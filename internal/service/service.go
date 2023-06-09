package service

import (
	"github.com/F7icK/pc-accessories-store/internal/clients/repository"
	"github.com/F7icK/pc-accessories-store/internal/service/storage"
	"github.com/F7icK/pc-accessories-store/internal/types"
)

type Storage interface {
	GetProduct(productID string) (*types.ProductResp, error)
	NewProduct(newProduct *types.Product, productProperty []types.ProductPropertyResp) (*types.ProductResp, error)
	UpdateProduct(product *types.Product, productProperty []types.ProductPropertyResp) (*types.ProductResp, error)
	DeleteProduct(productID string) error
	GetCategories() ([]types.CategoriesResp, error)
	GetProducts(filter *types.ReqFilterProducts) ([]types.ProductResp, error)

	NewCategory(category *types.Category) (*types.Category, error)
}

type Service struct {
	Storage
}

func NewService(db *repository.Repository) *Service {
	return &Service{
		Storage: storage.NewStorageService(db.Storage),
	}
}
