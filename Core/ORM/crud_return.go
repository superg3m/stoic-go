package ORM

type CrudReturn struct {
	err    error
	errMsg string
	isBad  bool
}

func CreateCRUD() CrudReturn {
	ret := CrudReturn{
		err:    nil,
		errMsg: "",
		isBad:  false,
	}

	return ret
}

func (c *CrudReturn) GetError() error {
	return c.err
}

func (c *CrudReturn) GetErrorMsg() string {
	return c.errMsg
}

func (c *CrudReturn) setError(err error) {
	c.makeBad()
	c.errMsg = err.Error()
}

func (c *CrudReturn) makeBad() {
	c.isBad = true
}

func (c *CrudReturn) IsBad() bool {
	return c.isBad == true
}
