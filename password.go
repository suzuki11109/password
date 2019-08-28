package password

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Generate(length int) (plain string, hashed string) {
	plain = randomStringBytes(length)
	h, _ := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	hashed = string(h)
	return
}

func IsMatch(hashed, input string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(input),
	) == nil
}
