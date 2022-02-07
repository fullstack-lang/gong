package main

import "github.com/fullstack-lang/gong/test/go/models"

var Stage models.StageStruct



func Unmarshall(stage *models.StageStruct) {

	// map of identifiers
	map_Astruct_Identifiers := make(map[*models.Astruct]string)
	_ = map_Astruct_Identifiers
	var __Astruct__000000_A1 *models.Astruct
	map_Astruct_Identifiers[__Astruct__000000_A1] = "__Astruct__000000_A1"
	var __Astruct__000001_a2 *models.Astruct
	map_Astruct_Identifiers[__Astruct__000001_a2] = "__Astruct__000001_a2"

	map_Bstruct_Identifiers := make(map[*models.Bstruct]string)
	_ = map_Bstruct_Identifiers
	var __Bstruct__000002_B3_with_a_very_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_name *models.Bstruct
	map_Bstruct_Identifiers[__Bstruct__000002_B3_with_a_very_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_name] = "__Bstruct__000002_B3_with_a_very_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_name"
	var __Bstruct__000003_B1 *models.Bstruct
	map_Bstruct_Identifiers[__Bstruct__000003_B1] = "__Bstruct__000003_B1"
	var __Bstruct__000004_B2 *models.Bstruct
	map_Bstruct_Identifiers[__Bstruct__000004_B2] = "__Bstruct__000004_B2"

	// initializers of values

	//Init Astruct A1
	__Astruct__000000_A1.Name = "A1"
	__Astruct__000000_A1.Aenum = "ENUM_VAL2"
	__Astruct__000000_A1.CFloatfield = 3.400000
	__Astruct__000000_A1.Intfield = 4444

	//Init Astruct a2
	__Astruct__000001_a2.Name = "a2"
	__Astruct__000001_a2.Aenum = ""
	__Astruct__000001_a2.CFloatfield = 0.200000
	__Astruct__000001_a2.Intfield = 0

	__Bstruct__000002_B3_with_a_very_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_long_name.Name = "B3 with a very long long long long long long long long long long long long long long long long long long long long long long long long long long long long long long name"
	__Bstruct__000003_B1.Name = "B1"
	__Bstruct__000004_B2.Name = "B2"

	// initializers of pointers{{PointersInitializers}}

	return
}

