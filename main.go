package main

import (
	"context"
	"flag"
	"fmt"

	logicServer "github.com/EAHITechnology/cage/server"
	"github.com/EAHITechnology/cage/utils"
	"github.com/EAHITechnology/raptor/config"
	"github.com/EAHITechnology/raptor/elog"
	"github.com/EAHITechnology/raptor/server"
)

var (
	configFilePath = flag.String("config", "./conf/server_config.toml", "config file path")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	s, err := server.NewServer(ctx, "", *configFilePath)
	if err != nil {
		fmt.Println("NewServer err:", err.Error())
		return
	}

	// injection lazy loading function
	InitConfigFunc := func(ctx context.Context, config config.ConfigParser) {
		if err := config.Unmarshal(&utils.Config); err != nil {
			elog.Elog.Errorf("config Unmarshal err:", err.Error())
			return
		}
	}

	InitServerFunc := func(ctx context.Context, config config.ConfigParser) {
		if err := logicServer.InitServer(ctx); err != nil {
			elog.Elog.Errorf("InitServer err:", err.Error())
			return
		}
		elog.Elog.Info("InitServer success")
	}

	InitBackend := func(ctx context.Context, config config.ConfigParser) {
		logicServer.InitBackend(ctx)
	}

	s.AfterInit(InitConfigFunc, InitServerFunc, InitBackend)

	if err := s.Run(ctx, cancel); err != nil {
		fmt.Println("NewServer err:", err.Error())
		return
	}
}
