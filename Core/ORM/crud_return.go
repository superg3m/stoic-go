package ORM

import "github.com/superg3m/stoic-go/Core/Utility"

type CrudReturn struct {
	Utility.ErrorHandler
}

func CreateCRUD() CrudReturn {
	ret := CrudReturn{}
	return ret
}

func (c *CrudReturn) IsBad() bool {
	return c.GetErrors() != nil
}
