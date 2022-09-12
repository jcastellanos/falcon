package ports

import "github.com/jcastellanos/falcon/core/models"

type Notifier interface {
	Notify(alert models.Alert, guard models.Guard) (bool, error)
}
