package main

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

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

	// GABUNG KE ...
	err = consumer.Subscribe("mahasiswa", nil)
	if err != nil {
		panic(err)
	}

}
