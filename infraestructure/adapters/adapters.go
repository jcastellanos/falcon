package adapters

import (
	"fmt"
	"github.com/jcastellanos/falcon/core/models"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type HttpMonitorAdapter struct {

}

func NewHttpMonitorAdapter() HttpMonitorAdapter {
	return HttpMonitorAdapter {

	}
}

func (a HttpMonitorAdapter) Ping(monitor models.Monitor) (bool, error) {
	fmt.Println("Ping")
	ctx, cancel := context.WithTimeout(context.Background(), 3000 * time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, monitor.Url, nil)
	if err != nil {
		return false, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != monitor.Response {
		return false, fmt.Errorf("status error. Expected[%d] Returned [%d]", monitor.Response, resp.StatusCode)
	}
	return true, nil
}
