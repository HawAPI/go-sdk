package hawapi

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

type FormatterHandler struct {
	slog.Handler
	logger *log.Logger
	attrs  []slog.Attr
}

func NewFormattedHandler(out io.Writer, opts *slog.HandlerOptions) *FormatterHandler {
	return &FormatterHandler{
		Handler: slog.NewJSONHandler(out, opts),
		logger:  log.New(out, "", 0),
	}
}

func (h *FormatterHandler) Handle(_ context.Context, r slog.Record) error {
	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error
	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Local().Format("2006/01/02 15:04:05")
	msg := r.Message

	level := r.Level.String() + ":"
	switch r.Level {
	case slog.LevelDebug:
		level = color.BlueString(level)
	case slog.LevelInfo:
		level = color.GreenString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	h.logger.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}
