package execx

type Logger interface {
	Print(...interface{})
}

type nopLogger struct{}

func (nopLogger) Print(...interface{}) {}
