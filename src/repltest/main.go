package main

import (
	"errors"
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

func main() {
	repl.REPL(new(TestHandler))
}
