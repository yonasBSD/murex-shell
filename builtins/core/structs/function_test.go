package structs

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test/count"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestAliases(t *testing.T) {
	alias := "TestAlias"

	count.Tests(t, 3)

	lang.InitEnv()

	if lang.GlobalAliases.Exists(alias) {
		t.Fatalf("Expecting '%s' not to exist, yet it does", alias)
	}

	p := lang.NewTestProcess()
	p.Name.Set("alias")
	p.Parameters.DefineParsed([]string{alias + "=foobar"})
	err := cmdAlias(p)
	if err != nil {
		t.Fatalf("Error calling cmdAlias(): %s", err.Error())
	}

	if !lang.GlobalAliases.Exists(alias) {
		t.Fatalf("Expecting '%s' to be created, it did not", alias)
	}

	p = lang.NewTestProcess()
	p.Name.Set("!alias")
	p.IsNot = true
	p.Parameters.DefineParsed([]string{alias})
	err = cmdUnalias(p)
	if err != nil {
		t.Fatalf("Error calling cmdAlias() for the 2nd time: %s", err.Error())
	}

	if lang.GlobalAliases.Exists(alias) {
		t.Fatalf("Expecting '%s' to be destroyed, it still exists", alias)
	}
}

func TestFunction(t *testing.T) {
	fn := "TestFunction"

	count.Tests(t, 3)

	lang.InitEnv()

	if lang.MxFunctions.Exists(fn) {
		t.Fatalf("Expecting '%s' not to exist, yet it does", fn)
	}

	p := lang.NewTestProcess()
	p.Name.Set("function")
	p.Parameters.DefineParsed([]string{fn, "{ test }"})
	err := cmdFunc(p)
	if err != nil {
		t.Fatalf("Error calling cmdFunc(): %s", err.Error())
	}

	if !lang.MxFunctions.Exists(fn) {
		t.Fatalf("Expecting '%s' to be created, it did not", fn)
	}

	p = lang.NewTestProcess()
	p.Name.Set("!function")
	p.IsNot = true
	p.Parameters.DefineParsed([]string{fn})
	err = cmdUnfunc(p)
	if err != nil {
		t.Fatalf("Error calling cmdFunc() for the 2nd time: %s", err.Error())
	}

	if lang.MxFunctions.Exists(fn) {
		t.Fatalf("Expecting '%s' to be destroyed, it still exists", fn)
	}
}

func TestPrivate(t *testing.T) {
	fn := "TestPrivate"
	mod := fmt.Sprintf("GoTest-%d", rand.Int())

	count.Tests(t, 4)

	lang.InitEnv()

	if lang.PrivateFunctions.ExistsString(fn, mod) {
		t.Fatalf("Expecting '%s/%s' not to exist, yet it does", mod, fn)
	}

	p := lang.NewTestProcess()
	p.Name.Set("function")
	p.FileRef = &ref.File{
		Source: &ref.Source{
			Module: mod,
		},
	}

	p.Parameters.DefineParsed([]string{fn, "{ test }"})
	err := cmdPrivate(p)
	if err != nil {
		t.Fatalf("Error calling cmdPrivate(): %s", err.Error())
	}

	if !lang.PrivateFunctions.ExistsString(fn, mod) {
		t.Fatalf("Expecting '%s/%s' to be created, it did not", mod, fn)
	}

	err = lang.PrivateFunctions.Undefine(fn, p.FileRef)
	if err != nil {
		t.Fatalf("Expecting no errors. Got %v", err)
	}

	if lang.PrivateFunctions.ExistsString(fn, mod) {
		t.Fatalf("Expecting '%s/%s' to be removed, it was not", mod, fn)
	}

	err = lang.PrivateFunctions.Undefine(fn, p.FileRef)
	if err == nil {
		t.Fatalf("Expecting and error")
	}
}
