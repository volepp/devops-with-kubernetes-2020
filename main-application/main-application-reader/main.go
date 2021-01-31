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
		filePath := "/usr/src/app/files/hash.txt"
		text, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Fprint(w, "Error")
			log.Println("Couldn't load text: ", err)
			return
		}
		fmt.Fprint(w, string(text))
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
