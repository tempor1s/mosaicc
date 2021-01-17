package server

import "github.com/labstack/echo/v4"

// Server represents the internal server structure and its deps
type Server struct {
	Echo *echo.Echo
}

// New will create a new server instance
func New() *Server {
	return &Server{
		Echo: echo.New(),
	}
}

// Start will start the http server
func (s *Server) Start(port string) {
	// register the routes
	s.Routes()

	// start the server
	s.Echo.Logger.Fatal(s.Echo.Start(port))
}
