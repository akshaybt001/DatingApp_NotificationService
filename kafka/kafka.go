package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	email "github.com/akshaybt001/DatingApp_NotificationService/internal/helper/EmailSend"
)

type MatchUser struct {
	Name    string `json:"Name"`
	Email   string `json:"Email"`
	Message string `json:"Message"`
}

func StartConsumingLikeUser() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true

	consumer, err := sarama.NewConsumer([]string{"apache-kafka-service:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("MatchUser", 0, sarama.OffsetNewest)
	fmt.Println("offset ", sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)

	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var matchUser MatchUser
			err := json.Unmarshal(msg.Value, &matchUser)
			fmt.Println("message received")
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}
			go func(user MatchUser) {
				message := fmt.Sprintf("You have match with %s .\n Hurry up visit the app for chat !!", user.Name)
				if err := email.SendEmail(user.Email, message); err != nil {
					log.Println(err)
				}
			}(matchUser)
		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}
