package common

import (
	"log/slog"
	"sync"
)

type taskFunc func() error

func Await(logger *slog.Logger, tasks ...taskFunc) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(t taskFunc) {
			defer wg.Done()
			err := t()
			if err != nil {
				logger.Error("task crashed:", err)
			}
		}(task)
	}

	wg.Wait()
}
