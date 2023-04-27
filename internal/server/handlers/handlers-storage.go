package handlers

import (
	"net/http"

	"github.com/F7icK/pc-accessories-store/internal/types"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetCategories(c echo.Context) error {
	return nil
}

func (h *Handlers) AddCategory(c echo.Context) error {
	newCategory := new(types.ReqCategory)

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

	return c.JSON(http.StatusCreated, category)
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

func (h *Handlers) AddProduct(c echo.Context) error {
	newProduct := new(types.ReqProduct)

	if err := c.Bind(newProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := h.s.NewProduct(&types.Product{
		Name:       newProduct.Name,
		Price:      newProduct.Price,
		CategoryID: newProduct.CategoryID,
	}, newProduct.Properties)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, product)
}

func (h *Handlers) UpdateProduct(c echo.Context) error {
	id := c.QueryParam("id")

	newProduct := new(types.ReqProduct)

	if err := c.Bind(newProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := h.s.UpdateProduct(&types.Product{
		ID:         id,
		Name:       newProduct.Name,
		Price:      newProduct.Price,
		CategoryID: newProduct.CategoryID,
	},
		newProduct.Properties)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func (h *Handlers) DeleteProduct(c echo.Context) error {
	productID := c.QueryParam("id")

	if err := h.s.DeleteProduct(productID); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
