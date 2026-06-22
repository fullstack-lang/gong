//go:build !js

package main

import (
	"github.com/fullstack-lang/gong/test/test1/go/cmd/test1/commands"
)

func main() {
	commands.Execute()
}
