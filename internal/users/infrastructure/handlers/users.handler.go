package handlers

import (
	"simple-go-api/internal/infrastructure/middlewares"
	application "simple-go-api/internal/users/application/use_cases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserHttpHandler struct {
	CreateUserUseCase *application.CreateUserUseCase
}

func NewUserHttpHandler(createUserUC *application.CreateUserUseCase) *UserHttpHandler {
	return &UserHttpHandler{
		CreateUserUseCase: createUserUC,
	}
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHttpHandler) CreateUserHandler(c *gin.Context) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var req CreateUserRequest
	uuid := uuid.New()

	if err := c.BindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.CreateUserUseCase.Execute(uuid.String(), req.Name, req.Email)

	if err != nil {
		logger.Error("Failed to create user", zap.Error(err))
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, user)
	logger.Info("User created successfully", zap.String("user", user.ID))
}

func (h *UserHttpHandler) RegisterRoutes(rg *gin.RouterGroup) {
	usersGroup := rg.Group("/users")
	{
		usersGroup.POST("", middlewares.TestMiddleware(), h.CreateUserHandler)
	}
}
