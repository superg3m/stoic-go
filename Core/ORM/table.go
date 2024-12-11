package ORM

import (
	"github.com/superg3m/stoic-go/Core/Utility"
)

type MemberAttributeMap map[string]Attribute // Key: StructMemberName

var tempTableName string
var globalTable map[string]MemberAttributeMap // Key: TableName

func RegisterTableName(tableName string) {
	tempTableName = tableName
}

func RegisterTableColumn(memberName string, columnName string, flags ORM_FLAG) {
	globalTable[tempTableName][memberName] = Attribute{
		ColumnName: columnName,
		Flags:      flags,
	}

	tempTableName = ""
}

func GetAttributes(tableName string) (map[string]Attribute, bool) {
	attributes, exists := globalTable[tableName]
	Utility.Assert(exists)
	return attributes, exists
}

func getAttribute(tableName string, memberName string) (Attribute, bool) {
	attribute, exists := globalTable[tableName][memberName]
	Utility.Assert(exists)
	return attribute, exists
}
