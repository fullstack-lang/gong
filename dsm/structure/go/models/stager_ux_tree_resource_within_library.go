package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeResourceWithinLibrary(
	library *Library,
	resource *Resource,
	parentNode *tree.Node,
) {
	resourceNode := &tree.Node{
		Name:            resource.GetName(),
		IsExpanded:      slices.Contains(library.ResourcesWhoseNodeIsExpanded, resource),
		IsNodeClickable: true,
		IsInEditMode:    resource.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, resourceNode)

	addRenameButton(resource, resourceNode, stager)

	resourceNode.OnNameChange = stager.onNameChange(resource)
	resourceNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, resource, &library.ResourcesWhoseNodeIsExpanded)
	resourceNode.OnClick = onNodeClicked(stager, resource)

	for _, participant := range stager.rm_Resource_Participants[resource] {
		stager.treeParticipantWithinLibrary(library, participant, resourceNode)
	}
}
