// generated code - do not edit
package portrayalcatalogue

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gcobase"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *gcobase.AbstractObject_Type
var _ *gml.DynamicFeatureType

// insertion point for enums declaration

// insertion point for gongstructs declarations

// MD_PortrayalCatalogueReference_Type Named source named complex type "MD_PortrayalCatalogueReference_Type"
// Information identifing the portrayal catalogue used
type MD_PortrayalCatalogueReference_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "portrayalCatalogueCitation" of complex type CI_Citation_PropertyType
	PortrayalCatalogueCitation []*gml.CI_Citation_PropertyType `xml:"portrayalCatalogueCitation,omitempty"`
}

// MD_PortrayalCatalogueReference_PropertyType Named source named complex type "MD_PortrayalCatalogueReference_PropertyType"
type MD_PortrayalCatalogueReference_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_PortrayalCatalogueReference" of complex type MD_PortrayalCatalogueReference_Type
	MD_PortrayalCatalogueReference *MD_PortrayalCatalogueReference_Type `xml:"MD_PortrayalCatalogueReference,omitempty"`
}
