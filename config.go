package main

import (
	"errors"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func getUrls() (urls []string, err error) {
	//get the root node in the config file
	config, error := yaml.ReadFile(*file)
	if error != nil {
		err = errors.New(fmt.Sprintf("Error retrieving root node: %s", error))
		return
	}

	//get the count of urls
	count, error := config.Count("urls")
	if error != nil {
		err = errors.New(fmt.Sprintf("Error getting urls from config: %s", error))
		return
	}

	for x := 0; x < count; x++ {
		url, error := config.Get(fmt.Sprintf("urls[%d]", x))
		if error != nil {
			err = errors.New(fmt.Sprintf("Error retrieving url: %s", error))
			return
		}

		urls = append(urls, url)
	}

	return
}
