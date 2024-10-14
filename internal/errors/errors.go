package errors

import (
	"errors"
	"fmt"
)

/*
Пакет errors содержит функции для работы с ошибками.
Распространено мнение, будто в Go ошибки чрезмерно упрощены, но это не так.
В Go ошибки являются обычными значениями, однако есть некоторые особенности, которые
делают их удобными в использовании.

Функция errors.New создает новую ошибку с текстом.
Функция errors.Is проверяет, является ли ошибка равной другой ошибке.
Функция errors.As пытается извлечь из ошибки другую ошибку.

Пакет также содержит функцию Join, которая объединяет ошибки в одну.

вся магия происходит в функциях Is и As, которые позволяют работать с ошибками
Is проверяет, является ли ошибка равной другой ошибке. Равенство определяется по нескольким критериям:
- если обе ошибки nil, то они равны
- если обе ошибки не nil, то они равны, если это одна и та же ошибка
- если обе ошибки не nil, то они равны, если первая ошибка реализует интерфейс {Unwrap() error} и рекурсивно возвращает
  ошибку, равную второй ошибке
- если обе ошибки не nil, то они равны, если первая ошибка реализует интерфейс {Unwrap() []error} и рекурсивно
  содержит в себе вторую ошибку
- если обе ошибки не nil, то они равны, если первая ошибка реализует интерфейс {Is(error) bool} и метод возвращает true

Все условия выполняются рекурсивно, пока не будет найдено равенство или не закончатся ошибки

As пытается извлечь из ошибки другую ошибку, для этого она проверяет те же условия, что и Is и в дополнение к этому
проверяет, реализует ли ошибка интерфейс {As(interface{}) bool} и вызывает метод As с переданным аргументом
Результатом работы As является true/false, в зависимости от того, удалось ли извлечь ошибку или нет
В случае успеха, в переданный аргумент записывается извлеченная ошибка
*/

var (
	EntityNotFound = errors.New("entity %d not found")
	EntityNotSaved = errors.New("entity %d not saved")
	EntityExists   = errors.New("entity %d already exists")
)

type entityError struct {
	err error
	id  int
}

func (err *entityError) Error() string {
	return fmt.Sprintf(err.err.Error(), err.id)
}

func (err *entityError) Unwrap() error {
	return err.err
}

func NewEntityNotFound(id int) error {
	return &entityError{
		err: EntityNotFound,
		id:  id,
	}
}

func NewEntityNotSaved(id int) error {
	return &entityError{
		err: EntityNotSaved,
		id:  id,
	}
}

func NewEntityExists(id int) error {
	return &entityError{
		err: EntityExists,
		id:  id,
	}
}
