package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	// iterating each key, value range of a map
	for _, value := range c {
		result = append(result, value)
	}
	return result
}

func (c Catalog) GetBook(id int) (Book, error) {
	b, ok := c[id]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", id)
	}
	return b, nil
}

func (b Book) NetPriceCents() int {
	discount := b.DiscountPercent * b.PriceCents / 100
	return b.PriceCents - discount
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("%d is an invalid price", price)
	}
	b.PriceCents = price
	return nil
}
