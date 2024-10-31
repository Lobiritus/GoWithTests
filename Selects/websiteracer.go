package selects

import (
	"net/http"
	"time"
)

type WebSiteRacerError string

func (ws WebSiteRacerError) Error() string {
	return string(ws)
}

const (
	ErrTimeLimitException = WebSiteRacerError("time limit exception")
	tenSecondTimeout      = time.Duration(10 * time.Second)
)

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-Ping(a):
		return a, nil
	case <-Ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", ErrTimeLimitException
	}
}

func Ping(url string) chan struct{} {
	result := make(chan struct{})
	go func() {
		http.Get(url)
		close(result)
	}()
	return result
}
