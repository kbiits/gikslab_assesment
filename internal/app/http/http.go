package http

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/kbiits/gikslab_assesment/internal/app"
	"github.com/kbiits/gikslab_assesment/internal/app/config"
	"github.com/kbiits/gikslab_assesment/internal/app/errors"
	"github.com/kbiits/gikslab_assesment/internal/app/handler"
	"github.com/kbiits/gikslab_assesment/internal/app/jwt"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type httpServer struct {
	fiber    *fiber.App
	handlers *handler.Handlers
	jwt      *jwt.Jwt
	config   *config.Config
}

func NewHttpApp(cfg *config.Config, handlers *handler.Handlers, jwt *jwt.Jwt) (app.App, error) {
	return &httpServer{
		fiber: fiber.New(fiber.Config{
			JSONDecoder: func(data []byte, v interface{}) error {
				return json.Unmarshal(data, v)
			},
			JSONEncoder: func(v interface{}) ([]byte, error) {
				return json.Marshal(v)
			},
			ErrorHandler: errors.HttpErrorHandler,
		}),
		config:   cfg,
		handlers: handlers,
		jwt:      jwt,
	}, nil
}

func (h *httpServer) Start() (err error) {
	h.prepareRoutes()

	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// graceful shutdown impl

	go func() {
		err = h.fiber.Listen(fmt.Sprintf(":%v", h.config.Http.Port))
		if err != nil {
			return
		}
	}()

	<-chanSignal
	log.Println("Gracefully shutdown the app")
	err = h.fiber.Shutdown()
	if err != nil {
		log.Printf("error graceful shutdown, err : %v\n", err)
	}

	return
}
