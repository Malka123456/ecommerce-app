package api

import (
	"ecommerce-app/config"
	"ecommerce-app/internal/api/rest"
	"ecommerce-app/internal/api/rest/handlers"
	"ecommerce-app/internal/domain"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()


  db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil{
		log.Fatalf("db connection error")
	}

	log.Println("db connected")

	//run migration
	db.AutoMigrate(&domain.User{})
	


	rh := &rest.RestHandler{
		App: app,
		DB:  db,

	}

	setupRoutes(rh)

	
	app.Listen(config.ServerPort)
}			


func setupRoutes(rh *rest.RestHandler){
	  // user handler 
		handlers.SetupUserRoutes(rh)   

		//transaction
		//catalog

}