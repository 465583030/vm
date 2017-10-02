package exec

import (
	"github.com/open-compiler-project/vm/internal/runtime"
)

type Instruction = uint16

const (
	Nop     Instruction = 0x0000
	Malloc  Instruction = 0x0001
	Calloc  Instruction = 0x0002
	Realloc Instruction = 0x0003
	Alloca  Instruction = 0x0004
	New     Instruction = 0x0005
	Free    Instruction = 0x0006
	Cmp     Instruction = 0x000a
	Cpy     Instruction = 0x000b
	Set     Instruction = 0x000c
	Swap    Instruction = 0x0010
	Dup     Instruction = 0x0011
	Drop    Instruction = 0x0012
	Jmp     Instruction = 0x0020
	Bz      Instruction = 0x0021
	Bnz     Instruction = 0x0022
	Beq     Instruction = 0x0023
	Bne     Instruction = 0x0024
	Blt     Instruction = 0x0025
	Ble     Instruction = 0x0026
	Blz     Instruction = 0x0027
	Bgt     Instruction = 0x0028
	Bge     Instruction = 0x0029
	Bgz     Instruction = 0x002a
	Bt      Instruction = 0x002b
	Bf      Instruction = 0x002c
	Bnl     Instruction = 0x002f
	Call    Instruction = 0x0030
	Return  Instruction = 0x0031
	Add     Instruction = 0x0100 // type
	Sub     Instruction = 0x0200 // type
	Mul     Instruction = 0x0300 // type
	Div     Instruction = 0x0400 // type
	Rem     Instruction = 0x0500 // type
	And     Instruction = 0x0600 // type
	Or      Instruction = 0x0700 // type
	Xor     Instruction = 0x0800 // type
	Not     Instruction = 0x0900 // type
	Store   Instruction = 0x1000 // type idx
	GStore  Instruction = 0x2000 // type idx
	Load    Instruction = 0x3000 // type idx
	GLoad   Instruction = 0x4000 // type idx
	Const   Instruction = 0x5000 // type idx
	Cast    Instruction = 0x6000 // from to
	Syscall Instruction = 0xe000 // num
	Exit    Instruction = 0xf000
)

var instrMap map[Instruction]func(*runtime.Runtime)

func InstrcutionMapping() map[Instruction]func(*runtime.Runtime) {
	return instrMap
}
