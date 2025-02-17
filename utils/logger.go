package utils

import (
	"GINOWEN/global"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once      sync.Once
	zapLogger *zap.Logger
)

// InitLogger 初始化 zap.Logger，使用 lumberjack 控制日志文件大小
func InitLogger() {
	once.Do(func() {
		// 设置编码配置
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		}

		// 定义日志级别
		logLevel := zapcore.DebugLevel

		// 控制台输出
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleOutput := zapcore.Lock(os.Stdout)

		// 设置 lumberjack 日志轮转
		fileOutput := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "logs/application.log", // 日志文件路径
			MaxSize:    50,                     // 单个日志文件最大大小（单位：MB）
			MaxBackups: 5,                      // 最多保留的旧日志文件数量
			MaxAge:     30,                     // 最多保留日志天数
			Compress:   true,                   // 是否压缩旧日志文件
		})

		// 创建多输出核心（控制台和文件）
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleOutput, logLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileOutput, logLevel),
		)

		// 创建 logger
		zapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		global.OWEN_LOG = zapLogger
	})
}

// GetLogger 返回 zap.Logger 实例
func GetLogger() *zap.Logger {
	if zapLogger == nil {
		InitLogger()
	}
	return zapLogger
}

// SafeStringField 返回字段，空字符串会显示为 "(empty:0)"
func SafeStringField(key, value string) zap.Field {
	if value == "" {
		return zap.String(key, "(empty:0)")
	}
	return zap.String(key, value)
}

// Info 打印信息日志
func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

// Warn 打印警告日志
func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

// Error 打印错误日志
func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

// Sync 刷新缓冲区，确保所有日志被写入
func Sync() {
	_ = GetLogger().Sync()
}
