package delivery

import (
	"enigmacamp.com/go_product/config"
	"enigmacamp.com/go_product/delivery/controller"
	"enigmacamp.com/go_product/manager"
	"fmt"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (p *appServer) initHandlers() {
	controller.NewProductController(p.engine, p.useCaseManager.ProductRegistrationUseCase())
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
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	usecaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s", c.Url)
	return &appServer{
		useCaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
