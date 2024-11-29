package ORM

import (
	"fmt"
	"reflect"

	"github.com/superg3m/stoic-go/Core"
)

type ORM_FLAG int

const (
	PRIMARY_KEY ORM_FLAG = 1 << iota // 1 << 0 = 1 (bit 0)
	NULLABLE                         // 1 << 1 = 2 (bit 1)
	UPDATABLE                        // 1 << 2 = 4 (bit 2)
)

type FieldMetadata struct {
	attributeName string   // Name of the field in the DB
	flags         ORM_FLAG // Bit flag to store Nullable, Updatable, etc.
}

type BaseStoicTable struct {
	tableName         string
	fieldMeta         map[string]FieldMetadata
	fieldOriginalData map[string]any
}

var globalTable BaseStoicTable

type I_CRUD interface {
	canUpdate()
	canCreate()
	canDelete()
}

func (base *BaseStoicTable) getFieldMetadata(fieldName string) (FieldMetadata, bool) {
	meta, exists := base.fieldMeta[fieldName]
	Core.Assert(exists)
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

func (b *BaseStoicTable) storeOriginalData(v interface{}) {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		b.fieldOriginalData[field.Name] = val.Field(i).Interface()
	}
}

func (b *BaseStoicTable) hasFieldChanged(v interface{}, fieldName string) bool {
	val := reflect.ValueOf(v).Elem()
	currentValue := val.FieldByName(fieldName).Interface()
	originalValue, exists := b.fieldOriginalData[fieldName]
	if !exists {
		return false
	}

	return originalValue != currentValue
}

func (b *BaseStoicTable) update(v interface{}) {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		meta, exists := b.fieldMeta[field.Name]
		if !exists {
			continue
		}

		if meta.isUpdatable() && b.hasFieldChanged(v, field.Name) {
			// actual update logic ...
		} else {
			errMsg := fmt.Sprintf("Field '%s' is not allowed to be updated.\n", field.Name)
			Core.AssertMsg(false, errMsg)
		}
	}
}

func (b *BaseStoicTable) create(v interface{}) {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		b.storeOriginalData(v)
		_, exists := b.fieldMeta[field.Name]
		if !exists {
			continue
		}
	}

	// actual create logic ...
}
