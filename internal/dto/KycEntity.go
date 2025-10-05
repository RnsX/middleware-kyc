package dto

type KycEntity struct {
	TempId     uint64                 `json:"temp_id"`
	PermId     uint64                 `json:"perm_id"`
	Properties map[string]interface{} `json:"properties"`
}
