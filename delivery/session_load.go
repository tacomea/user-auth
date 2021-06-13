package delivery

import (
	"github.com/gofiber/fiber/v2"
	"userCreation/domain"
)

type sessionLoadHandler struct {
	sessionUseCase domain.SessionUseCase
}

func SessionLoadHandler(c *fiber.App, su domain.SessionUseCase) {
	handler := &sessionLoadHandler{
		sessionUseCase: su,
	}

	c.Post("/session/load", handler.SessionLoad)
}

func (h *sessionLoadHandler) SessionLoad(c *fiber.Ctx) error {
	session := new(domain.Session)
	err := c.BodyParser(session)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "unexpected request: " + err.Error(),
		})
	}

	resSession, err := h.sessionUseCase.Load(session.ID)
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
	return c.JSON(resSession)
}
