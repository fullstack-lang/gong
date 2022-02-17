package main

import "github.com/fullstack-lang/gong/test2/go/models"

func init() {

	InjectionGateway["reference"] = func() {

		A := (&models.Astruct{Name: "toto"}).Stage()
		A.Name = "toto"
	}
}
