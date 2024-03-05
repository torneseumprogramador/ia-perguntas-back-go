package DTOs

import "http_gin/src/enums"

type PetDTO struct {
	Nome   string     `json:"nome"`
	DonoId string     `json:"donoId"`
	Tipo   enums.Tipo `json:"tipo"`
}
