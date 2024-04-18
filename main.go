package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		app := NewApp("localhost:6000", make([]string, 0))
		app.Run()
	}()

	go func() {
		defer wg.Done()
		app := NewApp("localhost:6001", make([]string, 0))
		app.Run()
	}()

	go func() {
		defer wg.Done()
		app := NewApp("localhost:6002", make([]string, 0))
		app.Run()
	}()

	wg.Wait()
}
