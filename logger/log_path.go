package logger

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	jwa "github.com/spf13/jwalterweatherman"
)

/*
userHomeDir 从变量中获取路径
*/
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

/*
日志文件判定
输出默认路径
*/
func absPathify(inPath string) string {
	// 输出路径
	jwa.INFO.Println("文件路径查找中:", inPath)
	// 如果文件是根目录	os.PathSeparator 不同文件不同的分隔符
	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = userHomeDir() + inPath[5:]
	}
	// 判断是否以"$"开头
	if strings.HasPrefix(inPath, "$") {
		end := strings.Index(inPath, string(os.PathSeparator))

		var value, suffix string
		if end == -1 {
			value = os.Getenv(inPath[1:])
		} else {
			value = os.Getenv(inPath[1:end])
			suffix = inPath[end:]
		}

		inPath = value + suffix
	}
	// IsAbs 判断是否是绝对路径
	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}
	// 加入当前工作目录
	p, err := filepath.Abs(inPath)

	// fmt.Println("当前工作目录为:", p)
	// fmt.Println("inPath:", inPath)
	jwa.INFO.Println("当前工作目录为:", p)
	if err == nil {
		return filepath.Clean(p)
	}

	jwa.ERROR.Println("无法自动发现这个目录")
	jwa.ERROR.Println(err)
	// 返回
	return p
}

// func init() {

// 	logDir, err := CreateDir("logs")
// 	if err != nil {
// 		fmt.Printf("日志目录操作失败:%v", err)
// 		panic(err)
// 	}
// 	info_log := LogName + ".log"
// 	info_log = filepath.Join(logDir, info_log)
// 	logger = *NewLogger(info_log, zapcore.InfoLevel, 128, 30, 7, true, "Main")
// }
