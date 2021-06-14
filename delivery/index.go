package delivery

import (
	"html/template"
	"log"
	"net/http"
	"userCreation/domain"
	"userCreation/token"
)

// standard library

type indexHandler struct {
	userUseCase    domain.UserUseCase
	sessionUseCase domain.SessionUseCase
}

func IndexHandler(uu domain.UserUseCase, su domain.SessionUseCase) {
	handler := &indexHandler{
		userUseCase:    uu,
		sessionUseCase: su,
	}
	http.HandleFunc("/", handler.Index)
}

func (h *indexHandler) Index(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")

	html, err := template.ParseFiles("templates/index.gohtml")

	cookie, err := r.Cookie("session")
	if err == nil {
		sessionId, err := token.ParseToken(cookie.Value)
		if err != nil {
			msg = "cookie modified"
		} else {
			session, err := h.sessionUseCase.Load(sessionId)
			if err == nil {
				msg = "from cookie, email: " + session.Email
			}
		}
	}

	err = html.Execute(w, msg)
	if err != nil {
		log.Println("Error in WriteString: ", err)
	}
}

// fiber

//func IndexHandler(c *fiber.App) {
//	c.Get("/", Index)
//}
//
//func Index(c *fiber.Ctx) error {
//	sr := repository.NewSyncMapSessionRepository()
//	su := usecase.NewSessionUsecase(sr)
//
//	msg := c.FormValue("msg")
//	html := `<!doctype html>
//<html lang="en">
//<head>
//<meta charset="utf-8">
//<title>Document</title>
//</head>
//<body>
//<p>%s</p>
//<h1>Account Creation</h1>
//<form action="/register" method="post">
//	<label>Email</label>
//	<input type="email" name="email" />
//	<label>password</label>
//	<input type="password" name="password"/>
//	<input type="submit" />
//</form>
//<h1>Login</h1>
//<form action="/login" method="post">
//	<label>user name</label>
//	<input type="email" name="email" />
//	<label>password</label>
//	<input type="password" name="password"/>
//	<input type="submit" />
//</form>
//<h1>Logout</h1>
//<form action="/logout" method="POST">
//	<input type="submit" value="logout">
//</form>
//</body>
//</html>
//`
//
//	cookie := c.Cookies("session")
//	if cookie == "" {
//		sessionId, err := token.ParseToken(cookie)
//		if err != nil {
//			msg = "cookie modified"
//		} else {
//			session, err := su.Load(sessionId)
//			if err != nil {
//				msg = "from cookie, username: " + session.Email
//			}
//		}
//	}
//
//	html = fmt.Sprintf(html, msg)
//	err := c.Send([]byte(html))
//	if err != nil {
//		c.Status(400)
//		return c.JSON(fiber.Map{
//			"message": "internal server error: " + err.Error(),
//		})
//	}
//
//	return nil
//}
