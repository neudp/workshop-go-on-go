package appContext

import (
	"fmt"
	logingContext "goOnGo/internal/swapi-func/infrastructure/app-context/loging-context"
	swapiContext "goOnGo/internal/swapi-func/infrastructure/app-context/swapi-context"
	"goOnGo/internal/swapi-func/infrastructure/environment"
	"goOnGo/internal/swapi-func/model/logging"
	"strings"
)

type AppContext struct {
	// Неизменяемая часть, мы можем безопасно сохранять окружение, поскольку оно не должно изменяться
	// в процессе исполнения программы
	env *environment.Environment

	// Изменяемая часть, мы можем изменять уровень логирования в процессе исполнения программы
	minLogLevel logging.Level

	// Медленная загрузка (lazy loading) - контексты кэшируются и сохраняются до тех пор,
	// пока контекст не будет изменен
	loggingCtx *logingContext.LoggingContext
	swapiCtx   *swapiContext.SwapiContext
}

func New(overrides ...environment.Override) (*AppContext, error) {
	env, err := environment.Read(overrides...)

	if err != nil {
		return nil, err
	}

	minLogLevel, err := parseLogLevel(env.MinLogLevel())

	return &AppContext{
		env:         env,
		minLogLevel: minLogLevel,
	}, nil
}

func (ctx *AppContext) LoggingContext() *logingContext.LoggingContext {
	if ctx.loggingCtx == nil {
		ctx.loggingCtx = logingContext.New(ctx.minLogLevel)
	}

	return ctx.loggingCtx
}

func (ctx *AppContext) SwapiContext() *swapiContext.SwapiContext {
	if ctx.swapiCtx == nil {
		ctx.swapiCtx = swapiContext.New(ctx.env.SwapiURL(), ctx.LoggingContext())
	}

	return ctx.swapiCtx
}

// При копировании контекста мы сохраняем неизменяемую часть, а изменяемую часть копируем

func (ctx *AppContext) Clone() *AppContext {
	return &AppContext{
		env:         ctx.env,
		minLogLevel: ctx.minLogLevel, // если бы это была бы структура, то мы бы вызвали метод Clone()
	}

	// Не копируя ленивые контексты (lazy contexts), мы создаем новый чистый контекст.
	// Кто бы не использовал старый контекст, он будет использовать старые инициализации.
	// Все новые процессы будут использовать новые инициализации.
	// Таким образом, мы достигаем чистоты наших функций в рамках 1 контекста.
	// Сам факт копирования контекста не нарушает принцип чистоты функций,
	// поскольку в следующем исполнении программы все процессы будут исполняться с другими входными данными,
	// но все запущенные процессы будут исполняться с прежними входными данными
}

/*
Функциональный подход отдает предпочтение созданию новых объектов вместо изменения существующих.
Это реализует одно из важных правил функционального программирования - чистоту функций.
Чистая функция - это функция, которая не изменяет состояние программы и не зависит от состояния
программы. Таким образом, функция всегда возвращает одинаковый результат при одинаковых входных данных.
Как следствие, у нас нет методов для изменения полей структуры. У нас есть функции, которые создают новый объект
*/

func ChangeMinLogLevel(ctx *AppContext, minLogLevel logging.Level) *AppContext {
	newCtx := ctx.Clone()
	newCtx.minLogLevel = minLogLevel

	return newCtx
}

/*
Есть лишь одно исключение, когда мы допускаем изменение состояния объекта - исключительно синхронное исполнение
Если структура используется в однопоточном режиме, то можно изменять состояние объекта, но функция обязана возвращать
значение, а вызывающая сторона должна использовать это значение. Это сохраняет принцип чистоты функций и позволяет
изменить поведение в будущем если исполнение программы станет многопоточным

Это следует рассматривать как метод оптимизации. В функциональной парадигме все структуры должны быть неизменяемыми
*/

func ChangeMinLogLevelSync(ctx *AppContext, minLogLevel logging.Level) *AppContext {
	ctx.minLogLevel = minLogLevel

	return ctx
}

func parseLogLevel(level string) (logging.Level, error) {
	switch strings.ToUpper(level) {
	case "INFO", "INF", "I", "0":
		return logging.Info, nil
	case "ERROR", "ERR", "E", "1":
		return logging.Error, nil
	default:
		return -1, fmt.Errorf("unknown log level: %s", level)
	}
}
