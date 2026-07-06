package main

import (
	"log"

	"github.com/fullstack-lang/gong/test/test3/go/models"
)

func main() {
	stage := models.NewStage("")
	stage.MarshallFile("data/stage.go", "github.com/fullstack-lang/gong/test/test3/go/models", "main")
	log.Println("Done")
}
