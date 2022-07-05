package api

import (
	"enigmacamp.com/go_product/delivery/api/response"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) Success(c *gin.Context, data interface{}) {
	response.NewSuccessJsonResponse(c, data).Send()
}

func (b *BaseApi) Failed(c *gin.Context, err error) {
	response.NewErrorJsonResponse(c, err).Send()
}
