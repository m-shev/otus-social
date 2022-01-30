package reader

import (
	"github.com/m-shev/otus-social/feed/internal/configuration"
	"github.com/segmentio/kafka-go"
	"log"
)

func NewPostReader(config configuration.Broker, logger *log.Logger) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     config.BrokerUrls,
		Topic:       config.PostTopic.Name,
		Logger:      kafka.LoggerFunc(logger.Printf),
		ErrorLogger: kafka.LoggerFunc(logger.Printf),
		GroupID:     config.PostTopic.GroupId,
	})
}
