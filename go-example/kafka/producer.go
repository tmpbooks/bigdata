package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     "kltao",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	var value string
	for {
		_, err := fmt.Scanf("%s", &value)
		if err != nil {
			fmt.Println("input err: ", err)
		} else {
			msg.Value = sarama.ByteEncoder(value)
			fmt.Println(value)

			partition, offset, err := producer.SendMessage(msg)
			if err != nil {
				fmt.Println("Send message Fail")
			}
			fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
		}
	}
}
