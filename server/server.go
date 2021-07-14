package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// EchoAPI struct
type EchoAPI struct {
	e *echo.Echo
}

// NewServer Instance of Echo
func NewServer() *EchoAPI {

	return &EchoAPI{
		e: echo.New(),
	}
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Start server functionality
func (s *EchoAPI) Start(port string) {
	// logger
	s.e.Use(middleware.Logger())
	// recover
	s.e.Use(middleware.Recover())
	//CORS
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// price endpoint
	// s.e.GET("/price", controllers.GetPrice)
	// Unauthenticated route
	s.e.GET("/", accessible)
	// Start Server
	s.e.Logger.Fatal(s.e.Start(port))
}

// Close server functionality
func (s *EchoAPI) Close() {
	s.e.Close()
}
