package main

import (
	"fmt"
	"github.com/jcastellanos/falcon/core/constants"
	"github.com/jcastellanos/falcon/core/usecases"
	"github.com/jcastellanos/falcon/core/utils"
	"github.com/jcastellanos/falcon/infraestructure/adapters"
	"log"
)

func main() {
	fmt.Println("Running falcon")
	if !validateEnviromentVariables() {
		log.Fatal("Environment variables doesn't exist.")
	}
	alertCase := usecases.NewAlertCase(adapters.NewCSVAlertReader("persons.csv", "guards.csv"))
	alertCase.Load()
	alertCase.AddNotifier(adapters.NewTeamsNotifierAdapter())
	alertCase.AddNotifier(adapters.NewAmazonConnectNotifierAdapter())
	monitorReader := adapters.NewCSVMonitorReader("monitors.csv")
	monitorCase := usecases.NewMonitorCase(adapters.NewHttpMonitorAdapter(),
		adapters.NewLocalAlerterAdapter(alertCase), monitorReader)
	monitorCase.Load()
	monitorCase.StartMonitoring()
}

func validateEnviromentVariables() bool {
	zone := utils.GetConfig(constants.AWS_ZONE_KEY)
	flow := utils.GetConfig(constants.AWS_CONNECT_CONTACT_FLOW_ID_KEY)
	instance := utils.GetConfig(constants.AWS_CONNECT_INSTANCE_ID_KEY)
	phone := utils.GetConfig(constants.AWS_CONNECT_SOURCE_PHONE_NUMBER_KEY)
	return zone != "" && flow != "" && instance != "" && phone != ""
}
