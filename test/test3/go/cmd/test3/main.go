//go:build !js

package main

import (
	"log"

	"github.com/fullstack-lang/gong/test/test3/go/cmd/test3/commands"
)

func main() {
	log.SetPrefix("test3: ")
	log.SetFlags(log.Lmicroseconds)

	commands.Execute()
}
