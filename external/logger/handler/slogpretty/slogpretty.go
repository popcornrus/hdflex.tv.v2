package slogpretty

import (
	"context"
	"fmt"
	"io"
	stdLog "log"
	"strings"

	"github.com/fatih/color"
	"log/slog"
)

type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type PrettyHandler struct {
	opts PrettyHandlerOptions
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

func (opts PrettyHandlerOptions) NewPrettyHandler(
	out io.Writer,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	// Set color based on log level
	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	// Extract attributes from the log record
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	// Add any additional attributes from the handler
	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	// Build the message components
	var b strings.Builder
	if len(fields) > 0 {
		for _, v := range fields {
			b.WriteString(fmt.Sprintf("%v", v))
		}
	}

	// Format the time and message
	timeStr := r.Time.Format("[15:04:05.000Z07:00]")
	msg := color.CyanString(r.Message)

	chain := fmt.Sprintf("\n%s", formatChain(b.String()))

	// Print the log message with the required format
	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(chain),
	)

	return nil
}

func formatChain(chain string) string {
	chains := strings.Split(chain, "->")

	for i, c := range chains {
		c = strings.TrimSpace(c)
		if i == 0 {
			chains[i] = fmt.Sprintf("\t├─> %s", c)
		} else if i == len(chains)-2 {
			chains[i] = fmt.Sprintf("\t└%s> %s", strings.Repeat("─", 2*i), c)
		} else if i < len(chains)-1 {
			chains[i] = fmt.Sprintf("\t├%s> %s", strings.Repeat("─", 2*i), c)
		}
	}

	chain = strings.Join(chains, "\n")
	return chain
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// TODO: implement
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
