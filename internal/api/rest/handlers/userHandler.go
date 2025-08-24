package handlers

import (
	"github.com/gofiber/fiber/v2"
	"ecommerce-app/internal/api/rest"
	"net/http"
)


type userHandler struct {
	//srv UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

		app := rh.App
      
    // create an instance of user service & inject to handler 

		handler := userHandler{}

		// public endpoints
		app.Post("/register", handler.Register)
		app.Post("/login", handler.Login)


		// Private endpoints
		app.Get("/verify", handler.GetVerificationCode)
		app.Post("/verify", handler.Verify)
		app.Post("/profile", handler.CreateProfile)
		app.Get("/profile", handler.GetProfile)

		app.Post("/cart", handler.AddToCart)
		app.Post("/cart", handler.GetCart)
		app.Get("/order", handler.GetOrders)
		app.Get("/order/:id", handler.GetOrder)

		app.Post("/become-deller", handler.BecomeSeller)
}

func (h *userHandler) Register(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
      "message": "register",
	})
}


func (h *userHandler) Login(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Login",
	})
}

func (h *userHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get verification Code",
	})
}

func (h *userHandler) Verify(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Verify",
	})
}

func (h *userHandler) CreateProfile(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Create Profile",
	})
}

func (h *userHandler) GetProfile(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get Profile",
	})
}

func (h *userHandler) AddToCart(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
	})
}


func (h *userHandler) GetCart(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get Cart",
	})
}

func (h *userHandler) CreateOrder(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Create Order",
	})
}

func (h *userHandler) GetOrders(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get Orders",
	})
}

func (h *userHandler) GetOrder(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get Order by id",
	})
}

func (h *userHandler) BecomeSeller(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Become seller",
	})
}