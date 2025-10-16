package main

import (
	"context"
	"fmt"

	"github.com/iapershin/teleconnect/internal/cli"
	"github.com/iapershin/teleconnect/internal/log"
	"github.com/iapershin/teleconnect/internal/telep"
)

func main() {
	ctx := context.Background()

	if err := telep.Version(ctx); err != nil {
		log.LogError(fmt.Sprintf("Error: %v", err))
		return
	}

	if err := cli.Execute(ctx); err != nil {
		log.LogError(fmt.Sprintf("Error: %v", err))
		return
	}
}
