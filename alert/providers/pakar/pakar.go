package pakar

import (
	"encoding/json"
	"net/http"
)

const alertsAPI = "https://www.oref.org.il/WarningMessages/alert/alerts.json"

type Provider struct {
	client  *http.Client
	headers map[string]string
}

type Response struct {
	Data  []string `json:"data"`
	Id    int64    `json:"id"`
	Title string   `json:"title"`
}

func NewProvider() *Provider {
	return &Provider{
		client: http.DefaultClient,
		headers: map[string]string{
			"Referer":          "https://www.oref.org.il/12481-he/Pakar.aspx",
			"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
			"X-Requested-With": "XMLHttpRequest",
			"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
		},
	}
}

func (p *Provider) GetAllAlerts() ([]string, error) {

	var response Response

	request, err := http.NewRequest(http.MethodGet, alertsAPI, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range p.headers {
		request.Header.Set(k, v)
	}

	resp, err := p.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
