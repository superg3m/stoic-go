package ORM

import (
	"github.com/superg3m/stoic-go/Core/Utility"
	"reflect"
	"slices"
	"strings"
	"time"
)

type ModelPayload struct {
	TableName string

	PrimaryKeyMemberNames []string
	UniqueMemberNames     []string
	MemberNames           []string
	ColumnNames           []string
	Flags                 map[string][]string
	Types                 map[string]string

	Values []any

	Pointers           []any
	PrimaryKeyPointers []any
	UniquePointers     []any
}

func getModelPayload[T InterfaceCRUD](model T) ModelPayload {
	stackModel := Utility.DereferencePointer(model)
	tableName := Utility.GetTypeName(stackModel)

	var memberNames []string
	var columnNames []string
	var pointers []any
	flags := make(map[string][]string)
	typeNames := make(map[string]string)
	types := Utility.GetStructMemberTypes(stackModel)

	var (
		primaryKeyMemberNames []string
		uniqueMemberNames     []string
		primaryKeyPointers    []any
		uniquePointers        []any
		correctedValues       []any
	)

	v := reflect.ValueOf(model)
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := range t.NumField() {
		field := t.Field(i)
		memberName := field.Name

		dbTag, ok := field.Tag.Lookup("db")
		if !ok {
			delete(types, memberName)
			continue
		}

		parts := strings.Split(dbTag, ",")
		columnName := strings.TrimSpace(parts[0])
		columnNames = append(columnNames, columnName)
		memberNames = append(memberNames, memberName)
		pointers = append(pointers, v.Field(i).Addr().Interface())

		var flagAccumulator []string
		for _, part := range parts[1:] {
			normalizedFlag := strings.TrimSpace(part)
			Utility.AssertMsg(slices.Contains(acceptableFlagStrings, normalizedFlag), "Flag: %s not allowed", normalizedFlag)

			if normalizedFlag == "KEY" {
				primaryKeyMemberNames = append(primaryKeyMemberNames, memberName)
				primaryKeyPointers = append(primaryKeyPointers, v.Field(i).Addr().Interface())
			} else if normalizedFlag == "UNIQUE" {
				uniqueMemberNames = append(uniqueMemberNames, memberName)
				uniquePointers = append(uniquePointers, v.Field(i).Addr().Interface())
			}

			flagAccumulator = append(flagAccumulator, normalizedFlag)

		}
		flags[columnName] = flagAccumulator

		if types[memberName] == "Time" {
			value := reflect.ValueOf(model).Elem().FieldByName(memberName)
			correctedValue := value.Interface().(time.Time).Format(time.DateTime)
			correctedValues = append(correctedValues, correctedValue)
		} else if types[memberName] == "*Time" {
			value := reflect.ValueOf(model).Elem().FieldByName(memberName)
			if value.IsNil() {
				correctedValues = append(correctedValues, nil)
			} else {
				correctedValue := value.Interface().(*time.Time).Format(time.DateTime)
				correctedValues = append(correctedValues, correctedValue)
			}
		} else {
			value := reflect.ValueOf(model).Elem().FieldByName(memberName)
			correctedValues = append(correctedValues, value.Interface())
		}

	}

	Utility.AssertMsg(len(memberNames) > 0, "%s is missing db tags", tableName)

	return ModelPayload{
		TableName: tableName,

		MemberNames:           memberNames,
		PrimaryKeyMemberNames: primaryKeyMemberNames,
		UniqueMemberNames:     uniqueMemberNames,
		ColumnNames:           columnNames,
		Flags:                 flags,
		Types:                 typeNames,

		Values: correctedValues,

		Pointers:           pointers,
		PrimaryKeyPointers: primaryKeyPointers,
		UniquePointers:     uniquePointers,
	}
}

func getColumnNames(tableName string, memberNames []string) []string {
	var ret []string
	attributes := GetAttributes(tableName)
	for _, memberName := range memberNames {
		ret = append(ret, attributes[memberName].ColumnName)
	}

	return ret
}
