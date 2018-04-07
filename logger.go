package prometheus

import (
	"log"
	"os"
)

const debug = "debug"

var defaultLogger Logger

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	defaultLogger = &logger{}
}

// Logger interface
type Logger interface {
	Panic(v ...interface{})
	Debugf(format string, v ...interface{})
}

// logger default
type logger struct{}

func (*logger) Panic(v ...interface{}) { log.Panic(v...) }

func (l *logger) Debugf(format string, v ...interface{}) {
	if os.Getenv("LEVEL") != debug {
		return
	}
	log.Printf(format, v...)
}
