package usecases

import (
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/ports"
)

type AlertCase struct {
	schedule models.GuardSchedule
	notifiers []ports.Notifier
}

func NewAlertCase() AlertCase {
	return AlertCase { }
}

func (a *AlertCase) Load() {
	person := models.Person {
		Username: "jcastellanos",
		Phone:    "3175338977",
		Email:    "juancastellanosm@gmail.com",
	}
	guard := models.Guard {
		ApplicationId: 	"1",
		Primary:       	person,
		ChannelWebhook: "https://grupoasd.webhook.office.com/webhookb2/e5833100-ddee-4ee3-bce5-ec4531bc1242@48de1fb0-71ca-41a5-b236-d3182d042c09/IncomingWebhook/80bb6a8e58d5480f9ed9dd656faa8f77/38628351-d1b8-4bc4-8f53-f83e3aafb46a",
	}
	a.schedule.AppendGuard(guard)
}

func (a *AlertCase) AddNotifier(notifier ports.Notifier) {
	a.notifiers = append(a.notifiers, notifier)
}

func (a *AlertCase) Alert(alert models.Alert) {
	for _, guard := range a.schedule.GetGuards() {
		if alert.ApplicationId == guard.ApplicationId {
			for _, notifier := range a.notifiers {
				notifier.Notify(alert, guard)
			}
		}
	}
}