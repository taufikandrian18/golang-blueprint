package utility

import (
	"encoding/base64"
	"math/rand"

	"golang.org/x/crypto/scrypt"
)

func GenerateSalt() string {
	salt := make([]byte, 16)
	rand.Read(salt)
	return base64.URLEncoding.EncodeToString(salt)
}

func HashPassword(password, salt string) string {
	saltedPassword := []byte(password + salt)
	hashedPassword, _ := scrypt.Key(saltedPassword, []byte(salt), 16384, 8, 1, 32)
	return base64.URLEncoding.EncodeToString(hashedPassword)
}
