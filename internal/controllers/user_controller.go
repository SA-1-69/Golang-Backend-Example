package controllers

import (
    "errors"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "gorm.io/gorm"

    "github.com/SA/Golong-Backend-Example/internal/dto"
    "github.com/SA/Golong-Backend-Example/internal/models"
    "github.com/SA/Golong-Backend-Example/internal/utils"
)

// UserController manages user endpoints.
type UserController struct {
    db       *gorm.DB
    validate *validator.Validate
}

// NewUserController creates a new UserController.
func NewUserController(db *gorm.DB) *UserController {
    return &UserController{
        db:       db,
        validate: validator.New(),
    }
}

// GetProfile returns the current user's profile.
// @Summary Get current user profile
// @Description Returns authenticated user profile
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.UserResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/profile [get]
func (h *UserController) GetProfile(c *gin.Context) {
    userID, ok := utils.GetUserIDFromContext(c)
    if !ok {
        utils.JSONError(c, http.StatusUnauthorized, "authorization required", "user id missing from token")
        return
    }

    user, err := h.findByID(userID)
    if err != nil {
        utils.JSONError(c, http.StatusNotFound, "profile not found", err.Error())
        return
    }
    if user == nil {
        utils.JSONError(c, http.StatusNotFound, "profile not found", "user not found")
        return
    }

    utils.JSONSuccess(c, http.StatusOK, mapUserToResponse(user))
}

// UpdateProfile updates the current user's profile.
// @Summary Update current user profile
// @Description Update authenticated user profile information
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.UpdateUserRequest true "Update User Payload"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /users/profile [put]
func (h *UserController) UpdateProfile(c *gin.Context) {
    userID, ok := utils.GetUserIDFromContext(c)
    if !ok {
        utils.JSONError(c, http.StatusUnauthorized, "authorization required", "user id missing from token")
        return
    }

    var payload dto.UpdateUserRequest
    if err := c.ShouldBindJSON(&payload); err != nil {
        utils.JSONError(c, http.StatusBadRequest, "invalid request payload", err.Error())
        return
    }

    if err := h.validate.Struct(payload); err != nil {
        utils.JSONError(c, http.StatusBadRequest, "validation error", err.Error())
        return
    }

    user, err := h.findByID(userID)
    if err != nil {
        utils.JSONError(c, http.StatusBadRequest, "update failed", err.Error())
        return
    }
    if user == nil {
        utils.JSONError(c, http.StatusBadRequest, "update failed", "user not found")
        return
    }

    if payload.Email != "" && payload.Email != user.Email {
        existing, err := h.findByEmail(payload.Email)
        if err != nil {
            utils.JSONError(c, http.StatusBadRequest, "update failed", err.Error())
            return
        }
        if existing != nil {
            utils.JSONError(c, http.StatusBadRequest, "update failed", "email already in use")
            return
        }
        user.Email = payload.Email
    }

    if payload.Name != "" {
        user.Name = payload.Name
    }
    if payload.Password != "" {
        hashedPassword, err := utils.HashPassword(payload.Password)
        if err != nil {
            utils.JSONError(c, http.StatusBadRequest, "update failed", err.Error())
            return
        }
        user.Password = hashedPassword
    }

    user.UpdatedAt = time.Now().UTC()
    if err := h.db.Save(user).Error; err != nil {
        utils.JSONError(c, http.StatusBadRequest, "update failed", err.Error())
        return
    }

    utils.JSONSuccess(c, http.StatusOK, mapUserToResponse(user))
}

// DeleteUser deletes the current user account.
// @Summary Delete current user
// @Description Delete authenticated user account
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /users/profile [delete]
func (h *UserController) DeleteUser(c *gin.Context) {
    userID, ok := utils.GetUserIDFromContext(c)
    if !ok {
        utils.JSONError(c, http.StatusUnauthorized, "authorization required", "user id missing from token")
        return
    }

    user, err := h.findByID(userID)
    if err != nil {
        utils.JSONError(c, http.StatusBadRequest, "delete failed", err.Error())
        return
    }
    if user == nil {
        utils.JSONError(c, http.StatusBadRequest, "delete failed", "user not found")
        return
    }

    if err := h.db.Delete(user).Error; err != nil {
        utils.JSONError(c, http.StatusBadRequest, "delete failed", err.Error())
        return
    }

    utils.JSONSuccess(c, http.StatusNoContent, nil)
}

// GetAllUsers returns a list of all users.
// @Summary Get all users
// @Description Returns all registered users
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {array} dto.UserResponse
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func (h *UserController) GetAllUsers(c *gin.Context) {
    var users []models.User
    if err := h.db.Find(&users).Error; err != nil {
        utils.JSONError(c, http.StatusInternalServerError, "failed to load users", err.Error())
        return
    }

    responses := make([]dto.UserResponse, 0, len(users))
    for _, user := range users {
        responses = append(responses, *mapUserToResponse(&user))
    }

    utils.JSONSuccess(c, http.StatusOK, responses)
}

// GetUserByID returns a user by ID.
// @Summary Get user by ID
// @Description Returns a user profile identified by the path ID
// @Tags users
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [get]
func (h *UserController) GetUserByID(c *gin.Context) {
    idParam := c.Param("id")
    if idParam == "" {
        utils.JSONError(c, http.StatusBadRequest, "invalid user id", "path parameter id is required")
        return
    }

    userID, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        utils.JSONError(c, http.StatusBadRequest, "invalid user id", "id must be a number")
        return
    }

    user, err := h.findByID(uint(userID))
    if err != nil {
        utils.JSONError(c, http.StatusBadRequest, "failed to load user", err.Error())
        return
    }
    if user == nil {
        utils.JSONError(c, http.StatusNotFound, "user not found", "no user exists with the given id")
        return
    }

    utils.JSONSuccess(c, http.StatusOK, mapUserToResponse(user))
}

func (h *UserController) findByID(id uint) (*models.User, error) {
    var user models.User
    err := h.db.First(&user, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

func (h *UserController) findByEmail(email string) (*models.User, error) {
    var user models.User
    err := h.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

func mapUserToResponse(user *models.User) *dto.UserResponse {
    return &dto.UserResponse{
        ID:        user.ID,
        Name:      user.Name,
        Email:     user.Email,
        CreatedAt: user.CreatedAt.Format(time.RFC3339),
        UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
    }
}
