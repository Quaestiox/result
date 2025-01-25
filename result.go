package result

import (
	"fmt"
	"os"
	"reflect"
)

type ResultIF interface {
	OK() any
	ERR() any
	IsOk() bool
	IsErr() bool
}

type Okey[T any] struct {
	Value T
}

type Error[E any] struct {
	Value E
}

func Err[E any](err E) Result[any, E] {
	errStruct := Error[E]{Value: err}
	return Result[any, E]{O: nil, E: &errStruct}
}

func Ok[T any](ok T) Result[T, any] {
	okStruct := Okey[T]{Value: ok}
	return Result[T, any]{O: &okStruct, E: nil}
}

func (ok Okey[T]) OK() any {
	return ok.Value
}

func (ok Okey[T]) ERR() any {
	return nil
}

func (ok Okey[T]) IsOk() bool {
	if ok.OK() != nil {
		return true
	}

	return false
}

func (ok Okey[T]) IsErr() bool {
	if ok.ERR() != nil {
		return true
	}

	return false
}
func (err Error[E]) ERR() any {
	return err.Value
}

func (err Error[E]) OK() any {
	return nil
}

func (err Error[E]) IsOk() bool {
	if err.OK() != nil {
		return true
	}

	return false
}

func (err Error[E]) IsErr() bool {
	if err.ERR() != nil {
		return true
	}

	return false
}

type Result[T, E any] struct {
	O *Okey[T]
	E *Error[E]
}

func (r Result[T, E]) OK() any {
	if r.O.OK() != nil {
		return r.O.Value
	}

	return nil
}

func (r Result[T, E]) ERR() any {
	if r.E.ERR() != nil {
		return r.E.Value
	}
	return nil
}

func (r Result[T, E]) IsOk() bool {
	return r.O != nil
}

func (r Result[T, E]) IsErr() bool {
	return r.E != nil
}

func (r Result[T, E]) Ok(ok T) Result[T, E] {
	okStruct := Okey[T]{Value: ok}
	return Result[T, E]{O: &okStruct, E: nil}
}

func (r Result[T, E]) Err(err E) Result[T, E] {
	errStruct := Error[E]{Value: err}
	return Result[T, E]{O: nil, E: &errStruct}
}

func AsRes(fn any, args ...any) Result[any, any] {

	fnType := reflect.TypeOf(fn)
	fnValue := reflect.ValueOf(fn)
	outc := fnType.NumOut()
	inc := fnType.NumIn()

	if fnType.Kind() != reflect.Func {
		fmt.Println("provided is not a function.")
		os.Exit(1)
	}

	if outc != 2 {
		fmt.Println("only support two out parameters.")
		os.Exit(1)
	}

	if len(args) != inc {
		fmt.Printf("error: expected %d arguments, got %d.\n", inc, len(args))
		os.Exit(1)
	}

	paras := make([]reflect.Value, inc)

	for i, v := range args {
		paras[i] = reflect.ValueOf(v)
	}

	outs := fnValue.Call(paras)

	if outs[1].IsNil() {
		return Ok(outs[0].Interface())
	} else {
		return Err(outs[1].Interface())
	}

}
