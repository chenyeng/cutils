package rool

import (
	"github.com/chenyeng/cutils/panErr"
	"sync"
)

type routinePool struct {
	pool  chan struct{}
	errCh chan error
	errs  []error
	reCh  chan []interface{}
	wg    sync.WaitGroup
	re    []interface{}
}

func NewRoutinePool(size int) *routinePool {
	t := &routinePool{
		pool:  make(chan struct{}, size),
		errCh: make(chan error),
		reCh:  make(chan []interface{}),
		wg:    sync.WaitGroup{},
	}

	// 初始化的时候新建一个协程回收错误
	go func() {
		for {
			if err, ok := <-t.errCh; ok {
				if err != nil {
					t.errs = append(t.errs, err)
				}
				t.wg.Done()
			} else {
				break
			}
		}
	}()

	// 初始化一个协程用于处理返回值
	go func() {
		for {
			if resp, ok := <-t.reCh; ok {
				if resp != nil {
					t.re = append(t.re, resp)
				}
				t.wg.Done()
			} else {
				break
			}
		}
	}()
	return t
}

// GO 从协程池中取出一个协程进行处理
func (r *routinePool) GO(f func()) {
	r.pool <- struct{}{}
	r.wg.Add(1)
	go func() {
		err := panErr.F(f)
		r.errCh <- err
		<-r.pool
	}()
}

func (r *routinePool) GOArgs(f interface{}, args ...interface{}) {
	r.pool <- struct{}{}
	r.wg.Add(2)
	go func() {
		err, resp := panErr.FArgs(f, args...)
		r.errCh <- err
		r.reCh <- resp
		<-r.pool
	}()
}

// Wait 等待所有协程处理完毕
func (r *routinePool) Wait() {
	r.wg.Wait()
}

// Done 关闭通道并等待所有协程处理完毕
func (r *routinePool) Done() ([]interface{}, []error) {
	close(r.pool)
	r.Wait()
	return r.re, r.errs
}
