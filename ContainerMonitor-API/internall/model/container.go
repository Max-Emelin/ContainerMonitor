package model

import "time"

type Container struct {
	Id          uint      `json:"id" db:"id"`
	IPAddress   string    `json:"ip_address" db:"ip_address"`
	PingTime    time.Time `json:"ping_time" db:"ping_time"`
	LastChecked time.Time `json:"last_checked" db:"last_checked"`
}
