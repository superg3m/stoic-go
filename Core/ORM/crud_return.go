package ORM

type CrudReturn struct {
	errMSG string
	isBad  bool
}

func CreateCRUD() CrudReturn {
	ret := CrudReturn{
		isBad:  false,
		errMSG: "",
	}

	return ret
}

func (c *CrudReturn) GetErrorMsg() string {
	return c.errMSG
}

func (c *CrudReturn) setErrorMsg(errMSG string) {
	c.makeBad()
	c.errMSG = errMSG
}

func (c *CrudReturn) makeBad() {
	c.isBad = true
}

func (c *CrudReturn) IsBad() bool {
	return c.isBad == true
}
