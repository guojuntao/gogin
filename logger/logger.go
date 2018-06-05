package logger

var l *Logger

func init() {
	time := true
	file := true
	debug := true
	trace := true
	colors := true
	pid := false

	l = NewStdLogger(time, file, debug, trace, colors, pid)
}

func GetLogger() *Logger {
	return l
}
