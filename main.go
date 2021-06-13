package main

import (
	"log"
	"net/http"
	"userCreation/delivery"
	"userCreation/repository"
	"userCreation/usecase"
)

//func main() {
//	ur := repository.NewSyncMapUserRepository()
//	uu := usecase.NewUserUsecase(ur)
//
//	sr := repository.NewSyncMapSessionRepository()
//	su := usecase.NewSessionUsecase(sr)
//
//	c := fiber.New()
//
//	c.Use(cors.New(cors.Config{
//		AllowCredentials: true,
//	}))
//
//	delivery.UserCreateHandler(c, uu)
//	delivery.UserDeleteHandler(c, uu)
//	delivery.UserCheckHandler(c, uu)
//
//	delivery.SessionStoreHandler(c, su)
//	delivery.SessionDeleteHandler(c, su)
//	delivery.SessionLoadHandler(c, su)
//
//	delivery.IndexHandler(c)
//	delivery.RegisterHandler(c)
//	delivery.LoginHandler(c)
//	delivery.LogoutHandler(c)
//
//	//log.Fatalln(c.Listen(":80"))
//}

func main() {
	ur := repository.NewSyncMapUserRepository()
	uu := usecase.NewUserUsecase(ur)
	sr := repository.NewSyncMapSessionRepository()
	su := usecase.NewSessionUsecase(sr)

	delivery.IndexHandler(uu, su)
	delivery.RegisterHandler(uu, su)
	delivery.LoginHandler(uu, su)
	delivery.LogoutHandler(uu, su)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln("Error in ListenAndServe()", err)
	}
}
