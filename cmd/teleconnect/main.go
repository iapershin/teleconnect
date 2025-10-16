package main

import (
	"fmt"

	"github.com/iapershin/teleconnect/internal/cli"
	"github.com/iapershin/teleconnect/internal/log"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.LogError(fmt.Sprintf("Error: %v", err))
	}
}
