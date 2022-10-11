package tickets

import (
	"desafio-goweb-sabrinagarcia/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAll(c *gin.Context) ([]domain.Ticket, error)
	GetTicketByDestination(c *gin.Context, destination string) ([]domain.Ticket, error)
	GetCountByDestination(c *gin.Context, destination string) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(c *gin.Context) ([]domain.Ticket, error) {
	return s.repository.GetAll(c)
}

func (s *service) GetTicketByDestination(c *gin.Context, destination string) ([]domain.Ticket, error) {
	return s.repository.GetTicketByDestination(c, destination)
}


func (s *service) GetCountByDestination(c *gin.Context, destination string) (int, error) {
	return s.repository.GetCountByDestination(c, destination)
}