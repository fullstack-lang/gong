package spectypes

import (
	"github.com/fullstack-lang/gong/dsm/reqif/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type increaseButtonProxy struct {
	stager      *models.Stager
	targetValue *int
}

func (p *increaseButtonProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button,
) {
	*p.targetValue = *p.targetValue + 1

	p.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(p.stager)
	p.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(p.stager)
}

type decreaseButtonProxy struct {
	stager      *models.Stager
	targetValue *int
}

func (p *decreaseButtonProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button,
) {
	*p.targetValue = *p.targetValue - 1

	p.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(p.stager)
	p.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(p.stager)
}
