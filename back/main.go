package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
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