package ORM

import (
	"fmt"
	"github.com/superg3m/stoic-go/Core/Utility"
)

type MemberAttributeMap map[string]Attribute // Key: StructMemberName

var tempTableName string
var globalTable map[string]MemberAttributeMap // Key: TableName

func init() {
	Utility.LogDebug("ORM_TABLE")
	globalTable = make(map[string]MemberAttributeMap)
}

func RegisterTableName(tableName string) {
	tempTableName = tableName
}

func RegisterTableColumn(memberName string, columnName string, flags ORM_FLAG) {
	if globalTable[tempTableName] == nil {
		globalTable[tempTableName] = make(MemberAttributeMap)
	}

	globalTable[tempTableName][memberName] = Attribute{
		ColumnName: columnName,
		Flags:      flags,
	}
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
