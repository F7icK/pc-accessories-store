package server

import (
	"github.com/F7icK/pc-accessories-store/internal/server/handlers"
	"github.com/labstack/echo/v4"
)

func NewRouter(h *handlers.Handlers) *echo.Echo {
	route := echo.New()
	route.GET("/ping", h.Ping)

	storage := route.Group("/storage")
	storage.GET("/products", h.GetProducts)

	storage.GET("/product", h.GetProduct)
	storage.POST("/product", h.AddProduct)
	storage.PUT("/product", h.UpdateProduct)
	storage.DELETE("/product", h.DeleteProduct)

	storage.GET("/categories", h.GetCategories)
	storage.POST("/category", h.AddCategory)

	return route
}
