package logger

import (
	"runtime"
)

/*
funcName 输出运行的方法名称
*/
func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

/*
funcCallerName 输出调用运行的方法名称
*/
func FuncCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
