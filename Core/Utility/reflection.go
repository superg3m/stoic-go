package Utility

import (
	"reflect"
)

func TypeIsStructure(structure any) bool {
	typeKind := reflect.TypeOf(structure).Kind()

	return typeKind == reflect.Struct
}

func TypeIsPointer(structure any) bool {
	typeKind := reflect.TypeOf(structure).Kind()

	return typeKind == reflect.Ptr
}

func GetStructMemberNames(structure StackAny) []string {
	AssertMsg(TypeIsStructure(structure), "structure is not of type structure")

	value := reflect.ValueOf(structure)
	typeStruct := value.Type()

	var fieldNames []string
	for i := 0; i < typeStruct.NumField(); i++ {
		field := typeStruct.Field(i)

		if field.Anonymous {
			continue
		}

		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames
}

func GetStructMemberPointer(structure HeapAny) []any {
	value := reflect.ValueOf(structure)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	} else {
		AssertMsg(false, "structure must be a pointer")
	}

	typeStruct := value.Type()
	var pointers []any

	for i := 0; i < typeStruct.NumField(); i++ {
		field := typeStruct.Field(i)

		if field.Anonymous {
			continue
		}

		if value.Field(i).CanAddr() {
			pointers = append(pointers, value.Field(i).Addr().Interface())
		}
	}

	return pointers
}

func GetStructMemberTypes(structure StackAny) map[string]string {
	AssertMsg(TypeIsStructure(structure), "structure is not of type structure")

	value := reflect.ValueOf(structure)
	typeStruct := value.Type()

	ret := make(map[string]string, typeStruct.NumField())
	for i := 0; i < typeStruct.NumField(); i++ {
		field := typeStruct.Field(i)

		if field.Anonymous {
			continue
		}

		ret[field.Name] = field.Type.Name()
	}

	return ret
}

func GetTypeName(s StackAny) string {
	val := reflect.ValueOf(s)

	return val.Type().Name()
}

func GetStructValues(structure StackAny) []any {
	AssertMsg(TypeIsStructure(structure), "structure is not of type structure")

	val := reflect.ValueOf(structure)
	typ := val.Type()

	var values []any
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if field.Anonymous {
			continue
		}

		values = append(values, val.Field(i).Interface())
	}

	return values
}
