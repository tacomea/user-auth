package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"userCreation/delivery"
	"userCreation/domain"
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

var (
	schema			= "%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local"
	username		= os.Getenv("MYSQL_USER")
	password		= os.Getenv("MYSQL_PASSWORD")
	userDbName		= os.Getenv("MYSQL_DATABASE")
	dataSourceName	= fmt.Sprintf(schema, username, password, userDbName)
)

func connect() *gorm.DB {
	connection, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Panic("Could not connect to the database")
	}

	connection.AutoMigrate(&domain.User{})
	connection.AutoMigrate(&domain.Session{})

	return connection
}

func main() {
	// MySQL
	db := connect()
	ur := repository.NewUserRepositoryMySQL(db)
	sr := repository.NewSessionRepositoryMySQL(db)

	// sync.Map
	//ur := repository.NewSyncMapUserRepository()
	//sr := repository.NewSyncMapSessionRepository()

	uu := usecase.NewUserUsecase(ur)
	su := usecase.NewSessionUsecase(sr)

	delivery.IndexHandler(uu, su)
	delivery.RegisterHandler(uu, su)
	delivery.LoginHandler(uu, su)
	delivery.LogoutHandler(uu, su)
	delivery.EnterHandler(uu, su)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln("Error in ListenAndServe()", err)
	}
}
