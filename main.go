package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Properties struct {
	Properties []Property `json:"properties"`
}

type Property struct {
	Infra Infrastructure `json:"data"`
	Auth Authentication `json:"data"`
	Sec Security `json:"data"`
}

type Infrastructure struct {
	ResourceID ResourceID `json:"infrastructure"`
}

type ResourceID struct {
	Id string `json:"id"`
	Arn string `json:"arn"`
	Ari string `json:"ari"`
	Grn string `json:"grn"`
}

type Authentication struct {
	Authentication []Auth `json:"authentication"`
}

type Auth struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port int `json:"port"`
}

type Security struct {
	Security []Sec `json:"security"`
}

type Sec struct {
	Iam string `json:"iam"`
}

func main() {

	file, err := os.Open("testcase.json")
	if err != nil {
		log.Fatalf("unable to read file: %s", err)
	}
	fmt.Printf("Successfully Opened %s", file.Name())

	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var properties Properties

	json.Unmarshal(byteValue, &properties)

	// Iterate through levels of JSON file to find infrastructure block, then print the value of either Ari, Arn, or Grn
	for i := 0; i < len(properties.Properties); i++ {
		fmt.Println("Resource ID: ")
		fmt.Println("Ari: ", properties.Properties[i].Infra.ResourceID.Ari)
		fmt.Println("Arn: ", properties.Properties[i].Infra.ResourceID.Arn)
		fmt.Println("Grn: ", properties.Properties[i].Infra.ResourceID.Grn)
		fmt.Println("ID: ", properties.Properties[i].Infra.ResourceID.Id)
	}
}