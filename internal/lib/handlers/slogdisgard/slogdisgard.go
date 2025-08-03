package slogdisgard

import (
	"context"
	"log/slog"
)

func NewDisgardLogger() *slog.Logger {
	return slog.New(NewDisgardHandler())
}

type DiscardHandler struct{}

func NewDisgardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	return h
}

func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
