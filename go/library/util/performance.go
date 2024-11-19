package util

import "time"

func Sleep(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}
