package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/KeihakuOh/Career-Connect/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
	db   *sql.DB
	cfg  *config.Config
}

func New(cfg *config.Config, db *sql.DB) *Server {
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	s := &Server{
		echo: e,
		db:   db,
		cfg:  cfg,
	}

	// ルート設定
	s.setupRoutes()

	return s
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}

func (s *Server) setupRoutes() {
	// ヘルスチェック
	s.echo.GET("/health", s.healthCheck)

	// API情報
	s.echo.GET("/api", s.apiInfo)

	// DB接続チェック
	s.echo.GET("/api/db-check", s.dbCheck)
}

func (s *Server) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func (s *Server) apiInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name":    "LabCareer API",
		"version": "0.0.1",
		"env":     s.cfg.AppEnv,
	})
}

func (s *Server) dbCheck(c echo.Context) error {
	if s.db == nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"status":  "error",
			"message": "Database not connected",
		})
	}

	var result int
	err := s.db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// health_checkテーブルから最新のステータスを取得
	var status string
	var checkedAt time.Time
	err = s.db.QueryRow("SELECT status, checked_at FROM health_check ORDER BY id DESC LIMIT 1").Scan(&status, &checkedAt)

	response := map[string]interface{}{
		"status":  "connected",
		"message": "Database connection successful",
	}

	if err == nil {
		response["last_check"] = map[string]interface{}{
			"status": status,
			"time":   checkedAt.Format(time.RFC3339),
		}
	}

	return c.JSON(http.StatusOK, response)
}
