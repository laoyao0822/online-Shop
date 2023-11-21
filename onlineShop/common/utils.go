package common

import (
	"time"
)

func GetUnix() int64 {
	return time.Now().Unix()
}
