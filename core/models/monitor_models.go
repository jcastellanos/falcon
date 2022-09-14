package models

type SystemMonitor struct {
	Monitors	[]Monitor
}

type Monitor struct {
	ApplicationId	string
	ApplicationName	string
	Url 			string
	Response 		int
	TimeoutMillis	int
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
