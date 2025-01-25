package main

import (
	"errors"
	"fmt"
	"github.com/Quaestiox/result"
)

func needPositive(num int) result.ResultIF {
	if num < 0 {
		return result.Err(errors.New("error input"))
	}
	return result.Ok(num)
}

func needPositive2(num int) (r result.Result[int, error]) {
	if num < 0 {
		return r.Err(errors.New("error input"))
	}
	return r.Ok(num)
}

func normal(num int) (int, error) {
	if num < 0 {
		return num, errors.New("error input")
	}
	return num, nil
}

func main() {
	if res := needPositive(-1); res.IsOk() {
		fmt.Println(res.OK())
	} else {
		fmt.Println(res.ERR())
	}

	if res := needPositive2(1); res.IsOk() {
		fmt.Println(res.OK())
	} else {
		fmt.Println(res.ERR())
	}

	if res := result.AsRes(normal, -1); res.IsOk() {
		fmt.Println("ok")
		fmt.Println(res.OK())
	} else {
		fmt.Println("err")
		fmt.Println(res.ERR())
	}
}
