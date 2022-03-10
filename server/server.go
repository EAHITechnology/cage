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
	logGroup.Get("/test_0", handler.TestHandler)
	logGroup.Post("/test_1", authMiddle, handler.TestHandler)

	// handler
	enet.HttpWeb.Get("/test", handler.TestHandler)

	//========业务初始化区========
	go logic.InitBackend(ctx)

	//===========end============
	go enet.HttpWeb.Run()
	return nil
}
