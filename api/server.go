package api

import (
	db "backend-intern/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	query *db.Queries
	route *gin.Engine
}

func NewServer(query *db.Queries) *Server {
	server := &Server{
		query: query,
		route: gin.Default(),
	}
	// Add routes here
	server.route.POST("/ads", server.CreateAds)
	server.route.GET("/ads", server.ListAds)

	return server
}

// errorResponse is a helper function to generate error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Run starts the server
func (server *Server) Start(address string) error {
	return server.route.Run(address)
}
