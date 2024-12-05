package Utility

import "reflect"

func GetStructFieldNames[T any](structure T) []string {
	var ret []string
	val := reflect.ValueOf(structure).Elem()
	for i := 0; i < val.NumField(); i++ {
		ret = append(ret, val.Type().Field(i).Name)
	}

	return ret
}

func GetStructValues[T any](structure T) []any {
	val := reflect.ValueOf(structure)
	values := make([]any, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		values[i] = val.Field(i).Interface()
	}

	return values
}
