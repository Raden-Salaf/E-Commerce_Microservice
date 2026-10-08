package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Raden-Salaf/E-commerce_Microservices/product-service/internal/repository"
	"github.com/Raden-Salaf/E-commerce_Microservices/product-service/internal/response"
)

type ProductHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// GetAll menangani GET /api/products
func (h *ProductHandler) GetAll(c *gin.Context) {
	products := h.repo.FindAll()
	c.JSON(http.StatusOK, response.OK(products))
}

// GetByID menangani GET /api/products/:id
func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("id harus berupa angka"))
		return
	}

	product, err := h.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, response.Fail("produk tidak ditemukan"))
			return
		}
		c.JSON(http.StatusInternalServerError, response.Fail("terjadi kesalahan pada server"))
		return
	}

	c.JSON(http.StatusOK, response.OK(product))
}
