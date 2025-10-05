package adapter

import (
	"RainmanwareKYC/internal/dto"
	"encoding/json"
	"log"
	"time"
)

type TyronKycAdapter struct {
}

func (kyc *TyronKycAdapter) CheckEntity(Entity *dto.KycEntity) KycResult {
	time.Sleep(5 * time.Second)
	jsonObj, _ := json.Marshal(Entity)
	log.Printf("[Tyron KYC Adapter] Screening entity :: %s", jsonObj)
	return KycResult{Status: SUCCESS, ResultMessage: "No hit"}
}

func (kyc *TyronKycAdapter) GetEntityById(Id uint64) *dto.KycEntity {
	return nil
}

func (kyc *TyronKycAdapter) SetEntityId(NewId uint64, OldId uint64) KycResult {
	return KycResult{Status: SUCCESS, ResultMessage: "Successfully updated id"}
}
