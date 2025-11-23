package router

import (
	"simple-go-api/internal/users/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine      *gin.Engine
	userHandler *handlers.UserHttpHandler
}

func NewRouter(r *gin.Engine, userHandler *handlers.UserHttpHandler) *Router {
	return &Router{
		engine:      r,
		userHandler: userHandler,
	}
}

func (r *Router) SetupRoutes() {
	// API v1 group
	v1 := r.engine.Group("/api/v1")

	// User routes
	r.userHandler.RegisterRoutes(v1)
}

func (r *Router) Run(port string) error {
	return r.engine.Run(port)
}
