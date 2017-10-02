package internal

import "github.com/open-compiler-project/vm/internal/reflect"

type NullPtr Ptr

var Null NullPtr = NullPtr{}

func (NullPtr) Type() reflect.Type {
	return reflect.Null
}

func IsNull(val reflect.Value) Bool {
	return val == Null
}
