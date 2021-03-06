// Package logger is the basic library of writing log.
// There are four level in logger, Error(0), Warn(1), Info(2) and Trace(3).
// If show level is set to 'n', only level smaller or equal to 'n' will
// print to standard output.
package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace is for debug only
	Trace *log.Logger

	// Info is for normal log
	Info *log.Logger

	// Warn if for protential error
	Warn *log.Logger

	// Error if for critical error
	Error *log.Logger

	logLevel = map[string]int{
		"ERROR": 0,
		"WARN":  1,
		"INFO":  2,
		"TRACE": 3,
	}
	levelCount = 4
	logPrefix  = ""
	showLevel  = 1
)

const (
	levelError = iota
	levelWarn
	levelInfo
	levelTrace
)

func init() {
	Init("", os.Stdout, os.Stdout, os.Stdout, ioutil.Discard)
}

// Init will init logger package with specific prefix and outputs.
// First parameter is prefix, and after second will be output of different level in order of:
// 	ERROR, WARN, INFO, TRACE.
// If parameter less then 5, level without output will use ioutil.Discard
func Init(
	prefix string,
	handler ...io.Writer) {
	logPrefix = prefix

	for len(handler) < levelCount {
		handler = append(handler, ioutil.Discard)
	}

	Error = log.New(handler[logLevel["ERROR"]],
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warn = log.New(handler[logLevel["WARN"]],
		"[WARN] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(handler[logLevel["INFO"]],
		"[INFO] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Trace = log.New(handler[logLevel["TRACE"]],
		"[TRACE] ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// SetPrefix will set log start with [prefix]
func SetPrefix(prefix string) {
	logPrefix = prefix
	output := createOutputIO(showLevel)
	Init(logPrefix, output...)
}

// SetLevel will set minimum level output to stdout.
// Level can be one of "ERROR", "WARN", "INFO", "TRACE".
// If input is not one of above, level will set to INFO
func SetLevel(level string) {
	var ok bool
	showLevel, ok = logLevel[level]
	if !ok {
		showLevel = 1
	}
	output := createOutputIO(showLevel)
	Init(logPrefix, output...)
}

func createOutputIO(level int) []io.Writer {
	output := make([]io.Writer, levelCount)
	for idx := range output {
		if idx <= level {
			output[idx] = os.Stdout
		} else {
			output[idx] = ioutil.Discard
		}
	}
	return output
}
