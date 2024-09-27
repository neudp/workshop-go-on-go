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

type Nested struct {
	Key string `json:"key"`
}

type Dto struct {
	StringValue  string `json:"string"`
	NumberValue  int    `json:"number"`
	BooleanValue bool   `json:"boolean"`
	ArrayValue   []int  `json:"array"`
	ObjectValue  Nested `json:"object"`
}

func (dto *Dto) IsValid() bool {
	return dto.StringValue != "" && dto.NumberValue != 0
}

func Required(getter func() string) bool {
	value := getter()

	return value != ""
}

func main() {
	dto := new(Dto) // = &Dto{}

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

	if !Required(func() string {
		return dto.StringValue
	}) {
		fmt.Println("StringValue is required")
	}

	body, err := json.MarshalIndent(dto, "", "    ")

	fmt.Printf("%s\n", body)
}
