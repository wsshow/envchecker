package log

import (
	"os"

	"golang.org/x/exp/slog"
)

type LevelHandler struct {
	level   slog.Level
	handler slog.Handler
}

func NewLevelHandler(level slog.Level, h slog.Handler) *LevelHandler {
	if lh, ok := h.(*LevelHandler); ok {
		h = lh.Handler()
	}
	return &LevelHandler{level, h}
}

func (h *LevelHandler) Enabled(level slog.Level) bool {
	return h.level >= level
}

func (h *LevelHandler) Handle(r slog.Record) error {
	return h.handler.Handle(r)
}

func (h *LevelHandler) With(attrs []slog.Attr) slog.Handler {
	return NewLevelHandler(h.level, h.handler.With(attrs))
}

func (h *LevelHandler) Handler() slog.Handler {
	return h.handler
}

func New() *slog.Logger {
	th := slog.HandlerOptions{
		AddSource: true,
		Level:     nil,
		ReplaceAttr: func(a slog.Attr) slog.Attr {
			return a
		}}.NewTextHandler(os.Stdout)

	logger := slog.New(NewLevelHandler(slog.DebugLevel, th))
	return logger
}
