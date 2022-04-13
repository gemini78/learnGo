package functions

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(n)
	return randomNumber
}
