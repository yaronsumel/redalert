package alert

import (
	"strings"
	"time"

	"github.com/yaronsumel/redalert/alert/models"
	"github.com/yaronsumel/redalert/alert/providers/dummy"
	"github.com/yaronsumel/redalert/alert/providers/pakar"
)

type Provider interface {
	GetAllAlerts() ([]string, error)
}

func Watch(providerName, areaFilter string, alertChan chan models.Alert, d time.Duration) {

	var defaultProvider Provider

	switch providerName {
	case "dummy":
		defaultProvider = dummy.NewProvider()
	case "pakar":
		fallthrough
	default:
		defaultProvider = pakar.NewProvider()
	}

	for {

		time.Sleep(d)

		alerts, err := defaultProvider.GetAllAlerts()
		if err != nil || len(alerts) == 0 {
			continue
		}

		for _, v := range alerts {
			if strings.Contains(v, areaFilter) {
				alertChan <- models.Alert{
					Name: v,
					Time: time.Now(),
				}
			}
		}

	}

}
