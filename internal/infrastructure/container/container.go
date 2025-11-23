package container

import (
	"simple-go-api/internal/infrastructure/config"
	"simple-go-api/internal/infrastructure/router"
	application "simple-go-api/internal/users/application/use_cases"
	"simple-go-api/internal/users/domain"
	"simple-go-api/internal/users/infrastructure/handlers"
	"simple-go-api/internal/users/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

// Container holds all application dependencies
type Container struct {
	Config *config.Config

	// Infrastructure
	Engine *gin.Engine
	Router *router.Router

	// Repositories
	UserRepository domain.UserRepository

	// Use Cases
	CreateUserUseCase *application.CreateUserUseCase

	// Handlers
	UserHandler *handlers.UserHttpHandler
}

// New creates and initializes the dependency injection container
func New(cfg *config.Config) (*Container, error) {
	c := &Container{
		Config: cfg,
	}

	if err := c.initInfrastructure(); err != nil {
		return nil, err
	}

	c.initRepositories()
	c.initUseCases()
	c.initHandlers()
	c.initRouter()

	return c, nil
}

// initInfrastructure initializes infrastructure components (DB, HTTP engine, etc.)
func (c *Container) initInfrastructure() error {
	gin.SetMode(c.Config.Server.Mode)
	c.Engine = gin.Default()
	return nil
}

// initRepositories initializes all repositories
func (c *Container) initRepositories() {
	c.UserRepository = repositories.NewInMemoryUserRepository()
}

// initUseCases initializes all use cases with their dependencies
func (c *Container) initUseCases() {
	c.CreateUserUseCase = application.NewCreateUserUseCase(c.UserRepository)
}

// initHandlers initializes all HTTP handlers with their dependencies
func (c *Container) initHandlers() {
	c.UserHandler = handlers.NewUserHttpHandler(c.CreateUserUseCase)
}

// initRouter initializes the router with all handlers
func (c *Container) initRouter() {
	c.Router = router.NewRouter(c.Engine, c.UserHandler)
}
