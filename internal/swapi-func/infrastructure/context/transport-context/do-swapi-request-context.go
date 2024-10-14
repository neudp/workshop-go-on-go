package transportContext

import "goOnGo/internal/swapi-func/logging"

type DoSwapiRequestContext struct {
	swapiURL   string
	loggingCtx logging.Context
}

func NewDoSwapiRequestContext(swapiURL string, loggingCtx logging.Context) *DoSwapiRequestContext {
	return &DoSwapiRequestContext{
		swapiURL:   swapiURL,
		loggingCtx: loggingCtx,
	}
}

func (ctx *DoSwapiRequestContext) LogInfo(message string) {
	logging.Info(ctx.loggingCtx, message)
}

func (ctx *DoSwapiRequestContext) LogError(message string) {
	logging.Error(ctx.loggingCtx, message)
}

func (ctx *DoSwapiRequestContext) SwapiURL() string {
	return ctx.swapiURL
}
