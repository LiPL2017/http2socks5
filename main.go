package main

import (
	"flag"
	"http2socks5/config"
	"http2socks5/server"
	"log"
)

var (
	configFile = flag.String("config", "config.yaml", "config file path")
	configType = flag.String("type", "yaml", "config type")
)

func main() {
	flag.Parse()
	icfg, err := config.InitConfig(*configFile, *configType)
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	log.Fatal(server.RunServer(icfg))
}
