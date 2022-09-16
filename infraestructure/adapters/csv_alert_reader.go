package adapters

import (
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/utils"
)

type CSVAlertReader struct {
	personsFilename string
	guardsFilename  string
}

func NewCSVAlertReader(personsFilename string, guardsFilename string) CSVAlertReader {
	return CSVAlertReader {
		personsFilename: personsFilename,
		guardsFilename:  guardsFilename,
	}
}

func (a CSVAlertReader) Read() (models.GuardSchedule, error) {
	guardSchedule :=  models.GuardSchedule{}
	personsRecords, err := utils.CSVReadData(a.personsFilename)
	if err != nil {
		return guardSchedule, err
	}
	persons := make(map[string][]string)
	for _, pRecord := range personsRecords {
		persons[pRecord[0]] = pRecord
	}
	guardsFilename, err := utils.CSVReadData(a.guardsFilename)
	if err != nil {
		return guardSchedule, err
	}
	for _, gRecord := range guardsFilename {
		primaryPersonId := gRecord[1]
		person := models.Person {
			Id:		  persons[primaryPersonId][0],
			Username: persons[primaryPersonId][1],
			Name:     persons[primaryPersonId][2],
			Phone:    persons[primaryPersonId][3],
			Email:    persons[primaryPersonId][4],
		}
		guard := models.Guard {
			ApplicationId: gRecord[0],
			Primary:       	person,
			ChannelWebhook: gRecord[2],
		}
		guardSchedule.AppendGuard(guard)
	}
	return guardSchedule, nil
}
