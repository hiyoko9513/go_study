package main

import "fmt"

/*
type error interface {
    Error() string
}
*/

type CustomError struct {
	ErrCode int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &CustomError{Message: "エラーが発生しました", ErrCode: 001}
}

func Do() (int, error) {
	err := RaiseError()
	return 1, err
}

func main() {
	i, err := Do()
	if err != nil {
		e, ok := err.(*CustomError)
		if ok {
			fmt.Println(e.ErrCode)
		}
		fmt.Println(err)
	}
	fmt.Println(i)
}
