package logger

import (
	"os"
	"user-api/pkg/setting"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLoggerZap(config setting.LoggerSetting) *LoggerZap {
	log_level := config.LogLevel
	var level zapcore.Level
	switch log_level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   //days
		Compress:   config.Compress, // disabled by default

	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)
	return &LoggerZap{
		zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	}

}

func getEncoderLog() zapcore.Encoder {
	encode_config := zap.NewProductionEncoderConfig()
	encode_config.EncodeTime = zapcore.ISO8601TimeEncoder
	encode_config.TimeKey = "time"
	encode_config.LevelKey = "level"

	return zapcore.NewJSONEncoder(encode_config)
}
