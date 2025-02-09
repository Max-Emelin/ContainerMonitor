package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	serverPort := os.Getenv("SERVER_PORT")
	backendURL := fmt.Sprintf("http://app:%s/api/containers/ping-result", serverPort)
	resp, err := http.Post(backendURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Error sending data to API %s", err)
		return
	}
	defer resp.Body.Close()

	logrus.Debugf("Result sent for Ip Address = %s.", result.IPAddress)
}
