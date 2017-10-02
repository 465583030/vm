package runtime

import (
	"github.com/Spriithy/gods/stacks/arraystack"
	"github.com/open-compiler-project/vm/internal"
	"github.com/open-compiler-project/vm/internal/reflect"
)

// Stack holds the elements in a slice
type Stack = arraystack.Stack

func Dump(stack *Stack) {
	print("[")
	for _, x := range stack.Values() {
		print(x.(reflect.Value).(internal.I32), " ")
	}
	println("\b]")
}
