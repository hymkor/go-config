package config

import (
	"fmt"
	"strings"

	"github.com/hymkor/go-ignorecase-sorted"
)

const (
	True  = "true"
	Yes   = "yes"
	Blank = ""
	One   = "1"

	False = "false"
	No    = "no"
	Zero  = "0"

	ReversePrifix = "no"
)

type Standard interface {
	Set(string) error
	String() string
	Usage() string
	Default() string
}

var All ignoreCaseSorted.Dictionary[Standard]

type KeyNotFoundError struct {
	Key string
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("%s: configuration name not found", e.Key)
}

func Set(name, value string) error {
	config1, ok := All.Get(name)
	if !ok {
		if !strings.HasPrefix(name, ReversePrifix) {
			return &KeyNotFoundError{Key: name}
		}
		_name := name[len(ReversePrifix):]
		if len(_name) >= 1 && _name[0] == '-' {
			_name = _name[1:]
		}
		config1, ok = All.Get(_name)
		if !ok {
			return &KeyNotFoundError{Key: name}
		}
		name = _name
		switch strings.ToLower(value) {
		case Blank, Yes, True, One:
			value = False
		case No, False, Zero:
			value = True
		}
	}
	err := config1.Set(value)
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	return nil
}
