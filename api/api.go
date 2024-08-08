package api

import (
	"github.com/VijayVPatil/Helping-Hand/controller"
	"github.com/VijayVPatil/Helping-Hand/store"
)

type ApiRoute struct {
	Server controller.ServerOperations
}

func (api *ApiRoute) StartApp(server controller.DbServer) {

	api.Server = &server
	api.Server.NewServer(store.DbConnect{})
}
