package panErr

import (
	"fmt"
	"reflect"
)

// FArgs 处理带有参数并且具有返回值的函数，利用反射执行，效率可能偏低
func FArgs(f interface{}, arg ...interface{}) (err error, resp []interface{}) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%+v", e)
		}
	}()
	rv := reflect.ValueOf(f)
	var par []reflect.Value
	for _, item := range arg {
		par = append(par, reflect.ValueOf(item))
	}
	re := rv.Call(par)
	for _, item := range re {
		resp = append(resp, item.Interface())
	}
	return err, resp
}

// F 处理不带有参数的函数，
func F(f func()) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%+v", e)
		}
	}()
	f()
	return err
}
