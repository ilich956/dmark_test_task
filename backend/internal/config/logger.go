package config

import (
	"log/slog"
	"os"
	"path/filepath"
)

// initializes and returns a new slog.Logger
func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					source.File = filepath.Base(source.File)
				}
			}
			return a
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	return logger
}
