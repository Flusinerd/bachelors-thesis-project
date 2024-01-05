package utils

import (
	c "api/config"
	"fmt"
	"time"
)

func TimeoutWrapper(originalFunc func() interface{}, duration time.Duration, config *c.Config) (interface{}, error) {
	startTime := time.Now()
	resultChannel := make(chan interface{})
	errChannel := make(chan error)
	timeFactor := int64(config.Server.ResponseDelayFactor)
	durationWithFactor := time.Duration(duration.Nanoseconds() * timeFactor)

	go func() {
		result := originalFunc()
		resultChannel <- result
	}()

	result := <-resultChannel

	if time.Since(startTime) < durationWithFactor {
		time.Sleep(durationWithFactor - time.Since(startTime))
	} else {
		err := fmt.Errorf("timeout after %v", durationWithFactor)
		errChannel <- err
		return nil, err
	}

	return result, nil
}
