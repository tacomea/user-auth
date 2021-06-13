package delivery

import (
	"github.com/gofiber/fiber/v2"
	"userCreation/domain"
)

type userCheckHandler struct {
	userUseCase domain.UserUseCase
}

func UserCheckHandler(c *fiber.App, uu domain.UserUseCase) {
	handler := &userCheckHandler{
		userUseCase: uu,
	}

	c.Post("/user/check", handler.UserCheck)
}

func (h *userCheckHandler) UserCheck(c *fiber.Ctx) error {
	user := new(domain.User)
	err := c.BodyParser(user)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "unexpected request: " + err.Error(),
		})
	}

	resUser, err := h.userUseCase.Check(user.Email)
	//emptyUser := domain.User{}
	//if resUser == emptyUser {
	//	c.Status(400)
	//	return c.JSON(fiber.Map{
	//		"message": "user not found",
	//	})
	//}
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "internal server error: " + err.Error(),
		})
	}
	return c.JSON(resUser)
}
