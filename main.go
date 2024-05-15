package main

import (
	"os"
	"log"
	"fmt"
)

func main() {

	type Properties struct {
		Properties string `json:"properties"`
	}

	file, err := os.ReadFile("testcase.json")
	if err != nil {
		log.Fatal("unable to read file: %v", err)
	}
	fmt.Println(string(file))

}