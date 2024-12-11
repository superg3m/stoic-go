package Utility

import (
	"fmt"
	"os"
)

func AssertOnError(err error) {
	if err != nil {
		LogFatal(fmt.Sprint("[Developer Error]: ", err))
		printCallStack()
		os.Exit(-1)
	}
}

func AssertOnErrorMsg(err error, msg string) {
	if err != nil {
		LogFatal(fmt.Sprint("[Developer Error]: ", msg))
		printCallStack()
		os.Exit(-1)
	}
}

func Assert(condition bool) {
	if !condition {
		LogFatal("[Assert Triggered]")
		printCallStack()
		os.Exit(-1)
	}
}

func AssertMsg(condition bool, msg string) {
	if !condition {
		LogFatal(fmt.Sprint("[Assert Triggered]: ", msg))
		printCallStack()
		os.Exit(-1)
	}
}
