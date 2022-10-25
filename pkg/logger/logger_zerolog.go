package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
)

// ZeroLogger -.
type ZeroLogger struct {
	logger *zerolog.Logger
}

var _ Logger = (*ZeroLogger)(nil)

// New -.
func New(level string) *ZeroLogger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &ZeroLogger{
		logger: &logger,
	}
}

// Debug -.
func (l *ZeroLogger) Debug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

// Info -.
func (l *ZeroLogger) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

// Warn -.
func (l *ZeroLogger) Warn(message string, args ...interface{}) {
	l.log(message, args...)
}

// Error -.
func (l *ZeroLogger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}

	l.msg("error", message, args...)
}

// Fatal -.
func (l *ZeroLogger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)

	os.Exit(1)
}

func (l *ZeroLogger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		l.logger.Info().Msgf(message, args...)
	}
}

func (l *ZeroLogger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
