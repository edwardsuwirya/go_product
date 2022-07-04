package controller

import (
	"enigmacamp.com/go_product/model"
	"enigmacamp.com/go_product/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	router  *gin.Engine
	usecase usecase.ProductRegistrationUseCase
}

func (cc *ProductController) registerNewProduct(ctx *gin.Context) {
	productId := ctx.PostForm("productId")
	productName := ctx.PostForm("productName")
	newProduct := model.Product{
		ProductId:   productId,
		ProductName: productName,
	}
	err := cc.usecase.Register(&newProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when creating product",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": newProduct,
	})
}

func NewProductController(r *gin.Engine, usecase usecase.ProductRegistrationUseCase) *ProductController {
	controller := ProductController{
		router:  r,
		usecase: usecase,
	}
	r.POST("/product", controller.registerNewProduct)
	return &controller
}
