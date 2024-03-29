package logging

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/vikash/gofr/pkg/gofr/http/middleware"

	"golang.org/x/term"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

type logger struct {
	level      level
	normalOut  io.Writer
	errorOut   io.Writer
	isTerminal bool
}

type logEntry struct {
	Level   level       `json:"level"`
	Time    time.Time   `json:"time"`
	Message interface{} `json:"message"`
}

func (l *logger) logf(level level, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	out := l.normalOut
	if level >= ERROR {
		out = l.errorOut
	}

	entry := logEntry{
		Level: level,
		Time:  time.Now(),
	}

	switch {
	case len(args) == 1 && format == "":
		entry.Message = args[0]
	case len(args) != 1 && format == "":
		entry.Message = args
	case format != "":
		entry.Message = fmt.Sprintf(format+"", args...) // TODO - this is stupid. We should not need empty string.
	}

	if l.isTerminal {
		l.prettyPrint(entry, out)
	} else {
		_ = json.NewEncoder(out).Encode(entry)
	}
}

func (l *logger) Debug(args ...interface{}) {
	l.logf(DEBUG, "", args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.logf(DEBUG, format, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.logf(INFO, "", args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logf(INFO, format, args...)
}

func (l *logger) Log(args ...interface{}) {
	l.logf(INFO, "", args...)
}

func (l *logger) Logf(format string, args ...interface{}) {
	l.logf(INFO, format, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.logf(ERROR, "", args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logf(ERROR, format, args...)
}

func (l *logger) prettyPrint(e logEntry, out io.Writer) {
	// Giving special treatment to framework's request log in terminal display. This does not add any overhead
	// in running the server. Decent tradeoff for the interface to struct conversion anti-pattern.
	if rl, ok := e.Message.(middleware.RequestLog); ok {
		fmt.Fprintf(out, "\u001B[%dm%s\u001B[0m [%s] %d  %8dµs %s %s \n", e.Level.color(), e.Level.String()[0:4],
			e.Time.Format("15:04:05"), rl.Response, rl.ResponseTime, rl.Method, rl.URI)
	} else {
		fmt.Fprintf(out, "\u001B[%dm%s\u001B[0m [%s] %v\n", e.Level.color(), e.Level.String()[0:4], e.Time.Format("15:04:05"), e.Message)
	}
}

func NewLogger(level level) Logger {
	l := &logger{
		normalOut: os.Stdout,
		errorOut:  os.Stderr,
		level:     level,
	}

	l.isTerminal = checkIfTerminal(l.normalOut)

	return l
}

// TODO - Do we need this? Only used for CMD log silencing.
func NewSilentLogger() Logger {
	l := &logger{
		normalOut: ioutil.Discard,
		errorOut:  ioutil.Discard,
	}

	return l
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return term.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}
