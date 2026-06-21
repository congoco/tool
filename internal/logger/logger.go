package logger

import (
	"log/slog"
	"os"
)

func New(cfgLogLevel string) (*slog.Logger, error) {
	var level slog.Level

	err := level.UnmarshalText(
		[]byte(cfgLogLevel),
	)
	if err != nil {
		return nil, err
	}

	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: level,
		},
	)

	logger := slog.New(handler)

	return logger, nil
}
