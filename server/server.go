package server

import (
	"golang.org/x/net/context"

	"github.com/EAHITechnology/cage/handler"
	"github.com/EAHITechnology/cage/logic"
	"github.com/EAHITechnology/raptor/enet"
)

func InitServer(ctx context.Context) error {
	// group handler
	logGroup := enet.HttpWeb.NewGroup("/demo")
	logGroup.Get("/gettest", handler.TestGetHandler)
	logGroup.Post("/posttest", authMiddle, handler.TestPostHandler)

	// handler
	enet.HttpWeb.Get("/get_test", handler.TestGetHandler)

	go enet.HttpWeb.Run()
	return nil
}

func InitBackend(ctx context.Context) {
	go logic.InitBackend(ctx)
}
