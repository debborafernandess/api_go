package controllers

import (
	"go-api/src/models"
	"go-api/src/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecases.ProductUsecase
}

func NewProductController(usecase usecases.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		panic(err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) FindProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := models.Response{
			Message: "Precisamos de um ID para fazer a busca",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi((id))
	if err != nil {
		response := models.Response{
			Message: "Precisamos de um ID numerico para fazer a busca",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.FindProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product models.Product
	err := ctx.BindJSON((&product))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	newProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newProduct)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := models.Response{
			Message: "Precisamos de um ID para fazer a busca",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi((id))
	if err != nil {
		response := models.Response{
			Message: "Precisamos de um ID numerico para fazer a busca",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUseCase.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	response := models.Response{
		Message: "Produto removido",
	}

	ctx.JSON(http.StatusOK, response)
}
