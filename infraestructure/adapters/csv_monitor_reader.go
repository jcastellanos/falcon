package adapters

import (
	"encoding/csv"
	"github.com/jcastellanos/falcon/core/models"
	"os"
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

func callOrDefault(call func(string) (int, error), param string, def int) int {
	res, err := call(param)
	if err != nil {
		res = def
	}
	return res
}

func (a CSVMonitorReader) Read() ([]models.Monitor, error) {
	records, err := readData(a.filename)
	if err != nil {
		return nil, err
	}
	monitors := make([] models.Monitor, 0)
	for _, record := range records {
		response := callOrDefault(strconv.Atoi, record[3], 200)
		timeout := callOrDefault(strconv.Atoi, record[4], 3000)
		monitor := models.Monitor{
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



func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}