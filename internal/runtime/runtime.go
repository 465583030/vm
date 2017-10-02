package runtime

import (
	"fmt"

	"github.com/open-compiler-project/vm/internal"
	"github.com/open-compiler-project/vm/internal/exec"
	"github.com/open-compiler-project/vm/internal/reflect"
)

type Runtime struct {
	Types      map[int]reflect.Type
	FramesInfo map[int]exec.FrameInfo
	Frame      *exec.Frame
	Globals    []reflect.Value
	GlobNames  []string
	Stack      exec.Stack

	IntPool    []internal.I64
	UintPool   []internal.U64
	F32Pool    []internal.F32
	F64Pool    []internal.F64
	StringPool []string

	Excode int
}

func (r *Runtime) Fatalf(str string, args ...interface{}) {
	fmt.Printf("fatal: "+str+"\n", args...)
	r.trace()
	fmt.Printf("runtime: exit code 0x%02x (%d)", r.Excode, r.Excode)
	r.stop()
}

func (r *Runtime) Errorf(str string, args ...interface{}) {
	fmt.Printf("error: "+str+"\n", args...)
	r.trace()
}

func (r *Runtime) trace() {}

func (r *Runtime) stop() {}
