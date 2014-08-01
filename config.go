package main

import (
	"flag"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

var file = flag.String("file", "config.yaml", "The path to the config file to use")

func getUrls() ([]string, error) {
	flag.Parse()

	//get the root node in the config file
	config, err := yaml.ReadFile(*file)
	if err != nil {
		return nil, err
	}

	//get the count of urls
	count, err := config.Count("urls")
	if err != nil {
		return nil, err
	}

	urls := []string{}

	for x := 0; x < count; x++ {
		url, err := config.Get(fmt.Sprintf("urls[%d]", x))
		if err != nil {
			return nil, err
		}

		urls = append(urls, url)
	}

	return urls, nil
}
