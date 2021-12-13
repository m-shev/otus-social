package api

import (
	"github.com/m-shev/otus-social/dialog/internal/configuration"
	dbconn "github.com/m-shev/otus-social/dialog/internal/db-connector"
	"github.com/m-shev/otus-social/dialog/internal/services/dialog"
	"log"
)

type Api struct {
	config        configuration.Configuration
	dialogService *dialog.Service
}

func NewApi(config configuration.Configuration) *Api {
	dialogService := createDialogService(config.DialogDb)

	return &Api{
		config:        config,
		dialogService: dialogService,
	}
}

func createDialogService(config configuration.DbConfig) *dialog.Service {
	dialogConn := dbconn.NewDbConnector()
	err := dialogConn.AddDb(config)

	handleErr(err)

	conn, err := dialogConn.GetConnection(config.DbId)

	handleErr(err)

	return dialog.NewDialogService(conn)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
