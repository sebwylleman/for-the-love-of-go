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

func TestMyBuilder(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := mb.Contents.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := mb.Contents.Len()
	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d", mb.Contents.String(), wantLen, gotLen)
	}
}

func TestStringUppercaser(t *testing.T) {
	t.Parallel()
	var su mytypes.StringUppercaser
	su.Contents.WriteString("hello")
	want := "HELLO"
	got := su.ToUpper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
