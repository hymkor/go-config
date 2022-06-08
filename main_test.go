package config_test

import (
	"errors"
	"testing"

	"github.com/hymkor/go-config"
)

func TestAll(t *testing.T) {
	confString := config.String("sss", "", "string test")
	confBool := config.Bool("bbb", false, "bool test")
	confInt := config.Int("iii", 0, "int test")
	confNoBool := config.Bool("nnn", true, "non bool test")
	var confCall string
	config.Call("call", "(not called)", "call test", func(v string) (string, error) {
		confCall = v
		return "(called)", nil
	})

	// test default value
	if *confString != "" {
		t.Fatalf("string default error %v", *confString)
		return
	}
	if *confBool != false {
		t.Fatalf("bool default error %v", *confBool)
		return
	}
	if *confInt != 0 {
		t.Fatalf("int default error %v", *confInt)
		return
	}
	if *confNoBool != true {
		t.Fatalf("no bool default error %v", *confNoBool)
	}

	// test setter
	if err := config.Set("sss", "_sss_"); err != nil {
		t.Fatalf("string set: %s", err.Error())
		return
	}
	if err := config.Set("bbb", "yes"); err != nil {
		t.Fatalf("boolean set: %s", err.Error())
		return
	}
	if err := config.Set("iii", "3"); err != nil {
		t.Fatalf("int set: %s", err.Error())
		return
	}
	if err := config.Set("no-nnn", ""); err != nil {
		t.Fatalf("non bool set: %s", err.Error())
		return
	}
	if err := config.Set("call", "callvalue"); err != nil {
		t.Fatalf("call set: %s", err.Error())
		return
	}

	// test setter error pattern (not found)
	var perr *config.KeyNotFoundError
	if err := config.Set("notexistkey", "true"); !errors.As(err, &perr) {
		t.Fatal("key not found did not occured")
		return
	}

	// test setter
	if err := config.Set("iii", "x"); err == nil {
		t.Fatalf("int could set: %s", "x")
		return
	}

	// error test
	if err := config.Set("bbb", "yesno"); err == nil {
		t.Fatalf("boolean could set: %s", "yesno")
		return
	}

	if *confString != "_sss_" {
		t.Fatalf("string get: %v", *confString)
		return
	}
	if *confBool != true {
		t.Fatalf("bool get: %v", *confBool)
		return
	}
	if *confInt != 3 {
		t.Fatalf("int get: %v", *confInt)
		return
	}
	if *confNoBool != false {
		t.Fatalf("non bool get: %v", *confNoBool)
		return
	}
	if confCall != "callvalue" {
		t.Fatalf("call get: %v", confCall)
		return
	}
}
