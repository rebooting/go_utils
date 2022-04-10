package go_utils

import (
	"fmt"
	"log"
	"runtime"
)

func GetCaller() string {
	pfn, fileName, line, ok := runtime.Caller(1)
	if ok {
		details := runtime.FuncForPC(pfn)
		return fmt.Sprintf("%s:%s,%d", details.Name(), fileName, line)
	}
	return ""
}

func Log(fileDetails string, message string) {
	log.Print(fileDetails + ", " + message)
}
