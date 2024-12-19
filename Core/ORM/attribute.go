package ORM

type ORM_FLAG int

const ( // ORM_Flags
	KEY            ORM_FLAG = 1 << iota // 1 << 0 = 1 (bit 0)
	UPDATABLE                           // 1 << 1 = 2 (bit 1)
	AUTO_INCREMENT                      // 1 << 2 = 4 (bit 2)
	UNIQUE                              // 1 << 3 = 8 (bit 3)
)

type Attribute struct {
	MemberName string
	ColumnName string // The name of the column in the database
	TypeName   string
	Flags      ORM_FLAG // Bit flag to store Nullable, Updatable, etc.
}

func (attribute *Attribute) isPrimaryKey() bool {
	return attribute.Flags&KEY != 0
}

func (attribute *Attribute) isUpdatable() bool {
	return attribute.Flags&UPDATABLE != 0
}

func (attribute *Attribute) isAutoIncrement() bool {
	return attribute.Flags&AUTO_INCREMENT != 0
}

func (attribute *Attribute) isUnique() bool {
	return attribute.Flags&UNIQUE != 0
}
