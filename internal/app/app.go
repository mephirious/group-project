package app

import (
	"github.com/mephirious/group-project/internal/adapters/http"
)

const (
	addr = ":8080"
)

type App struct {
	SimpleServer *http.Server
}

func New(server *http.Server) App {
	return App{
		SimpleServer: server,
	}
}
