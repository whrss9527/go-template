package util

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("1234567890abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func StringsContains(array []string, val string) bool {

	for i := 0; i < len(array); i++ {
		if array[i] == val {

			return true
		}
	}
	return false
}
