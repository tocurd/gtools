package gtools

import (
	"bytes"
	"fmt"
	"path"
	"time"

	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Log struct {
	LogDir     string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	TimeLayout string
	Caller     int
	LogLevel   logrus.Level
	lumberjack *lumberjack.Logger

	PanicColor string
	FatalColor string
	ErrorColor string
	WarnColor  string
	InfoColor  string
	DebugColor string
	TraceColor string
}

type formatter struct {
	Caller int
}

func (formatter formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	caller := entry.Caller
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	newLog := fmt.Sprintf("[%s] %s:%d [%s] %s\n", timestamp, caller.File, caller.Line, entry.Level, entry.Message)

	levelColor := map[logrus.Level]string{
		0: "eeeeee", // Panic
		1: "eeeeee", // Fatal
		2: "f92672", // Error
		3: "yellow", // Warn
		4: "eeeeee", // Info
		5: "2db7f5", // Debug
		6: "eeeeee", // Trace
	}
	color.Printf("<fg=%s>[%s] %s:%d [%s]</> %s\n", levelColor[entry.Level], timestamp, caller.File, caller.Line, entry.Level, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

type hook struct {
	logger *Log
}

func (h *hook) Fire(entry *logrus.Entry) error {
	filename := path.Join(h.logger.LogDir, h.logger.getFileName(h.logger.Filename))
	if filename != h.logger.lumberjack.Filename {
		if err := h.logger.lumberjack.Rotate(); err != nil {
			return err
		}
	}
	return nil
}

func (h *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

/**
 * @description: 获取日志文件的名字
 * @param {string} baseName
 * @return {*}
 */
func (l *Log) getFileName(baseName string) string {
	timeLayout := "2006-01-02"
	if l.TimeLayout != "" {
		timeLayout = l.TimeLayout
	}
	var dateTime = time.Unix(time.Now().Unix(), 0).Format(timeLayout)
	return fmt.Sprintf("%s-%s.log", baseName, dateTime)
}

/**
 * @description: 创建日志
 * @param {*}
 * @return {*}
 */
func (l *Log) New() *logrus.Logger {
	var log = logrus.New()

	l.lumberjack = &lumberjack.Logger{}
	if l.LogDir == "" {
		l.LogDir = "./logs/"
	}
	if l.Filename == "" {
		l.Filename = "log"
	}
	if l.MaxSize == 0 {
		l.MaxSize = 100
	}
	if l.MaxBackups == 0 {
		l.MaxBackups = 100
	}
	if l.MaxAge == 0 {
		l.MaxAge = 31
	}

	l.lumberjack.Filename = path.Join(l.LogDir, l.getFileName(l.Filename))
	l.lumberjack.MaxSize = l.MaxSize
	l.lumberjack.MaxBackups = l.MaxBackups
	l.lumberjack.MaxAge = l.MaxAge
	l.lumberjack.Compress = true
	l.lumberjack.LocalTime = true

	log.SetOutput(l.lumberjack)
	log.SetFormatter(&formatter{Caller: l.Caller})
	log.SetReportCaller(true)

	log.AddHook(&hook{logger: l})
	if l.LogLevel == 0 {
		l.LogLevel = logrus.DebugLevel
	}
	log.SetLevel(l.LogLevel)
	return log
}
