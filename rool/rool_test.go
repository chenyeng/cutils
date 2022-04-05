package rool

import (
	"fmt"
	"testing"
)

func TestRool(t *testing.T) {
	r := NewRoutinePool(10)
	for i := 0; i < 100; i++ {
		i := i
		if i%2 == 0 {
			r.GO(func() {
				panic(i)
			})
		} else {
			if i%3 == 0 {
				r.GOArgs(func() {
					panic(i)
				})
			} else {
				r.GOArgs(func(i int) int {
					return i
				}, i)
			}
		}
	}
	resp, err := r.Done()
	fmt.Printf("resp: %+v err: %+v", resp, err)
}
