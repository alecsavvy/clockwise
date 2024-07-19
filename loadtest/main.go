package main

import (
	"embed"
	"io/fs"
	"net/http"
	"sync"
	"text/template"

	"github.com/alecsavvy/clockwise/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

//go:embed templates/*
var embeddedFiles embed.FS
var stats_templ *template.Template
var block_templ *template.Template
var health_templ *template.Template
var error_templ *template.Template

// config
var interval = 250
var parallelRequests = 5
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

	error_templ, err = template.ParseFS(embeddedFiles, "templates/error_template.html")
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
			var g errgroup.Group
			for i := 0; i < parallelRequests; i++ {
				g.Go(func() error {
					return testSequence(stats)
				})
			}
			if err := g.Wait(); err != nil {
				logger.Error("something blew up", "error", err)
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
		e.Use(errorMiddleware)

		err = e.Start("0.0.0.0:8080")
		if err != nil {
			logger.Error("server crashed", "error", err)
			return
		}
	}()

	wg.Wait()
}

func errorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			return error_templ.ExecuteTemplate(c.Response().Writer, "error", map[string]interface{}{
				"Error": err.Error(),
			})
		}
		return nil
	}
}
