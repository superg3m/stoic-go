package ORM

import (
	"fmt"
	"github.com/superg3m/stoic-go/core/Utility"
	"reflect"
)

type ORM_FLAG int

const ( // ORM_Flags
	PRIMARY_KEY ORM_FLAG = 1 << iota // 1 << 0 = 1 (bit 0)
	NULLABLE                         // 1 << 1 = 2 (bit 1)
	UPDATABLE                        // 1 << 2 = 4 (bit 2)
)

type FieldMetadata struct {
	attributeName string   // Name of the field in the Database
	flags         ORM_FLAG // Bit flag to store Nullable, Updatable, etc.
}

type InterfaceCRUD interface {
	canCreate() bool
	canUpdate() bool
	canDelete() bool
}

type BaseStoicTable struct {
	InterfaceCRUD
	tableName string
	fieldMeta map[string]FieldMetadata // Key is memeberFieldName
	isCreated bool
}

var globalTable BaseStoicTable

func (base *BaseStoicTable) getFieldMetadata(fieldName string) (FieldMetadata, bool) {
	meta, exists := base.fieldMeta[fieldName]
	Utility.Assert(exists)
	return meta, exists
}

func (f *FieldMetadata) isNullable() bool {
	return f.flags&NULLABLE != 0
}

func (f *FieldMetadata) isUpdatable() bool {
	return f.flags&UPDATABLE != 0
}

func RegisterTableName(tableName string) {
	globalTable.tableName = tableName
}

func RegisterTableColumn(structMemberName string, attributeName string, flags ORM_FLAG) {
	globalTable.fieldMeta[structMemberName] = FieldMetadata{
		attributeName: attributeName,
		flags:         flags,
	}
}

func (b *BaseStoicTable) update(v interface{}) {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldMeta, exists := b.fieldMeta[field.Name]
		if !exists {
			continue
		}

		if fieldMeta.isUpdatable() {
			// actual update logic ...
		} else {
			errMsg := fmt.Sprintf("Field '%s' is not allowed to be updated.\n", field.Name)
			Utility.AssertMsg(false, errMsg)
		}
	}
}

func (b *BaseStoicTable) create(v interface{}) {
	Utility.Assert(b.canCreate())

	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		_, exists := b.fieldMeta[field.Name]
		if !exists {
			continue
		}
	}

	b.isCreated = true

	createInDatabase

	// actual create logic ...
}
