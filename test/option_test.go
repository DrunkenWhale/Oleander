package test

import (
	"Orchid/tpe"
	"testing"
)

type test1 struct {
	testFiled string
	author    string
}

func TestOption(t *testing.T) {
	opt1 := tpe.NewOption(114514)
	t.Log(opt1.IsEmpty())
	t.Log(opt1.NonEmpty())
	t.Log(opt1.Get())
	temp := tpe.NewOption[string]("1919810")
	t.Log(temp.IsEmpty())
	temp = tpe.EmptyOption[string]()
	t.Log(temp)
	t.Log(tpe.Some(114514))
	t.Log(tpe.None[any]())
	forever := test1{"succeed", "health"}
	t.Log(tpe.Some[test1](forever))
}
