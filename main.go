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
	monitorCase := usecases.NewMonitorCase(adapters.NewHttpMonitorAdapter(),
		adapters.NewLocalAlerterAdapter(alertCase))
	monitorCase.Load()
	monitorCase.StartMonitoring()
}
