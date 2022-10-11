package tickets

import (
	"context"
	"fmt"
	"desafio-goweb-sabrinagarcia/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetCountByDestination(ctx context.Context, destination string) (int, error)
	//GetAverageByDestination(ctx context.Context, destination string) (int, error)
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

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) GetCountByDestination(ctx context.Context, destination string) (int, error) {
	var ticketCount int = 0

	for _, t := range r.db {
		if t.Country == destination {
			ticketCount++
		}
	}

	if ticketCount == 0 {
		return 0, fmt.Errorf("there are no tickets to this destination")
	}

	return ticketCount, nil
}

// func (r *repository) GetAverageByDestination(ctx context.Context, destination string) (int, error) {
// 	var ticketAvg int
// 	var count int
// 	var ticketSum int

// 	for _, t := range r.db {
// 		if t.Country == destination {
// 			count++

// 		}
// 	}
// }