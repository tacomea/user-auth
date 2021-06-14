package delivery

import (
	"html/template"
	"log"
	"net/http"
	"userCreation/domain"
)

type enterHandler struct {
	userUseCase    domain.UserUseCase
	sessionUseCase domain.SessionUseCase
}

func EnterHandler(uu domain.UserUseCase, su domain.SessionUseCase) {
	handler := &enterHandler{
		userUseCase:    uu,
		sessionUseCase: su,
	}
	http.HandleFunc("/enter", handler.Enter)
}

func (h *enterHandler) Enter(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")

	html, err := template.ParseFiles("templates/enter.gohtml")
	if err != nil {
		_, err := w.Write([]byte("500: internal server error"))
		if err != nil {
			log.Println("Error in WriteString: ", err)
		}
		return
	}

	err = html.Execute(w, msg)
	if err != nil {
		log.Println("Error in WriteString: ", err)
	}
}
