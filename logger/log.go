package logger

import (
	"fmt"
	"log"
)

var (
	defaultLogger *Logger
	errorLogger *Logger
	mapLogger map[string]*Logger

	loglevel int
)

const (
	DefaultFlag = log.Ldate | log.Lmicroseconds
	DefaultFileNameFormat = "%s_logs"
	DefaultErrorFileNameFormat = "%s_errors"
	Depth = 2
)

const (
	ERROR int = 0 + iota
	WARNING
	IMPORTANT
	INFO
	DEBUG
)

type Config struct {
	FilePath	string	`yaml:"file-path" json:"file_path"`
	//Flag		int		`yaml:"flag" json:"flag"`
	Level		int		`yaml:"level" json:"level"`
	Stdout		bool	`yaml:"stdout" json:"stdout"`
}

func Open(config *Config) {
	if config.Level > DEBUG || config.Level < ERROR {
		log.Fatalf("loglevel must between %d and %d", ERROR, DEBUG)
	}
	loglevel = config.Level

	if config.FilePath != "" {
		defaultLogger = NewLoggerWithOptions(fmt.Sprintf(DefaultFileNameFormat, config.FilePath), Depth, config.Stdout)
		errorLogger = NewLoggerWithOptions(fmt.Sprintf(DefaultErrorFileNameFormat, config.FilePath), Depth, false)
	} else {
		defaultLogger = NewLoggerWithOptions("", Depth, config.Stdout)
	}
}

func Close() {
	if defaultLogger != nil {
		defer defaultLogger.Close()
	}
	if errorLogger != nil {
		defer errorLogger.Close()
	}
	defer func() {
		for _,logger := range mapLogger {
			logger.Close()
		}
	}()
}

func Fatal(v... interface{}) {
	if errorLogger != nil {
		errorLogger.Fatal(v...)
	}
	defaultLogger.Fatal(v...)
}

func Fatalf(format string, v... interface{}) {
	if errorLogger != nil {
		errorLogger.Fatalf(format, v...)
	}
	defaultLogger.Fatalf(format, v...)
}

func Fatalln(v... interface{}) {
	if errorLogger != nil {
		errorLogger.Fatalln(v...)
	}
	defaultLogger.Fatalln(v...)
}

func Error(v... interface{}) {
	if errorLogger != nil {
		errorLogger.Error(v...)
	}
	defaultLogger.Error(v...)
}

func Errorf(format string, v... interface{}) {
	if errorLogger != nil {
		errorLogger.Errorf(format, v...)
	}
	defaultLogger.Errorf(format, v...)
}

func Errorln(v... interface{}) {
	if errorLogger != nil {
		errorLogger.Errorln(v...)
	}
	defaultLogger.Errorln(v...)
}

func Warning(v... interface{}) {
	if loglevel < WARNING {
		return
	}
	defaultLogger.Warning(v...)
}

func Warningf(format string, v... interface{}) {
	if loglevel < WARNING {
		return
	}
	defaultLogger.Warningf(format, v...)
}

func Warningln(v... interface{}) {
	if loglevel < WARNING {
		return
	}
	defaultLogger.Warningln(v...)
}

func Important(v... interface{}) {
	if loglevel < IMPORTANT {
		return
	}
	defaultLogger.Info(v...)
}

func Importantf(format string, v... interface{}) {
	if loglevel < IMPORTANT {
		return
	}
	defaultLogger.Infof(format, v...)
}

func Importantln(v... interface{}) {
	if loglevel < IMPORTANT {
		return
	}
	defaultLogger.Infoln(v...)
}

func Info(v... interface{}) {
	if loglevel < INFO {
		return
	}
	defaultLogger.Info(v...)
}

func Infof(format string, v... interface{}) {
	if loglevel < INFO {
		return
	}
	defaultLogger.Infof(format, v...)
}

func Infoln(v... interface{}) {
	if loglevel < INFO {
		return
	}
	defaultLogger.Infoln(v...)
}

func Debug(v... interface{}) {
	if loglevel < DEBUG {
		return
	}
	defaultLogger.Info(v...)
}

func Debugf(format string, v... interface{}) {
	if loglevel < DEBUG {
		return
	}
	defaultLogger.Infof(format, v...)
}

func Debugln(v... interface{}) {
	if loglevel < DEBUG {
		return
	}
	defaultLogger.Infoln(v...)
}
