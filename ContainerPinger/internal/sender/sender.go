package sender

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type PingResult struct {
	IPAddress   string    `json:"ip_address"`
	PingTime    time.Time `json:"ping_time"`
	LastChecked time.Time `json:"last_checked"`
}

func SendPingResult(ip string, pingTime time.Time) {
	result := PingResult{
		IPAddress:   ip,
		PingTime:    pingTime,
		LastChecked: time.Now(),
	}

	jsonData, _ := json.Marshal(result)
	resp, err := http.Post(os.Getenv("BACKEND_URL"), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Error sending data to API %s", err)
		return
	}
	defer resp.Body.Close()

	logrus.Debugf("Result sent for Ip Address = %s.", result.IPAddress)
}
