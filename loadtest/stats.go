// webserver for the loadtest that shows the state of the loadtest
// this is all kept in memory so likely won't scale well lol
package main

import (
	"github.com/labstack/echo/v4"
)

type Stats struct {
	TotalSuccessfulTransactions int
	TotalFailedTransactions     int
	TotalNodeCount              int
	TransactionIntervalMillis   int
	ParallelRequests            int
	RPS                         int
	NodeTotalAttempted          map[string]int
	NodeTotalSuccesses          map[string]int
	NodeTotalFailures           map[string]int
	updateChan                  chan StatsUpdate
}

type StatsUpdate struct {
	Node     string
	WasError bool
}

func NewStats() *Stats {
	stats := &Stats{
		TotalSuccessfulTransactions: 0,
		TotalNodeCount:              len(discprovUrls),
		TransactionIntervalMillis:   interval,
		ParallelRequests:            parallelRequests,
		RPS:                         calculateRPS(),
		NodeTotalAttempted:          map[string]int{},
		NodeTotalSuccesses:          map[string]int{},
		NodeTotalFailures:           map[string]int{},
		updateChan:                  make(chan StatsUpdate, statsBuffer),
	}
	go stats.runUpdater()
	return stats
}

func calculateRPS() int {
	intervalSeconds := float64(interval) / 1000
	return int(float64(parallelRequests) / intervalSeconds)
}

// run the update listener
func (s *Stats) runUpdater() {
	for update := range s.updateChan {
		if update.WasError {
			s.NodeTotalAttempted[update.Node]++
			s.NodeTotalFailures[update.Node]++
			s.TotalFailedTransactions++
		} else {
			s.NodeTotalAttempted[update.Node]++
			s.NodeTotalSuccesses[update.Node]++
			s.TotalSuccessfulTransactions++
		}
	}
}

// send update down channel
func (s *Stats) recordStat(node string, wasError bool) {
	s.updateChan <- StatsUpdate{Node: node, WasError: wasError}
}

// route that serves the stats ui
func (s *Stats) statsHandler(c echo.Context) error {
	return stats_templ.ExecuteTemplate(c.Response().Writer, "stats", s)
}
