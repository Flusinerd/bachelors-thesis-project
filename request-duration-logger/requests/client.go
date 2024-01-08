package requests

import (
	"net/http"
	"sync"
	"time"
)

type RequestConfig struct {
	TotalRequestsCount uint
	ParallelCount      uint
	Url                string
	Method             string
	Body               string
}

type RequestTimingResult struct {
	TotalTime  time.Duration
	StartedAt  time.Time
	FinishedAt time.Time
	StatusCode int
}

type requestClient struct {
	config  RequestConfig
	results []RequestTimingResult
}

type RequestClient interface {
	Run() []RequestTimingResult
}

func NewRequestClient(config RequestConfig) RequestClient {
	return &requestClient{
		config:  config,
		results: make([]RequestTimingResult, config.TotalRequestsCount),
	}
}

func (client *requestClient) Run() []RequestTimingResult {
	client.results = client.doRequests()
	return client.results
}

func (client *requestClient) doRequests() []RequestTimingResult {
	semaphore := make(chan struct{}, client.config.ParallelCount)
	var wg sync.WaitGroup
	var results = make([]RequestTimingResult, client.config.TotalRequestsCount)

	for i := uint(0); i < client.config.TotalRequestsCount; i++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(index uint) {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			start := time.Now()
			res := client.makeRequest()
			results[index] = RequestTimingResult{
				TotalTime:  time.Since(start),
				StatusCode: res.StatusCode,
				StartedAt:  start,
				FinishedAt: time.Now(),
			}
		}(i)
	}

	wg.Wait()
	close(semaphore)

	return results
}

func (client *requestClient) makeRequest() *http.Response {
	req, err := http.NewRequest(client.config.Method, client.config.Url, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	return res
}
