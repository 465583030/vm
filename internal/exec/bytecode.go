package exec

import (
	"github.com/open-compiler-project/vm/internal/runtime"
)

type Instruction = uint16

const (
	Nop     Instruction = 0x0000
	Malloc  Instruction = 0x0010
	Calloc  Instruction = 0x0011
	Realloc Instruction = 0x0012
	Alloca  Instruction = 0x0013
	New     Instruction = 0x0014
	Free    Instruction = 0x0015
	Cmp     Instruction = 0x001a
	Cpy     Instruction = 0x001b
	Set     Instruction = 0x001c

	Swap Instruction = 0x001
	Dup  Instruction = 0xf002
	Drop Instruction = 0xf003

	Call   Instruction = 0xff00
	Return Instruction = 0xff01

	Add Instruction = 0x0100
	Sub Instruction = 0x0200
	Mul Instruction = 0x0300
	Div Instruction = 0x0400
	Rem Instruction = 0x0500
	And Instruction = 0x0600
	Or  Instruction = 0x0700
	Xor Instruction = 0x0800
	Not Instruction = 0x0900

	Store   Instruction = 0x1000
	GStore  Instruction = 0x2000
	Load    Instruction = 0x3000
	GLoad   Instruction = 0x4000
	Const   Instruction = 0x5000
	Cast    Instruction = 0x6000
	Syscall Instruction = 0xe000
	Exit    Instruction = 0xf000
)

var instrMap map[Instruction]func(*runtime.Runtime)

func InstrcutionMapping() map[Instruction]func(*runtime.Runtime) {
	return instrMap
}
