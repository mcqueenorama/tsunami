package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var cmdGet = &Command{
	Run:       runGet,
	UsageLine: "get",
	Short:     "get a url",
	Long:      "get a url",
}

func runGet(cmd *Command, args []string) bool {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	return true
}
