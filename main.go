package main

import (
	"github.com/open-compiler-project/vm/internal"
)

func main() {
	val := internal.Null

	ptr := internal.PtrTo(val)

	println(internal.IsNull(ptr))
	println(val.Type().String(), ptr.Type().String())
}
