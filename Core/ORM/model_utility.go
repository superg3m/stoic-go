package ORM

import "github.com/superg3m/stoic-go/Core/Utility"

func getModelMemberNames[T InterfaceCRUD](model T) []string {
	stackModel := Utility.DereferencePointer(model)
	return Utility.GetStructMemberNames(stackModel, excludeList...)
}

func getModelTableName[T InterfaceCRUD](model T) string {
	stackModel := Utility.DereferencePointer(model)
	return Utility.GetTypeName(stackModel)
}

func getModelValues[T InterfaceCRUD](model T) []any {
	stackModel := Utility.DereferencePointer(model)
	return Utility.GetStructValues(stackModel, excludeList...)
}

func getModelTypes[T InterfaceCRUD](model T) map[string]string {
	stackModel := Utility.DereferencePointer(model)
	return Utility.GetStructMemberTypes(stackModel, excludeList...)
}
