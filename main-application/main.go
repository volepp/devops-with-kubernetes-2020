package main

import(
	"log"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	str := generateString(16)

	for {
		// Log automatically prints out the timestamp
		log.Println(str)
		time.Sleep(5*time.Second)
	}
}

func generateString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, length)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}

	return string(str)
}
