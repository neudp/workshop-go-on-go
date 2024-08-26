package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	switch os.Args[1] {
	case "basic":
		basic()
	case "custom":
		custom()
	case "cookies":
		cookies()
	default:
		fmt.Println("Invalid argument")
	}
}

func basic() {
	response, err := http.Get("https://golang.org")

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(response.Status)
}

func custom() {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "https://golang.org", nil)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	request.Header.Add("User-Agent", "Custom")

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(response.Status)
}

func cookies() {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "https://google.com", nil)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	cookie := &http.Cookie{
		Name:  "cookie",
		Value: "value",
	}

	request.AddCookie(cookie)

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(response.Status)
	fmt.Printf("%+v\n", response.Cookies())
}
