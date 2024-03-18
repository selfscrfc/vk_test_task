package logger

import (
	"fmt"
	"github.com/lmittmann/tint"
	"io"
	"log/slog"
	"os"
	"time"
)

func NewLogger(inFile bool) *slog.Logger {
	var w io.Writer
	if inFile {
		if os.Getenv("IN_DOCKER") != "TRUE" {
			filename := fmt.Sprintf("logs/%s-log", time.Now().Format(time.RFC3339))
			f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			w = io.MultiWriter(os.Stderr, f)
		} else {
			w = os.Stderr
		}
	} else {
		w = os.Stderr
	}

	level := slog.LevelDebug
	if _, ok := os.LookupEnv("DEBUG_LEVEL"); ok {
		level = slog.LevelInfo
	}

	logger := slog.New(tint.NewHandler(w, &tint.Options{Level: level}))
	slog.SetDefault(slog.New(tint.NewHandler(w, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	})))

	return logger
}

func Fatalf(logger *slog.Logger, format string, args ...any) {
	logger.Error("FATAL: " + fmt.Sprintf(format, args...))
	os.Exit(1)
}
