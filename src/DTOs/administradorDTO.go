package DTOs

type AdministradorDTO struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
	Super bool   `json:"super"`
}
