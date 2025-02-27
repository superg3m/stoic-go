package ORM

import (
	"github.com/superg3m/stoic-go/Core/Utility"
)

type MemberAttributeMap map[string]Attribute // Key: StructMemberName

var tempTableName string
var tempTypes map[string]string
var tempAutoIncrementFound bool
var globalTable map[string]MemberAttributeMap // Key: TableName

func init() {
	globalTable = make(map[string]MemberAttributeMap)
	tempAutoIncrementFound = false
}

func RegisterTableName(table InterfaceCRUD) {
	stackModel := Utility.DereferencePointer(table)
	meta := Utility.GetMemberValue(stackModel, "Meta")
	tempTableName = Utility.GetTypeName(stackModel)
	tempTypes = Utility.GetStructMemberTypes(meta, excludeList...)
}

func RegisterTableColumn(memberName string, columnName string, flags ...ORM_FLAG) {
	if globalTable[tempTableName] == nil {
		globalTable[tempTableName] = make(MemberAttributeMap)
	}

	var finalFlag ORM_FLAG
	for _, flag := range flags {
		finalFlag |= flag
	}

	attribute := Attribute{
		MemberName: memberName,
		ColumnName: columnName,
		TypeName:   tempTypes[memberName],
		Flags:      finalFlag,
	}

	if !tempAutoIncrementFound && (attribute.isAutoIncrement() && !attribute.isPrimaryKey()) {
		Utility.AssertMsg(false, "Stoic-Go does not support models where auto increment is not primary key as well")
	}

	if attribute.isAutoIncrement() {
		tempAutoIncrementFound = true
	}

	globalTable[tempTableName][memberName] = attribute
}

func GetAttributes(tableName string) map[string]Attribute {
	attributes, exists := globalTable[tableName]
	Utility.Assert(exists)
	return attributes
}
