package tickets

import (
	"context"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if r.db == nil || len(r.db) == 0 {
		return []domain.Ticket{}, domain.ErrEmptyListOfTickets
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if r.db == nil || len(r.db) == 0 {
		return []domain.Ticket{}, domain.ErrEmptyListOfTickets
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}
	if len(ticketsDest) == 0 {
		return []domain.Ticket{}, domain.ErrDestinationNotFound
	}
	return ticketsDest, nil
}
