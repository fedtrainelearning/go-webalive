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
		fmt.Println("BOOM 1! ", err)
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
			fmt.Println("BOOM 3! ", err)
			return
		}

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
