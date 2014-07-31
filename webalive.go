package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://smsprodweb1/eapplication",
	"http://smsprodweb1/estudent",
	"http://smsprodweb2/eapplication",
	"http://smsprodweb2/estudent",
}

func main() {
	for _, url := range urls {
		fmt.Println("Attempting url: ", url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("         Error: ", err)
			continue
		}

		fmt.Println("   Status code: ", resp.StatusCode)
		resp.Body.Close()
	}
}
