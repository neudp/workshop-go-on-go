package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
Для работы с JSON в Go используется пакет encoding/json.
Как и в случае с переменными окружения, этот работает с помощью рефлексии.
*/

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

	// Unmarshal - функция, которая десериализует JSON в структуру.
	// Marshal - функция, которая сериализует структуру в JSON.
	//
	// Важно: поля структуры должны быть экспортируемыми
	if err := json.Unmarshal(content, dto); err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Printf("%+v\n", dto)
}
