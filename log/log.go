package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// TraceLogger logs the stack trace messages.
	TraceLogger *log.Logger
	// InfoLogger logs informative (user readable) messages.
	InfoLogger *log.Logger
	// WarningLogger logs warning messages.
	WarningLogger *log.Logger
	// ErrorLogger logs error messages.
	ErrorLogger *log.Logger
)

// Tracef logs formatted trace messages.
func Tracef(format string, values ...interface{}) {
	TraceLogger.Printf(format, values...)
}

// Traceln logs trace messages.
func Traceln(values ...interface{}) {
	TraceLogger.Println(values...)
}

// Infof logs formatted information messages.
func Infof(format string, values ...interface{}) {
	InfoLogger.Printf(format, values...)
}

// Infoln logs information messages.
func Infoln(values ...interface{}) {
	InfoLogger.Println(values...)
}

// Warnf logs formatted warnings.
func Warnf(format string, values ...interface{}) {
	WarningLogger.Printf(format, values...)
}

// Warnln logs warnings.
func Warnln(values ...interface{}) {
	WarningLogger.Println(values...)
}

// Error logs fatal errors.
func Error(values ...interface{}) {
	ErrorLogger.Fatal(values)
}

// Errorf logs formatted fatam messages.
func Errorf(format string, values ...interface{}) {
	ErrorLogger.Fatalf(format, values)
}

func init() {
	initLogging(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

func initLogging(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	TraceLogger = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger = log.New(infoHandle, "", 0)

	WarningLogger = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLogger = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
