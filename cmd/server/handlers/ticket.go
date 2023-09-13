package handler

import (
	"errors"
	"net/http"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

type ResponseBody struct {
	Message string      `json:"Message"`
	Result  interface{} `json:"Result"`
}

// GetTicketsByCountry returns the tickets with destination in path param "dest"
func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrDestinationNotFound):
				body := ResponseBody{Message: err.Error(), Result: 0}
				c.JSON(http.StatusNotFound, body)
			default:
				body := ResponseBody{Message: "internal server error", Result: 0}
				c.JSON(http.StatusInternalServerError, body)
			}
			return
		}
		body := ResponseBody{Message: "total of tickets", Result: tickets}
		c.JSON(200, body)
	}
}

// AverageDestination returns the average people with destination in path param "dest"
func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrDestinationNotFound):
				body := ResponseBody{Message: err.Error(), Result: 0}
				c.JSON(http.StatusNotFound, body)
			default:
				body := ResponseBody{Message: "internal server error", Result: 0}
				c.JSON(http.StatusInternalServerError, body)
			}
			return
		}
		body := ResponseBody{Message: "average of tickets", Result: avg}
		c.JSON(http.StatusOK, body)
	}
}
