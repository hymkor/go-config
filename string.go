package config

type _ConfigString struct {
	value    string
	usage    string
	defValue string
}

func (c *_ConfigString) Set(value string) error {
	c.value = value
	return nil
}

func (c *_ConfigString) Usage() string {
	return c.usage
}

func (c *_ConfigString) String() string {
	return c.value
}

func (c *_ConfigString) Default() string {
	return c.defValue
}

func String(name, value, usage string) *string {
	c := &_ConfigString{
		value: value,
		usage: usage,
	}
	All.Set(name, c)
	return &c.value
}
