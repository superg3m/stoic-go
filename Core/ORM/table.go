package ORM

import (
	"github.com/superg3m/stoic-go/Core/Utility"
)

type MemberAttributeMap map[string]Attribute // Key: StructMemberName

var tempTableName string
var tempAutoIncrementFound bool
var globalTable map[string]MemberAttributeMap // Key: TableName

func init() {
	globalTable = make(map[string]MemberAttributeMap)
	tempAutoIncrementFound = false
}

func registerTableName(tableName string) {
	tempTableName = tableName
	globalTable[tempTableName] = make(MemberAttributeMap)
}

func registerTableColumn(memberName, columnName, typeName string, flags []string) {
	attribute := Attribute{
		ColumnName: columnName,
		TypeName:   typeName,
		Flags:      flags,
	}

	if !tempAutoIncrementFound && (attribute.isAutoIncrement() && !attribute.isPrimaryKey()) {
		Utility.AssertMsg(false, "Stoic-Go does not support models where auto increment is not primary key as well")
	}

	if attribute.isAutoIncrement() {
		tempAutoIncrementFound = true
	}

	globalTable[tempTableName][memberName] = attribute
}

func RegisterModel(model InterfaceCRUD) {
	payload := getModelPayload(model)
	registerTableName(payload.TableName)
	for i, memberName := range payload.MemberNames {
		columnName := payload.ColumnNames[i]
		flags := payload.Flags[columnName]
		typeName := payload.Types[memberName]

		registerTableColumn(memberName, columnName, typeName, flags)
	}
}

func GetAttributes(tableName string) map[string]Attribute {
	attributes, exists := globalTable[tableName]
	Utility.Assert(exists)
	return attributes
}
