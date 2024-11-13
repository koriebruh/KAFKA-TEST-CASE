package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

// HASIL NAMA YG DI KIRIM OLEH PRODUCER AKAN DI DI ADD KE DB DAN DI BUKANA
// NIM SECARA OTOMATIS DAN DI PILIHKAN DOSEN WALI SECARA RANDOM
func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "go_clientbg",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// GABUNG KE TOPIC ...
	err = consumer.Subscribe("mahasiswa", nil)
	if err != nil {
		panic(err)
	}

	for {
		// IN HERE GO LISTEN
		message, err := consumer.ReadMessage(1 * time.Second)
		if err == nil {
			fmt.Printf("Recive message : %s", message.Value)
		}
	}

}
