package main

import (
	"flag"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"net/http"
)

var file = flag.String("file", "config.yaml", "The path to the config file to use")

func main() {
	flag.Parse()

	//get the root node in the config file
	config, err := yaml.ReadFile(*file)
	if err != nil {
		fmt.Println("Error retrieving root node ", err)
		return
	}

	//get the count of urls
	count, err := config.Count("urls")
	if err != nil {
		fmt.Println("Error getting urls from config: ", err)
		return
	}

	fmt.Println("Number of urls found in config: ", count)

	//create a channel
	c := make(chan string)

	//iterate the urls
	for x := 0; x < count; x++ {
		url, err := config.Get(fmt.Sprintf("urls[%d]", x))
		if err != nil {
			fmt.Println("Error getting the url out of the config list: ", err)
			return
		}

		fmt.Println("Spawning routine url: ", url)

		//send requests concurrently
		go sendRequest(url, c)
	}

	//loop until a response is received from all urls
	for x := 0; x < count; x++ {
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
