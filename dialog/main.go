package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/dialog/api"
	"github.com/m-shev/otus-social/dialog/internal/configuration"
	"github.com/m-shev/otus-social/dialog/internal/migration"
	"log"
	"net/http"
	"time"
)

func main() {
	config := configuration.GetConfig()

	time.Sleep(config.Server.StartDelay * time.Second)

	runMigrate(config.Db)
	r := router(config)
	addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Printf("server start at %s\n", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Print("server start error: ", err)
	}
}

func router(conf configuration.Configuration) *gin.Engine {
	a := api.NewApi(conf)

	handler := gin.New()

	handler.Use(gin.Logger())

	handler.POST("dialog", a.CreateDialog)
	handler.POST("dialog/member", a.AddDialogMember)
	handler.GET("dialog/:dialogId", a.GetDialogById)

	return handler
}

func runMigrate(dbConfig configuration.Db) {
	h := migration.NewMigrationHelper()
	dialog := migration.DbConfig{
		User:          dbConfig.DialogDb.User,
		Password:      dbConfig.DialogDb.Password,
		Host:          dbConfig.DialogDb.Host,
		Port:          dbConfig.DialogDb.Port,
		DbName:        dbConfig.DialogDb.DbName,
		MigrationPath: dbConfig.DialogDb.MigrationPath,
	}
	h.Up(dialog)
	//h.Force(1, dialog)
}
