package adapter

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type KafkaRequestManager struct {
	consumer               *kafka.Consumer
	screeningRequestTopic  string
	screeningResponseTopic string
	requestHandler         RequestHandler
}

func (rm *KafkaRequestManager) SetRequestHandler(handler RequestHandler) {
	rm.requestHandler = handler
}

func NewKafkaRequestManager(brokers, groupID string) (*KafkaRequestManager, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to create consumer %w", err)
	}
	return &KafkaRequestManager{
		consumer:               c,
		screeningRequestTopic:  "kyc-requests",
		screeningResponseTopic: "kyc-responses",
	}, nil
}

func (rm *KafkaRequestManager) Start(ctx context.Context) error {
	err := rm.consumer.SubscribeTopics([]string{rm.screeningRequestTopic}, nil)
	if err != nil {
		return fmt.Errorf("Failed to subscribe to topic %s: %w", rm.screeningRequestTopic, err)
	}
	run := true
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for run {
		select {
		case <-ctx.Done():
			log.Println("[Kafka consumer] Context canceled, shutting down")
			run = false
		case sig := <-sigchan:
			log.Printf("[Kafka consumer] Caught signal %v, shutting down...", sig)
		default:
			ev := rm.consumer.Poll(100)
			if ev == nil {
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				if rm.requestHandler != nil {
					go rm.requestHandler(EntityCheckRequest{Payload: e})
				}
			case kafka.Error:
				log.Printf("[Kafka consumer] Error: %v", e)
				if e.IsFatal() {
					run = false
				}
			default:
				// ignore
			}
		}
	}
	rm.consumer.Close()
	log.Println("[Kafka consumer] Closed consumer")
	return nil
}
