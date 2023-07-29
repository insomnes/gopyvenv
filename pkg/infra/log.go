package infra

import (
	"fmt"
)

func DebugLog(msg string) {
	if !debugOn {
		return
	}
	fmt.Println(msg)
}
