package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := "1234567890"
	card, err := creditcard.New(want)
	if err != nil {
		t.Fatal(err)
	}
	got := card.Number()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestNewInvalidReturnsError(t *testing.T) {
	t.Parallel()
	_, err := creditcard.New("")
	if err == nil {
		t.Fatal("want error for invalid card number, got nil")
	}
}
