package api

import (
	"ecommerce-app/internal/api/rest"
	"ecommerce-app/internal/api/rest/handlers"
	"ecommerce-app/config"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,

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