package main

import (
	"fmt"

	"github.com/insomnes/gopyvenv/pkg/infra"
	"github.com/insomnes/gopyvenv/pkg/virtenv"
)

func main() {
	if !infra.Enabled {
		fmt.Print("")
		return
	}
	fmt.Print(virtenv.GetCommand(infra.VenvDirs))
}
