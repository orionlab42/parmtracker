package log

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Used colors for the unix console
const (
	// output colors
	ColorDefault = "0"
	ColorRed     = "31"
	ColorGreen   = "32"
	ColorYellow  = "33"
	ColorCyan    = "36"
	// log levels
	LevelTrace = 0
	LevelDebug = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelError = 4
)

// Log wrapper, inherit from Logger
type Logger struct {
	log   *log.Logger
	level int
}

var instance *Logger
var once sync.Once

// Creates a new Logger
func NewLogger() *Logger {
	return &Logger{
		nil,
		LevelDebug,
	}
}

// Transforming Logger into a Singleton
func GetInstance() *Logger {
	once.Do(func() {
		instance = NewLogger()
	})
	return instance
}

// Set the log file to also log into a file
func (l *Logger) SetLogFile(filename string) {
	// open log file, and do not close it!
	file, e := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatal(e)
	}
	l.log = log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds)
}

// Get level
func (l *Logger) GetLevel() int {
	return l.level
}

// Set level
func (l *Logger) SetLevel(level int) {
	l.level = level
}

// Print to log
func (l *Logger) printLog(msg string) {
	if l.log != nil {
		l.log.Println(msg)
	}
}

// Print to stdout, only if in verbose mode
func (l *Logger) printStdout(msg string, color string) {
	// Add color to a string for printing it out in the console
	var colorized = "\x1b[" + color + "m" + time.Now().UTC().Format("2006-01-02 15:04:05") + " " + msg + "\x1b[0m"
	fmt.Println(colorized)
}

// Log Trade
func (l *Logger) Trace(prefix string, msg string) {
	if LevelTrace >= l.level {
		logMsg := "TRACE [" + prefix + "] " + msg
		l.printLog(logMsg)
		l.printStdout(logMsg, ColorCyan)
	}
}

// Log Debug
func (l *Logger) Debug(prefix string, msg string) {
	if LevelDebug >= l.level {
		logMsg := "DEBUG [" + prefix + "] " + msg
		l.printLog(logMsg)
		l.printStdout(logMsg, ColorGreen)
	}
}

// Log Info
func (l *Logger) Info(prefix string, msg string) {
	if LevelInfo >= l.level {
		logMsg := "INFO  [" + prefix + "] " + msg
		l.printLog(logMsg)
		l.printStdout(logMsg, ColorDefault)
	}
}

// Log Warning
func (l *Logger) Warn(prefix string, msg string) {
	if LevelWarn >= l.level {
		logMsg := "WARN  [" + prefix + "] " + msg
		l.printLog(logMsg)
		l.printStdout(logMsg, ColorYellow)
	}
}

// Log Error
func (l *Logger) Error(prefix string, msg string) {
	if LevelError >= l.level {
		logMsg := "ERROR [" + prefix + "] " + msg
		l.printLog(logMsg)
		l.printStdout(logMsg, ColorRed)
	}
}

// Log Trace formatted
func (l *Logger) Tracef(prefix string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	l.Trace(prefix, msg)
}

// Log Debug formatted
func (l *Logger) Debugf(prefix string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	l.Debug(prefix, msg)
}

// Log Info formatted
func (l *Logger) Infof(prefix string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	l.Info(prefix, msg)
}

// Log Warning formatted
func (l *Logger) Warnf(prefix string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	l.Warn(prefix, msg)
}

// Log Error formatted
func (l *Logger) Errorf(prefix string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	l.Error(prefix, msg)
}
