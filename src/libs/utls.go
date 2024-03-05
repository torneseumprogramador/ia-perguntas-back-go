package libs

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

func IsCrypto(senha string) bool {
	return len(senha) == 60
}

func Crypto(senha string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func CryptoEq(senhaUnCrypto, senhaHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaUnCrypto))
	return err == nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
