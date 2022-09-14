package main

import (
	"fmt"
	"github.com/jcastellanos/falcon/core/usecases"
	"github.com/jcastellanos/falcon/infraestructure/adapters"
)

func main() {
	fmt.Println("Running falcon")
	alertCase := usecases.NewAlertCase()
	alertCase.Load()
	alertCase.AddNotifier(adapters.NewTeamsNotifierAdapter())
	monitorReader := adapters.NewCSVMonitorReader("monitors.csv")
	monitorCase := usecases.NewMonitorCase(adapters.NewHttpMonitorAdapter(),
		adapters.NewLocalAlerterAdapter(alertCase), monitorReader)
	monitorCase.Load()
	monitorCase.StartMonitoring()
}
