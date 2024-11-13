package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	topic := "student"

	//EXAMPLE SEND MESSAGE STATIC
	for i := 0; i < 10; i++ {
		fmt.Printf("send message to kafka %d \n", i)
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},

			Value: []byte(fmt.Sprintf("Message %d nih bg", i)),
		}
		err := producer.Produce(msg, nil)
		if err != nil {
			panic(err)
		}

		producer.Flush(5 * 1000)
	}

}
