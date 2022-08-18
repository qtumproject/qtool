package server

import (
	"os"

	"github.com/qtumproject/qtool/qtool-api/handlers"
	"github.com/qtumproject/qtool/qtool-api/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Server struct {
	address string
	debug   bool
	echo    *echo.Echo
}

func NewServer(debug bool, address string) (*Server, error) {

	return &Server{
		address: address,
		debug:   debug,
	}, nil
}

func (s *Server) Start() error {

	logger, err := log.NewLogger(
		log.WithDebugLevel(s.debug),
		log.WithWriter(os.Stdout),
	)
	logger.Info("Starting qtool-server...")
	if err != nil {
		return err
	}

	s.echo = echo.New()
	e := s.echo
	// e.Debug = s.debug
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logger.WithFields(logrus.Fields{
				"URI":       values.URI,
				"status":    values.Status,
				"method":    values.Method,
				"remote IP": values.RemoteIP,
			}).Info("request received")

			return nil
		},
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.POST("/privatekey", handlers.PrivateKeyHandler)
	// e.POST("/publickey", handlers.PublicKey)
	e.POST("/address", handlers.AddressHandler)
	e.POST("/script", handlers.ScriptPubKeyHandler)
	e.Debug = true

	err = e.Start(s.address)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	return s.echo.Close()
}
