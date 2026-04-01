package specobjects

import (
	m "github.com/fullstack-lang/gong/dsm/reqif/go/models"
	"github.com/fullstack-lang/gong/dsm/reqif/go/spectypes"
)

func AddIdentifierAndNameToTitle(stager *m.Stager, specObjectType *m.SPEC_OBJECT_TYPE, markDownContent *string, specObject *m.SPEC_OBJECT) {

	specObjectTypeRendering := spectypes.GetSpecObjectTypeRendering(stager.GetStage(), specObjectType)

	if specObjectTypeRendering.ShowIdentifier {
		*markDownContent += specObject.IDENTIFIER
	}

	if specObjectTypeRendering.ShowIdentifier &&
		specObjectTypeRendering.ShowName {
		*markDownContent += " - "
	}

	if specObjectTypeRendering.ShowName {
		*markDownContent += specObject.Name
	}

	if specObjectTypeRendering.ShowIdentifier ||
		specObjectTypeRendering.ShowName {
		*markDownContent += " "
	}
}
