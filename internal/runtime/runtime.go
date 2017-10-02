package runtime

import (
	"fmt"

	"os"

	"github.com/Spriithy/gods/stacks/arraystack"
	"github.com/open-compiler-project/vm/internal"
	"github.com/open-compiler-project/vm/internal/reflect"
)

type Runtime struct {
	Types      map[int]reflect.Type
	FramesInfo map[int]*FrameInfo
	Frame      *Frame
	Globals    []reflect.Value
	GlobNames  []string
	Stack      *Stack

	IntPool    []internal.Int
	UIntPool   []internal.UInt
	F32Pool    []internal.F32
	F64Pool    []internal.F64
	StringPool []string

	Excode int
}

func (r *Runtime) Fatalf(str string, args ...interface{}) {
	fmt.Printf("fatal: "+str+"\n", args...)
	r.trace()
	fmt.Printf("runtime: exit code 0x%02x (%d)\n", r.Excode, r.Excode)
	r.stop()
}

func (r *Runtime) Errorf(str string, args ...interface{}) {
	fmt.Printf("error: "+str+"\n", args...)
	r.trace()
}

func (r *Runtime) trace() {}

func (r *Runtime) stop() {
	os.Exit(r.Excode + 1)
}

func InitRuntime(finfo map[int]*FrameInfo) *Runtime {
	types := map[int]reflect.Type{
		0x00: reflect.Bool,
		0x01: reflect.I8,
		0x02: reflect.U8,
		0x03: reflect.I16,
		0x04: reflect.U16,
		0x05: reflect.I32,
		0x06: reflect.U32,
		0x07: reflect.I64,
		0x08: reflect.U64,
		0x09: reflect.Int,
		0x0a: reflect.UInt,
		0x0b: reflect.F32,
		0x0c: reflect.F64,
		0x0d: reflect.Null,
		0x0e: reflect.Ptr,
	}

	return &Runtime{
		Types:      types,
		FramesInfo: finfo,
		Frame:      NewFrame(finfo[0], nil),
		Stack:      arraystack.New(),
		Globals:    []reflect.Value{internal.Null, internal.Null},
		GlobNames:  []string{"x", "y"},
		F32Pool:    []internal.F32{},
		F64Pool:    []internal.F64{},
		IntPool:    []internal.Int{13, 8},
		UIntPool:   []internal.UInt{},
		StringPool: []string{},
	}
}
