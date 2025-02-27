package Utility

import (
	"fmt"
	"reflect"
	"slices"
)

func UpdateMemberValue(structure HeapAny, memberName string, data HeapAny) {
	v := reflect.ValueOf(structure)

	AssertMsg(TypeIsPointer(structure), "must be a pointer to a struct")

	structValue := v.Elem()
	field := structValue.FieldByName(memberName)
	AssertMsg(field.IsValid(), "field '%s' not found in struct", memberName)

	switch field.Kind() {
	case reflect.Int:
		field.SetInt(reflect.ValueOf(data).Int())
	case reflect.Int8:
		field.SetInt(int64(reflect.ValueOf(data).Int()))
	case reflect.Int16:
		field.SetInt(int64(reflect.ValueOf(data).Int()))
	case reflect.Int32:
		field.SetInt(int64(reflect.ValueOf(data).Int()))
	case reflect.Int64:
		field.SetInt(reflect.ValueOf(data).Int())
	case reflect.Uint:
		field.SetUint(uint64(reflect.ValueOf(data).Uint()))
	case reflect.Uint8:
		field.SetUint(uint64(reflect.ValueOf(data).Uint()))
	case reflect.Uint16:
		field.SetUint(uint64(reflect.ValueOf(data).Uint()))
	case reflect.Uint32:
		field.SetUint(uint64(reflect.ValueOf(data).Uint()))
	case reflect.Uint64:
		field.SetUint(reflect.ValueOf(data).Uint())
	case reflect.Float32:
		field.SetFloat(reflect.ValueOf(data).Float())
	case reflect.Float64:
		field.SetFloat(reflect.ValueOf(data).Float())
	case reflect.String:
		field.SetString(reflect.ValueOf(data).String())
	case reflect.Bool:
		field.SetBool(reflect.ValueOf(data).Bool())
	default:
		AssertMsg(false, "unsupported field type: %v", field.Kind())
	}
}

func GetMemberValue(structure StackAny, member string) any {
	AssertMsg(TypeIsStructure(structure), "structure is not of type structure")

	val := reflect.ValueOf(structure)
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == member {
			return val.Field(i).Interface()
		}
	}

	Assert(false)
	return nil
}

func TypeIsStructure(structure any) bool {
	typeKind := reflect.TypeOf(structure).Kind()

	return typeKind == reflect.Struct
}

func TypeIsPointer(structure any) bool {
	typeKind := reflect.TypeOf(structure).Kind()

	return typeKind == reflect.Ptr
}

func GetStructMemberNames(structure StackAny, excludeList ...string) []string {
	AssertMsg(TypeIsStructure(structure), "structure is not of type structure")

	value := reflect.ValueOf(structure)
	typeStruct := value.Type()

	var fieldNames []string
	for i := 0; i < typeStruct.NumField(); i++ {
		field := typeStruct.Field(i)

		if field.Anonymous || slices.Contains(excludeList, field.Name) {
			continue
		}

		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames
}

func GetStructMemberPointer(structure any, excludeList ...string) []any {
	value := reflect.ValueOf(structure)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	} else {
		ptr := reflect.New(reflect.TypeOf(structure))
		ptr.Elem().Set(value)
		value = ptr.Elem()
	}

	typeStruct := value.Type()
	var pointers []any

	for i := 0; i < typeStruct.NumField(); i++ {
		field := typeStruct.Field(i)

		if slices.Contains(excludeList, field.Name) {
			continue
		}

		if value.Field(i).CanAddr() {
			pointers = append(pointers, value.Field(i).Addr().Interface())
		}
	}

	return pointers
}

func GetStructMemberTypes(structure StackAny, excludeList ...string) map[string]string {
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

func GetStructValues(structure StackAny, excludeList ...string) []any {
	AssertMsg(TypeIsStructure(structure), "structure is not of type structure")

	val := reflect.ValueOf(structure)
	typ := val.Type()

	var values []any
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if field.Anonymous || slices.Contains(excludeList, field.Name) {
			continue
		}

		values = append(values, val.Field(i).Interface())
	}

	return values
}

func DereferencePointer(p any) any {
	v := reflect.ValueOf(p)
	AssertMsg(v.Kind() == reflect.Ptr, "argument must be a pointer")

	return v.Elem().Interface()
}

func SetNil[T any](target T) (T, error) {
	v := reflect.ValueOf(target)

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Func, reflect.Chan:
		return reflect.Zero(v.Type()).Interface().(T), nil
	default:
		return target, fmt.Errorf("type %s cannot be set to nil", v.Kind())
	}
}

func Copy[T any](source T, dest T) {
	v := reflect.ValueOf(source)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	destVal := reflect.ValueOf(dest).Elem()
	AssertMsg(v.Type().AssignableTo(destVal.Type()), "types are not compatible: %v to %v", v.Type(), destVal.Type())

	destVal.Set(v)
}
