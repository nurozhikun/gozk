package zutils

import (
	"time"
)

func HoursBeforNow(t time.Time) int {
	diff := time.Now().Unix() - t.Unix()
	d := time.Duration(diff) * time.Second
	return int(d.Hours())
}
