package api

import (
	db "backend-intern/db/sqlc"
	"backend-intern/util"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	query *db.Queries
	route *gin.Engine
	redis *redis.Client
}

func NewServer(query *db.Queries, config util.Config) *Server {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	server := &Server{
		query: query,
		route: gin.Default(),
		redis: rdb,
	}
	// Add routes here
	server.route.POST("/ads", server.CreateAds)
	server.route.GET("/ads", server.ListAds)
	server.route.GET("/ads/random", server.CreateRandomAds)

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
