package logger

/*
INFO
*/
func Info(args ...interface{}) {
	mylog.Info(FuncCallerName(), args)
}

/*
DEBUG
*/
func Debug(args ...interface{}) {
	mylog.Debug(FuncCallerName(), args)
}

/*
Warn
*/
func Warn(args ...interface{}) {
	mylog.Warn(FuncCallerName(), args)
}

/*
DPanic
*/
func DPanic(args ...interface{}) {
	mylog.DPanic(FuncCallerName(), args)
}

/*
Error
*/
func Error(args ...interface{}) {
	mylog.Error(FuncCallerName(), args)
}

/*
Error
*/
func Panic(args ...interface{}) {
	mylog.Panic(FuncCallerName(), args)
}

/*
Fatal
*/
func Fatal(args ...interface{}) {
	mylog.Fatal(FuncCallerName(), args)
}

// ------format-------

/*
Infof 加入了format
*/
func Infof(template string, args ...interface{}) {
	mylog.Infof(template, args)
}

/*
DEBUGf加入了format
*/
func Debugf(template string, args ...interface{}) {
	mylog.Debugf(template, args)
}

/*
Warnf 加入了format
*/
func Warnf(template string, args ...interface{}) {
	mylog.Debugf(template, args)
}

/*
DPanicf 加入了format
*/
func DPanicf(template string, args ...interface{}) {
	mylog.DPanicf(template, args)
}

/*
Errorf 加入了format
*/
func Errorf(template string, args ...interface{}) {
	mylog.Errorf(template, args)
}

/*
Errorf 加入了format
*/
func Panicf(template string, args ...interface{}) {
	mylog.Panicf(template, args)
}

/*
Fatalf 加入了format
*/
func Fatalf(template string, args ...interface{}) {
	mylog.Fatalf(template, args)
}
