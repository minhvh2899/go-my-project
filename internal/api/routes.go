// internal/api/routes.go
package api

import (
	"my-project/internal/api/handlers"
	"my-project/internal/repository"
	"my-project/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Order routes
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	productRoutes := router.Group("/products")
	{
		productRoutes.POST("", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.GetProduct)
		productRoutes.GET("", productHandler.ListProducts)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

}
