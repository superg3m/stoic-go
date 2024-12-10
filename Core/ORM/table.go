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

func RegisterTableColumn(structMemberName string, attributeName string, flags ORM_FLAG) {
	globalTable[tempTableName][structMemberName] = Attribute{
		name:  attributeName,
		flags: flags,
	}
}

func getAttribute(tableName string, structMemberName string) (Attribute, bool) {
	meta, exists := globalTable[tableName][structMemberName]
	Utility.Assert(exists)
	return meta, exists
}
