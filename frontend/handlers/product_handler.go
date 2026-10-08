package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Raden-Salaf/ecommerce-microservices/frontend/services"
	"github.com/gin-gonic/gin"
)

// Home menampilkan katalog produk.
func Home(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"title":   "Terjadi Kesalahan",
			"message": "Gagal memuat produk. Silakan coba lagi.",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "Toko UMKM",
		"products": products,
	})
}

// ProductDetail menampilkan satu produk.
func ProductDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title":   "Permintaan Tidak Valid",
			"message": "ID produk tidak valid.",
		})
		return
	}

	product, err := services.GetProductByID(id)
	if err != nil {
		if errors.Is(err, services.ErrProductNotFound) {
			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"title":   "Produk Tidak Ditemukan",
				"message": "Produk yang kamu cari tidak ada.",
			})
			return
		}
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"title":   "Terjadi Kesalahan",
			"message": "Gagal memuat produk. Silakan coba lagi.",
		})
		return
	}

	c.HTML(http.StatusOK, "product_detail.html", gin.H{
		"title":   product.Name,
		"product": product,
	})
}
