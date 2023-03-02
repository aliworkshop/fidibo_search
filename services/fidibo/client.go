package fidibo

import (
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	Search(keyword string) (*BookResponse, error)
}

type config struct {
	BaseURL string
	Timeout time.Duration
}

type client struct {
	cl      *http.Client
	baseURL url.URL
	timeout time.Duration
}

func NewFidiboClient(registry *viper.Viper) (Client, error) {
	c, err := readConfig(registry)
	if err != nil {
		return nil, err
	}

	u, errParse := url.Parse(c.BaseURL)
	if errParse != nil {
		return nil, errParse
	}

	if c.Timeout == 0 {
		c.Timeout = 2 * time.Second
	}
	cl := &http.Client{Timeout: c.Timeout}

	return &client{
		cl:      cl,
		baseURL: *u,
		timeout: c.Timeout,
	}, nil
}

func readConfig(registry *viper.Viper) (config, error) {
	c := new(config)
	err := registry.Unmarshal(c)
	if err != nil {
		return *c, err
	}
	return *c, nil
}
