package runners

import (
	"context"
	"sync"
	"time"
)

type ExternalRoutinesHandler struct {
	once *sync.Once

	enders  []chan bool
	tickers []*time.Ticker
	wg      *sync.WaitGroup
}

func (h *ExternalRoutinesHandler) Shutdown(context context.Context) error {
	for _, ticker := range h.tickers {
		ticker.Stop()
	}

	for _, ender := range h.enders {
		ender <- true
	}

	h.wg.Wait()

	return nil
}

// Run routine wait group for external service grabbing
func ExternalRoutine() *ExternalRoutinesHandler {
	erh := &ExternalRoutinesHandler{
		wg:      &sync.WaitGroup{},
		enders:  []chan bool{},
		tickers: []*time.Ticker{},
	}

	return erh
}
