package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/internal/config"
	"github.com/m-shev/otus-social/internal/migration"
	"log"
)

func main() {
	logger := log.Default()
	conf := config.GetConfig()
	m := migration.NewManager(conf.Db, logger)
	m.Up()
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