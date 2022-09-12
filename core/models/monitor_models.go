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
	Id				string
	ApplicationId	string
	ApplicationName	string
	Url 			string
	Response 		int
	TimeoutMillis	int
	Retry			int
}

type MonitorAlert struct {
	Id				string
	ApplicationId	string
	ApplicationName string
	Url 			string
	Subject 		string
	Message			string
	Priority		string
}
