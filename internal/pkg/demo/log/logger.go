package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var debugLogger *zap.Logger
var sugaredDebugLogger *zap.SugaredLogger

var infoLogger *zap.Logger
var sugaredInfoLogger *zap.SugaredLogger

var warnLogger *zap.Logger
var sugaredWarnLogger *zap.SugaredLogger

var errorLogger *zap.Logger
var sugaredErrorLogger *zap.SugaredLogger

// https://github.com/uber-go/zap
// https://github.com/uber-go/zap/blob/master/FAQ.md
func init() {
	// if debug is enable
	debugLogger, _ = zap.NewDevelopment()
	defer func(debugLogger *zap.Logger) {
		err := debugLogger.Sync()
		if err != nil {

		}
	}(debugLogger)
	sugaredDebugLogger = debugLogger.Sugar()

	infoWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.info.log",
		MaxSize:    200, // megabytes
		MaxBackups: 10,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
	infoZapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		infoWriteSyncer,
		zap.InfoLevel,
	)
	infoLogger = zap.New(infoZapCore)
	sugaredInfoLogger = infoLogger.Sugar()

	warnWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.warn.log",
		MaxSize:    200, // megabytes
		MaxBackups: 10,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
	warnZapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		warnWriteSyncer,
		zap.WarnLevel,
	)
	warnLogger = zap.New(warnZapCore)
	sugaredWarnLogger = warnLogger.Sugar()

	errorWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.error.log",
		MaxSize:    200, // megabytes
		MaxBackups: 10,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
	errorZapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		errorWriteSyncer,
		zap.ErrorLevel,
	)
	errorLogger = zap.New(errorZapCore)
	sugaredErrorLogger = errorLogger.Sugar()
}

func Debug(msg string, fields ...zap.Field) {
	debugLogger.Debug(msg, fields...)
}
func Debugf(template string, args ...interface{}) {
	sugaredDebugLogger.Debugf(template, args)
}

func Info(msg string, fields ...zap.Field) {
	infoLogger.Info(msg, fields...)
}
func Infof(template string, args ...interface{}) {
	sugaredInfoLogger.Infof(template, args)
}

func Warn(msg string, fields ...zap.Field) {
	warnLogger.Warn(msg, fields...)
}
func Warnf(template string, args ...interface{}) {
	sugaredWarnLogger.Infof(template, args)
}

func Error(msg string, fields ...zap.Field) {
	errorLogger.Error(msg, fields...)
}
func Errorf(template string, args ...interface{}) {
	sugaredErrorLogger.Errorf(template, args)
}
