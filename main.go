package main

import (
	"log"

	"github.com/nightfury1204/demo-macaron/cmds"
)

func main() {
	if err := cmds.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
