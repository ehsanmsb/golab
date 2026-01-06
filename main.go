package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	fmt.Println("Hello, I'm Golab")

	args := os.Args
	if len(args) <= 2 {
		log.Fatal("Your run command must have at least two argument")
	}
	if args[1] != "--url" {
		log.Fatalf("App hasn't %s arg", args[1])
	}
	if _, err := url.ParseRequestURI(args[2]); err != nil {
		log.Fatal("Your url address is not valid")
	}

	var resp *http.Response
	var err error

	resp, err = http.Get(args[2])
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The url status is %d\n", resp.StatusCode)
	fmt.Printf("The response is %v", string(body))

}
