package manager

import (
	"enigmacamp.com/go_product/config"
	"enigmacamp.com/go_product/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type InfraManager interface {
	DbConn() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg config.Config
}

func (i *infraManager) initDb() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application Failed to run", err)
		}
	}()

	if err != nil {
		panic(err)
	}
	i.db = db
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connected")
}
func (i *infraManager) DbConn() *gorm.DB {
	return i.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initDb()
	return &infra
}
