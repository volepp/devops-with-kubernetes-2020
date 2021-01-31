package main

import(
	"log"
	"fmt"
	"time"
	"math/rand"
	"net/http"
	"io"
	"os"
	"os/signal"
	"syscall"
)

var str string

func getStatus(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/status" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		io.WriteString(w, fmt.Sprintf("%s %s", time.Now(), str))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	str = generateString(16)
	port := 3000

	http.HandleFunc("/status", getStatus)

	log.Printf("Listening to port %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		os.Exit(1)
	}()

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
