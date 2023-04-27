package types

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string         `json:"name" gorm:"column:name"`
	Price      int            `json:"price" gorm:"column:price"`
	CategoryID string         `json:"category_id" gorm:"column:category_id"`
	CreatedAt  time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

type Category struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"name" gorm:"column:name"`
	ParentID  string         `json:"parent_id" gorm:"column:parent_id;type:uuid;default:null"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

type ProductProperty struct {
	ProductID  string         `json:"product_id" gorm:"column:product_id"`
	PropertyID string         `json:"property_id" gorm:"column:property_id"`
	Value      string         `json:"value" gorm:"column:value"`
	CreatedAt  time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

type ProductPropertyResp struct {
	Name  string `json:"name" gorm:"column:name"`
	Value string `json:"value" gorm:"column:value"`
}

type Property struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"name" gorm:"column:name"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

type ProductResp struct {
	ID             string                `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name           string                `json:"name" gorm:"column:name"`
	Price          int                   `json:"price" gorm:"column:price"`
	CategoryID     string                `json:"category_id" gorm:"column:category_id"`
	Category       string                `json:"category" gorm:"column:category"`
	ParentCategory string                `json:"parent_category" gorm:"column:parent_category"`
	Properties     []ProductPropertyResp `json:"properties" gorm:"-"`
	CreatedAt      time.Time             `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time             `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt        `json:"deleted_at" gorm:"column:deleted_at"`
}

type ReqCategory struct {
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

type ReqProduct struct {
	Name       string                `json:"name"`
	Price      int                   `json:"price"`
	CategoryID string                `json:"category_id"`
	Properties []ProductPropertyResp `json:"properties"`
}

type CategoriesResp struct {
	ID         string `json:"id" gorm:"column:id"`
	Name       string `json:"name" gorm:"column:name"`
	ParentName string `json:"parent_name" gorm:"column:parent_name"`
}
