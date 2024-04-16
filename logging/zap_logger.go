package logging

import (
	"fmt"
	"sync"
	"time"

	"github.com/HosseinRouhi79/log-pkg/config"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	cfg    *config.LoggerConfig
	logger *zap.SugaredLogger
}

var zapOnce sync.Once

var zapLogLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"fatal": zap.FatalLevel,
}

var zapFinLogger *zap.SugaredLogger

func (l *zapLogger) getLevel(zapMap map[string]zapcore.Level) zapcore.Level {
	level, ok := zapMap[l.cfg.Level]
	if !ok {
		return zap.DebugLevel
	}
	return level
}

func (z *zapLogger) Init() {

	zapOnce.Do(func() {
		fileName := fmt.Sprintf("%s%s-%s.%s", z.cfg.FilePath, time.Now().Format("2006-Jan-02"), uuid.New(), "log")
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    1,
			MaxAge:     5,
			MaxBackups: 10,
			Compress:   true,
			LocalTime:  true,
		})
		zEncoder := zap.NewProductionEncoderConfig()
		zEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(zapcore.NewJSONEncoder(zEncoder), w, z.getLevel(zapLogLevelMap))

		logger := zap.New(core, zap.AddStacktrace(zap.ErrorLevel)).Sugar()
		fmt.Println(logger)

		zapFinLogger = logger.With("AppName", "myApp", "LoggerName", "Zap")
	})

	z.logger = zapFinLogger

}

func newZapLogger(cfg *config.LoggerConfig) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)

	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

func prepareLogInfo(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{})
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub

	return logParamsToZapParams(extra)
}

func logParamsToZapParams(keys map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0, len(keys))

	for k, v := range keys {
		params = append(params, string(k))
		params = append(params, v)
	}

	return params
}
