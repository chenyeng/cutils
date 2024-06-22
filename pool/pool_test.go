package pool

import (
	"context"
	"errors"
	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"sync/atomic"
	"testing"
)

func TestPool(t *testing.T) {
	PatchConvey("normal case", t, func() {
		ctx := context.Background()
		const cnt = int32(1000)
		t := atomic.Int32{}
		pool := NewPool("normal", cnt)
		for i := int32(0); i < cnt; i++ {
			pool.Go(ctx, func() error {
				t.Add(1)
				return nil
			})
		}
		So(pool.Close(), ShouldBeNil)
		So(t.Load(), ShouldEqual, cnt)
	})

	PatchConvey("panic case", t, func() {
		ctx := context.Background()
		const cnt = int32(10)
		pool := NewPool("panic case", cnt)
		pool.Go(ctx, func() error {
			panic("aha!")
		})
		So(pool.Close(), ShouldNotBeNil)
	})

	PatchConvey("err case", t, func() {
		ctx := context.Background()
		const cnt = int32(10)
		pool := NewPool("err case", cnt)
		pool.Go(ctx, func() error {
			return errors.New("aha")
		})
		So(pool.Close(), ShouldNotBeNil)
	})

	PatchConvey("cancel case", t, func() {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		const cnt = int32(10)
		pool := NewPool("err case", cnt)
		i := atomic.Int32{}
		pool.Go(ctx, func() error {
			i.Add(1)
			return nil
		})
		So(pool.Close(), ShouldNotBeNil)
		So(i.Load(), ShouldBeZeroValue)
	})
}
