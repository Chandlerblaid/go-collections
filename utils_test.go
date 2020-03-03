package go_collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	l := []interface{}{1, 2, 3}
	rev := []interface{}{3, 2, 1}
	Reverse(l)
	if !reflect.DeepEqual(rev, l) {
		t.Errorf("Reverse(), want: %v, got: %v", rev, l)
	}
}

func TestShuffle(t *testing.T) {
	l := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ogLen := len(l)
	Shuffle(l)
	if !reflect.DeepEqual(ogLen, len(l)) {
		t.Errorf("length, want: %v, got: %v", ogLen, len(l))
	}
	fmt.Printf("got: %v", l)
}
