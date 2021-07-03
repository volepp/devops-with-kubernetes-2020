package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var websiteUrl *url.URL

func main() {
	port := 3000

	var err error
	websiteUrl, err = url.ParseRequestURI(os.Getenv("WEBSITE_URL"))
	if err != nil {
		log.Fatalf("The given URL is not valid: %s", err)
	}

	http.HandleFunc("/", handleRoot)

	log.Printf("Listening to port %v", port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Printf("Error listening to port %v: %s", port, err)
	}

}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	reqPath := req.URL.Path

	siteUrl := websiteUrl.String() + reqPath

	res, err := http.Get(siteUrl)
	if err != nil {
		log.Printf("Error getting the page: %s", err)
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading site: %s", err)
	}

	w.Write(content)
}
