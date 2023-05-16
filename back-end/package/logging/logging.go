package logging

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)

	logFile := &lumberjack.Logger{
		Filename:  "log/log.log",
		MaxSize:    2,
		MaxBackups: 2,
		MaxAge:     10, //days
	}
	l.AddHook(&WriteHook{
		Writer:    []io.Writer{logFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	warnFile := &lumberjack.Logger{
		Filename:   "log/warn.log",
		MaxSize:    2,
		MaxBackups: 2,
		MaxAge:     10,
	}
	l.AddHook(&WriteHook{
		Writer: []io.Writer{warnFile},
		LogLevels: []logrus.Level{
			logrus.WarnLevel, logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel,
		},
	})

	errFile := &lumberjack.Logger{
		Filename:  "log/error.log",
		MaxSize:    2,
		MaxBackups: 2,
		MaxAge:     10,
	}
	l.AddHook(&WriteHook{
		Writer:    []io.Writer{errFile},
		LogLevels: []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel},
	})

	l.SetLevel(logrus.TraceLevel)
	l.SetOutput(io.Discard)

	e = logrus.NewEntry(l)
}

type WriteHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (h *WriteHook) Fire(entry *logrus.Entry) error {
	line := fmt.Sprintf("level: %s, time: %s, file: %s, line %d, msg: %s\n",
		entry.Level,
		entry.Time.Format("15:04:05 02/01/2006"),
		path.Base(entry.Caller.File),
		entry.Caller.Line,
		entry.Message,
	)

	for _, w := range h.Writer {
		w.Write([]byte(line))
	}
	return nil
}

func (h *WriteHook) Levels() []logrus.Level {
	return h.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

func GetLogger() Logger {
	Init()
	return Logger{e}
}
