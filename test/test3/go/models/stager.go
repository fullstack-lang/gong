// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"time"

	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	probeForm ProbeIF,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.probeForm = probeForm

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	stager.splitStage.Commit()

	stager.probeForm.AddNotification(time.Now(), `La rue assourdissante autour de moi hurlait.
Longue, mince, en grand deuil, douleur majestueuse,
Une femme passa, d’une main fastueuse
Soulevant, balançant le feston et l’ourlet;

Agile et noble, avec sa jambe de statue.
Moi, je buvais, crispé comme un extravagant,
Dans son oeil, ciel livide où germe l’ouragan,
La douceur qui fascine et le plaisir qui tue.

Un éclair… puis la nuit! – Fugitive beauté
Dont le regard m’a fait soudainement renaître,
Ne te verrai-je plus que dans l’éternité?

Ailleurs, bien loin d’ici! trop tard! jamais peut-être!
Car j’ignore où tu fuis, tu ne sais où je vais,
O toi que j’eusse aimée, ô toi qui le savais!`)

	callbacks := &BeforeCommitImplementation{
		stager: stager,
	}
	stager.stage.OnInitCommitFromBackCallback = callbacks
	callbacks.BeforeCommit(stage)

	return
}

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {

}
