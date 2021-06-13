package delivery

import (
	"github.com/gofiber/fiber/v2"
	"userCreation/domain"
)

type userDeleteHandler struct {
	userUseCase domain.UserUseCase
}

func UserDeleteHandler(c *fiber.App, uu domain.UserUseCase) {
	handler := userDeleteHandler{
		userUseCase: uu,
	}

	c.Post("/user/delete", handler.UserDelete)
}

func (h *userDeleteHandler) UserDelete(c *fiber.Ctx) error {
	user := new(domain.User)
	err := c.BodyParser(user)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "unexpected request" + err.Error(),
		})
	}

	err = h.userUseCase.Delete(user.Email)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "internal server error" + err.Error(),
		})
	}
	return c.JSON("Delete OK")
}
