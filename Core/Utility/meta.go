package Utility

import (
	"reflect"
)

func GetStructFieldNames(structure interface{}) []string {
	val := reflect.ValueOf(structure)
	typ := val.Type()

	var fieldNames []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if field.Anonymous {
			continue
		}

		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			fieldNames = append(fieldNames, dbTag)
		}
	}

	return fieldNames
}

func GetStructValues(structure interface{}) ([]interface{}, error) {
	val := reflect.ValueOf(structure)
	typ := val.Type()

	var values []interface{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if field.Anonymous {
			continue
		}

		values = append(values, val.Field(i).Interface())
	}

	return values, nil
}
