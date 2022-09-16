package adapters

import (
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/utils"
	"strconv"
)

type CSVMonitorReader struct {
	filename string
}

func NewCSVMonitorReader(filename string) CSVMonitorReader {
	return CSVMonitorReader {
		filename: filename,
	}
}

func (a CSVMonitorReader) Read() ([]models.Monitor, error) {
	records, err := utils.CSVReadData(a.filename)
	if err != nil {
		return nil, err
	}
	monitors := make([] models.Monitor, 0)
	for _, record := range records {
		response := utils.CallOrDefault(strconv.Atoi, record[3], 200)
		timeout := utils.CallOrDefault(strconv.Atoi, record[4], 3000)
		monitor := models.Monitor {
			ApplicationId:   record[0],
			ApplicationName: record[1],
			Url:             record[2],
			Response:        response,
			TimeoutMillis:   timeout,
		}
		monitors = append(monitors, monitor)
	}
	return monitors, nil
}