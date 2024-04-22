package server

import "github.com/labstack/echo/v4"

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
	e.engine.Logger.Fatal(e.engine.Start(":8080"))
}

func (e *echoServerImpl) WithRouter(fn func(EchoServer)) {
	fn(e)
}

type HandlerFunc func(Context) error

func (e *echoServerImpl) POST(relativePath string, handlerFunc HandlerFunc) {
	e.engine.POST(relativePath, Convert(handlerFunc))
}
