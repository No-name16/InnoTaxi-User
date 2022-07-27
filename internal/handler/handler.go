package handler

import (
	"github.com/No-name16/InnoTaxi-User/internal/entity"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(number, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		api.POST("/", h.GetUsersId)
	}

	return router
}
