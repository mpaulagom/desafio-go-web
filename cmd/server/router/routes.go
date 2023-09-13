package router

import (
	handler "github.com/bootcamp-go/desafio-go-web/cmd/server/handlers"
	"github.com/gin-gonic/gin"
)

type TicketRouter struct {
	ginEng  *gin.Engine
	handler *handler.Service
}

// NewRouter creates a new router injecting the corresponding handlers
func NewRouter(g *gin.Engine, handlerS *handler.Service) TicketRouter {
	router := TicketRouter{
		ginEng:  g,
		handler: handlerS,
	}
	return router
}

// MapRoutes creates the routes and maps it to the handler method that implements it
func (tr TicketRouter) MapRoutes() {
	tr.ginEng.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	apiTickets := tr.ginEng.Group("/tickets")
	apiTickets.GET("/getByCountry/:dest", tr.handler.GetTicketsByCountry())
	apiTickets.GET("/getAverage/:dest", tr.handler.AverageDestination())
}
