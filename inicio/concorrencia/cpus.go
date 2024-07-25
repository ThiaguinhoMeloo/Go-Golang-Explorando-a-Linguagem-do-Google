package concorrencia

import (
	"fmt"
	"runtime"
)

func NumCPU() {
	fmt.Println(runtime.NumCPU())
}
