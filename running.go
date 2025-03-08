package gorunning

import (
	"os"
	"os/signal"
	"syscall"
)

func Loop(Quit, Hup func(), pid *Pid) error {
	if pid == nil {
		pid = NewPid(os.Getpid(), pidFilePath)
	}
	if err := pid.Create(); err != nil {
		return err
	}
	defer pid.Close()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT,
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGTRAP)
loop:
	for {
		switch <-c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if Quit != nil {
				Quit()
			}
			break loop
		case syscall.SIGHUP:
			if Hup != nil {
				Hup()
			}

		}
	}
	return nil
}
