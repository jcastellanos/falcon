package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/usecases"
	"log"
	"net/http"
)

type GinHandler struct {
	alertCase usecases.AlertCase
}

func NewGinHandler(alertCase usecases.AlertCase) GinHandler {
	return GinHandler {
		alertCase,
	}
}

func (a *GinHandler) AlertWebhook(c *gin.Context) {
	log.Println("AlertWebhook")
	headerToken := c.Request.Header["Authorization"]
	if headerToken == nil || len(headerToken) == 0 || headerToken[0] != "Bearer token-123456-secret" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid authentication",
		})
		return
	}
	json, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(json))
	a.alertCase.Alert(models.Alert {
		Id:              "1",
		ApplicationId:   "1",
		ApplicationName: "Test",
		Url:             "N/A",
		Subject:         "Alarma - Error en metrica de New Relic",
		Message:         string(json),
		Priority:        "N/A",
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "Trigger alert",
	})
}
