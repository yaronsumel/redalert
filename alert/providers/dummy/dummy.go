package dummy

import (
	"math/rand"
)

var cities = []string{
	"ראשון לציון",
	"בני ברק",
	"פתח תקווה",
	"תל אביב",
	"רמת גן",
	"גבעתיים",
	"הרצליה",
}

type Provider struct {
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) GetAllAlerts() ([]string, error) {
	if rand.Intn(10)%2 == 0 {
		return cities[rand.Intn(len(cities)):], nil
	}
	return nil, nil
}
