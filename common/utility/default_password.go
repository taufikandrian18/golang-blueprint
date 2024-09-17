package utility

import (
	"encoding/base64"
	"fmt"
	"math/rand"

	sqlc "gitlab.com/wit-id/project-latihan/src/repository/pgbo_sqlc"
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

func GenerateDefaultPassword(emp sqlc.Employee) string {
	return fmt.Sprint(
		emp.DateOfBirth.Time.Format("02012006"),
	)
}
