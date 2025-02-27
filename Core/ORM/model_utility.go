package ORM

import (
	"github.com/superg3m/stoic-go/Core/Utility"
	"reflect"
	"time"
)

type ModelPayload struct {
	TableName string

	PrimaryKeyMemberNames []string
	UniqueMemberNames     []string
	MemberNames           []string
	ColumnNames           []string

	Values []any

	Pointers           []any
	PrimaryKeyPointers []any
	UniquePointers     []any
}

func getModelPayload[T InterfaceCRUD](model T) ModelPayload {
	stackModel := Utility.DereferencePointer(model)
	meta := Utility.GetMemberValue(stackModel, "Meta")

	tableName := Utility.GetTypeName(stackModel)
	memberNames := Utility.GetStructMemberNames(meta, excludeList...)
	columnNames := getColumnNames(tableName, memberNames)
	pointers := Utility.GetStructMemberPointer(meta, excludeList...)

	var (
		primaryKeyMemberNames []string
		uniqueMemberNames     []string

		primaryKeyPointers []any
		uniquePointers     []any
	)

	attributes := GetAttributes(tableName)

	for i, memberName := range memberNames {
		attribute := attributes[memberName]
		pointer := pointers[i]
		if attribute.isPrimaryKey() {
			primaryKeyMemberNames = append(primaryKeyMemberNames, memberName)
			primaryKeyPointers = append(primaryKeyPointers, pointer)
		}

		if attribute.isUnique() {
			uniqueMemberNames = append(uniqueMemberNames, memberName)
			uniquePointers = append(uniquePointers, pointer)
		}
	}

	var formattedTimeValues []any
	types := Utility.GetStructMemberTypes(meta, excludeList...)

	for i, memberName := range memberNames {
		fieldType, exists := types[memberName]
		Utility.Assert(exists)
		if fieldType == "time.Time" || fieldType == "*time.Time" {
			value := reflect.ValueOf(model).Elem().FieldByName(memberName)
			if value.IsValid() && value.Kind() == reflect.Struct && value.Type() == reflect.TypeOf(time.Time{}) {
				formattedTime := value.Interface().(time.Time).Format(time.DateTime)
				formattedTimeValues = append(formattedTimeValues, formattedTime)
			}
		} else {
			originalValue := Utility.GetStructValues(meta, excludeList...)
			formattedTimeValues = append(formattedTimeValues, originalValue[i])
		}
	}

	return ModelPayload{
		TableName: tableName,

		MemberNames:           memberNames,
		PrimaryKeyMemberNames: primaryKeyMemberNames,
		UniqueMemberNames:     uniqueMemberNames,
		ColumnNames:           columnNames,

		Values: formattedTimeValues,

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
