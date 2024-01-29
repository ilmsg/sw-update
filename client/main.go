package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	requestUrl := "https://github.com/ilmsg/sw-update/raw/main/version.txt"
	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", respBody)
	// fmt.Printf("client status code: %d\n", resp.StatusCode)
}
