package models

import (
	"cmp"
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibraryBSinDiagram(currentDiagram *Diagram, library *Library, parentNode *tree.Node) {
	subLibraries := library.SubLibraries
	slices.SortFunc(subLibraries, func(a, b *Library) int {
		return cmp.Compare(a.Name, b.Name)
	})

	for _, subLibrary := range subLibraries {
		confNode := TreeNodeAndShapeConfigurationWithoutLink[
			*Library, Library,
			*Library, Library,
			*LibraryShape, LibraryShape,
			*Diagram,
		]{
			diagram:                     currentDiagram,
			parentNode:                  parentNode,
			element:                     subLibrary,
			parentElement:               library,
			elementsWhoseNodeIsExpanded: &currentDiagram.LibrariesWhoseNodeIsExpanded,
			shapes:                      &currentDiagram.Library_Shapes,
			shapesMap:                   currentDiagram.map_Library_LibraryShape,
		}
		libNode := addNodeToTreeWithoutLink(stager, confNode)

		confSubLibraries := ItemAndShapeButtonConfiguration[
			Library, *Library,
			Library, *Library,
			LibraryShape, *LibraryShape,
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Library, *Library,
				Library, *Library,
			]{
				parentNode:                         libNode,
				sliceForNewAddedItem:               &subLibrary.SubLibraries,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
				parentNodeExpansionSliceEncoding:   &currentDiagram.LibrariesWhoseNodeIsExpanded,
				parentElement:                      subLibrary,
			},
			receivingDiagram:      currentDiagram,
			sliceForNewAddedShape: &currentDiagram.Library_Shapes,
		}
		addCreateItemAndShapeButton(stager, confSubLibraries)

		stager.treeLibraryBSinDiagram(currentDiagram, subLibrary, libNode)
	}
}
