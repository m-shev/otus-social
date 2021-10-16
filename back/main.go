package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/internal/config"
	"github.com/m-shev/otus-social/internal/connector"
	"github.com/m-shev/otus-social/internal/migration"
	"github.com/m-shev/otus-social/internal/repositories/user"
	"log"
)

func main() {
	logger := log.Default()
	conf := config.GetConfig()
	m := migration.NewManager(conf.Db, logger)
	m.Up()
	conn := connector.NewDbConnector(conf.Db, logger)
	userRepository := user.NewRepository(conn)

	createdUser, err := userRepository.Create(user.CreateUserForm{
		Name:     "alex",
		Surname:  "shev",
		Age:      20,
		City:     "Moscow",
		Email:    "sobaka@some.ru",
		Password: "qwerty",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(createdUser)

	//router := makeRouter()
	//server := &http.Server{
	//	Addr: "0.0.0.0:3005",
	//	Handler: router,
	//}
	//
	//if err := server.ListenAndServe(); err != nil {
	//	log.Print("server start error: ", err)
	//}
}

func makeRouter() *gin.Engine {
	handler := gin.New()

	return handler
}