package panErr

import (
	"fmt"
	"testing"
)

func TestFArgs(t *testing.T) {
	err, _ := FArgs(func(a int, b []int64) {
		panic(int64(a) + b[0] + b[1] + b[2])
	}, 1, []int64{1, 2, 3})
	if err != nil {
		fmt.Printf("err is :%+v\n", err)
	} else {
		t.Error("err")
	}

	_, resp := FArgs(func(a int, b []int64) int64 {
		return int64(a) + b[0] + b[1] + b[2]
	}, 1, []int64{1, 2, 3})
	if resp != nil {
		fmt.Printf("resp is: %+v,%T", resp[0].(int64), resp[0])
	} else {
		t.Error("err")
	}
}

func TestF(t *testing.T) {
	err := F(func() {
		panic("err happen")
	})
	if err != nil {
		fmt.Printf("errrrrr :%+v", err)
	} else {
		t.Error(err)
	}
}
