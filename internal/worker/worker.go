package worker

import (
	"sync"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type Worker interface {
	Start() error
}

func StartWorkers(logger *zerolog.Logger, workers []Worker) {
	wg := &sync.WaitGroup{}

	for _, w := range workers {
		wg.Add(1)

		go func(w Worker, wg *sync.WaitGroup) {
			defer wg.Done()

			if err := w.Start(); err != nil {
				logger.Error().Err(err).Msg("failed to start worker")
			}
		}(w, wg)
	}

	wg.Wait()
}
