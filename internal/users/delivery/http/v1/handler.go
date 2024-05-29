package v1

import (
	error2 "github.com/ariocp/go-app/pkg/responce"
	"net/http"

	"github.com/ariocp/go-app/internal/users/entities"
	"github.com/ariocp/go-app/internal/users/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) SignUp(c *gin.Context) {
	var inp entities.User

	if err := c.BindJSON(&inp); err != nil {
		error2.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(inp)
	if err != nil {
		error2.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var inp signInInput

	if err := c.BindJSON(&inp); err != nil {
		error2.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(inp.Username, inp.Password)
	if err != nil {
		error2.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
