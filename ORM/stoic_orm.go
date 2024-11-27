package ORM

import (
	"fmt"
	"reflect"
)

type ORM_FLAG int

const (
	PRIMARY_KEY ORM_FLAG = 1 << iota // 1 << 0 = 1 (bit 0)
	NULLABLE                         // 1 << 1 = 2 (bit 1)
	UPDATABLE                        // 1 << 2 = 4 (bit 2)
)

type FieldMetadata struct {
	field string   // Name of the field in the DB
	flags ORM_FLAG // Bit flag to store Nullable, Updatable, etc.
}

type BaseStoicTable struct {
	tableName    string
	fieldMeta    map[string]FieldMetadata
	originalData map[string]interface{}
}

type I_CRUD interface {
	canUpdate()
	canCreate()
	canDelete()
}

func (base *BaseStoicTable) getFieldMetadata(fieldName string) (FieldMetadata, bool) {
	meta, exists := base.fieldMeta[fieldName]
	return meta, exists
}

func (f *FieldMetadata) isNullable() bool {
	return f.flags&NULLABLE != 0
}

func (f *FieldMetadata) isUpdatable() bool {
	return f.flags&UPDATABLE != 0
}

func RegisterTableName(tableName string) {

}

func RegisterTableColumn(structMemberName string, attributeName string, flags ORM_FLAG) {

}

func (b *BaseStoicTable) storeOriginalData(v interface{}) {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		b.originalData[field.Name] = val.Field(i).Interface()
	}
}

func (b *BaseStoicTable) hasFieldChanged(v interface{}, fieldName string) bool {
	val := reflect.ValueOf(v).Elem()
	currentValue := val.FieldByName(fieldName).Interface()
	originalValue, exists := b.originalData[fieldName]
	if !exists {
		return false
	}
	return originalValue != currentValue
}

func (b *BaseStoicTable) update(v interface{}) error {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		meta, exists := b.fieldMeta[field.Name]
		if !exists {
			continue
		}

		// Only check updatable fields
		if meta.isUpdatable() && b.hasFieldChanged(v, field.Name) {
			fmt.Printf("Field '%s' has been updated.\n", field.Name)
		} else {
			fmt.Printf("Field '%s' has not been updated.\n", field.Name)
		}
	}

	// Proceed with actual update logic here (e.g., DB update)
	return nil
}

func (b *BaseStoicTable) create(v interface{}) error {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		_, exists := b.fieldMeta[field.Name]
		if !exists {
			continue
		}
	}

	// Proceed with actual create logic here (e.g., DB insert)
	return nil
}
