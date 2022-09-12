package models

type Alert struct {
	Id				string
	ApplicationId	string
	ApplicationName string
	Url 			string
	Subject 		string
	Message			string
	Priority		string
}

type GuardSchedule struct {
	guards []Guard
}

func (a *GuardSchedule) AppendGuard(guard Guard) {
	a.guards = append(a.guards, guard)
}

func (a *GuardSchedule) GetGuards() []Guard {
	return a.guards
}

type Guard struct {
	ApplicationId	string
	Primary			Person
	Secundary		Person
	Shadow			Person
	Leader			Person
	ChannelWebhook 	string
}

type Person struct {
	Username	string
	Phone		string
	Email		string
}
