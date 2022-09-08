package usecases

import (
	"fmt"
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/ports"
	"time"
)

type AlertCase struct {
	systemMonitor models.SystemMonitor
	httpMonitor ports.HttpMonitor
	notifiers []ports.Notifier
}

func NewAlertCase(httpMonitor ports.HttpMonitor) AlertCase {
	return AlertCase {
		systemMonitor: models.NewSystemMonitor(),
		httpMonitor: httpMonitor,
	}
}

func (a *AlertCase) Load() {
	monitor1 := models.Monitor{
		Url:          "https://localhost/",
		Response:     200,
		Timeout:      3000,
		GuardChannel: "https://grupoasd.webhook.office.com/webhookb2/e5833100-ddee-4ee3-bce5-ec4531bc1242@48de1fb0-71ca-41a5-b236-d3182d042c09/IncomingWebhook/80bb6a8e58d5480f9ed9dd656faa8f77/38628351-d1b8-4bc4-8f53-f83e3aafb46a",
		PrimaryGuard: models.Guard {
			Username: "jcastellanos",
			Phone: "3175338977",
			Email: "juancastellanosm@gmail.com",
		},
	}
	a.systemMonitor.Append(monitor1)
}

func (a *AlertCase) AddNotifier(notifier ports.Notifier) {
	a.notifiers = append(a.notifiers, notifier)
}

func (a *AlertCase) StartMonitoring() {
	for _, monitor := range a.systemMonitor.GetMonitors() {
		a.monitoring(monitor, 0)
	}
}

func (a *AlertCase) monitoring(monitor models.Monitor, retry int) {
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
		for _, notifier := range a.notifiers {
			notifier.Notify(monitor)
		}
	}

}