package environment

import (
	"fmt"
	"os"
	"reflect"
)

var setters = map[reflect.Kind]func(string, *reflect.Value) error{
	reflect.String:  setString,
	reflect.Int:     setInt,
	reflect.Int8:    setInt,
	reflect.Int16:   setInt,
	reflect.Int32:   setInt,
	reflect.Int64:   setInt,
	reflect.Uint:    setUint,
	reflect.Uint8:   setUint,
	reflect.Uint16:  setUint,
	reflect.Uint32:  setUint,
	reflect.Uint64:  setUint,
	reflect.Float32: setFloat,
	reflect.Float64: setFloat,
	reflect.Bool:    setBool,
	reflect.Pointer: setPointer,
}

func Read(a interface{}) (err error) {
	// Получаем тип переменной
	aPointer := reflect.TypeOf(a)

	// Проверяем, что переменная является указателем
	if aPointer.Kind() != reflect.Ptr {
		return fmt.Errorf("expected a pointer to a struct, got %v", aPointer.Kind())
	}

	// Получаем тип за указателем
	aType := aPointer.Elem()

	// Проверяем, что за указателем находится структура
	if aType.Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct, got a pointer to %v", aType.Elem().Kind())
	}

	// Получаем значение переменной
	aValue := reflect.ValueOf(a).Elem()

	// Проходим по полям структуры
	for i := 0; i < aType.NumField(); i++ {
		// Получаем тип и значение поля
		field := aType.Field(i)
		value := aValue.Field(i)

		// Получаем тег поля
		// Теги являются частью метаданных типа, что может вызвать непонимание
		// Однако семантически теги являются частью декларации типа, а не его значения
		// мы читаем это как Поле А типа Б с тегом "С"
		tag := field.Tag.Get("env")

		// Если тег не пустой, то пытаемся прочитать дополнительные теги
		// переменную окружения и устанавливаем значение поля
		if tag != "" {
			defaultValue := field.Tag.Get("default")
			isRequired := field.Tag.Get("required") == "true"

			// Получаем функцию установки значения в зависимости от типа поля
			setter, ok := setters[field.Type.Kind()]
			if !ok {
				return fmt.Errorf("unsupported type: %v", field.Type.Kind())
			}

			// Получаем значение переменной окружения
			envValue := os.Getenv(tag)

			// Если переменная окружения не установлена, то устанавливаем значение по умолчанию
			if envValue == "" {
				envValue = defaultValue
			}

			// Если переменная окружения не установлена и она обязательна, то возвращаем ошибку
			if envValue == "" && isRequired {
				return fmt.Errorf("required environment variable %s is not set", tag)
			}

			// Устанавливаем значение поля если переменная окружения установлена
			if envValue != "" {
				// Важно! Передаем указатель на значение, чтобы установить его значение
				// Так же важно чтобы пале было экспортируемым, иначе будет ошибка

				// в сеттерах мы преобразуем строку в нужный тип и устанавливаем его в поле
				if err = setter(envValue, &value); err != nil {
					return fmt.Errorf("failed to set field %s: %w", field.Name, err)
				}
			}
		}
	}

	return nil
}
