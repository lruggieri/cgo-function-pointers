package main

type Logger struct {
	traceFp, debugFp, infoFp, warningFp, errorFp func(format string)
}
func (l *Logger) Trace(message string) {
	l.traceFp(message)
}
func (l *Logger) Debug(message string) {
	l.debugFp(message)
}
func (l *Logger) Info(message string) {
	l.infoFp(message)
}
func (l *Logger) Warning(message string) {
	l.warningFp(message)
}
func (l *Logger) Error(message string) {
	l.errorFp(message)
}

func NewLogger(traceFp, debugFp, infoFp, warningFp, errorFp func(format string)) *Logger {
	return &Logger{traceFp, debugFp, infoFp, warningFp, errorFp}
}
