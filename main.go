package main

import (
	"time"
	"os/signal"
	"syscall"
	"os"
	log "github.com/sirupsen/logrus"
)

// main function
func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02T15:04:05.000-07:00"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go setLogger(c)
	i:=0
	for i<10{
		log.Info("Printing Info ",i)
		log.Debug("Printing Debug",i)
		time.Sleep(5 * time.Second)
		i+=1
	}
}

func setLogger( types chan os.Signal){
	for {
		// Block until a signal is received.
		log.Info(<-types)
		if log.GetLevel()==log.InfoLevel{
			log.SetLevel(log.DebugLevel)
		}else if log.GetLevel()==log.DebugLevel{
			log.SetLevel(log.InfoLevel)
		}
	}
}