package mrutils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateUUIDv4() string {
	// UUID memiliki panjang 16 byte (128 bit)
	uuid := make([]byte, 16)

	// Mengisi UUID dengan nilai random dari crypto/rand
	_, err := rand.Read(uuid)
	if err != nil {
		return ""
	}

	// Mengatur versi UUID ke 4 (4-bit pertama dari byte ke-7 harus diatur ke 0100)
	uuid[6] = (uuid[6] & 0x0F) | 0x40

	// Mengatur varian UUID sesuai dengan RFC 4122 (2-bit pertama dari byte ke-9 harus diatur ke 10)
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	// Format UUID ke dalam string dengan format standar "xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx"
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func GeneratePassword(length ...int) string {
	const charset = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+`

	const defaultPasswordLength = 12
	passwordLength := defaultPasswordLength
	if len(length) > 0 && length[0] > 0 {
		passwordLength = length[0]
	}

	password := make([]byte, passwordLength)

	// Mengisi password dengan karakter acak
	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password)
}
