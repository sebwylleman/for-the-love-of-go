package creditcard

import (
	"errors"
)

// This code defines a creditcard package in Go. It creates a card struct to store a credit card's account number

type card struct {
	account_number string
}

func New(account_number string) (card, error) {
	if account_number == "" {
		return card{}, errors.New("Account number must not be empty")
	}
	return card{account_number}, nil
}

func (c *card) Number() string {
	return c.account_number
}
