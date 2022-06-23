package config

import (
	"fmt"
	"strconv"
)

type _ConfigInt struct {
	value    int
	usage    string
	defValue int
}

func (c *_ConfigInt) Raw() interface{} {
	return c.value
}

func (c *_ConfigInt) Set(value string) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	c.value = v
	return nil
}

func (c *_ConfigInt) Usage() string {
	return c.usage
}

func (c *_ConfigInt) String() string {
	return fmt.Sprintf("%d", c.value)
}

func (c *_ConfigInt) Default() string {
	return fmt.Sprintf("%d", c.defValue)
}

func Int(name string, value int, usage string) *int {
	c := &_ConfigInt{
		value: value,
		usage: usage,
	}
	All.Set(name, c)
	return &c.value
}
