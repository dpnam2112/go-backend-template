package handlers

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/dpnam2112/go-backend-template/internal/dto"
	"github.com/dpnam2112/go-backend-template/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserHandler struct
type UserHandler struct {
	uowFactory *repositories.UnitOfWorkFactory
	Logger     *slog.Logger
	UserRepo   *repositories.UserRepository
}

// NewUserHandler initializes a new UserHandler
func NewUserHandler(userRepo *repositories.UserRepository, uowFactory *repositories.UnitOfWorkFactory, logger *slog.Logger) *UserHandler {
	return &UserHandler{UserRepo: userRepo, uowFactory: uowFactory, Logger: logger}
}

// GetUser retrieve a specific user object
// @Summary Retrieve user information based on the given ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} dto.APIResponse[dto.UserResponse]
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/users/:id [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		h.Logger.Error("error parsing ID", "err", err, "id", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User's ID must be a valid UUID."})
		return
	}

	user, err := h.UserRepo.GetUserByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	h.Logger.Debug("Successfully retrieved the user", "user_id", user.ID)

	c.JSON(http.StatusOK, dto.APIResponse[dto.UserResponse]{
		Status: http.StatusOK,
		Data: &dto.UserResponse{
			ID:   user.ID,
			Name: user.Name,
		},
	})
}

// CreateUser creates a new user entity
// @Summary Create a new user object
// @Description Create a new user with name and email
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dto.CreateUserRequest true "User Data"
// @Success 200 {object} dto.APIResponse[dto.UserResponse]
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

	// Bind and validate JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txCtx := context.Background()

	// Create a new business transaction.
	uow, err := h.uowFactory.Create(txCtx)
	if err != nil {
		log.Fatalln("failed to create a transaction")
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	defer uow.Rollback(txCtx)

	// Transaction context.
	// Create a new repository which is bound with the unit of work.
	uowUserRepo := h.UserRepo.WithUnitOfWork(uow)
	user, err := uowUserRepo.CreateUser(context.Background(), req.Name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "An error ocurred."})
		return
	}

	// Commit the business transaction.
	uow.Commit(txCtx)

	c.JSON(http.StatusOK, dto.APIResponse[dto.UserResponse]{
		Status: http.StatusOK,
		Data: &dto.UserResponse{
			ID:   user.ID,
			Name: user.Name,
		},
	})
}
