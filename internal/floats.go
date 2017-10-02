package internal

import (
	"unsafe"

	"github.com/open-compiler-project/vm/internal/reflect"
)

type (
	F32 float32
	F64 float64
)

func (F32) Type() reflect.Type { return reflect.F32 }
func (F64) Type() reflect.Type { return reflect.F64 }

func F32FromU32Bits(u32 U32) F32 {
	return *(*F32)(unsafe.Pointer(&u32))
}

func (f32 F32) ToU32Bits() U32 {
	return *(*U32)(unsafe.Pointer(&f32))
}

func F64FromU64Bits(u64 U64) F64 {
	return *(*F64)(unsafe.Pointer(&u64))
}

func (f64 F64) ToU64Bits() U64 {
	return *(*U64)(unsafe.Pointer(&f64))
}
