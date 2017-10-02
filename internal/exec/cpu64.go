package exec

import (
	"github.com/open-compiler-project/vm/internal"
	"github.com/open-compiler-project/vm/internal/reflect"
	"github.com/open-compiler-project/vm/internal/runtime"
)

type VM struct {
	runtime *runtime.Runtime
	code    []Instruction
}

func InitVM(code []Instruction, runtime *runtime.Runtime) *VM {
	return &VM{runtime, code}
}

func (vm *VM) Start() {
	types := vm.runtime.Types
	stack := vm.runtime.Stack
	frame := vm.runtime.Frame
	ip := &frame.IP
	*ip = 0
	for {
		if *ip >= len(vm.code) {
			return
		}

		switch vm.code[*ip] {
		case Const:
			*ip++
			t := vm.code[*ip]
			switch types[int(t)] {
			case reflect.I32:
				stack.Push(internal.I32(vm.runtime.IntPool[int(vm.code[1+*ip])]))
				*ip += 2
			default:
				vm.runtime.Fatalf("(%s).const unsupported", types[int(t)].String())
				return
			}
		case Add:
			*ip++
			t := vm.code[*ip]
			switch types[int(t)] {
			case reflect.I32:
				a, _ := stack.Pop()
				b, _ := stack.Pop()
				stack.Push(a.(internal.I32) + b.(internal.I32))
				*ip++
			default:
				vm.runtime.Fatalf("(%s).add unsupported", types[int(t)].String())
				return
			}
		}
		runtime.Dump(stack)
	}
}
