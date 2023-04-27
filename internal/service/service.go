package service

import (
	"github.com/F7icK/pc-accessories-store/internal/clients/repository"
	"github.com/F7icK/pc-accessories-store/internal/service/storage"
	"github.com/F7icK/pc-accessories-store/internal/types"
)

type Storage interface {
	GetProduct(productID string) (*types.ProductResp, error)
	NewProduct(product *types.Product) (*types.Product, error)

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
