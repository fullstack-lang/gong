package main

import (
	"fmt"
	"log"
	"path/filepath"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

func genAngular(modelPkg *gong_models.ModelPkg) {

	{

		directory, err :=
			filepath.Abs(
				filepath.Join(*pkgPath,
					fmt.Sprintf("../../ng/projects/%s/src/lib", modelPkg.Name)))
		gong_models.MatTargetPath = directory
		if err != nil {
			log.Panic("Problem with frontend target path " + err.Error())
		}
	}
}
