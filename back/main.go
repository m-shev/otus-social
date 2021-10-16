package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/internal/config"
	"log"
	"net/http"
)

func main() {
	conf := config.GetConfig()
	fmt.Print(conf)
	router := makeRouter()
	server := &http.Server{
		Addr: "0.0.0.0:3005",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Print("server start error: ", err)
	}
}

func makeRouter() *gin.Engine {
	handler := gin.New()

	return handler
}