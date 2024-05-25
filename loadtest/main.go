package main

import (
	"github.com/alecsavvy/clockwise/utils"
)

func main() {
	logger := utils.NewLogger(nil)
	logger.Info("MOSH PIT")
	select {}
}
