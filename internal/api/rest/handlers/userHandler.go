package handlers

import (
	"ecommerce-app/internal/api/rest"
	"ecommerce-app/internal/dto"
	"ecommerce-app/internal/repository"
	"ecommerce-app/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)


type userHandler struct {
	//srv UserService
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

		app := rh.App
      
    // create an instance of user service & inject to handler 
		svc := service.UserService{
			Repo: repository.NewUserRepository(rh.DB),
		}
		handler := userHandler{
			svc: svc,
		}

		// public endpoints
		app.Post("/register", handler.Register)
		app.Post("/login", handler.Login)


		// Private endpoints
		app.Get("/verify", handler.GetVerificationCode)
		app.Post("/verify", handler.Verify)
		app.Post("/profile", handler.CreateProfile)
		app.Get("/profile", handler.GetProfile)

		app.Post("/cart", handler.AddToCart)
		app.Get("/cart", handler.GetCart)
		app.Get("/order", handler.GetOrders)
		app.Get("/order/:id", handler.GetOrder)

		app.Post("/become-seller", handler.BecomeSeller)
}

func (h *userHandler) Register(ctx *fiber.Ctx) error {
	// to create user
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)
  if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide valid inputs",
		})
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Error on signup",
		})
	}


	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
      "message": token,
	})
}


func (h *userHandler) Login(ctx *fiber.Ctx) error {

	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)
  if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide valid inputs",
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "please provide correct user id password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Login",
		"token": token,
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