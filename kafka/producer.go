package kafka

import (
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func Produce(msg string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Set up producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Kafka producer:", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln("Failed to close Kafka producer:", err)
		}
	}()

	// Set up signal handler to handle interrupt
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Produce messages
	go func() {
		_, _, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: "messages",
			Value: sarama.StringEncoder(msg),
		})
		if err != nil {
			log.Println("Failed to produce message:", err)
		} else {
			log.Println("Produced message:", msg)
		}
	}()

	<-sigchan
}
