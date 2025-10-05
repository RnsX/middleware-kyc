package adapter

import "RainmanwareKYC/internal/dto"

type KycResultStatus int

const (
	SUCCESS KycResultStatus = iota
	FAILED
	ERROR
)

type KycResult struct {
	ResultMessage string
	Status        KycResultStatus
}

type KycSystemAdapter interface {
	CheckEntity(Entity *dto.KycEntity) KycResult
	GetEntityById(Id uint64) *dto.KycEntity
	SetEntityId(NewId uint64) KycResult
}
