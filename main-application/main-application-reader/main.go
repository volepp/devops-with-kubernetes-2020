package main

import(
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
)

func getStatus(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/status" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		hashFilePath := "/usr/src/app/files/hash.txt"
		hashText, err := ioutil.ReadFile(hashFilePath)
		if err != nil {
			fmt.Fprint(w, "Error")
			log.Println("Couldn't load hashText: ", err)
			return
		}
		pingpongFile := "/usr/src/app/files/pingpong.txt"
		pingpongText, err := ioutil.ReadFile(pingpongFile)
		if err != nil {
			fmt.Fprint(w, "Error")
			log.Println("Couldn't load ping pong text:", err)
			return
		}
		fmt.Fprint(w, fmt.Sprintf("%s\nPing / Pongs: %s", string(hashText), string(pingpongText)))
	}
}

func main() {
	port := 3000

	http.HandleFunc("/status", getStatus)

	log.Printf("Listening to port %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
