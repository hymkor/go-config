package config

import (
	"fmt"
	"strings"
)

type _ConfigBool struct {
	value    *bool
	usage    string
	defValue bool
}

func (c *_ConfigBool) Set(value string) error {
	value = strings.ToLower(strings.TrimSpace(value))
	if c.value == nil {
		c.value = new(bool)
	}
	switch value {
	case "1", "yes", "true", "":
		*c.value = true
	case "0", "no", "false":
		*c.value = false
	default:
		return fmt.Errorf("%v: not a boolean value", value)
	}
	return nil
}

func (c *_ConfigBool) Usage() string {
	return c.usage
}

func (c *_ConfigBool) String() string {
	if c.value != nil && *c.value {
		return "true"
	} else {
		return "false"
	}
}

func (c *_ConfigBool) Default() string {
	if c.defValue {
		return True
	} else {
		return False
	}
}

func BoolVar(p *bool, name string, value bool, usage string) {
	c := &_ConfigBool{
		value: p,
		usage: usage,
	}
	*c.value = value
	All.Set(name, c)
}

func Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	BoolVar(p, name, value, usage)
	return p
}
