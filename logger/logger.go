package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/logger"
)

type Logger struct {
	filePath string
	depth int

	useFile bool
	stdout bool

	logger *logger.Logger
	file *os.File
	today string

	lock sync.Mutex
}

func NewLogger(filePath string, depth int) *Logger {
	o := &Logger{}
	o.depth = depth
	o.init(filePath, false)
	return o
}

func NewLoggerWithOptions(filePath string, depth int, stdout bool) *Logger {
	o := &Logger{}
	o.depth = depth
	o.init(filePath, stdout)
	return o
}

func (o *Logger) init(filePath string, stdout bool) {
	o.filePath = filePath
	if o.filePath != "" {
		o.useFile = true
	}

	o.stdout = stdout

	o.lock = sync.Mutex{}
	o.checkDate(time.Now().Format("20060102"))
}

// FATAL
func (o *Logger) Fatal(v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.FatalDepth(o.depth, v...)
}
func (o *Logger) Fatalf(format string, v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.FatalDepth(o.depth, fmt.Sprintf(format, v...))
}
func (o *Logger) Fatalln(v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.FatalDepth(o.depth, v...)
}

// ERROR
func (o *Logger) Error(v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.ErrorDepth(o.depth, v...)
}
func (o *Logger) Errorf(format string, v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.ErrorDepth(o.depth, fmt.Sprintf(format, v...))
}
func (o *Logger) Errorln(v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.ErrorDepth(o.depth, v...)
}

// WARN
func (o *Logger) Warning(v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.WarningDepth(o.depth, v...)
}
func (o *Logger) Warningf(format string, v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.WarningDepth(o.depth, fmt.Sprintf(format, v...))
}
func (o *Logger) Warningln(v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.WarningDepth(o.depth, v...)
}

// INFO
func (o *Logger) Info(v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.InfoDepth(o.depth, v...)
}
func (o *Logger) Infof(format string, v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.InfoDepth(o.depth, fmt.Sprintf(format, v...))
}
func (o *Logger) Infoln(v ...interface{}) {
	o.checkDate(time.Now().Format("20060102"))
	o.logger.InfoDepth(o.depth, v...)
}

func (o *Logger) checkDate(today string) {
	n := time.Now().Format("20060102")
	if o.today != n {
		o.lock.Lock()
		defer o.lock.Unlock()
		if o.today != n {
			o.changeFile(n)
		}
	}
}

func (o *Logger) changeFile(today string) {
	o.today = today
	fullPath := fmt.Sprintf("%s_%s.log", o.filePath, o.today)

	splitLogPath := strings.Split(fullPath, "/")
	err := os.MkdirAll(strings.Join(splitLogPath[:len(splitLogPath)-1], "/"), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	oldFile := o.file
	defer func() {
		if oldFile != nil {
			oldFile.Close()
		}
	}()

	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	o.logger = logger.Init(o.filePath, o.stdout, true, file)
	o.file = file
}

func (o *Logger) Close() {
	o.file.Close()
}