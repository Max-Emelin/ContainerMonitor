package ping

import (
	"fmt"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func Ping(ip string) (time.Time, error) {
	pinger, err := probing.NewPinger(ip)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to create pinger for %s: %w", ip, err)
	}

	pinger.Count = 1
	pinger.Timeout = 3 * time.Second
	pinger.SetPrivileged(false)

	start := time.Now()
	err = pinger.Run()
	if err != nil {
		return time.Time{}, fmt.Errorf("ping failed for %s: %w", ip, err)
	}

	return start, nil
}
