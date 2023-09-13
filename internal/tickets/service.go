package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

// Service is a service for tickets
type ServiceTickets struct {
	repo Repository
}

func NewService(r Repository) *ServiceTickets {
	return &ServiceTickets{
		repo: r,
	}
}

// GetTotalTickets returns the total of tickets to the city with name in destination variable
func (s *ServiceTickets) GetTotalTickets(ctx context.Context, destination string) (totalTickets int, err error) {
	ticketsR, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return
	}
	totalTickets = len(ticketsR)
	return
}

// AverageDestination calculates the average of tickets to the city with name in destination variable
func (s *ServiceTickets) AverageDestination(ctx context.Context, destination string) (avg float64, err error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return
	}
	allTickets, err := s.repo.GetAll(ctx)
	if err != nil {
		return
	}
	avg = float64(len(tickets)) / float64(len(allTickets))
	return
}
