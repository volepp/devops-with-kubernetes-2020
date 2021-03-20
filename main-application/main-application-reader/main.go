package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getStatus(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		message := os.Getenv("MESSAGE")

		hashFilePath := "/usr/src/app/files/hash.txt"
		hashText, err := ioutil.ReadFile(hashFilePath)
		if err != nil {
			fmt.Fprint(w, "Error")
			log.Println("Couldn't load hashText: ", err)
			return
		}
		// pingpongFile := "/usr/src/app/files/pingpong.txt"
		// pingpongText, err := ioutil.ReadFile(pingpongFile)
		// if err != nil {
		// 	fmt.Fprint(w, "Error")
		// 	log.Println("Couldn't load ping pong text:", err)
		// 	return
		// }
		nrPongs := fetchNrPongs()
		fmt.Fprint(w, fmt.Sprintf("%s\n%s\nPing / Pongs: %v",
			message, string(hashText), nrPongs))
	}
}

func fetchNrPongs() int {
	path := fmt.Sprintf("http://dwk-pingpong-svc.%s:80/pingpong/pongs", os.Getenv("NAMESPACE"))

	res, err := http.Get(path)
	if err != nil {
		log.Println("Couldn't fetch pongs:", err)
		return -1
	}

	resContent, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if err != nil {
		log.Println("Couldn't read pongs from the response:", err)
		return -1
	}
	nrPongs, err := strconv.Atoi(string(resContent))
	if err != nil {
		log.Println("Couldn't convert pongs to int:", err)
		return -1
	}

	return nrPongs
}

func main() {
	port := 3000

	http.HandleFunc("/", getStatus)

	log.Printf("Listening to port %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
