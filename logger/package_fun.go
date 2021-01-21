package logger

/*
INFO
*/
func Info(args ...interface{}) {
	h.Info(args)
}

/*
DEBUG
*/
func Debug(args ...interface{}) {
	h.Debug(args)
}

/*
Warn
*/
func Warn(args ...interface{}) {
	h.Debug(args)
}

/*
DPanic
*/
func DPanic(args ...interface{}) {
	h.DPanic(args)
}

/*
Error
*/
func Error(args ...interface{}) {
	h.Error(args)
}

/*
Error
*/
func Panic(args ...interface{}) {
	h.Panic(args)
}

/*
Fatal
*/
func Fatal(args ...interface{}) {
	h.Fatal(args)
}
