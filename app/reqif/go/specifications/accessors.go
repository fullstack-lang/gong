package specifications

import (
	"log"

	"github.com/fullstack-lang/gong/app/reqif/go/models"
)

func GetSelectedSpecification(stage *models.Stage) (specification *models.SPECIFICATION) {
	for specificationRendering := range *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION_Rendering](stage) {
		if specificationRendering.IsSelected {
			for specification_ := range *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION](stage) {
				if specification_.GetIdentifier() == specificationRendering.GetName() {
					specification = specification_
				}
			}
		}
	}

	return
}

func SetSelectedSpecification(stage *models.Stage, specification *models.SPECIFICATION) {

	for specificationRendering := range *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION_Rendering](stage) {
		specificationRendering.IsSelected = false
		if specificationRendering.GetName() == specification.GetIdentifier() {
			specificationRendering.IsSelected = true
		}
	}
}

func GetSpecificationRendering(
	stage *models.Stage,
	spectObjectType *models.SPECIFICATION,
) (
	specificationRendering *models.SPECIFICATION_Rendering) {

	// SPECIFICATION_Rendering instances Names are the identifiers of the
	// SPECIFICATION instance. Since, by ReqIF design, those identifiers are unique,
	// we can use the gong map of those instances
	map_ := models.GongGetMap[*models.SPECIFICATION_Rendering](stage)

	var ok bool
	if specificationRendering, ok = map_[spectObjectType.GetIdentifier()]; !ok {
		log.Panic("GetSpecificationRendering: Unknown specification", spectObjectType.GetName())
	}

	return
}
