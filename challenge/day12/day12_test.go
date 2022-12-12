package day12

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	for c := 'a'; c <= 'z'; c++ {
		fmt.Println(c)
	}
}
