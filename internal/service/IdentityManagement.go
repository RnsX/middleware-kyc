package service

import (
	"RainmanwareKYC/internal/dto"
	"fmt"
	"hash/fnv"
)

type IdentityManager interface {
	GetId(entity dto.KycEntity) uint64
}

type TyronIdentityManager struct {
}

func (im *TyronIdentityManager) GetId(entity dto.KycEntity) uint64 {
	required := []string{"Name", "Surname", "DOB", "SSN"}
	key := ""

	for _, p := range required {
		prop, exists := entity.Properties[p]
		if !exists {
			fmt.Printf("Entity does not contain all required properties for generating id")
			return 0
		}
		key += prop.(string)
	}

	h := fnv.New64()
	h.Write([]byte(key))
	return h.Sum64()
}
