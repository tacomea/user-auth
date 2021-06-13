package delivery

import (
	"github.com/gofiber/fiber/v2"
	"userCreation/domain"
)

type sessionDeleteHandler struct {
	sessionUseCase domain.SessionUseCase
}

func SessionDeleteHandler(c *fiber.App, uu domain.SessionUseCase) {
	handler := sessionDeleteHandler{
		sessionUseCase: uu,
	}

	c.Post("/session/delete", handler.SessionDelete)
}

func (h *sessionDeleteHandler) SessionDelete(c *fiber.Ctx) error {
	Session := new(domain.Session)
	err := c.BodyParser(Session)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "unexpected request" + err.Error(),
		})
	}

	err = h.sessionUseCase.Delete(Session.ID)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "internal server error" + err.Error(),
		})
	}
	return c.JSON("Delete OK")
}
