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
		Username: "xxx",
		Phone:    "xxx",
		Email:    "xxxx",
	}
	guard1 := models.Guard {
		ApplicationId: 	"1",
		Primary:       	person,
		ChannelWebhook: "XXXX",
	}
	guard2 := models.Guard {
		ApplicationId: 	"2{",
		Primary:       	person,
		ChannelWebhook: "XXXX",
	}
	a.schedule.AppendGuard(guard1)
	a.schedule.AppendGuard(guard2)
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