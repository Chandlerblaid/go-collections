package go_collections

import (
	"reflect"
	"testing"
)

func Test_stack_Empty(t *testing.T) {
	y := NewStack(nil)
	y.list = SimpleList{}
	if got := y.Empty(); !got {
		t.Errorf("Empty() = %v, want true", got)
	}

	y.list = []interface{}{1, 2, 3, 4, 5}
	if got := y.Empty(); got {
		t.Errorf("Empty() = %v, want false", got)
	}
}

func Test_stack_Peek(t *testing.T) {
	y := NewStack(nil)
	p := 5
	y.list = []interface{}{1, 2, 3, 4, p}
	if got := y.Peek(); !reflect.DeepEqual(got, p) {
		t.Errorf("Peek() = %v, want %v", got, p)
	} else if y.Size() != 5 {
		t.Errorf("y.Size() = %v, want %v", y.Size(), p)
	}
}

func Test_stack_Peek_empty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Peek didn't panic")
		}
	}()
	y := Stack{}
	_ = y.Peek()
}

func Test_stack_Pop(t *testing.T) {
	y := NewStack(nil)
	p := 5
	y.list = []interface{}{1, 2, 3, 4, p}
	exp := []interface{}{1, 2, 3, 4}
	if got := y.Pop(); !reflect.DeepEqual(got, p) {
		t.Errorf("Pop() = %v, want %v", got, p)
	} else if !reflect.DeepEqual(exp, y.list) {
		t.Errorf("after Pop() = %v, want %v", exp, y)
	} else if y.Size() != 4 {
		t.Errorf("y.Size() = %v, want %v", y.Size(), 4)
	}
}

func Test_stack_Pop_empty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Pop didn't panic")
		}
	}()
	y := Stack{}
	_ = y.Pop()
}

func Test_stack_Put(t *testing.T) {
	y := NewStack(nil)
	p := 5
	y.list = []interface{}{1, 2, 3, 4}
	exp := []interface{}{1, 2, 3, 4, p}
	if y.Put(p); !reflect.DeepEqual(exp, y) {
		t.Errorf("Pop() = %v, want %v", exp, p)
	} else if y.Size() != p {
		t.Errorf("y.Size() = %v, want %v", y.Size(), 4)
	}
}

func Test_stack_Size(t *testing.T) {
	y := NewStack(nil)
	p := 5
	y.list = []interface{}{1, 2, 3, 4, p}
	if got := y.Size(); !reflect.DeepEqual(got, p) {
		t.Errorf("Size() = %v, want %v", got, p)
	}

	y = Stack{}
	if got := y.Size(); got != 0 {
		t.Errorf("Size() = %v, want %v", got, 0)
	}
}
