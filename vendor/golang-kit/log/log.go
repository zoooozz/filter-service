package log

import (
	"golang-kit/config"
	"math"
	"path"

	log "golang/log4go"
)

var (
	logger log.Logger
)

func Init(c *config.Log) {
	if c.Dir != "" {
		logger = log.Logger{}
		log.LogBufferLength = 10240
		// new info writer
		iw := log.NewFileLogWriter(path.Join(c.Dir, "info.log"), false)
		iw.SetRotateDaily(true)
		iw.SetRotateSize(math.MaxInt32)
		iw.SetFormat("[%D %T] [%L] [%S] %M")
		logger["info"] = &log.Filter{
			Level:     log.INFO,
			LogWriter: iw,
		}
		// new warning writer
		ww := log.NewFileLogWriter(path.Join(c.Dir, "warning.log"), false)
		ww.SetRotateDaily(true)
		ww.SetRotateSize(math.MaxInt32)
		ww.SetFormat("[%D %T] [%L] [%S] %M")
		logger["warning"] = &log.Filter{
			Level:     log.WARNING,
			LogWriter: ww,
		}
		// new error writer
		ew := log.NewFileLogWriter(path.Join(c.Dir, "error.log"), false)
		ew.SetRotateDaily(true)
		ew.SetRotateSize(math.MaxInt32)
		ew.SetFormat("[%D %T] [%L] [%S] %M")
		logger["error"] = &log.Filter{
			Level:     log.ERROR,
			LogWriter: ew,
		}
	}
}

func Close() {
	if logger != nil {
		logger.Close()
	}
}

func Info(format string, args ...interface{}) {
	if logger != nil {
		logger.Info(format, args...)
	}
}

func Warn(format string, args ...interface{}) {
	if logger != nil {
		logger.Warn(format, args...)
	}
}

func Error(format string, args ...interface{}) {
	if logger != nil {
		logger.Error(format, args...)
	}
}
