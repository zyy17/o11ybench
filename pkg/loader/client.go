package loader

import (
	"net/http"
	"time"
)

type ClientOptions struct {
	TimeoutInMillisecond int
}

func NewClient(opts *ClientOptions) (*http.Client, error) {
	client := http.DefaultClient

	// Override the default settings.
	client.Transport = &http.Transport{
		ResponseHeaderTimeout: time.Duration(opts.TimeoutInMillisecond) * time.Millisecond,
	}

	return client, nil
}
