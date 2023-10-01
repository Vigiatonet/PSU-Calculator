package main

import (
	"fmt"

	"github.com/Vigiatonet/PSU-Calculator/src/config"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println(cfg)
}
