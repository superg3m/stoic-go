package ORM

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Database"
	"github.com/superg3m/stoic-go/Core/Utility"
	"reflect"
)

type InterfaceCanCRUD interface {
	canCreate() bool
	canUpdate() bool
	canDelete() bool
}

type StoicModel struct {
	InterfaceCanCRUD

	db        *sqlx.DB
	tableName string
	isCreated bool
}

func (model *StoicModel) Update() error {
	Utility.AssertMsg(model.isCreated, fmt.Sprintf("%s Model must be created first before attempting to update!", model.tableName))

	if !model.InterfaceCanCRUD.canUpdate() {
		return fmt.Errorf("update not allowed for this model")
	}

	val := reflect.ValueOf(model).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldMeta, exists := getAttribute(model.tableName, field.Name)
		if !exists {
			continue
		}

		Utility.AssertMsg(fieldMeta.isUpdatable(), fmt.Sprintf("field '%s' is not updatable", field.Name))
	}

	err := Database.UpdateRecord(model.db, model.tableName, model)
	Utility.AssertOnError(err)

	return nil
}

func (model *StoicModel) Create() error {
	if !model.InterfaceCanCRUD.canCreate() {
		return fmt.Errorf("canCreate() returned false")
	}

	val := reflect.ValueOf(model).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		_, exists := getAttribute(model.tableName, field.Name)
		if !exists {
			continue
		}
	}

	model.isCreated = true

	err := Database.InsertRecord(model.db, model.tableName, model)
	Utility.AssertOnError(err)

	return nil
}

func (model *StoicModel) Delete() error {
	if !model.InterfaceCanCRUD.canDelete() {
		return fmt.Errorf("canDelete() returned false")
	}

	Utility.AssertMsg(model.isCreated, fmt.Sprintf("%s Model must be created first before attempting to delete!", model.tableName))

	err := Database.DeleteRecord(model.db, model.tableName, model)
	Utility.AssertOnError(err)

	// actual create logic ...

	return nil
}
