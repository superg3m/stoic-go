package Utility

import (
	"fmt"
	"github.com/fatih/color"
	"path/filepath"
	"runtime"
)

var (
	COLOR_RED     = color.New(color.FgRed).SprintFunc()
	COLOR_BLUE    = color.New(color.FgBlue).SprintFunc()
	COLOR_GREEN   = color.New(color.FgGreen).SprintFunc()
	COLOR_YELLOW  = color.New(color.FgYellow).SprintFunc()
	COLOR_MAGENTA = color.New(color.FgMagenta).SprintFunc()
	COLOR_CYAN    = color.New(color.FgCyan).SprintFunc()

	COLOR_BG_RED = color.New(color.FgWhite, color.BgRed).SprintFunc()
)

func printCallStack() {
	pc := make([]uintptr, 32)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])

	var allFrames []runtime.Frame
	for {
		frame, more := frames.Next()
		allFrames = append(allFrames, frame)
		if !more {
			break
		}
	}

	LogPrint(COLOR_CYAN("Call Stack:"))
	for i := 0; i < len(allFrames)-1; i++ {
		frame := allFrames[i]
		relativePath, pathErr := filepath.Rel(".", frame.File)
		if pathErr != nil {
			relativePath = frame.File
		}

		fmt.Printf("%s:%d\n", COLOR_MAGENTA(relativePath), frame.Line)
	}
}

func log(msg string) {
	fmt.Println(msg)
}

func LoggerInit() {
	// stdout OR filePtr
}

// TODO (Jovanni): need to fix this!
func LogOnError(err error, format string, args ...any) {
	if err != nil {
		LogFatal(fmt.Sprintf("Error: %s\n", err))
		panic(err)
	}
}

func LogPrint(msg string) {
	log(msg)
}

func LogSuccess(msg string) {
	log(COLOR_GREEN(msg))
}

func LogDebug(msg string) {
	log(COLOR_BLUE(msg))
}

func LogWarn(msg string) {
	log(COLOR_MAGENTA(msg))
}

func LogError(msg string) {
	log(COLOR_RED(msg))
}

func LogFatal(msg string) {
	log(COLOR_BG_RED(msg))
}
