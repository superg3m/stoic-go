package ORM

import "slices"

var acceptableFlagStrings = []string{"KEY", "UPDATABLE", "AUTO_INCREMENT", "UNIQUE"}

type Attribute struct {
	ColumnName string // The name of the column in the database
	TypeName   string
	Flags      []string // Bit flag to store Nullable, Updatable, etc.
}

func (attribute *Attribute) isPrimaryKey() bool {
	return slices.Contains(attribute.Flags, "KEY")
}

func (attribute *Attribute) isUpdatable() bool {
	return slices.Contains(attribute.Flags, "UPDATABLE")
}

func (attribute *Attribute) isAutoIncrement() bool {
	return slices.Contains(attribute.Flags, "AUTO_INCREMENT")
}

func (attribute *Attribute) isUnique() bool {
	return slices.Contains(attribute.Flags, "UNIQUE")
}
