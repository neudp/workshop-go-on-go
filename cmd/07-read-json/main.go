package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dto struct {
	StringValue  string `json:"string"`
	NumberValue  int    `json:"number"`
	BooleanValue bool   `json:"boolean"`
	ArrayValue   []int  `json:"array"`
	ObjectValue  struct {
		Key string `json:"key"`
	} `json:"object"`
}

func main() {
	dto := new(Dto)

	content, err := os.ReadFile("example.json")

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	if err := json.Unmarshal(content, dto); err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Printf("%+v\n", dto)
}
