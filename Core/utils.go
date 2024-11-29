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

var (
	COLOR_RED     = color.New(color.FgRed).SprintFunc()
	COLOR_BLUE    = color.New(color.FgBlue).SprintFunc()
	COLOR_GREEN   = color.New(color.FgGreen).SprintFunc()
	COLOR_YELLOW  = color.New(color.FgYellow).SprintFunc()
	COLOR_MAGENTA = color.New(color.FgMagenta).SprintFunc()
	COLOR_CYAN    = color.New(color.FgCyan).SprintFunc()

	COLOR_BG_RED = color.New(color.FgWhite, color.BgRed).SprintFunc()
)

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

func GetBit(number int, bit_to_check int) int {
	return ((number & (1 << bit_to_check)) >> bit_to_check)
}

func SetBit(number *int, bit_to_set int) {
	*number |= (1 << bit_to_set)
}

func UnsetBit(number *int, bit_to_unset int) {
	*number &= (^(1 << bit_to_unset))
}

func ToggleBit(number *int, bit_to_toggle int) {
	*number ^= (1 << bit_to_toggle)
}

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

		fmt.Printf("  %s\n    %s:%d\n",
			COLOR_YELLOW(frame.Function),
			COLOR_MAGENTA(relativePath), frame.Line)
	}
}

func AssertOnError(err error) {
	if err != nil {
		LogFatal(fmt.Sprint("[Developer Error]: ", err))
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
