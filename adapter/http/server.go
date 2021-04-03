package http

import (
	"database/sql"
	"log"

	"github.com/labstack/echo"
)

type (
	server struct {
		port       string
		echoServer *echo.Echo
	}
	// Server interface
	Server interface {
		Serve() error
	}
)

// NewServer はインスタンスを生成します
func NewServer(port string, dbConnection *sql.DB) *server {
	echoServer := NewEchoServer(dbConnection)
	return &server{
		port,
		echoServer,
	}
}

func (s *server) Serve() {
	log.Fatal(s.echoServer.Start(":" + s.port))
}
