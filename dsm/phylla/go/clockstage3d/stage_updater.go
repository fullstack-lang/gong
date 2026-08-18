package clockstage3d

import (
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
)

type Clock3DStageUpdater struct{}

func NewClock3DStageUpdater() *Clock3DStageUpdater {
	return &Clock3DStageUpdater{}
}

func (u *Clock3DStageUpdater) UpdateClock3DStage(stager *models.Stager) {
	u.ux_3d_clock(stager)
}
