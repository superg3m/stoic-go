package ORM

type ORM_FLAG int

const ( // ORM_Flags
	PRIMARY_KEY ORM_FLAG = 1 << iota // 1 << 0 = 1 (bit 0)
	NULLABLE                         // 1 << 1 = 2 (bit 1)
	UPDATABLE                        // 1 << 2 = 4 (bit 2)
	UNIQUE                           // 1 << 4 = 8 (bit 3)
)

type Attribute struct {
	ColumnName string   // The name of the column in the database
	Flags      ORM_FLAG // Bit flag to store Nullable, Updatable, etc.
}

func (attribute *Attribute) isNullable() bool {
	return attribute.Flags&NULLABLE != 0
}

func (attribute *Attribute) isUpdatable() bool {
	return attribute.Flags&UPDATABLE != 0
}

func (attribute *Attribute) isUnique() bool {
	return attribute.Flags&UNIQUE != 0
}

func (attribute *Attribute) isPrimaryKey() bool {
	return attribute.Flags&PRIMARY_KEY != 0
}
