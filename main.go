package main

import (
	"flag"
	"fmt"
	"github.com/liuhuanjs/MyDB/config"
	"github.com/liuhuanjs/MyDB/printer"
	"github.com/liuhuanjs/MyDB/server"
	"log"
	"os"
)

// Flag name
const (
	nmVersion = "V"
	nmHost    = "host"
	nmPort    = "port"
)

var (
	version = flag.Bool(nmVersion, false, "print version information and exit")
	host    = flag.String(nmHost, "127.0.0.1", "mydb server host")
	port    = flag.String(nmPort, "6666", "mydb server port")
)

var (
	cfg *config.Config
	svr *server.Server
)

func main() {
	flag.Parse()
	if *version {
		fmt.Println(printer.GetMyDBInfo())
		os.Exit(0)
	}

	loadConfig()

	//crete and run server
	createServer()
	runServer()
}

func loadConfig() {
	cfg = config.GetGlobalConfig()

}

func createServer() {
	var err error
	svr, err = server.NewServer(cfg)
	if err != nil {
		log.Fatalf("create new server failed , exit now\n")
		os.Exit(1)
	}
}

func runServer() {
	err := svr.Run()
	if err != nil {
		log.Fatalf("running server failed , exit now\n")
		os.Exit(1)
	}
}
