package Utility

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"net/mail"
	"os"
	"path/filepath"
	"time"
)

type StackAny any
type HeapAny any

func NewTime(newTime time.Time) *time.Time {
	ret := new(time.Time)
	*ret = newTime

	return ret
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

func CastAny[T any](v any) T {
	var result any

	switch any(new(T)).(type) {
	case *int:
		result = cast.ToInt(v)
	case *bool:
		result = cast.ToBool(v)
	case *float64:
		result = cast.ToFloat64(v)
	case *string:
		result = cast.ToString(v)

	case *[]int:
		result = cast.ToIntSlice(v)
	case *[]bool:
		result = cast.ToBoolSlice(v)
	case *[]float64:
		result = cast.ToFloat64Slice(v)
	case *[]string:
		result = cast.ToStringSlice(v)
	}

	return result.(T)
}

func findFileAtDepth(filename string, maxDepth int) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	for depth := 0; depth <= maxDepth; depth++ {
		filePath := filepath.Join(currentDir, filename)
		if _, err := os.Stat(filePath); err == nil {
			return filePath, nil // File found
		} else if !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("error checking file %s: %w", filePath, err)
		}

		currentDir = filepath.Dir(currentDir)
	}

	return "", fmt.Errorf("file %s not found within %d directories", filename, maxDepth)
}

func GetSiteSettings() map[string]any {
	filename := "siteSettings.json"
	var MaxSearchDepth = 3

	filePath, err := findFileAtDepth(filename, MaxSearchDepth)
	AssertOnError(err)

	byteData, err := os.ReadFile(filePath)
	AssertOnError(err)

	var settings map[string]any
	err = json.Unmarshal(byteData, &settings)
	AssertOnError(err)

	return settings
}

func Sha256HashString(str string) string {
	h := sha256.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}
