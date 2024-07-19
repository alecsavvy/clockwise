// webserver for the loadtest that shows the state of the loadtest
// this is all kept in memory so likely won't scale well lol
package main

import (
	"time"

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
	TimeElapsed                 int
	startTime                   time.Time
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
		RPS:                         0,
		NodeTotalAttempted:          map[string]int{},
		NodeTotalSuccesses:          map[string]int{},
		NodeTotalFailures:           map[string]int{},
		updateChan:                  make(chan StatsUpdate, statsBuffer),
		startTime:                   time.Now(),
	}
	go stats.runUpdater()
	return stats
}

func (s *Stats) GetRPS() {
	timeElapsed := time.Since(s.startTime).Seconds()
	s.TimeElapsed = int(timeElapsed)
	s.RPS = int(float64(s.TotalSuccessfulTransactions) / timeElapsed)
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
func (s *Stats) recordStat(node string, err error) {
	s.updateChan <- StatsUpdate{Node: node, WasError: err != nil}
}

// route that serves the stats ui
func (s *Stats) statsHandler(c echo.Context) error {
	s.GetRPS()
	return stats_templ.ExecuteTemplate(c.Response().Writer, "stats", s)
}
