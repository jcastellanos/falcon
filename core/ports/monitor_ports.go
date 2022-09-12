package ports

import "github.com/jcastellanos/falcon/core/models"

type HttpMonitor interface {
	Ping(monitor models.Monitor) (bool, error)
}

type Alerter interface {
	ThrowAlert(monitorAlert models.MonitorAlert) (bool, error)
}
