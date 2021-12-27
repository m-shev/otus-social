package message

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/m-shev/otus-social/dialog/internal/configuration"
	"github.com/m-shev/otus-social/dialog/internal/consistent"
	"github.com/m-shev/otus-social/dialog/internal/db-connector"
	"strconv"
	"time"
)

const queryTimeout = time.Second * 5

type Service struct {
	getConnection func(id string) (*sql.DB, error)
	ring          *consistent.Ring
}

func NewMessageService(dbConfigs []configuration.DbConfig) (*Service, error) {
	conn, err := createDbConn(dbConfigs)

	if err != nil {
		return nil, err
	}

	ring, err := createRing(dbConfigs)

	if err != nil {
		return nil, err
	}

	return &Service{getConnection: conn.GetConnection, ring: ring}, nil
}

func createDbConn(dbConfigs []configuration.DbConfig) (*dbconn.Connector, error) {
	conn := dbconn.NewDbConnector()
	err := conn.AddDbList(dbConfigs)

	return conn, err
}

func createRing(dbConfigs []configuration.DbConfig) (*consistent.Ring, error) {
	shards := make([]consistent.ShardConfig, 0)

	for _, v := range dbConfigs {
		shards = append(shards, consistent.ShardConfig{NodeId: v.DbId, TargetId: ""})
	}

	ring, err := consistent.NewRing(shards)

	return ring, err
}

func (s *Service) getShardConn(dialogId int64) (*sql.DB, error) {
	id := strconv.FormatInt(dialogId, 10)
	dbId := s.ring.GetNode(id)

	conn, err := s.getConnection(dbId)

	return conn, err
}

func (s *Service) CreateMessage(form CreateMessageForm) (Message, error) {
	db, err := s.getShardConn(form.DialogId)

	if err != nil {
		return Message{}, err
	}

	uuid, err := s.generateUuid(db)
	fmt.Println(uuid)
	if err != nil {
		return Message{}, err
	}

	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)
	query := `insert into message(message_id, dialog_id, author_id, content) values(UUID_TO_BIN(?), ?, ?, ?)`

	_, err = db.ExecContext(ctx, query, uuid, form.DialogId, form.AuthorId, form.Content)

	if err != nil {
		return Message{}, err
	}

	return s.getMessage(db, uuid)
}

func (s *Service) getMessage(db *sql.DB, messageId string) (Message, error) {
	query := `select 
       bin_to_uuid(message_id) message_id, 
       dialog_id, author_id, 
       content, 
       create_at
	from message where message_id = uuid_to_bin(?)`

	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)

	row := db.QueryRowContext(ctx, query, messageId)

	m := Message{}

	err := row.Scan(&m.MessageId, &m.DialogId, &m.AuthorId, &m.Content, &m.CreateAt)

	return m, err
}

func (s *Service) generateUuid(db *sql.DB) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)
	query := `select uuid()`

	var uuid string

	row := db.QueryRowContext(ctx, query)

	err := row.Scan(&uuid)

	return uuid, err
}
