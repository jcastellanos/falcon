package models

type SystemMonitor struct {
	monitors	[]Monitor
}

func NewSystemMonitor() SystemMonitor {
	return SystemMonitor {}
}

func (s *SystemMonitor) Append(monitor Monitor) {
	s.monitors = append(s.monitors, monitor)
}

func (s *SystemMonitor) GetMonitors() []Monitor {
	return s.monitors
}

type Monitor struct {
	Url 			string
	Response 		int
	Timeout			int
	Retry			int
	GuardPhones		string
	GuardChannel	string
}
