package main

import (
	"fmt"

	"gopyvenv/virtenv"
)

func main() {
	venvDirs := []string{"venv", ".venv"}
	fmt.Print(virtenv.GetCommand(venvDirs))
}
