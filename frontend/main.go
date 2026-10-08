package main

import (
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Raden-Salaf/ecommerce-microservices/frontend/handlers"
)

// formatRupiah mengubah 15000 menjadi "Rp 15.000".
func formatRupiah(n int) string {
	s := strconv.Itoa(n)
	out := ""
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			out += "."
		}
		out += string(ch)
	}
	return "Rp " + out
}

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"rupiah": formatRupiah,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", handlers.Home)
	r.GET("/products/:id", handlers.ProductDetail)

	r.Run(":3000")
}
