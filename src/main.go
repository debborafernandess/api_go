package main

import (
	"go-api/src/controllers"
	"go-api/src/db"
	"go-api/src/repositories"
	"go-api/src/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, error := db.ConnectDB()
	if error != nil {
		panic(error)
	}

	//repositories
	ProductRepository := repositories.NewProductRepository(dbConnection)

	//Usecases
	ProductUsecase := usecases.NewProductUsecase(ProductRepository)

	//controllers
	ProductController := controllers.NewProductController(ProductUsecase)

	//routes
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Olá o/",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:id", ProductController.FindProduct)
	server.POST("/products", ProductController.CreateProduct)
	server.DELETE("/product/:id", ProductController.DeleteProduct)

	server.Run(":8000")
}
