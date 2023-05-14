package main

import (
	"fmt"

	Hermes "github.com/realTristan/Hermes/nocache"
)

func main() {
	var cache, err = Hermes.InitWithJson("data.json")
	if err != nil {
		panic(err)
	}
	var result, _ = cache.Search("realTristan", 100, false, map[string]bool{
		"url":         true,
		"title":       true,
		"description": true,
	})
	for _, item := range result {
		fmt.Println(item["url"])
	}
}
