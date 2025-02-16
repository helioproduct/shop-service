package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	Log = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel)
}

func Error(err error, caller string) {
	if err != nil {
		Log.
			Err(err).
			Str("caller", caller).
			Str("error", err.Error()).
			Send()
	}
}

func Info(caller string, msg string, fields map[string]interface{}) {
	event := Log.Info().
		Str("caller", caller)
	for k, v := range fields {
		switch val := v.(type) {
		case string:
			event.Str(k, val)
		case int:
			event.Int(k, val)
		case uint64:
			event.Uint64(k, val)
		case float64:
			event.Float64(k, val)
		case bool:
			event.Bool(k, val)
		default:
			event.Interface(k, val)
		}
	}
	event.Msg(msg)
}

func Warn(caller string, msg string, fields map[string]interface{}) {
	event := Log.Warn().
		Str("caller", caller)
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(msg)
}

func Debug(caller string, msg string, fields map[string]interface{}) {
	event := Log.Debug().
		Str("caller", caller)
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(msg)
}

func Print(caller string, arg interface{}) {
	fmt.Printf("[caller: %s] %v\n", caller, arg)
}
