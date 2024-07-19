package main

import (
	"embed"
	"fmt"
	"io/fs"
	"math/rand/v2"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/alecsavvy/clockwise/sdk"
	"github.com/alecsavvy/clockwise/utils"
	"github.com/bxcodec/faker/v3"
	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

//go:embed templates/*
var embeddedFiles embed.FS
var stats_templ *template.Template
var block_templ *template.Template
var health_templ *template.Template

// config
var interval = 250
var parallelRequests = 10
var statsBuffer = 100000

func init() {
	var err error
	stats_templ, err = template.ParseFS(embeddedFiles, "templates/stats_template.html")
	if err != nil {
		panic(err)
	}

	block_templ, err = template.ParseFS(embeddedFiles, "templates/block_template.html")
	if err != nil {
		panic(err)
	}

	health_templ, err = template.ParseFS(embeddedFiles, "templates/health_template.html")
	if err != nil {
		panic(err)
	}
}

func main() {
	logger := utils.NewLogger(nil)

	var wg sync.WaitGroup
	wg.Add(2)

	stats := NewStats()

	go func() {
		defer wg.Done()
		for {
			start := time.Now()
		
			var g errgroup.Group
			for i := 0; i < parallelRequests; i++ {
				g.Go(func() error {
					return sendRandomRequest(logger, stats)
				})
			}
	
			if err := g.Wait(); err != nil {
				logger.Error("Error in request:", err)
			}
	
			elapsed := time.Since(start)
			sleepDuration := time.Duration(interval)*time.Millisecond - elapsed
	
			if sleepDuration > 0 {
				time.Sleep(sleepDuration)
			}
		}
	}()

	go func() {
		defer wg.Done()

		htmlFS, err := fs.Sub(embeddedFiles, "templates")
		if err != nil {
			logger.Error("could not read embedded template files", "error", err)
			return
		}
		htmlTemplates := echo.WrapHandler(http.FileServerFS(htmlFS))

		e := echo.New()
		e.HideBanner = true
		e.GET("/stats", stats.statsHandler)
		e.GET("/health_stats", getHealthStats)
		e.POST("/get_block", getBlock)
		e.GET("/", htmlTemplates)

		err = e.Start("0.0.0.0:8080")
		if err != nil {
			logger.Error("server crashed", "error", err)
			return
		}
	}()

	wg.Wait()
}

func sendRandomRequest(_ *utils.Logger, stats *Stats) error {
	node := randomDiscprov()
	sdk := sdk.NewSdk(fmt.Sprintf("%s/query", node))

	account, _ := generateWallet()

	handle := faker.Username()
	address := account.Address.Hex()
	bio := faker.Sentence()

	_, err := sdk.CreateUser(
		handle,
		address,
		bio,
	)

	wasError := err != nil
	stats.recordStat(node, wasError)

	if wasError {
		return err
	}
	return nil
}

var discprovUrls = []string{"http://node-0:26659", "http://node-1:26659", "http://node-2:26659", "http://node-3:26659", "http://node-4:26659", "http://node-5:26659", "http://node-6:26659"}

func randomDiscprov() string {
	randomIndex := rand.IntN(len(discprovUrls))
	return discprovUrls[randomIndex]
}
