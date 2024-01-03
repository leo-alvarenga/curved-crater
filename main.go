package main

import (
	"curved-crater/api"
	"curved-crater/repository"
	"curved-crater/utils"
	"fmt"
	"os"
	"strings"
)

const (
    databaseFile         string = "analytics.db"
    eventLogFilename     string = "curved_crater.events.log"
    analyticsLogFilename string = "curved_crater.analytics.log"
)

func main() {
    eventLogger, _ := utils.NewLogger(eventLogFilename, false)
    analyticsLogger, _ := utils.NewLogger(analyticsLogFilename, false)

    eventLogger.Log(utils.Low, "Starting...")
    
    analytics := repository.NewAnalyticsInstance(eventLogger, analyticsLogger)
    analytics.OpenDb(databaseFile, repository.OpenSqliteDb)

    defer analytics.CloseDb()

    addr := ":3000"
    if len(os.Args) > 1 {
        port := os.Args[1]

        if !strings.Contains(port, ":") {
            port = fmt.Sprintf(":%s", port)
        }

        addr = port
    }

    api.Api(addr, analytics)
}
