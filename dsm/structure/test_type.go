package main

import (
	"fmt"
	"github.com/fullstack-lang/gong/dsm/structure/go/models"
)

func main() {
	fmt.Println(models.GetPointerToGongstructName[*models.Part]())
}
