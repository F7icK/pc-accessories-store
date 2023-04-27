package handlers

import (
	"net/http"

	"github.com/F7icK/pc-accessories-store/internal/types"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetCategories(c echo.Context) error {
	return nil
}

type reqCategory struct {
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

func (h *Handlers) AddCategory(c echo.Context) error {
	newCategory := new(reqCategory)

	if err := c.Bind(newCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	category, err := h.s.NewCategory(&types.Category{
		Name:     newCategory.Name,
		ParentID: newCategory.ParentID,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, category)
}

func (h *Handlers) GetProducts(c echo.Context) error {
	return nil
}

func (h *Handlers) GetProduct(c echo.Context) error {
	productID := c.FormValue("id")

	product, err := h.s.GetProduct(productID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

type reqProduct struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
}

func (h *Handlers) AddProduct(c echo.Context) error {
	newProduct := new(reqProduct)

	if err := c.Bind(newProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := h.s.NewProduct(&types.Product{
		Name:       newProduct.Name,
		Price:      newProduct.Price,
		CategoryID: newProduct.CategoryID,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func (h *Handlers) UpdateProduct(c echo.Context) error {
	return nil
}

func (h *Handlers) DeleteProduct(c echo.Context) error {
	return nil
}
