package reflect

import "unsafe"

type (
	Type interface {
		String() string
		IsRef() bool
		Size() uint
	}

	baseType struct {
		name string
		size uint
	}

	refType struct {
		name string
		to   Type
	}
)

func (t *baseType) String() string { return t.name }
func (t *baseType) IsRef() bool    { return false }
func (t *baseType) Size() uint     { return t.size }

func (t *refType) String() string { return t.name }
func (t *refType) IsRef() bool    { return true }
func (t *refType) Size() uint     { return UInt.Size() }

var (
	// just to get specific platform int/uint sizes
	i int
	u uint

	Null Type = &refType{name: "null"}
	Ptr  Type = &refType{name: "ptr"}

	Bool Type = &baseType{name: "bool", size: 1}

	I8 Type = &baseType{name: "i8", size: 1}
	U8 Type = &baseType{name: "u8", size: 1}

	I16 Type = &baseType{name: "i16", size: 2}
	U16 Type = &baseType{name: "u16", size: 2}

	I32 Type = &baseType{name: "i32", size: 4}
	U32 Type = &baseType{name: "u32", size: 4}

	I64 Type = &baseType{name: "i64", size: 8}
	U64 Type = &baseType{name: "u64", size: 8}

	Int  Type = &baseType{name: "int", size: uint(unsafe.Sizeof(&i))}
	UInt Type = &baseType{name: "uint", size: uint(unsafe.Sizeof(&u))}

	F32 Type = &baseType{name: "f32", size: 4}
	F64 Type = &baseType{name: "f64", size: 8}
)

func PtrTo(t Type) Type {
	return &refType{name: "*" + t.String(), to: t}
}
