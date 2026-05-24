//go:build !wasm

package main

import (
	"flag"
	"log"
	"strconv"
)

func main() {
	// 1. Parse arguments for the OS environment
	flag.Parse()

	// 2. Initialize the Gong app and get the Gin router
	r, _ := setupApp()

	// 3. Start the standard TCP server
	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
