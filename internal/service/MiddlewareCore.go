package service

import (
	"RainmanwareKYC/internal/adapter"
	"RainmanwareKYC/internal/dto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type MiddlewareCore struct {
	RequestManager  adapter.RequestManagerAdapter
	KycSystem       adapter.KycSystemAdapter
	IdManagerModule IdentityManager
}

func (mw *MiddlewareCore) EntityCheckRequest(request adapter.EntityCheckRequest) {
	msg := request.Payload.(*kafka.Message)

	log.Printf("[Middleware Core] Received message: topic=%s partition=%d offset=%d key=%s value=%s",
		*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset,
		string(msg.Key), string(msg.Value),
	)

	var entity dto.KycEntity

	err := json.Unmarshal(msg.Value, &entity)

	if err != nil {
		fmt.Println("(check entity request) Error while deserializing to KycEntity")
		return
	}
	tempId := mw.IdManagerModule.GetId(entity)

	if tempId == 0 {
		fmt.Println("(check entity request) Unable to generate temporary id")
		return
	}
	entity.TempId = tempId
	result := mw.KycSystem.CheckEntity(&entity)

	log.Printf("Middleware Core] Entity Check Result :: %s", result.ResultMessage)
	// TODO: now after this publish result message to 'kafka' or appropriate request manager
}

func (mw *MiddlewareCore) Start(ctx context.Context) error {
	log.Println("Starting Request Manager for Middleware Core")
	return mw.RequestManager.Start(ctx)
}

// temporary implementation. Constructor should take in all components that need to be instantiatetd
func NewMiddlewareCoreDefault() *MiddlewareCore {
	kafkaAdapter, err := adapter.NewKafkaRequestManager("localhost:9092", "example-group")

	if err != nil {
		fmt.Printf("Error while creating new kafka request manager")
		return nil
	}

	mw := &MiddlewareCore{
		RequestManager:  kafkaAdapter,
		KycSystem:       &adapter.TyronKycAdapter{},
		IdManagerModule: &TyronIdentityManager{},
	}

	mw.RequestManager.SetRequestHandler(mw.EntityCheckRequest)

	return mw
}
