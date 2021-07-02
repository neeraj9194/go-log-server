package main

import (
    "flag"
    "github.com/neeraj9194/go-log-server/config"
    "github.com/neeraj9194/go-log-server/src"
)

var (
	configFile = flag.String("c", "./config/config.yaml", "Config file path")
	serverPort string
)

func main() {
    flag.Parse()
    conf := config.LoadConfig(*configFile)
    serverPort = conf.Port
    src.RunServer(serverPort)
}
