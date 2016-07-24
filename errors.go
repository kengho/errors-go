package errors

import (
	"log"
	"runtime"
)

// http://stackoverflow.com/a/24809646
func HandleErr(err error) () {
  if err != nil {
    _, fn, line, _ := runtime.Caller(1)
    log.Fatalf("%v [%s:%d]", err, fn, line)
  }
}