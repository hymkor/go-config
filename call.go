package config

type _ConfigCall struct {
	f        func(string) (string, error)
	value    string
	defValue string
	usage    string
}

func (c *_ConfigCall) Set(value string) error {
	var err error
	c.value, err = c.f(value)
	return err
}

func (c *_ConfigCall) Usage() string {
	return c.usage
}

func (c *_ConfigCall) String() string {
	return c.value
}

func (c *_ConfigCall) Default() string {
	return c.defValue
}

func Call(name, value, usage string, f func(string) (string, error)) {
	c := &_ConfigCall{
		f:        f,
		value:    value,
		defValue: value,
		usage:    usage,
	}
	All.Set(name, c)
}
