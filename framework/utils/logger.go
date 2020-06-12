package utils

import (
	"adminframe/framework/config"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

//初始化日志系统

func NewLogger()*zap.Logger{
	writer := getWriter()
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err :=  l.UnmarshalText([]byte(strings.ToUpper("INFO")))
	if err != nil {
		panic(fmt.Sprintf("Err Logger:%s",err))
	}
	core := zapcore.NewCore(encoder,writer,l)
	return zap.New(core,zap.AddCaller())
}

func getWriter()zapcore.WriteSyncer{
	logger := lumberjack.Logger{
		Filename:   config.LogSetting.FileName,
		MaxAge:     config.LogSetting.MaxAge,
		MaxSize:    config.LogSetting.MaxSize,
		MaxBackups: config.LogSetting.MaxBackups,
		Compress:   false,
	}
	return zapcore.AddSync(&logger)
}
func getEncoder()zapcore.Encoder{
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.TimeKey = "time"
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder.EncodeDuration = zapcore.SecondsDurationEncoder
	encoder.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoder)
}