package Core

import (
	"fmt"
	"net/mail"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/fatih/color"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func castAny[T any](v any) T {
	var result any

	switch any(new(T)).(type) {
	case *int:
		switch v := v.(type) {
		case string:
			intValue, err := strconv.Atoi(v)
			AssertOnError(err)
			result = intValue
		case float64:
			result = int(v)
		case bool:
			if v {
				result = 1
			} else {
				result = 0
			}
		default:
			panic(fmt.Sprintf("unsupported conversion to int from type: %T", v))
		}
	case *string:
		switch v := v.(type) {
		case int:
			result = strconv.Itoa(v)
		case float64:
			result = strconv.FormatFloat(v, 'f', -1, 64)
		case bool:
			result = strconv.FormatBool(v)
		default:
			panic(fmt.Sprintf("unsupported conversion to string from type: %T", v))
		}
	case *bool:
		switch v := v.(type) {
		case string:
			boolValue, err := strconv.ParseBool(v)
			AssertOnError(err)
			result = boolValue
		case int:
			result = v != 0
		case float64:
			result = v != 0.0
		default:
			panic(fmt.Sprintf("unsupported conversion to bool from type: %T", v))
		}
	case *float64:
		switch v := v.(type) {
		case string:
			floatValue, err := strconv.ParseFloat(v, 64)
			AssertOnError(err)
			result = floatValue
		case int:
			result = float64(v)
		case bool:
			if v {
				result = 1.0
			} else {
				result = 0.0
			}
		default:
			panic(fmt.Sprintf("unsupported conversion to float64 from type: %T", v))
		}
	default:
		panic(fmt.Sprintf("unsupported target type: %T", any(new(T))))
	}

	return result.(T)
}

func formatArgs(format string, args ...any) string {
	return fmt.Sprintf(format, args)
}

func printCallStack() {
	// Retrieve program counters for all active stack frames
	pc := make([]uintptr, 32)   // Pre-allocate for stack frames
	n := runtime.Callers(2, pc) // Skip runtime.Callers and printCallStack
	frames := runtime.CallersFrames(pc[:n])

	// Store all frames to calculate total depth
	var allFrames []runtime.Frame
	for {
		frame, more := frames.Next()
		allFrames = append(allFrames, frame)
		if !more {
			break
		}
	}

	// Set up color formatting
	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	functionName := color.New(color.FgYellow).SprintFunc()
	fileDetails := color.New(color.FgMagenta).SprintFunc()

	fmt.Println(header("Call Stack:"))
	for i := 0; i < len(allFrames)-1; i++ {
		frame := allFrames[i]
		// Get relative file path
		relativePath, pathErr := filepath.Rel(".", frame.File)
		if pathErr != nil {
			relativePath = frame.File // Fallback to absolute path
		}

		// Print stack frame information
		fmt.Printf("  %s\n    %s:%d\n",
			functionName(frame.Function),
			fileDetails(relativePath), frame.Line)
	}
}

func AssertOnError(err error) {
	if err != nil {
		header := color.New(color.FgRed, color.Bold).SprintFunc()
		fmt.Println(header("[Developer Error]: ", err))
		printCallStack()
		os.Exit(-1)
	}
}

func LoggerInit(fileName string) {

}

// TODO (Jovanni): need to fix this!
func LogOnError(err error, format string, args ...any) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		panic(err)
	}
}

// Asserts
// Logging
//
