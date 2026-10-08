package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Raden-Salaf/E-commerce_Microservices/product-service/internal/handler"
	"github.com/Raden-Salaf/E-commerce_Microservices/product-service/internal/repository"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "product-service",
			"status":  "ok",
		})
	})

	repo := repository.NewMemoryProductRepository()
	productHandler := handler.NewProductHandler(repo)

	api := r.Group("/api")
	{
		api.GET("/products", productHandler.GetAll)
		api.GET("/products/:id", productHandler.GetByID)
	}

	r.Run(":8081")
}
