package domain

import "errors"

var (
	ErrEmptyListOfTickets  = errors.New("empty list of tickets")
	ErrDestinationNotFound = errors.New("destination not found")
)

type Ticket struct {
	Id      string
	Name    string
	Email   string
	Country string
	Time    string
	Price   float64
}
