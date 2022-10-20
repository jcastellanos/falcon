package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jcastellanos/falcon/core/constants"
	"github.com/jcastellanos/falcon/core/usecases"
	"github.com/jcastellanos/falcon/core/utils"
	"github.com/jcastellanos/falcon/infraestructure/adapters"
	"github.com/jcastellanos/falcon/infraestructure/handlers"
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
	go monitorCase.StartMonitoring()

	r := gin.Default()
	hand := handlers.NewGinHandler(alertCase)
	r.POST("/alert", hand.AlertWebhook)
	r.RunTLS("0.0.0.0:8443", "./certs/certificate.crt", "./certs/private.key")
	//r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 ("localhost:8080")
}

func validateEnviromentVariables() bool {
	zone := utils.GetConfig(constants.AWS_ZONE_KEY)
	flow := utils.GetConfig(constants.AWS_CONNECT_CONTACT_FLOW_ID_KEY)
	instance := utils.GetConfig(constants.AWS_CONNECT_INSTANCE_ID_KEY)
	phone := utils.GetConfig(constants.AWS_CONNECT_SOURCE_PHONE_NUMBER_KEY)
	return zone != "" && flow != "" && instance != "" && phone != ""
}
