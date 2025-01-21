package api

import (
	"time"

	"github.com/Avon11/ShrinkRay/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Handler struct {
	Service *service.ShortCodeService
}

func NewHandler(service *service.ShortCodeService) *Handler {
	return &Handler{
		Service: service,
	}
}

func SetupAPIHandler(rdb *redis.Client) (*gin.Engine, error) {
	// Initialize the Gin router
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))

	service := service.NewCodeService(rdb)

	handler := NewHandler(service)

	r.GET("/api/healthz", handler.HealthCheck)
	r.GET("/api/get-url", handler.GetUrl)
	r.POST("/api/post-url", handler.PostUrl)

	return r, nil
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
