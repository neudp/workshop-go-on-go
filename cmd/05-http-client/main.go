package main

import (
	"fmt"
	"net/http"
	"os"
)

/*
Go имеет встроенный пакет net/http, который реализует все необходимое для работы с HTTP.
*/

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

/*
Для выполнения простых HTTP запросов используется функция http.Get().

Это прокси для http.DefaultClient.Get().
*/

func basic() {
	response, err := http.Get("https://golang.org")

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(response.Status)
}

/*
Для выполнения более сложных запросов используется структуры http.Client
и http.Request. Это позволяет управлять различными аспектами запроса, такими как
заголовки, методы, таймауты и т.д.
*/

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

/*
Для работы с куками используется структура http.Cookie.

Куки могут быть установлены в запросе с помощью метода AddCookie() структуры http.Request.

Получить куки из ответа можно с помощью метода Cookies() структуры http.Response.
*/

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
