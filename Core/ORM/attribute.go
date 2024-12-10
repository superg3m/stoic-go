package ORM

type ORM_FLAG int

const ( // ORM_Flags
	PRIMARY_KEY ORM_FLAG = 1 << iota // 1 << 0 = 1 (bit 0)
	NULLABLE                         // 1 << 1 = 2 (bit 1)
	UPDATABLE                        // 1 << 2 = 4 (bit 2)
)

type Attribute struct {
	name  string   // The name of the column in the database
	flags ORM_FLAG // Bit flag to store Nullable, Updatable, etc.
}

func (attribute *Attribute) isNullable() bool {
	return attribute.flags&NULLABLE != 0
}

func (attribute *Attribute) isUpdatable() bool {
	return attribute.flags&UPDATABLE != 0
}
