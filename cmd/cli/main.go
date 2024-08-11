package main

import (
	"log"

	"github.com/fatih/color"
	cmd "github.com/inveracity/go-embed-spa/internal/cli"
)

func main() {
	cli := cmd.New()
	err := cli.Run()
	if !cli.NoColor {
		color.Set(color.FgRed, color.Bold)
		defer color.Unset() // This is probably not needed
	}
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
