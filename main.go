package main

import (
	"fmt"

	"gopyvenv/virtenv"
)

func main() {
	fmt.Print(virtenv.GetCommand())
}
