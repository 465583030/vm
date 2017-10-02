package internal

import (
	"github.com/open-compiler-project/vm/internal/reflect"
)

type Ptr struct {
	typ reflect.Type
	to  reflect.Value
}

func (p Ptr) Type() reflect.Type {
	return p.typ
}

func PtrTo(val reflect.Value) reflect.Value {
	return Ptr{reflect.PtrTo(val.Type()), val}
}
