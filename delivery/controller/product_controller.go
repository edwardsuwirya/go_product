package controller

import (
	"enigmacamp.com/go_product/delivery/api"
	"enigmacamp.com/go_product/model"
	"enigmacamp.com/go_product/usecase"
	"enigmacamp.com/go_product/utils"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router  *gin.Engine
	usecase usecase.ProductRegistrationUseCase
	api.BaseApi
}

func (cc *ProductController) registerNewProduct(ctx *gin.Context) {
	var newProduct model.Product
	err := cc.ParseRequestBody(ctx, &newProduct)
	if err != nil {
		cc.Failed(ctx, utils.RequiredError())
		return
	}
	err = cc.usecase.Register(&newProduct)
	if err != nil {
		cc.Failed(ctx, err)
		return
	}
	cc.Success(ctx, newProduct)
}

func NewProductController(r *gin.Engine, usecase usecase.ProductRegistrationUseCase) *ProductController {
	controller := ProductController{
		router:  r,
		usecase: usecase,
	}
	r.POST("/product", controller.registerNewProduct)
	return &controller
}
