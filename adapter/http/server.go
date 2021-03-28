package http

import (
	"log"

	"github.com/jinzhu/gorm"
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
func NewServer(port string, dbConnection *gorm.DB) *server {
	echoServer := NewEchoServer(dbConnection)

	return &server{
		port,
		echoServer,
	}
}

func (s *server) Serve() {
	log.Fatal(s.echoServer.Start(":" + s.port))
}
