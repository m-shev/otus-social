package broker

import (
	"context"
	"encoding/json"
	"github.com/m-shev/otus-social/back/internal/config"
	"github.com/segmentio/kafka-go"
	"log"
	"net"
	"strconv"
)

type Broker struct {
	config config.Broker
	w      *kafka.Writer
	log    *log.Logger
}

func NewBroker(conf config.Broker, logger *log.Logger) (*Broker, error) {
	err := createTopic(conf.BrokerUrls[0], conf.PostTopic)

	if err != nil {
		return nil, err
	}

	b := &Broker{config: conf, log: logger, w: createWriter(conf, logger)}

	return b, nil
}

func (b *Broker) WriteJSON(key string, val interface{}) error {
	bytes, err := json.Marshal(val)

	if err != nil {
		return err
	}

	err = b.w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: bytes,
	})

	return err
}

func createWriter(conf config.Broker, log *log.Logger) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(conf.BrokerUrls...),
		Topic:        conf.PostTopic.Name,
		RequiredAcks: kafka.RequireAll,
		Balancer:     &kafka.LeastBytes{},
		Compression:  kafka.Snappy,
		Logger:       kafka.LoggerFunc(log.Printf),
		ErrorLogger:  kafka.LoggerFunc(log.Printf),
	}
}

func createTopic(url string, conf config.Topic) error {
	conn, err := getController(url)

	if err != nil {
		return err
	}

	topicConfig := kafka.TopicConfig{

		Topic:             conf.Name,
		NumPartitions:     conf.NumPartitions,
		ReplicationFactor: conf.ReplicationFactor,
	}

	return conn.CreateTopics(topicConfig)
}

func getController(url string) (con *kafka.Conn, err error) {
	conn, err := kafka.Dial("tcp", url)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	controller, err := conn.Controller()

	if err != nil {
		return nil, err
	}

	controllerConn, err := kafka.Dial("tcp",
		net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))

	return controllerConn, err
}
