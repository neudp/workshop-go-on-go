package logging

type Logger interface {
	Info(message string, labels ...*Label)
	Error(message string, labels ...*Label)
}
