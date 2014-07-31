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

	//iterate the urls
	for x := 0; x < count; x++ {
		url, err := config.Get(fmt.Sprintf("urls[%d]", x))
		if err != nil {
			fmt.Println("   Error getting the url out of the config list: ", err)
			return
		}

		fmt.Println("Attempting url: ", url)

		sendRequest(url)
	}
}

func sendRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(fmt.Sprintf("Url '%s' has error: %s", url, err))
		return
	}

	defer resp.Body.Close()

	fmt.Println(fmt.Sprintf("Url '%s' has status code: %s", url, resp.Status))
}
