package main

import(
	"fmt"
	"log"
	"time"
	"math/rand"
	"os"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	filePath := "/usr/src/app/files/hash.txt"
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Couldn't create file with path %s: %s\n", filePath, err)
	}
	defer f.Close()

	for {
		str := generateString(16)
		hash := fmt.Sprintf("%s %s", time.Now(), str)
		err := f.Truncate(0)
		if err != nil {
			log.Println("Couldn't truncate file: ", err)
		}
		f.Seek(0,0)
		f.WriteString(hash)
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
