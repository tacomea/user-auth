package delivery

import (
	"encoding/base64"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/url"
	"userCreation/domain"
	"userCreation/token"
)

type loginHandler struct {
	userUseCase    domain.UserUseCase
	sessionUseCase domain.SessionUseCase
}

func LoginHandler(uu domain.UserUseCase, su domain.SessionUseCase) {
	handler := &loginHandler{
		userUseCase:    uu,
		sessionUseCase: su,
	}
	http.HandleFunc("/login", handler.Login)
}

func (h *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	encodedEmail := base64.StdEncoding.EncodeToString([]byte(email))

	user, err := h.userUseCase.Check(encodedEmail)
	if err != nil {
		query := url.QueryEscape("username doesn't exist")
		http.Redirect(w, r, "/?msg="+query, http.StatusSeeOther)
		return
	}
	hashedPassword := user.Password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		query := url.QueryEscape("login failed")
		http.Redirect(w, r, "/?msg="+query, http.StatusSeeOther)
	} else {
		sessionId := uuid.NewString()
		err := h.sessionUseCase.Store(domain.Session{
			ID:    sessionId,
			Email: email,
		})
		t, err := token.CreateToken(sessionId)
		if err != nil {
			log.Println("Error in createToken(): ", err)
			query := url.QueryEscape("Server Error, Try Again")
			http.Redirect(w, r, "/?msg="+query, http.StatusInternalServerError)
			return
		}
		cookie := http.Cookie{
			Name:  "session",
			Value: t,
		}
		http.SetCookie(w, &cookie)
		query := url.QueryEscape("logged in")
		http.Redirect(w, r, "/?msg="+query, http.StatusSeeOther)
	}
}

//func LoginHandler(c *fiber.App) {
//	c.Post("/login", Login)
//}
//
//func Login(c *fiber.Ctx) error {
//	ur := repository.NewSyncMapUserRepository()
//	uu := usecase.NewUserUsecase(ur)
//	sr := repository.NewSyncMapSessionRepository()
//	su := usecase.NewSessionUsecase(sr)
//
//	email := c.FormValue("email")
//	password := c.FormValue("password")
//	user, err := uu.Check(email)
//	if err != nil {
//		query := url.QueryEscape("username doesn't exist")
//		err := c.Redirect("/?msg="+query, http.StatusSeeOther)
//		if err != nil {
//			c.Status(400)
//			return c.JSON(fiber.Map{
//				"message": "redirect failed: " + err.Error(),
//			})
//		}
//		return c.JSON("user doesn't exist")
//	}
//
//	hashedPassword := user.Password
//	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
//	if err != nil {
//		query := url.QueryEscape("login failed")
//		err := c.Redirect("/?msg="+query, http.StatusSeeOther)
//		if err != nil {
//			c.Status(400)
//			return c.JSON(fiber.Map{
//				"message": "redirect failed: " + err.Error(),
//			})
//		}
//		return c.JSON("login failed")
//	} else {
//		sessionId := uuid.NewString()
//		err := su.Store(domain.Session{
//			ID:    sessionId,
//			Email: email,
//		})
//		if err != nil {
//			c.Status(500)
//			return c.JSON(fiber.Map{
//				"message": "store() failed: " + err.Error(),
//			})
//		}
//
//		t, err := token.CreateToken(sessionId)
//		if err != nil {
//			log.Println("Error in createToken(): ", err)
//			query := url.QueryEscape("Server Error, Try Again")
//			err2 := c.Redirect("/?msg="+query, http.StatusInternalServerError)
//			if err2 != nil {
//				c.Status(400)
//				return c.JSON(fiber.Map{
//					"message": "redirect failed: " + err2.Error(),
//				})
//			}
//			return c.JSON(fiber.Map{
//				"message": "server error: " + err.Error(),
//			})
//		}
//		c.Cookie(&fiber.Cookie{
//			Name: "session",
//			Value: t,
//		})
//		query := url.QueryEscape("logged in")
//		err = c.Redirect("/?msg="+query, http.StatusSeeOther)
//		if err != nil {
//			c.Status(400)
//			return c.JSON(fiber.Map{
//				"message": "redirect failed: " + err.Error(),
//			})
//		}
//	}
//	return c.JSON(fiber.Map{
//		"message": "Login OK",
//	})
//}
