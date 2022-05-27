package hello

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func Hello() {
	go func() {
		log.Info("Hello ~, this is world !")
		time.Sleep(1 * time.Second)
	}()
	log.Info("Hi~")
}
