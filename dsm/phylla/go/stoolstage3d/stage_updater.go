package stoolstage3d

import (
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
)

type Stool3DStageUpdater struct{}

func NewStool3DStageUpdater() *Stool3DStageUpdater {
	return &Stool3DStageUpdater{}
}

func (u *Stool3DStageUpdater) UpdateStool3DStage(stager *models.Stager) {
	u.ux_3d_stool(stager)
}

func (u *Stool3DStageUpdater) GenerateStoolOBJ(stager *models.Stager, plant *models.PlantAbstract) string {
	return GenerateStoolOBJ(stager, plant)
}

func (u *Stool3DStageUpdater) GenerateStoolSTL(stager *models.Stager, plant *models.PlantAbstract) string {
	return GenerateStoolSTL(stager, plant)
}
