package main

import (
	"feedFetcher/logger"
	"feedFetcher/rest"
)

func main() {
	logger.SetUpLogger()

	rest.InitServer()
}
