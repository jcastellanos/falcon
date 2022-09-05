package usecases

import (
	"fmt"
	"github.com/jcastellanos/falcon/core/models"
)

type LoadSystemMonitorCase struct {
	systemMonitor models.SystemMonitor
}

func NewLoadSystemMonitorCase() LoadSystemMonitorCase {
	return LoadSystemMonitorCase {
		systemMonitor: models.NewSystemMonitor(),
	}
}

func (a *LoadSystemMonitorCase) Load() {
	monitor1 := models.Monitor{
		Url:          "https://www.google.com/",
		Response:     200,
		Timeout:      3000,
		Retry:        3,
		GuardPhones:  "3175338977",
		GuardChannel: "",
	}
	a.systemMonitor.Append(monitor1)
}

func (a *LoadSystemMonitorCase) StartMonitoring() {
	for _, monitor := range a.systemMonitor.GetMonitors() {
		fmt.Println(monitor.Url)
	}
}