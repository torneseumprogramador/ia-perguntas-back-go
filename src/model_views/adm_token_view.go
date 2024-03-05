package model_views

type AdmTokenView struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Nome  string `json:"nome"`
	Super bool   `json:"super"`
	Token string `json:"token"`
}
