package util

import (
	"context"
	"io"
	"net/http"
	"time"
)

func ForwardRequestToAnotherService(method string, requestBody io.ReadCloser, targetURL string) (*http.Response, error) {

	// targetURL := "http://your-target-service-url"

	req, err := http.NewRequest(method, targetURL, requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ForwardRequestToAnotherServiceV2(r *http.Request, targetURL string) (*http.Response, error) {

	timeout := 20 * 60 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		return nil, err
	}

	req.Header = r.Header
	req = req.WithContext(ctx)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
