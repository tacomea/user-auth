package delivery

import (
	"log"
	"net/http"
	"net/url"
	"userCreation/domain"
	"userCreation/token"
)

type logoutHandler struct {
	userUseCase    domain.UserUseCase
	sessionUseCase domain.SessionUseCase
}

func LogoutHandler(uu domain.UserUseCase, su domain.SessionUseCase) {
	handler := &logoutHandler{
		userUseCase:    uu,
		sessionUseCase: su,
	}
	http.HandleFunc("/logout", handler.Logout)
}

func (h *logoutHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		query := url.QueryEscape("You cannot when you are not logged in")
		http.Redirect(w, r, "/?msg="+query, http.StatusSeeOther)
		return
	}

	sessionId, err := token.ParseToken(cookie.Value)
	if err != nil {
		query := url.QueryEscape("Logout: Cookie Modified")
		http.Redirect(w, r, "/?msg="+query, http.StatusSeeOther)
		return
	}

	err = h.sessionUseCase.Delete(sessionId)
	if err != nil {
		log.Println("session was not deleted: ", err)
	}

	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	query := url.QueryEscape("successfully logged out")
	http.Redirect(w, r, "/?msg="+query, http.StatusSeeOther)
}

//func LogoutHandler(c *fiber.App) {
//	c.Post("/logout", Logout)
//}
//
//func Logout(c *fiber.Ctx) error {
//	sr := repository.NewSyncMapSessionRepository()
//	su := usecase.NewSessionUsecase(sr)
//
//	cookie := c.Cookies("session")
//
//	sessionId, err := token.ParseToken(cookie)
//	if err != nil {
//		query := url.QueryEscape("Logout: Cookie Modified")
//		err2 := c.Redirect("/?msg=" + query, http.StatusSeeOther)
//		if err2 != nil {
//			c.Status(400)
//			return c.JSON(fiber.Map{
//				"message": "redirect failed" + err.Error(),
//			})
//		}
//	}
//
//	err = su.Delete(sessionId)
//	if err != nil {
//		c.Status(500)
//		return c.JSON(fiber.Map{
//			"message": "Delete()" + err.Error(),
//		})
//	}
//
//	c.ClearCookie("session")
//	query := url.QueryEscape("successfully logged out")
//	err = c.Redirect("/?msg=" +query, http.StatusSeeOther)
//	if err != nil {
//		c.Status(500)
//		return c.JSON(fiber.Map{
//			"message": "redirect failed" + err.Error(),
//		})
//	}
//	return c.JSON("Logout OK")
//}
