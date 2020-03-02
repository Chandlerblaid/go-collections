package go_collections

import (
	"reflect"
	"testing"
)

func Test_simplelist_Empty(t *testing.T) {
	y := SimpleList{}
	if got := y.Empty(); !got {
		t.Errorf("Empty() = %v, want true", got)
	}

	y = SimpleList{1, 2, 3, 4, 5}
	if got := y.Empty(); got {
		t.Errorf("Empty() = %v, want false", got)
	}
}

func Test_simplelist_Peek(t *testing.T) {
	p := 5
	y := SimpleList{1, 2, 3, 4, p}
	if got := y.Peek(); !reflect.DeepEqual(got, p) {
		t.Errorf("Peek() = %v, want %v", got, p)
	} else if y.Size() != 5 {
		t.Errorf("y.Size() = %v, want %v", y.Size(), p)
	}
}

func Test_simplelist_PeekFront(t *testing.T) {
	p := 1
	y := SimpleList{p, 2, 3, 4, 5}
	if got := y.PeekFront(); !reflect.DeepEqual(got, p) {
		t.Errorf("PeekFront() = %v, want %v", got, p)
	} else if y.Size() != 5 {
		t.Errorf("y.Size() = %v, want %v", y.Size(), p)
	}
}

func Test_simplelist_Peek_empty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Peek didn't panic")
		}
	}()
	y := SimpleList{}
	_ = y.Peek()
}

func Test_simplelist_Pop(t *testing.T) {
	p := 5
	y := SimpleList{1, 2, 3, 4, p}
	exp := SimpleList{1, 2, 3, 4}
	if got := y.Pop(); !reflect.DeepEqual(got, p) {
		t.Errorf("Pop() = %v, want %v", got, p)
	} else if !reflect.DeepEqual(exp, y) {
		t.Errorf("after Pop() = %v, want %v", exp, y)
	} else if y.Size() != 4 {
		t.Errorf("y.Size() = %v, want %v", y.Size(), 4)
	}
}

func Test_simplelist_PopFront(t *testing.T) {
	p := 1
	y := SimpleList{p, 2, 3, 4, 5}
	exp := SimpleList{2, 3, 4, 5}
	if got := y.PopFront(); !reflect.DeepEqual(got, p) {
		t.Errorf("PopFront() = %v, want %v", got, p)
	} else if !reflect.DeepEqual(exp, y) {
		t.Errorf("after PopFront() = %v, want %v", exp, y)
	} else if y.Size() != 4 {
		t.Errorf("y.Size() = %v, want %v", y.Size(), 4)
	}
}

func Test_simplelist_Pop_empty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Pop didn't panic")
		}
	}()
	y := SimpleList{}
	_ = y.Pop()
}

func Test_simplelist_Put(t *testing.T) {
	p := 5
	y := SimpleList{1, 2, 3, 4}
	exp := SimpleList{1, 2, 3, 4, p}
	if y.Put(p); !reflect.DeepEqual(exp, y) {
		t.Errorf("Pop() = %v, want %v", exp, p)
	} else if y.Size() != p {
		t.Errorf("y.Size() = %v, want %v", y.Size(), 4)
	}
}

func Test_simplelist_Size(t *testing.T) {
	p := 5
	y := SimpleList{1, 2, 3, 4, p}
	if got := y.Size(); !reflect.DeepEqual(got, p) {
		t.Errorf("Size() = %v, want %v", got, p)
	}

	y = SimpleList{}
	if got := y.Size(); got != 0 {
		t.Errorf("Size() = %v, want %v", got, 0)
	}
}
