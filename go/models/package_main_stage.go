package models

const PackageStage = `package main

import (
	"time"

	"{{PkgPathRoot}}/models"
)

var __Dummy_time_variable time.Time

func Unmarshall(stage *models.StageStruct) {

}
`
