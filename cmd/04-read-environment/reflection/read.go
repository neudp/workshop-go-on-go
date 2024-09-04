package reflection

import "goOnGo/internal/environment"

/*
Рефлексия - это способность программы анализировать и модифицировать свое поведение во время выполнения.
В Go рефлексия реализована в пакете reflect.

Сама по себе рефлексия не является плохой практикой, но ее использование может привести к ухудшению
производительности и увеличению сложности кода. Поэтому перед использованием рефлексии стоит
проанализировать, действительно ли она необходима. Одним из важных критериев в решении использовать
рефлексию или нет, является сложность внутренней логики и частота использования. Рефлексия должна
использоваться только в случае, если другие способы решения задачи не подходят или слишком сложны.

В данном примере рефлексия используется для десериализации объекта из переменных окружения. Без рефлексии
решение этой задачи было бы гораздо сложнее и требовало бы большего количества кода. Поскольку
переменные окружения читаются один раз при запуске приложения, а их изменение гарантированно развитием приложения,
то использование рефлексии в данном случае оправдано.

enviroment.Read() как и библиотека от caarlos0, использует рефлексию для чтения тегов структуры
и заполнения полей структуры значениями переменных окружения.
*/

type Environment struct {
	Host string `env:"HOST" default:"localhost"`
	Port int    `env:"PORT" default:"8080"`
}

func ReadEnv() (*Environment, error) {
	env := new(Environment)

	if err := environment.Read(env); err != nil {
		return nil, err
	}

	return env, nil
}
