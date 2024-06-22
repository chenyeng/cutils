package pool

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/gopool"
	"runtime/debug"
	"sync"
)

type Pool interface {
	Go(ctx context.Context, f func() error)
	Close() error
}

func NewPool(name string, cap int32) Pool {
	p := gopool.NewPool(name, cap, gopool.NewConfig())
	return &pool{
		p:  p,
		wg: &sync.WaitGroup{},
	}
}

type pool struct {
	p   gopool.Pool
	wg  *sync.WaitGroup
	err error
}

func (p *pool) Close() error {
	p.wg.Wait()
	return p.err
}

func (p *pool) Go(ctx context.Context, f func() error) {
	if err := ctx.Err(); err != nil {
		p.err = err
	}
	if p.err != nil {
		return
	}
	p.wg.Add(1)
	p.p.CtxGo(ctx, func() {
		defer p.wg.Done()
		defer func() {
			if e := recover(); e != nil {
				p.err = fmt.Errorf("panic hanppend :%+v, stack :%s", e, string(debug.Stack()))
			}
		}()
		if err := f(); err != nil {
			p.err = err
		}

	})
}
