package models

// ensure that for every object that need a rendering configuration
// there is a rendering configuration
//
// In principle, the rendering configuration should be fields
// for OBJECT_TYPE, ATTRIBUTE_DEFINITIONxxx, ...
// however, since those objects are created from the parsing of the ReqIF
// file, there can be persisted in another file for the rendering conf
// only purpose
func (stager *Stager) enforceRenderingConfigurationSemantic() {

	stage := stager.stage

	enforceRendering(stage, stage.SPEC_OBJECT_TYPEs, stage.SPEC_OBJECT_TYPE_Renderings_mapString)
	enforceRendering(stage, stage.SPECIFICATIONs, stage.SPECIFICATION_Renderings_mapString)

	enforceRendering(stage, stage.ATTRIBUTE_DEFINITION_BOOLEANs, stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_mapString)
	enforceRendering(stage, stage.ATTRIBUTE_DEFINITION_DATEs, stage.ATTRIBUTE_DEFINITION_DATE_Renderings_mapString)
	enforceRendering(stage, stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_mapString)
	enforceRendering(stage, stage.ATTRIBUTE_DEFINITION_INTEGERs, stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_mapString)
	enforceRendering(stage, stage.ATTRIBUTE_DEFINITION_REALs, stage.ATTRIBUTE_DEFINITION_REAL_Renderings_mapString)
	enforceRendering(stage, stage.ATTRIBUTE_DEFINITION_STRINGs, stage.ATTRIBUTE_DEFINITION_STRING_Renderings_mapString)
	enforceRendering(stage, stage.ATTRIBUTE_DEFINITION_XHTMLs, stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_mapString)
}

// enforceRendering is a generic function that ensures rendering configurations exist for source objects
// and removes orphaned rendering configurations.
// It is indeed a fascinating syntax! It looks like magic because it blends two
// different concepts (structural types and interfaces) into one constraint line.
// You usually see this syntax (interface{ *T; Constraint }) in Generic Object Mappers (ORMs) and
// Serialization Libraries where the code needs to create a new object and immediately call a method on it (like .Scan() or .UnmarshalJSON()).
func enforceRendering[T Identifiable, R Gongstruct,
	PR interface {
		*R
		GongstructIF
	}](
	stage *Stage,
	sourceMap map[T]struct{},
	renderingCongMap map[string]PR,
) {

	// Create rendering for new instances
	for instance := range sourceMap {
		if _, ok := renderingCongMap[instance.GetIdentifier()]; !ok {
			var r R
			pr := PR(&r)
			pr.SetName(instance.GetIdentifier())
			pr.StageVoid(stage) // because the name is set with the identifier, the map is
		}
	}

	// Remove renderings for deleted instances
	for _, renderingConf := range renderingCongMap {
		found := false
		for instance := range sourceMap {
			if instance.GetIdentifier() == renderingConf.GetName() {
				found = true
				break
			}
		}
		if !found {
			renderingConf.UnstageVoid(stage)
		}
	}
	// stage.ResetMapStrings()
}
