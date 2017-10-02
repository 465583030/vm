package exec

import (
	"fmt"

	"github.com/open-compiler-project/vm/internal"
	"github.com/open-compiler-project/vm/internal/reflect"
)

type FrameInfo struct {
	Name       string
	StartIP    int
	LocalNames []string
	RetType    reflect.Type
}

type Frame struct {
	info   *FrameInfo
	locals []reflect.Value
	retVal reflect.Value
	IP     int
	parent *Frame
}

func NewFrame(finfo *FrameInfo, parent *Frame) *Frame {
	return &Frame{
		info:   finfo,
		locals: make([]reflect.Value, len(finfo.LocalNames)),
		retVal: internal.Null,
		parent: parent,
	}
}

func (f *Frame) Trace() string {
	trc := fmt.Sprintf(" \t%s\n", f.info.Name)
	if f.parent != nil {
		return "trace: last is latest\n>" + f.parent.Trace() + trc
	}
	return trc
}
