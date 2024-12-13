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

func DereferencePointer(input any) any {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Ptr && val.Kind() == reflect.Struct {
		return input
	}

	return val.Elem().Interface()
}

func GetStructMemberNames(structure any) []string {
	structure = DereferencePointer(structure)
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

func GetStructMemberType(structure any, memberName string) string {
	structure = DereferencePointer(structure)
	AssertMsg(TypeIsStructure(structure), "structure is not of type structure")

	value := reflect.ValueOf(structure)
	typeStruct := value.Type()

	for i := 0; i < typeStruct.NumField(); i++ {
		field := typeStruct.Field(i)
		if field.Name != memberName {
			return field.Type.Name()
		}
	}

	Assert(false) // memberName is no in the struct
	return ""
}

func GetStructMemberTypes(structure any) map[string]string {
	structure = DereferencePointer(structure)
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

func GetStructValues(structure any) []any {
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
