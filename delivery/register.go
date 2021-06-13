package delivery

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/url"
	"userCreation/domain"
)

type registerHandler struct {
	userUseCase    domain.UserUseCase
	sessionUseCase domain.SessionUseCase
}

func RegisterHandler(uu domain.UserUseCase, su domain.SessionUseCase) {
	handler := &registerHandler{
		userUseCase:    uu,
		sessionUseCase: su,
	}
	http.HandleFunc("/register", handler.Register)
}

func (h *registerHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error while hashing", err)
	}
	err = h.userUseCase.Create(domain.User{
		Email:    email,
		Password: encodedPassword,
	})

	query := url.QueryEscape("account successfully created")
	http.Redirect(w, r, "/?msg="+query, http.StatusSeeOther)
}

//func RegisterHandler(c *fiber.App) {
//	c.Post("/register", Register)
//}
//
//func Register(c *fiber.Ctx) error {
//	ur := repository.NewSyncMapUserRepository()
//	uu := usecase.NewUserUsecase(ur)
//
//	email := c.FormValue("email")
//	password := c.FormValue("password")
//
//	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		log.Println("Error while hashing", err)
//	}
//	err = uu.Create(domain.User{
//		Email: email,
//		Password: encodedPassword,
//	})
//	if err != nil {
//		c.Status(500)
//		return c.JSON(fiber.Map{
//			"message": "internal server error: " + err.Error(),
//		})
//	}
//
//	query := url.QueryEscape("account successfully created")
//	err = c.Redirect("/?msg="+query, http.StatusSeeOther)
//	if err != nil {
//		c.Status(500)
//		return c.JSON(fiber.Map{
//			"message": "internal server error: " + err.Error(),
//		})
//	}
//	return c.JSON("Register OK")
//}
