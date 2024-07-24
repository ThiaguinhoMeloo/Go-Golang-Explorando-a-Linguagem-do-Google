package functions

import "runtime/debug"

func Func3() {
	debug.PrintStack()
}

func Func2() {
	Func3()
}

func Func1() {
	Func2()
}
