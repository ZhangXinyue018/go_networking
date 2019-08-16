package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "http://www.google.com"
	response, err := http.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fmt.Printf("Status: %v\n", response.Status)
	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}
	os.Exit(0)
}
