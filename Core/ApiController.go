package Core

type I_CRUD inteface {
	create()
	delete()
	update()
}



func (b *Base) CanCreate(obj interface{}) {
	if specific, ok := obj.(SpecificCreator); ok {
		specific.Create() // Call the specific implementation
	} else {
		b.Create() // Fallback to base logic
	}
}