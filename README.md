# Result

Result type in Go for error handling like Haskell and Rust.

## Install

`go get github.com/Quaestiox/result`

## Example
```
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

func main() {
	if res := needPositive(-1); res.IsOk() {
		fmt.Println(res.OK())
	} else {
		fmt.Println(res.ERR())
	}

	if res := needPositive(1); res.IsOk() {
		fmt.Println(res.OK())
	} else {
		fmt.Println(res.ERR())
	}
}
```
