package main

import (
	"github.com/wilgun/joy-technologies-be/cmd/webservice"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	webservice.Start()
}
