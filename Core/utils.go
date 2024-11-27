package Core

import (
	"fmt"
	"net/mail"
	"strconv"
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
			PanicOnError(err, "failed to convert string to int: %v", v)
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
			PanicOnError(err, "failed to convert string to bool: %v", v)
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
			PanicOnError(err, "failed to convert string to float64: %v", v)
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

func PanicOnError(err error, format string, args ...any) {
	if err != nil {
		fmt.Printf("Developer Error: %s\n", formatArgs(format, args))
		panic(err)
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
