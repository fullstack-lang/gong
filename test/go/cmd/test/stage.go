package main

import "github.com/fullstack-lang/gong/test/go/models"

var Stage models.StageStruct



func Unmarshall(stage *models.StageStruct) {

	// map of identifiers
	var __Astruct__000000_A1 *models.Astruct
	var __Astruct__000001_a2 *models.Astruct

	var __Bstruct__000002_B2 *models.Bstruct
	var __Bstruct__000003_B3_with_a_very_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_name *models.Bstruct
	var __Bstruct__000004_B1 *models.Bstruct

	// initializers of values

	//Init Astruct A1
	__Astruct__000000_A1.Name = "A1"
	__Astruct__000000_A1.Aenum = "ENUM_VAL2"

	//Init Astruct a2
	__Astruct__000001_a2.Name = "a2"
	__Astruct__000001_a2.Aenum = ""

	__Bstruct__000002_B2.Name = "B2"
	__Bstruct__000003_B3_with_a_very_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_name.Name = "B3 with a very long long long long long long long long long long long long long long long long long long long long long long long long long long long long long long name"
	__Bstruct__000004_B1.Name = "B1"

	// initializers of pointers{{PointersInitializers}}

	return
}

