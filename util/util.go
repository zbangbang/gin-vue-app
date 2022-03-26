package util

import "math/rand"

func RandomString(n int) string {
	letters := []byte("qwertyuiopasdfghjklzxcvbnm")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
