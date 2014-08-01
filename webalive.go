package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls, err := getUrls()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Number of urls found in config: ", len(urls))

	//create a channel
	c := make(chan string)

	//iterate the urls
	for _, url := range urls {

		fmt.Println("Spawning routine url: ", url)

		//send requests concurrently
		go sendRequest(url, c)
	}

	//loop until a response is received from all urls
	for x := 0; x < len(urls); x++ {
		fmt.Println(<-c)
	}
}

func sendRequest(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("Url '%s' has error: %s", url, err)
		return
	}

	defer resp.Body.Close()

	c <- fmt.Sprintf("Response received for '%s' has status code: %s", url, resp.Status)
}
