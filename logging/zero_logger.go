package logging

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/HosseinRouhi79/log-pkg/config"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type zerologger struct {
	cfg    *config.LoggerConfig
	logger *zerolog.Logger
}

var zeroFinLogger *zerolog.Logger
var zeroOnce sync.Once

var zerologMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func (l *zerologger) getLogLevel() zerolog.Level {
	level, ok := zerologMap[l.cfg.Level]
	if !ok {
		return zerolog.DebugLevel
	}
	return level
}

func newZeroLogger(cfg *config.LoggerConfig) *zerologger {
	zero := &zerologger{cfg: cfg}
	zero.Init()
	return zero
}

func (z *zerologger) Init() {
	zeroOnce.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		fileName := fmt.Sprintf("%s%s-%s.%s", z.cfg.FilePath, time.Now().Format("2006-01-02 15:04:05"), uuid.New(), "log")

		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			panic("Could not open log file")
		}
		var logger = zerolog.New(file).
			With().
			Timestamp().
			Str("AppName", "MyApp").
			Str("LoggerName", "Zerolog").
			Logger()

		zerolog.SetGlobalLevel(z.getLogLevel())
		zeroFinLogger = &logger

	})
	z.logger = zeroFinLogger
}

func (l *zerologger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zerologger) Debugf(template string, args ...interface{}) {
	l.logger.
		Debug().
		Msgf(template, args...)
}

func (l *zerologger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Info().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zerologger) Infof(template string, args ...interface{}) {
	l.logger.
		Info().
		Msgf(template, args...)
}

func (l *zerologger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zerologger) Warnf(template string, args ...interface{}) {
	l.logger.
		Warn().
		Msgf(template, args...)
}

func (l *zerologger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Error().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zerologger) Errorf(template string, args ...interface{}) {
	l.logger.
		Error().
		Msgf(template, args...)
}

func (l *zerologger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zerologger) Fatalf(template string, args ...interface{}) {
	l.logger.
		Fatal().
		Msgf(template, args...)
}

func logParamsToZeroParams(keys map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0, len(keys))

	for k, v := range keys {
		params = append(params, string(k))
		params = append(params, v)
	}

	return params
}
