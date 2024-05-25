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
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:embed templates/*
var embeddedFiles embed.FS
var stats_templ *template.Template
var block_templ *template.Template

// config
var interval = 250
var parallelRequests = 20
var statsBuffer = 100

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
}

// sends random entity manager transations to random nodes
func main() {
	logger := utils.NewLogger(nil)

	var wg sync.WaitGroup
	wg.Add(2)

	stats := NewStats()

	go func() {
		defer wg.Done()
		for {
			for i := 0; i < parallelRequests; i++ {
				go sendRandomRequest(logger, stats)
			}
			time.Sleep(time.Duration(interval) * time.Millisecond)
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

func sendRandomRequest(logger *utils.Logger, stats *Stats) {
	node := randomDiscprov()
	sdk := sdk.NewSdk(fmt.Sprintf("%s/query", node))

	requestId := uuid.NewString()
	userId := randomIntID()
	entityId := randomIntID()
	signer := uuid.NewString()
	entityType := randomEntity()
	action := randomAction()
	metadata := "metadata"

	_, err := sdk.ManageEntity(
		requestId,
		userId,
		signer,
		entityType,
		entityId,
		metadata,
		action,
	)

	wasError := err != nil

	if wasError {
		logger.Error("error sending manage entity", "error", err)
	}
	stats.recordStat(node, wasError)
}

var discprovUrls = []string{"http://node-0:26659", "http://node-1:26659", "http://node-2:26659", "http://node-3:26659", "http://node-4:26659", "http://node-5:26659", "http://node-6:26659"}

func randomDiscprov() string {
	randomIndex := rand.IntN(len(discprovUrls))
	return discprovUrls[randomIndex]
}

var entities = []string{"User", "Track", "Playlist"}
var actions = []string{"Create", "Update", "Repost", "Follow", "Unfollow", "Unrepost", "Delete"}

func randomEntity() string {
	randomIndex := rand.IntN(len(entities))
	return entities[randomIndex]
}

func randomAction() string {
	randomIndex := rand.IntN(len(actions))
	return actions[randomIndex]
}

func randomIntID() int {
	return rand.Int()
}
