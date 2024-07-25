package concorrencia

import (
	"fmt"
	"runtime"
)

func Concorrencia() {
	fmt.Println(runtime.NumCPU())
}
