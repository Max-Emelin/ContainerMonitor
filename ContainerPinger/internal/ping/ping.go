package ping

import (
	"fmt"
	"net"
	"time"
)

func Ping(ip string) (string, error) {
	start := time.Now()
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:80", ip), 2*time.Second)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%dms", time.Since(start).Milliseconds()), nil
}
