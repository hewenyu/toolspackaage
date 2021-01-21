package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Hliog struct {
	LogPath string
	LogName string
	Level   zapcore.Level
	base    *zap.SugaredLogger
}

var (
	h *Hliog
)

/*
SetName 设置默认名称
*/
func SetName(logname string) { h.SetName(logname) }

/*
SetName 设置默认名称
*/
func (h *Hliog) SetName(logname string) {

	absPathify(".")

}

func SetPath(in string) { h.SetPath(in) }

/*
SetPath 设置路径
*/
func (h *Hliog) SetPath(in string) {
	if in != "" {
		h.LogPath = absPathify(in)
	} else {
		h.LogPath = absPathify("./")
	}
}

/**Logs
 * zapcore构造
 */
func newLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {

	var (
		encoder    zapcore.EncoderConfig
		loggerPath lumberjack.Logger
	)
	// 日志文件路径
	loggerPath = lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoder = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder), // 编码器配置
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&loggerPath)), //默认输出到文件
		atomicLevel, // 日志级别
	)
}

/*
New 创建日志
*/
func New() {
	h.New()
}

/*
New 创建日志
*/
func (h *Hliog) New() {
	// 日志初始化
	core := newLogger(h.LogPath, h.Level, 128, 30, 7, true)
	Logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development())
	h.base = Logger.Sugar()
	return
}
