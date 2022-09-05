package main

import (
	"fmt"
	"github.com/jcastellanos/falcon/core/usecases"
)

func main() {
	fmt.Println("Running falcon")
	loadSystemMonitor := usecases.NewLoadSystemMonitorCase()
	loadSystemMonitor.Load()
	loadSystemMonitor.StartMonitoring()
}
