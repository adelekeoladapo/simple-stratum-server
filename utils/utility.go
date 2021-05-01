package utils

import "math/rand"

func GenerateRandomString(length int) string  {
	b:= make([]rune, length)
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}