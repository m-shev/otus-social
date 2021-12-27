package api

import (
	"github.com/m-shev/otus-social/dialog/internal/configuration"
	dbconn "github.com/m-shev/otus-social/dialog/internal/db-connector"
	"github.com/m-shev/otus-social/dialog/internal/services/dialog"
	"github.com/m-shev/otus-social/dialog/internal/services/message"
	"log"
)

type Api struct {
	config         configuration.Configuration
	dialogService  *dialog.Service
	messageService *message.Service
}

func NewApi(config configuration.Configuration) *Api {
	dialogService := createDialogService(config.DialogDb)
	messageService := createMessageService(config.MessageDbList)

	return &Api{
		config:         config,
		dialogService:  dialogService,
		messageService: messageService,
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

func createMessageService(dbList []configuration.DbConfig) *message.Service {
	s, err := message.NewMessageService(dbList)

	handleErr(err)

	return s
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
