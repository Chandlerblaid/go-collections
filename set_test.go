package go_collections

import (
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {

}

func TestSet_Contains(t *testing.T) {
	type args struct {
		el interface{}
	}
	tests := []struct {
		name string
		s    Set
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.el); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Empty(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args struct {
		el interface{}
	}
	tests := []struct {
		name string
		s    Set
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSet_Size(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ToArray(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
