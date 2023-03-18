package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().Unix())
}

// RandomString generates a random string
func RandomString(n int) string {
	var stringBuilder strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		stringBuilder.WriteByte(c)
	}

	return stringBuilder.String()
}

// RandomPrefix will generate a random prefix with the length of 4
func RandomPrefix() string {
	return RandomString(4)
}
