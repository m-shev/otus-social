package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/m-shev/otus-social/back/internal/config"
	"github.com/m-shev/otus-social/back/internal/services/notifier"
	"github.com/segmentio/kafka-go"
	"log"
	"net"
	"strconv"
	"time"
)

type Broker struct {
	config           config.Broker
	createPostWriter *kafka.Writer
	log              *log.Logger
}

func NewPostQueue(conf config.Broker, logger *log.Logger) (*Broker, error) {
	err := createTopic(conf.BrokerUrls[0], conf.PostTopic)

	if err != nil {
		return nil, err
	}

	b := &Broker{
		config:           conf,
		log:              logger,
		createPostWriter: createWriter(conf.BrokerUrls, conf.PostTopic, logger),
	}

	return b, nil
}

func (b *Broker) SendPostCreated(m notifier.MessagePostCreate) error {
	return b.writePostCreated(strconv.Itoa(m.Post.Id), m)
}

func (b *Broker) writePostCreated(key string, val interface{}) error {

	bytes, err := json.Marshal(val)

	if err != nil {
		return err
	}
	start := time.Now()
	err = b.createPostWriter.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: bytes,
	})
	fmt.Println("finished operation", time.Since(start))
	return err
}

func createWriter(urls []string, topic config.Topic, log *log.Logger) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(urls...),
		Topic:        topic.Name,
		RequiredAcks: kafka.RequireAll,
		Balancer:     &kafka.RoundRobin{},
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
