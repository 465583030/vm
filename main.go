package main

import (
	"github.com/open-compiler-project/vm/internal/exec"
	"github.com/open-compiler-project/vm/internal/reflect"
	"github.com/open-compiler-project/vm/internal/runtime"
)

func main() {
	finfo := map[int]*runtime.FrameInfo{
		0: &runtime.FrameInfo{
			Name:       "main",
			LocalNames: []string{},
			RetType:    reflect.Int,
			StartIP:    0,
		},
	}

	rt := runtime.InitRuntime(finfo)
	vm := exec.InitVM(
		[]exec.Instruction{
			exec.Const, 0x05, 0x00,
			exec.Const, 0x05, 0x01,
			exec.Add, 0x05,
		}, rt)

	vm.Start()
}
