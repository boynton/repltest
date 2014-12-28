package main

import (
	"errors"
	"fmt"
	"github.com/boynton/repl"
	"strings"
)

type TestHandler struct {
	value string
}

//test incomplete lines by counting parens -- they must match.
func (th *TestHandler) Eval(expr string) (interface{}, bool, error) {
	whole := th.value + expr
	opens := len(strings.Split(whole, "("))
	closes := len(strings.Split(whole, ")"))
	if opens > closes {
		th.value = whole + " "
		return nil, true, nil
	} else if closes > opens {
		th.value = ""
		return nil, false, errors.New("Unbalanced ')'")
	} else {
		th.value = ""
		return whole, false, nil
	}
}

func (th *TestHandler) Reset() {
	th.value = ""
}
func (th *TestHandler) Start() []string {
	fmt.Println("Load and return persistent history here")
	return nil
}
func (th *TestHandler) Stop(history []string) {
	fmt.Printf("Save %d lines of history here\n", len(history))
}

func (th *TestHandler) Prompt() string {
	return "> "
}

func (th *TestHandler) Complete(expr string) (string, []string) {
	//expr is what has been typed so far. Return any more text we can autocomplete,
	//and the slice of options that remain possible
	return "", nil
}

func main() {
	repl.REPL(new(TestHandler))
}
