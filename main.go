package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/EAHITechnology/cage/server"
	"github.com/EAHITechnology/cage/utils"
	"github.com/EAHITechnology/raptor/config"
	"github.com/EAHITechnology/raptor/elog"
	"github.com/EAHITechnology/raptor/enet"

	"golang.org/x/net/context"
)

var (
	configFilePath = flag.String("config", "./conf/server_config.toml", "config file path")
)

func main() {
	flag.Parse()

	t, err := config.NewConfigParser(config.TomlConfigParserType)
	if err != nil {
		fmt.Println("NewConfigParser err:", err)
		os.Exit(0)
	}

	if err := t.Read(*configFilePath); err != nil {
		fmt.Println("read err:", err)
		os.Exit(0)
	}

	ctx, cancel := context.WithCancel(context.Background())
	if err := t.InitConfig(ctx); err != nil {
		fmt.Println("InitConfig err:", err)
		os.Exit(0)
	}

	if err := t.Unmarshal(&utils.Config); err != nil {
		fmt.Println("Unmarshal err:", err)
		os.Exit(0)
	}

	if err := server.InitServer(ctx); err != nil {
		fmt.Println("InitServer err:", err)
		os.Exit(0)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGPIPE, syscall.SIGUSR1)
	for {
		sig := <-sc
		if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGQUIT {
			elog.Elog.Warnf("get os signal, close server signal: %s", sig.String())
			if enet.HttpWeb != nil {
				enet.HttpWeb.Close()
			}
			cancel()
			break
		} else {
			elog.Elog.Warnf("ignore os signal: %s", sig.String())
		}
	}

}
