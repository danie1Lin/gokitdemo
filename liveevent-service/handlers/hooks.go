package handlers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"trussdemo/liveevent-service/svc"
)

func InterruptHandler(errc chan<- error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	terminateError := fmt.Errorf("%s", <-c)

	// Place whatever shutdown handling you want here

	errc <- terminateError
}

func SetConfig(cfg svc.Config) svc.Config {
	return cfg
}
