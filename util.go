package proxy

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// NewLogger creates a new logger with our preferred options
func NewLogger() (*zap.Logger, error) {
	config := zap.Config{
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		Encoding:          "json",
		ErrorOutputPaths:  []string{"stdout"},
		Level:             zap.NewAtomicLevel(),
		OutputPaths:       []string{"stdout"},
	}
	l, err := config.Build()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create logger")
	}
	return l, nil
}
