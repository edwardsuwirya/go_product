package delivery

import (
	"enigmacamp.com/go_product/model"
	"enigmacamp.com/go_product/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductApi struct {
	productRegistrationUseCase usecase.ProductRegistrationUseCase
}

func (p *ProductApi) createProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.PostForm("productId")
		productName := c.PostForm("productName")
		newProduct := model.Product{ProductId: productId, ProductName: productName}
		err := p.productRegistrationUseCase.Register(&newProduct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"message": newProduct,
		})
	}
}

func NewProductApi(productRoute *gin.RouterGroup, productReqistrationUseCase usecase.ProductRegistrationUseCase) {
	api := ProductApi{
		productRegistrationUseCase: productReqistrationUseCase,
	}
	productRoute.POST("", api.createProduct())
}
