package main

import (
	"inicio/functions"
)

func main() {
	approveds := []string{"Maria", "Pedro", "Guilherme", "Allan"} // slice
	functions.PrintApproved(approveds...)
}
