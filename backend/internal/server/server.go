package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/markbates/goth/gothic"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"

	"sample/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

// Google OAuth 認証開始のハンドラー
func (s *Server) googleAuthHandler(c *gin.Context) {
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", "google"))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}
