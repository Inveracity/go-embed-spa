package main

import (
	"fmt"
	"log"

	cmd "github.com/inveracity/go-embed-spa/internal/cli"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	cli := cmd.New()
	err := cli.Run()

	if err != nil {
		errWithStack := errors.Cause(err).(stackTracer)
		st := errWithStack.StackTrace()
		fmt.Printf("%+v\n", st[0:])
		log.Fatalf("%s", err.Error())
	}
}
