package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"Backend/admin_panel/internal/domain/user"
)

type UserHandler struct{ service user.Service }

func NewUserHandler(service user.Service) *UserHandler { return &UserHandler{service: service} }

func toResponse(user *user.User) gin.H {
	return gin.H{
		"id":        user.ID.Hex(),
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"isActive":  user.IsActive,
		"role":      user.Role,
	}
}

// POST /api/v1/users
func (h *UserHandler) Create(c *gin.Context) {
	var req user.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	user, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to create user", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "user": toResponse(user)})
}

// GET /api/v1/users/:id
func (h *UserHandler) Get(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}
	user, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": toResponse(user)})
}

// PUT /api/v1/users/:id
func (h *UserHandler) Update(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id", "details": err.Error()})
		return
	}
	var req user.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}
	user, err := h.service.Update(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to update user", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully", "user": toResponse(user)})
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id", "details": err.Error()})
		return
	}
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to delete user", "details": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *UserHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	items, total, err := h.service.List(c.Request.Context(), page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users", "details": err.Error()})
		return
	}
	data := make([]gin.H, len(items))
	for _, u := range items {
		data = append(data, toResponse(&u))
	}
	totalPages := (int(total) + size - 1) / size
	c.JSON(http.StatusOK, gin.H{"data": data, "meta": gin.H{
		"page": page, "pageSize": size, "totalItems": total, "totalPages": totalPages,
	}})
}
