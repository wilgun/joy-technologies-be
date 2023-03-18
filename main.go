package main

import (
	"github.com/wilgun/joy-technologies-be/cmd/webservice"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGSTOP)

	go func() {
		webservice.Start()
	}()

	for {
		<-c
		log.Println("terminating service...")
		os.Exit(0)
	}

}
