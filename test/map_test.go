package test

import (
	"Orchid/util"
	"testing"
)

func TestMap(t *testing.T) {
	arr := []int{1, 1, 4, 5, 1, 4, 1, 9, 1, 9, 8, 1, 0}
	arr = util.Map(arr, func(x int) int {
		return x * 7
	})
	t.Log(arr)
}
