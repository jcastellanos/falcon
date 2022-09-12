package usecases

import (
	"fmt"
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/ports"
	"time"
)

type MonitorCase struct {
	systemMonitor models.SystemMonitor
	httpMonitor ports.HttpMonitor
	alerter ports.Alerter
}

func NewMonitorCase(httpMonitor ports.HttpMonitor, alerter ports.Alerter) MonitorCase {
	return MonitorCase {
		systemMonitor: models.NewSystemMonitor(),
		httpMonitor: httpMonitor,
		alerter: alerter,
	}
}

func (a *MonitorCase) Load() {
	monitor1 := models.Monitor{
		Id:			  		"1",
		ApplicationId:		"1",
		ApplicationName:  	"Jenkins Calle 46",
		Url:          		"https://localhost/",
		Response:     		200,
		TimeoutMillis:      3000,
	}
	a.systemMonitor.Append(monitor1)
}

func (a *MonitorCase) StartMonitoring() {
	for _ = range time.Tick(time.Minute * 5) {
		for _, monitor := range a.systemMonitor.GetMonitors() {
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
			Id:       			monitor.Id,
			ApplicationId: 		monitor.ApplicationId,
			ApplicationName: 	monitor.ApplicationName,
			Url:      			monitor.Url,
			Subject:  			"Error al consultar la aplicación " + monitor.ApplicationName,
			Message:  			"Se ha presentado un error en el monitor de la aplicación",
			Priority: 			"CRITICAL",
		})
	}

}