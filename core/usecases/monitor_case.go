package usecases

import (
	"fmt"
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/ports"
	"log"
	"time"
)

type MonitorCase struct {
	systemMonitor models.SystemMonitor
	httpMonitor ports.HttpMonitor
	alerter ports.Alerter
	monitorReader ports.MonitorReader
}

func NewMonitorCase(httpMonitor ports.HttpMonitor, alerter ports.Alerter, monitorReader ports.MonitorReader) MonitorCase {
	return MonitorCase {
		systemMonitor: models.SystemMonitor{},
		httpMonitor: httpMonitor,
		alerter: alerter,
		monitorReader: monitorReader,
	}
}

func (a *MonitorCase) Load() {
	monitors, err := a.monitorReader.Read()
	if err != nil {
		log.Fatal(err)
	}
	a.systemMonitor.Monitors = monitors
}

func (a *MonitorCase) StartMonitoring() {
	// Is necessary to force execution for the first time because Tick starts X minutes after.
	for _, monitor := range a.systemMonitor.Monitors {
		go a.monitoring(monitor, 0)
	}
	for _ = range time.Tick(time.Minute * 5) {
		for _, monitor := range a.systemMonitor.Monitors {
			go a.monitoring(monitor, 0)
		}
	}
}

func (a *MonitorCase) monitoring(monitor models.Monitor, retry int) {
	if retry < 3 {
		res, err := a.httpMonitor.Ping(monitor)
		if !res {
			fmt.Println(err)
			time.Sleep(10 * time.Second)
			a.monitoring(monitor, retry + 1)
		} else {
			fmt.Println("Ping OK")
		}
	} else {
		fmt.Println("Error despues de los retry")
		a.alerter.ThrowAlert(models.MonitorAlert{
			ApplicationId: 		monitor.ApplicationId,
			ApplicationName: 	monitor.ApplicationName,
			Url:      			monitor.Url,
			Subject:  			"Error al consultar la aplicación " + monitor.ApplicationName,
			Message:  			"Se ha presentado un error en la aplicación.",
			Priority: 			"CRITICAL",
		})
	}

}