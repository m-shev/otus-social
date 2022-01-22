package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	createTopic()
	logger := log.Default()

	go produce(logger)
	consume(logger)
	time.Sleep(time.Minute * 5)
}

func createTopic() {

	conn, err := kafka.Dial("tcp", "localhost:49171")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()

	if err != nil {
		panic(err.Error())
	}

	var controllerConn *kafka.Conn

	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))

	if err != nil {
		panic(err.Error())
	}

	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             "test-topic-2",
			NumPartitions:     2,
			ReplicationFactor: 2,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)

	if err != nil {
		panic(err.Error())
	}

	partitions, err := controllerConn.ReadPartitions()

	for _, v := range partitions {
		fmt.Println(v)
	}

}

func consume(logger *log.Logger) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:49171", "localhost:49173"},
		Topic:       "test-topic-2",
		Logger:      kafka.LoggerFunc(logger.Printf),
		ErrorLogger: kafka.LoggerFunc(logger.Printf),
		GroupID:     "test-group-id",
	})

	for {
		msg, err := r.ReadMessage(context.Background())

		if err != nil {
			logger.Print(err.Error())
		}

		fmt.Println("received message: ", string(msg.Value))
	}
}

func produce(logger *log.Logger) {
	w := &kafka.Writer{
		Addr:        kafka.TCP("localhost:49171", "localhost:49173"),
		Topic:       "test-topic-2",
		Balancer:    &kafka.LeastBytes{},
		Compression: kafka.Snappy,
		Logger:      kafka.LoggerFunc(logger.Printf),
		ErrorLogger: kafka.LoggerFunc(logger.Printf),
	}

	i := 0
	for {
		i++
		key := strconv.Itoa(i)
		message := fmt.Sprintf("message №%d", i)
		err := w.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(key),
			Value: []byte(message),
		})

		if err != nil {
			logger.Print(err.Error())
		}

		time.Sleep(time.Second * 10)
	}
}
