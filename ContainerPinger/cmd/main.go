package main

import (
	"ContainerPinger/internal/docker"
	"ContainerPinger/internal/ping"
	"ContainerPinger/internal/sender"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(("./.env")); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	pingInterval, err := getPingInterval()
	if err != nil {
		logrus.Fatalf("error getting value of PING_INTERVAL_SEC in env file: %s", err.Error())
	}

	for {
		ips, err := docker.GetContainerIPs()
		if err != nil {
			logrus.Errorf("Error getting IP containers: %s", err.Error())
			time.Sleep(pingInterval)
			continue
		}

		for _, ip := range ips {
			pingTime, err := ping.Ping(ip)
			if err != nil {
				logrus.Errorf("Container %s is not responding: %s", ip, err.Error())
				continue
			}

			sender.SendPingResult(ip, pingTime)
		}

		time.Sleep(pingInterval)
	}
}

func getPingInterval() (time.Duration, error) {
	var value int
	_, err := fmt.Sscanf(os.Getenv("PING_INTERVAL_SEC"), "%d", &value)

	return time.Duration(value) * time.Second, err
}
