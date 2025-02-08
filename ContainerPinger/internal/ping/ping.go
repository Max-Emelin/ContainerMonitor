package ping

import (
	"fmt"
	"net"
	"time"

	"github.com/go-ping/ping"
)

func Ping1(ip string) (string, error) {
	start := time.Now()
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:80", ip), 2*time.Second)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds()), nil
}

func Ping(ip string) (time.Time, error) {
	start := time.Now()
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to create pinger for %s: %w", ip, err)
	}

	pinger.Count = 4                 // Количество ICMP-запросов
	pinger.Timeout = 5 * time.Second // Таймаут для всего пинга
	pinger.SetPrivileged(true)       // Для работы без sudo

	err = pinger.Run()
	if err != nil {
		return time.Time{}, fmt.Errorf("ping failed for %s: %w", ip, err)
	}

	return start, nil
}
