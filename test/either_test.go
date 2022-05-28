package test

import (
	"Orchid/tpe"
	"go/types"
	"testing"
)

func TestEither(t *testing.T) {
	t.Log(tpe.Left[string](types.Error{}))
	t.Log(tpe.Left[string](types.Error{}).Right())
	t.Log(tpe.Left[string](types.Error{}).Left())
	t.Log(tpe.Right[int](1145141919810))
	t.Log(tpe.Right[int](1145141919810).Left())
	t.Log(tpe.Right[int](1145141919810).Right())
}
