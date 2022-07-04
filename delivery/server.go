package delivery

import (
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
	repo := repository.NewProductRepository()
	usecase := usecase.NewProductRegistrationUseCase(repo)
	host := fmt.Sprintf("%s:%s", "0.0.0.0", "8888")
	return &appServer{
		productRegistrationUseCase: usecase,
		engine:                     r,
		host:                       host,
	}
}
