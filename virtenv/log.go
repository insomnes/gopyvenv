package virtenv

import (
	"fmt"
	"os"
)

var debugOn = os.Getenv("GOVENV_DEBUG") == "1"

func debugLog(msg string) {
	if !debugOn {
		return
	}
	fmt.Println(msg)
}
