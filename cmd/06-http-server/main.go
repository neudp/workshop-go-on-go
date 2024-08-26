package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			http.NotFound(writer, request)

			return
		}

		_, err := writer.Write([]byte("Hello, World!\n"))

		if err != nil {
			fmt.Println(err.Error())
		}
	})

	server.HandleFunc("/best", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello, BEST PATH!\n"))

		if err != nil {
			fmt.Println(err.Error())
		}
	})

	server.HandleFunc("/{path}", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(fmt.Sprintf("Hello, %s!\n", request.PathValue("path"))))

		if err != nil {
			fmt.Println(err.Error())
		}
	})

	err := http.ListenAndServe("localhost:8080", server)

	if err != nil {
		fmt.Println(err.Error())
	}
}
