package main
 
import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "sync"
 
    "github.com/IBM/sarama"
)
 
func Consumer() {
    // Set up consumer
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true
    config.Consumer.Offsets.Initial = sarama.OffsetNewest

    consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
    if err != nil {
        log.Fatalln("Failed to start Kafka consumer:", err)
    }
    defer func() {
        if err := consumer.Close(); err != nil {
            log.Fatalln("Failed to close Kafka consumer:", err)
        }
    }()
 
    // Set up signal handler to handle interrupt
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, os.Interrupt)
 
    // Set up consumer group
    consumerGroup, err := sarama.NewConsumerGroup(
        []string{"localhost:9092"},
        "messages", 
        config,
    )
    if err != nil {
        log.Fatalln("Failed to start Kafka consumer group:", err)
    }
    defer func() {
        if err := consumerGroup.Close(); err != nil {
            log.Fatalln("Failed to close Kafka consumer group:", err)
        }
    }()
 
    wg := &sync.WaitGroup{}
    wg.Add(1)
    go func() {
        for {
            err := <-consumerGroup.Errors()
            if err != nil {
                log.Println("Consumer group error:", err)
            }
        }
    }()
 
    // Consume messages
    go func() {
        partitionConsumer, err := consumer.ConsumePartition("messages", 0, sarama.OffsetNewest)
        if err != nil {
            log.Fatalln("Failed to start consumer:", err)
        }
        defer func() {
            if err := partitionConsumer.Close(); err != nil {
                log.Fatalln("Failed to close partition consumer:", err)
            }
        }()
 
        for {
            select {
            case msg := <-partitionConsumer.Messages():
                log.Printf("Received message: %s\n", msg.Value)
            case err := <-partitionConsumer.Errors():
                log.Println("Consumer error:", err)
            case <-sigchan:
                log.Println("Interrupt signal received. Shutting down consumer...")
                return
            }
        }
    }()
 
    <-sigchan
    fmt.Println("Shutting down consumer")
}

func main() {
    Consumer()
}