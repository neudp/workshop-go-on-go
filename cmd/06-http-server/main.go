package main

import (
	"fmt"
	"net/http"
)

/*
net/http так же содержит простой в использовании сервер. До версии 1.22 роутинг
не поддерживал паттерны, и для роутинга использовался пакет gorilla/mux.
*/

func main() {
	server := http.NewServeMux() // с версии 1.22 можно использовать http.NewServeMux()

	// Обработка запросов по пути "/"
	server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// Крайне неочевидный момент - "/" обратывает все запросы, для которых хендлер не был найден
		if request.URL.Path != "/" {
			http.NotFound(writer, request)

			return
		}

		// Отправка ответа
		_, err := writer.Write([]byte("Hello, World!\n"))

		if err != nil {
			fmt.Println(err.Error())
		}
	})

	// Обработка запросов по пути "/best"
	server.HandleFunc("/best", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello, BEST PATH!\n"))

		if err != nil {
			fmt.Println(err.Error())
		}
	})

	// Обработка запросов по пути "/{path}"
	// /best обработается хендлером выше, так как он более конкретный
	// Порядок объявления хендлеров не важен
	server.HandleFunc("/{path}", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(fmt.Sprintf("Hello, %s!\n", request.PathValue("path"))))

		if err != nil {
			fmt.Println(err.Error())
		}
	})

	// кроме HandleFunc, есть еще Handle, который позволяет использовать реализацию
	//интерфейса http.Handler, например, http.FileServer для обработки
	// статических файлов или http.StripPrefix для обработки путей с префиксом
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Запуск сервера на хосте localhost и порту 8080
	// Golang использует строку вида "host:port" для указания адреса сервера
	// чтобы сервер был доступен извне, можно указать "0.0.0.0:8080"
	// или просто ":8080"
	err := http.ListenAndServe("localhost:8080", server)

	if err != nil {
		fmt.Println(err.Error())
	}
}
