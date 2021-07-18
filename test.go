package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var url string
	fmt.Print("Введите URL: ")
	fmt.Fscan(os.Stdin, &url)
	// url := "https://golang.org/"
	HttpGetRequest(url)
}

func HttpGetRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Count for \"%s\": %d \n", url, strings.Count(string(body), "Go"))
}
