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

// ------format-------

/*
Infof 加入了format
*/
func Infof(template string, args ...interface{}) {
	h.Infof(template, args)
}

/*
DEBUGf加入了format
*/
func Debugf(template string, args ...interface{}) {
	h.Debugf(template, args)
}

/*
Warnf 加入了format
*/
func Warnf(template string, args ...interface{}) {
	h.Debugf(template, args)
}

/*
DPanicf 加入了format
*/
func DPanicf(template string, args ...interface{}) {
	h.DPanicf(template, args)
}

/*
Errorf 加入了format
*/
func Errorf(template string, args ...interface{}) {
	h.Errorf(template, args)
}

/*
Errorf 加入了format
*/
func Panicf(template string, args ...interface{}) {
	h.Panicf(template, args)
}

/*
Fatalf 加入了format
*/
func Fatalf(template string, args ...interface{}) {
	h.Fatalf(template, args)
}
