package stool

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

func (u *Stool3DStageUpdater) GenerateStool2DDXF(stager *models.Stager, plant *models.PlantAbstract) string {
	return GenerateStool2DDXF(stager, plant)
}

func (u *Stool3DStageUpdater) GenerateStool3DDXF(stager *models.Stager, plant *models.PlantAbstract) string {
	return GenerateStool3DDXF(stager, plant)
}

func (u *Stool3DStageUpdater) GenerateStoolHaussmannBlenderPy(stager *models.Stager, plant *models.PlantAbstract) string {
	return GenerateStoolHaussmannBlenderPy(stager, plant)
}

func (u *Stool3DStageUpdater) GenerateStoolManufacturingPDF(stager *models.Stager, plant *models.PlantAbstract) []byte {
	return GenerateStoolManufacturingPDF(stager, plant)
}
