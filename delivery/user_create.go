package delivery

import (
	"github.com/gofiber/fiber/v2"
	"userCreation/domain"
)

type userCreateHandler struct {
	userUseCase domain.UserUseCase
}

func UserCreateHandler(c *fiber.App, uu domain.UserUseCase) {
	handler := &userCreateHandler{
		userUseCase: uu,
	}

	c.Post("/user/create", handler.UserCreate)
}

func (h *userCreateHandler) UserCreate(c *fiber.Ctx) error {
	user := new(domain.User)
	err := c.BodyParser(user)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "unexpected request" + err.Error(),
		})
	}

	err = h.userUseCase.Create(*user)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "internal server error" + err.Error(),
		})
	}
	return c.JSON("Create OK")
}
