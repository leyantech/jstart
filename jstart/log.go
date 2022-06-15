package jstart

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// NOTE: "jstart" in prefix improves logging recognizability
	DEBUG = log.New(getDebugLoggerWriter(), "DEBUG jstart ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	INFO  = log.New(os.Stderr, "INFO jstart ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	WARN  = log.New(os.Stderr, "WARN jstart ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	ERROR = log.New(os.Stderr, "ERROR jstart ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
)

func getDebugLoggerWriter() io.Writer {
	if os.Getenv("JSTART_DEBUG") == "" {
		return ioutil.Discard
	} else {
		return os.Stderr
	}
}
