package internal

import "github.com/open-compiler-project/vm/internal/reflect"

type (
	I8 int8
	U8 uint8

	I16 int16
	U16 uint16

	I32 int32
	U32 uint32

	I64 int64
	U64 uint64

	Int  int
	UInt uint
)

func (*I8) Type() reflect.Type   { return reflect.I8 }
func (*U8) Type() reflect.Type   { return reflect.U8 }
func (*I16) Type() reflect.Type  { return reflect.I16 }
func (*U16) Type() reflect.Type  { return reflect.U16 }
func (*I32) Type() reflect.Type  { return reflect.I32 }
func (*U32) Type() reflect.Type  { return reflect.U32 }
func (*I64) Type() reflect.Type  { return reflect.I64 }
func (*U64) Type() reflect.Type  { return reflect.U64 }
func (*Int) Type() reflect.Type  { return reflect.Int }
func (*UInt) Type() reflect.Type { return reflect.UInt }
