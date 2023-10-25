package main

import (
	"log"
	"sesi-6/app/product"
	"sesi-6/config"
	"sesi-6/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Products Services",
		Prefork: true,
	})

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	db, err := database.ConnectGORMPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	dbSqlx, err := database.ConnectSqlxPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	product.RegisterServiceProduct(router, db, dbSqlx, nil)

	router.Listen(config.Cfg.App.Port)
}
