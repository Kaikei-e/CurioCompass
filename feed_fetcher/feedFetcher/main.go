package main

import (
	"feedFetcher/initialize"
	"feedFetcher/logger"
	"feedFetcher/rest"
)

func main() {
	logger.SetUpLogger()

	initialize.InitDBConnection()

	rest.InitServer()
}
