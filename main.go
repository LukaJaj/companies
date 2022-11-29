package main

import (
	"log"
	"os"

	"companies-crud/route"
	"companies-crud/service"
	"companies-crud/store"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	err = store.Migrate(db)
	if err != nil {
		panic(err)
	}

	store := store.NewStore(*db)
	service := service.New(store)
	route.RegisterRoutes(e, *service)

	e.HideBanner = true
	log.Fatal(e.Start(":9944"))
}
