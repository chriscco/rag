package logger 

import (
	"os"
	"io"
	"github.com/sirupsen/logrus"
)

type ILog interface {
	Debug(args ... interface{})
	Info(args ... interface{})
	Warn(args ... interface{})
	Fatal(args ... interface{})
	Error(args ... interface{})
}

type Logger struct {
	logger *logrus.Logger 
}
type LogEmailHook struct {} 

func (l *LogEmailHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
	}
}

func (l *LogEmailHook) Fire(entry *logrus.Entry) error {
	return nil
}

func NewLogger(filepath string) ILog {
	parselvl, err := logrus.ParseLevel("info")
	if err != nil {
		panic(err.Error()) 
	}
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err.Error())
	}

	log := &logrus.Logger {
		Out: io.MultiWriter(file, os.Stdout), 
		Level: parselvl, 
		Hooks: make(map[logrus.Level][]logrus.Hook), 
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
			TimestampFormat: "2006-01-02 15:04:05",
			DisableColors: true,
		},
	}
	log.AddHook(&LogEmailHook{})
	log.Infof("log booted\n")
	return &Logger{logger: log}
}

func (l *Logger) Debug(args ... interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}
