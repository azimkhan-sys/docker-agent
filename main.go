# This is the Guides to Downlaod the packages

#Packages that i Have comment I used For Testing

package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/docker/cli/cli"

	"github.com/docker/docker-agent/cmd/root"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	if err := root.Execute(ctx, os.Stdin, os.Stdout, os.Stderr, os.Args[1:]...); err != nil {
		cancel()
		if statusErr, ok := errors.AsType[cli.StatusError](err); ok {
			os.Exit(statusErr.StatusCode)
		}
		os.Exit(1)
	} else {
		cancel()
		os.Exit(0)
	}
}
