package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
net/http предоставляет простые интерфейсы для работы с запросами и ответами.
*/

type Dto struct {
	PathValue   string `json:"pathValue"`
	HeaderValue string `json:"headerValue"`
	URLValue    string `json:"urlValue"`
	BodyValue   string `json:"bodyValue"`
	BodyError   string `json:"bodyError"`
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/{pathValue}", func(writer http.ResponseWriter, request *http.Request) {
		// request.PathValue() возвращает значение переменной из пути
		pathValue := request.PathValue("pathValue")

		// request.Header содержит заголовки запроса в виде map[string][]string
		// с методами Get и Values для доступа к значениям
		header := request.Header.Get("X-Header")

		// request.URL содержит информацию о URL запроса
		// request.URL.Path содержит путь запроса
		// request.URL.Query() содержит параметры запроса в виде map[string][]string
		// с методами Get для доступа к значениям
		// и тд
		url := request.URL

		// request.Body содержит тело запроса
		// тело запроса может быть прочитано с помощью  io.ReadAl()
		// или других методов чтения байтовых стримов
		bodyReader := request.Body
		bodyBytes, err := io.ReadAll(bodyReader)
		_ = request.Body.Close()

		// writer - интерфейс для записи ответа
		// writer.Header().Set(name, value string) - устанавливает заголовок ответа
		// writer.WriteHeader(int) - устанавливает код ответа
		// writer.Write([]byte) - записывает байты в ответ
		//
		// Важно: если код ответа не был установлен, то он будет установлен в 200
		// Порядок вызова методов важен
		// Нельзя писать в writer.Header().Set() после writer.WriteHeader()
		// Нельзя писать в writer.WriteHeader() после writer.Write()
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		response := Dto{
			PathValue:   pathValue,
			HeaderValue: header,
			URLValue:    url.String(),
			BodyValue:   string(bodyBytes),
			BodyError:   "",
		}

		if err != nil {
			response.BodyError = err.Error()
		}

		// json.NewEncoder() - создает новый кодировщик JSON
		// encoder.Encode() - кодирует структуру в JSON и записывает в writer
		// это аналогично
		// body, _ := json.Marshal(response)
		// _, _ = writer.Write(body)
		_ = json.NewEncoder(writer).Encode(response)

		// writer.Write() - возвращает 2 значения: количество переданных байт и ошибка
		// это позволяет проверить, были ли переданы все данные
	})

	err := http.ListenAndServe("localhost:8080", server)

	if err != nil {
		fmt.Println(err.Error())
	}
}
