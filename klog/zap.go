package klog

import (
	"fmt"
	"github.com/hewenyu/toolspackage/systemctl_plus/path_make"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var level zapcore.Level

const (
	Director      = "logs"
	LogLevel      = "info"
	ShowLine      = true
	StacktraceKey = "stacktrace"
	EncodeLevel   = "LowercaseLevelEncoder"
	Format        = "console"
	LinkName      = "latest.log"
	InConsole     = true
	Prefix        = "[YUAN]"
	//LogPrefix     = "base"
)

func Zap() (logger *zap.Logger) {
	return logZap(LogLevel, Director)
}

func NewZap(dir string) (logger *zap.Logger) {
	return logZap(LogLevel, dir)
}

func logZap(loggerLevel string, dir string) (logger *zap.Logger) {
	
	if ok, _ := path_make.PathExists(dir); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", dir)
		_ = os.Mkdir(dir, os.ModePerm)
	}
	
	switch loggerLevel { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case LogLevel:
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(dir), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore(dir))
	}
	if ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig(encodeLevel string) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	
	// 日志参数
	switch encodeLevel {
	case "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder(format string) (config zapcore.Encoder) {
	
	switch format {
	case Format:
		config = zapcore.NewConsoleEncoder(getEncoderConfig(EncodeLevel))
	case "json":
		config = zapcore.NewJSONEncoder(getEncoderConfig(EncodeLevel))
	default:
		config = zapcore.NewConsoleEncoder(getEncoderConfig(EncodeLevel))
	}
	
	return
	
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(dir string) (core zapcore.Core) {
	writer, err := GetWriteSyncer(dir, LinkName, InConsole) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(Format), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(Prefix + "2006/01/02 - 15:04:05.000"))
}
