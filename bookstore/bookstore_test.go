package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Dune",
		Author: "F. Herbert",
		Copies: 10,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Dune",
		Author: "F. Herbert",
		Copies: 10,
	}
	want := 9
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if got != want {
		t.Errorf("got %d copies, want %d", got, want)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Dune",
		Author: "F. Herbert",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies. got nil")
	}
}
func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go Tools"},
	}
	want := bookstore.Book{
		ID:    2,
		Title: "The Power of Go Tools",
	}
	got, err := catalog.GetBook(2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go Tools"},
	}
	_, err := catalog.GetBook(3)
	if err == nil {
		t.Fatal("Want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "Dune",
		Author:          "Frank Herbert",
		Copies:          521,
		ID:              121423413,
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := b.NetPriceCents()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
