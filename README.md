# Result

Result type in Go for error handling like Haskell and Rust.

## Install

`go get github.com/Quaestiox/result`

## Example
```
// First way is using the ResultIF interface if you need return different types
func needPositive(num int) result.ResultIF {
	if num < 0 {
		return result.Err(errors.New("error input"))
	}
	return result.Ok(num)
}

// Second way is using the specific Result struct. In this way, you should use named result parameters.
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
