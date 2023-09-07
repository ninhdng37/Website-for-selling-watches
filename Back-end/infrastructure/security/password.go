package security

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	sha256Hash := sha256.New()

	// Write the data to the hash object
	sha256Hash.Write([]byte(password))

	// Get the final SHA-256 hash as a byte slice
	hashResult := sha256Hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashHex := hex.EncodeToString(hashResult)
	return hashHex
}

func VerifyPassword(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}
