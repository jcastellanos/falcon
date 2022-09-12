package adapters

import (
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/usecases"
)

type LocalAlerterAdapter struct {
	alertCase	usecases.AlertCase
}

func NewLocalAlerterAdapter(alertCase usecases.AlertCase) LocalAlerterAdapter {
	return LocalAlerterAdapter{
		alertCase: alertCase,
	}
}

func (a LocalAlerterAdapter) ThrowAlert(monitorAlert models.MonitorAlert) (bool, error) {
	alert := models.Alert{
		Id:              monitorAlert.Id,
		ApplicationId:   monitorAlert.ApplicationId,
		ApplicationName: monitorAlert.ApplicationName,
		Url:             monitorAlert.Url,
		Subject:         monitorAlert.Subject,
		Message:         monitorAlert.Message,
		Priority:        monitorAlert.Priority,
	}
	a.alertCase.Alert(alert)
	return true, nil
}
