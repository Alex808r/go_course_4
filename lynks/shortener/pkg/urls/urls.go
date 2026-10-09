package urls

import (
	"crypto/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"

// Shorten генерирует короткий идентификатор для URL.
func Shorten(src string) string {
	if src == "" {
		return ""
	}
	b := make([]byte, 5)
	_, _ = rand.Read(b)
	for i := range b {
		b[i] = alphabet[int(b[i])%len(alphabet)]
	}
	return string(b)
}
