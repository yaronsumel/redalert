package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/yaronsumel/redalert/alert"
	"github.com/yaronsumel/redalert/alert/models"
)

func main() {

	var (
		alertsChannel   = make(chan models.Alert)
		osChan          = make(chan os.Signal, 1)
		duration        = time.Second
		defaultProvider = "pakar"
		areaFilter      = flag.String("filter", "", "city filter or none")
	)

	flag.Parse()

	log.Println(fmt.Sprintf("redalert started data-provider:%s duration:%f(Seconds) filter:%s", defaultProvider, duration.Seconds(), *areaFilter))

	signal.Notify(osChan, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	go alert.Watch(defaultProvider, *areaFilter, alertsChannel, duration)

	go func() {
		for {
			select {
			case alert := <-alertsChannel:
				var (
					title = fmt.Sprintf("Red Alert For: %s", alert.Name)
					body  = fmt.Sprintf("%s\n\nAt: %s", title, alert.Time.String())
				)
				beeep.Alert(title, body, "")
			}
		}
	}()

	<-osChan

	log.Println("done")

}
