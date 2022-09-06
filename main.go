package main

import (
	"fmt"
	"github.com/jcastellanos/falcon/core/usecases"
	"github.com/jcastellanos/falcon/infraestructure/adapters"
)

func main() {
	fmt.Println("Running falcon")
	httpMonitor := adapters.NewHttpMonitorAdapter()
	loadSystemMonitor := usecases.NewLoadSystemMonitorCase(httpMonitor)
	loadSystemMonitor.Load()
	loadSystemMonitor.StartMonitoring()
}
