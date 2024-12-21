package ORM

import (
	"github.com/superg3m/stoic-go/Core/Utility"
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

	tableName := Utility.GetTypeName(stackModel)
	memberNames := Utility.GetStructMemberNames(stackModel, excludeList...)
	columnNames := getColumnNames(tableName, memberNames)
	values := Utility.GetStructValues(stackModel, excludeList...)
	pointers := Utility.GetStructMemberPointer(model, excludeList...)

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

	return ModelPayload{
		TableName: tableName,

		MemberNames:           memberNames,
		PrimaryKeyMemberNames: primaryKeyMemberNames,
		UniqueMemberNames:     uniqueMemberNames,
		ColumnNames:           columnNames,

		Values: values,

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
