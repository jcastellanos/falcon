package usecases

import (
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/ports"
	"log"
)

type AlertCase struct {
	schedule 	models.GuardSchedule
	notifiers 	[]ports.Notifier
	alertReader ports.AlertReader
}

func NewAlertCase(alertReader ports.AlertReader) AlertCase {
	return AlertCase {
		alertReader: alertReader,
	}
}

func (a *AlertCase) Load() {
	schedule, err := a.alertReader.Read()
	if err != nil {
		log.Fatal(err)
	}
	a.schedule = schedule
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