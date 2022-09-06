package ports

import "github.com/jcastellanos/falcon/core/models"

type HttpMonitor interface {
	Ping(monitor models.Monitor) (bool, error)
}

type Notifiers interface {
	Notify(monitor models.Monitor) (bool, error)
}
