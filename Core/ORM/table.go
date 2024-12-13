package ORM

import (
	"fmt"

	"github.com/superg3m/stoic-go/Core/Utility"
)

type MemberAttributeMap map[string]Attribute // Key: StructMemberName

var tempTableName string
var tempTableTypes map[string]string
var globalTable map[string]MemberAttributeMap // Key: TableName

func init() {
	Utility.LogDebug("ORM_TABLE")
	globalTable = make(map[string]MemberAttributeMap)
}

func RegisterTableName(table InterfaceCRUD, tableName string) {
	tempTableName = tableName
	tempTableTypes = Utility.GetStructMemberTypes(table)
}

func ensureAttributesAreValid(attribute Attribute) {
	msg := fmt.Sprintf("Not allowed to apply auto increment attribute when type is not numeric! Member %s is of type %s", attribute.MemberName, attribute.TypeStr)
	Utility.AssertMsg(attribute.TypeStr != "int" && attribute.isAutoIncrement(), msg)
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
		ColumnName: columnName,
		Flags:      finalFlag,
	}

	globalTable[tempTableName][memberName] = Attribute{
		MemberName: memberName,
		ColumnName: columnName,
		Flags:      finalFlag,
		TypeStr:    tempTableTypes[memberName],
	}

	ensureAttributesAreValid(attribute)
}

func GetAttributes(tableName string) (map[string]Attribute, bool) {
	attributes, exists := globalTable[tableName]
	Utility.Assert(exists)
	return attributes, exists
}

func getAttribute(tableName string, memberName string) (Attribute, bool) {
	attribute, exists := globalTable[tableName][memberName]
	Utility.AssertMsg(exists, fmt.Sprintf("Table: %s, Member: %s | Doesn't exist", tableName, memberName))
	return attribute, exists
}
