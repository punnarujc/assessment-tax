package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/punnarujc/assessment-tax/lib/config"
)

type EchoServer interface {
	Start()
	WithRouter(fn func(EchoServer))
	POST(relativePath string, handlerFunc HandlerFunc)
}

type echoServerImpl struct {
	engine echo.Echo
}

func New() EchoServer {
	return &echoServerImpl{
		engine: *echo.New(),
	}
}

func (e *echoServerImpl) Start() {
	port := config.New().GetPort()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := e.engine.Start(port); err != nil && err != http.ErrServerClosed {
			e.engine.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Print("\n\n\n--------- shutting down the server ---------\n\n")

	if err := e.engine.Shutdown(ctx); err != nil {
		e.engine.Logger.Fatal(err)
	}
}

func (e *echoServerImpl) WithRouter(fn func(EchoServer)) {
	fn(e)
}

type HandlerFunc func(Context) error

func (e *echoServerImpl) POST(relativePath string, handlerFunc HandlerFunc) {
	e.engine.POST(relativePath, Convert(handlerFunc))
}
