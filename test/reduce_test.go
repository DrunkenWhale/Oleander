package test

import (
	"Orchid/util"
	"log"
	"testing"
)

func TestReduce(t *testing.T) {
	arr := []int{1, 1, 4, 5, 1, 4, 1, 9, 1, 9, 8, 1, 0}
	brr := util.Reduce(arr, func(x int, y int) int { return x*7 + y })
	log.Println(brr)
}
