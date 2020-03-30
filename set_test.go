package go_collections

import (
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {}

func TestSet_Contains(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		el   string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.el); got != tt.want {
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

func TestSet_AddAll(t *testing.T) {
	tests := []struct {
		name string
		s    *Set
		el   []string
	}{
		{
			name: "[]interface{}",
			s:    NewSet([]string{"initial"}),
			el:   []string{"hi", "bye", "ok"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.AddAll(tt.el...)
			for _, e := range tt.el {
				if !tt.s.Contains(e) {
					t.Fatalf("Set should contain %v", e)
				}
			}
		})
	}
}

func TestSet_Merge(t *testing.T) {
	tests := []struct {
		name  string
		s     Set
		e     Set
		final Set
	}{
		{
			name:  "merge",
			s:     Set{"initial": nil},
			e:     Set{"hi": nil, "bye": nil, "ok": nil},
			final: Set{"initial": nil, "hi": nil, "bye": nil, "ok": nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Merge(tt.e)
			if !reflect.DeepEqual(tt.s, tt.final) {
				t.Fatalf("exp: %v, \ngot: %v", tt.final, tt.s)
			}
		})
	}
}
