package swapiContext

import (
	"goOnGo/internal/swapi-func/logging"
	"goOnGo/internal/swapi-func/transport"
	"net/http"
)

/*
SwapiContext - это не "объект" в понимании ООП, это композиция функций-зависимостей.
Он принимает в себя функции, которые производят зависимости, тем самым изолируя логику
от контекста исполнения.
*/

type SwapiContext struct {
	loggingCtx   logging.Context
	transportCtx transport.DoSwapiRequestContext
}

func New(
	loggingCtx logging.Context,
	transportCtx transport.DoSwapiRequestContext,
) *SwapiContext {
	return &SwapiContext{
		loggingCtx:   loggingCtx,
		transportCtx: transportCtx,
	}
}

func (ctx *SwapiContext) LogInfo(message string) {
	logging.Info(ctx.loggingCtx, message)
}

func (ctx *SwapiContext) LogError(message string) {
	logging.Error(ctx.loggingCtx, message)
}

func (ctx *SwapiContext) DoRequest(request *http.Request) (*http.Response, error) {
	return transport.DoSwapiRequest(ctx.transportCtx, request)
}
