package exec

import "github.com/open-compiler-project/vm/internal/reflect"

// Stack holds the elements in a slice
type Stack struct {
	elements []reflect.Value
	size     int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

func (stack *Stack) Peek() (reflect.Value, bool) {
	if stack.Empty() {
		return nil, false
	}
	return stack.elements[stack.size-1], true
}

func (stack *Stack) Push(val reflect.Value) {
	stack.growBy(1)
	stack.elements[stack.size] = val
	stack.size++
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (stack *Stack) Pop() (reflect.Value, bool) {
	if stack.Empty() {
		return nil, false
	}

	stack.size--
	return stack.elements[stack.size], true
}

// Empty returns true if list does not contain any elements.
func (stack *Stack) Empty() bool {
	return stack.size == 0
}

// Size returns number of elements within the list.
func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) resize(cap int) {
	newElements := make([]reflect.Value, cap, cap)
	copy(newElements, stack.elements)
	stack.elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (stack *Stack) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(stack.elements)
	if stack.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		stack.resize(newCapacity)
	}
	stack.shrink()
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (stack *Stack) shrink() {
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(stack.elements)
	if stack.size <= int(float32(currentCapacity)*shrinkFactor) {
		stack.resize(stack.size)
	}
}
