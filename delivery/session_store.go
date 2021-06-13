package delivery

import (
	"github.com/gofiber/fiber/v2"
	"userCreation/domain"
)

type sessionStoreHandler struct {
	sessionUseCase domain.SessionUseCase
}

func SessionStoreHandler(c *fiber.App, su domain.SessionUseCase) {
	handler := &sessionStoreHandler{
		sessionUseCase: su,
	}

	c.Post("/session/load", handler.SessionStore)
}

func (h *sessionStoreHandler) SessionStore(c *fiber.Ctx) error {
	session := new(domain.Session)
	err := c.BodyParser(session)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "unexpected request: " + err.Error(),
		})
	}

	err = h.sessionUseCase.Store(*session)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "internal server error" + err.Error(),
		})
	}
	return c.JSON("Store OK")
}
