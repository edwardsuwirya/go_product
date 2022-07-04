package delivery

import (
	"enigmacamp.com/go_product/config"
	"enigmacamp.com/go_product/delivery/controller"
	"enigmacamp.com/go_product/repository"
	"enigmacamp.com/go_product/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	productRegistrationUseCase usecase.ProductRegistrationUseCase
	engine                     *gin.Engine
	host                       string
}

func (p *appServer) initHandlers() {
	controller.NewProductController(p.engine, p.productRegistrationUseCase)
}

func (p *appServer) Run() {
	p.initHandlers()
	err := p.engine.Run(p.host)
	if err != nil {
		panic(err)
	}
}

func Server() *appServer {
	r := gin.Default()
	c := config.NewConfig()
	repo := repository.NewProductRepository()
	usecase := usecase.NewProductRegistrationUseCase(repo)
	host := fmt.Sprintf("%s", c.Url)
	return &appServer{
		productRegistrationUseCase: usecase,
		engine:                     r,
		host:                       host,
	}
}
