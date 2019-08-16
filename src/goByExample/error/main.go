package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string  {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg int) (int, error)  {
	if arg == 42 {
		return -1, errors.New("can not work with 42")
	}
	return arg + 3, nil
}

func f2(arg int) (int, error)  {
	if arg == 42 {
		return -1, &argError{arg, "can not work with it"}
	}
	return arg + 3, nil
}

func main() {
	s := []int{7, 42}

	for _, i := range s {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println(r)
		}
	}

	for _, i := range s {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	_, e := f2(42)
	// e.(...) 代表类型断言
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
