package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()
	if want != got {
		t.Errorf("twice %d: want %d, got %d", input, want,
			got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	input := mytypes.MyString("hello")
	want := 5
	got := input.Len()
	if want != got {
		t.Errorf("want length of %s = %d, got %d", input, want, got)
	}
}
func TestMyBuilderHello(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	want := "Hello, Gophers!"
	got := mb.Hello()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
