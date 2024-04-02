package bookstore

import (
	"errors"
	"fmt"
)

// This code project implements a simple bookstore in Go. It defines a Book struct to store information about books and a Catalog which is a map to store them by ID. It includes functions to:

// - Purchase a book (reducing stock)
// - Get a list of all books
// - Get a specific book by ID
// - Calculate the net price of a book after discount
// - Update a book's price (with validation)
// - Update a book's category (with validation)
// - Get a book's category

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Catalog map[int]Book

type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

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

func (b *Book) SetCategory(category Category) error {
	if !validCategory[category] {
		return fmt.Errorf("unknown category: %q", category)
	}
	b.category = Category(category)
	return nil
}

func (b Book) Category() Category {
	return b.category
}
