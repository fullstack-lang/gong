package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/tone/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Freqency__000000_B4 := (&models.Freqency{}).Stage(stage)
	__Freqency__000001_Bb4 := (&models.Freqency{}).Stage(stage)
	__Freqency__000002_C5 := (&models.Freqency{}).Stage(stage)
	__Freqency__000003_D5 := (&models.Freqency{}).Stage(stage)
	__Freqency__000004_Eb4 := (&models.Freqency{}).Stage(stage)
	__Freqency__000005_G_4 := (&models.Freqency{}).Stage(stage)
	__Freqency__000006_G4 := (&models.Freqency{}).Stage(stage)

	__Note__000000_1_C5 := (&models.Note{}).Stage(stage)
	__Note__000001_10_G4 := (&models.Note{}).Stage(stage)
	__Note__000002_2_B4 := (&models.Note{}).Stage(stage)
	__Note__000003_3_C5 := (&models.Note{}).Stage(stage)
	__Note__000004_4_G4 := (&models.Note{}).Stage(stage)
	__Note__000005_5_G4_ := (&models.Note{}).Stage(stage)
	__Note__000006_6_C5 := (&models.Note{}).Stage(stage)
	__Note__000007_7_B4 := (&models.Note{}).Stage(stage)
	__Note__000008_8_C5 := (&models.Note{}).Stage(stage)
	__Note__000009_9_D5 := (&models.Note{}).Stage(stage)

	__Player__000000_player := (&models.Player{}).Stage(stage)

	// Setup of values

	__Freqency__000000_B4.Name = `B4`

	__Freqency__000001_Bb4.Name = `Bb4`

	__Freqency__000002_C5.Name = `C5`

	__Freqency__000003_D5.Name = `D5`

	__Freqency__000004_Eb4.Name = `Eb4`

	__Freqency__000005_G_4.Name = `G#4`

	__Freqency__000006_G4.Name = `G4`

	__Note__000000_1_C5.Name = `1 C5`
	__Note__000000_1_C5.Start = 0.000000
	__Note__000000_1_C5.Duration = 0.500000
	__Note__000000_1_C5.Velocity = 0.000000
	__Note__000000_1_C5.Info = ``

	__Note__000001_10_G4.Name = `10. G4`
	__Note__000001_10_G4.Start = 7.000000
	__Note__000001_10_G4.Duration = 1.000000
	__Note__000001_10_G4.Velocity = 0.000000
	__Note__000001_10_G4.Info = ``

	__Note__000002_2_B4.Name = `2. B4`
	__Note__000002_2_B4.Start = 0.500000
	__Note__000002_2_B4.Duration = 0.500000
	__Note__000002_2_B4.Velocity = 0.000000
	__Note__000002_2_B4.Info = ``

	__Note__000003_3_C5.Name = `3. C5`
	__Note__000003_3_C5.Start = 1.000000
	__Note__000003_3_C5.Duration = 1.000000
	__Note__000003_3_C5.Velocity = 0.000000
	__Note__000003_3_C5.Info = ``

	__Note__000004_4_G4.Name = `4. G4`
	__Note__000004_4_G4.Start = 2.000000
	__Note__000004_4_G4.Duration = 1.000000
	__Note__000004_4_G4.Velocity = 0.000000
	__Note__000004_4_G4.Info = ``

	__Note__000005_5_G4_.Name = `5. G4#`
	__Note__000005_5_G4_.Start = 3.000000
	__Note__000005_5_G4_.Duration = 1.000000
	__Note__000005_5_G4_.Velocity = 0.000000
	__Note__000005_5_G4_.Info = ``

	__Note__000006_6_C5.Name = `6. C5`
	__Note__000006_6_C5.Start = 4.000000
	__Note__000006_6_C5.Duration = 0.500000
	__Note__000006_6_C5.Velocity = 0.000000
	__Note__000006_6_C5.Info = ``

	__Note__000007_7_B4.Name = `7. B4`
	__Note__000007_7_B4.Start = 4.500000
	__Note__000007_7_B4.Duration = 0.500000
	__Note__000007_7_B4.Velocity = 0.000000
	__Note__000007_7_B4.Info = ``

	__Note__000008_8_C5.Name = `8. C5`
	__Note__000008_8_C5.Start = 5.000000
	__Note__000008_8_C5.Duration = 1.000000
	__Note__000008_8_C5.Velocity = 0.000000
	__Note__000008_8_C5.Info = ``

	__Note__000009_9_D5.Name = `9. D5`
	__Note__000009_9_D5.Start = 6.000000
	__Note__000009_9_D5.Duration = 1.000000
	__Note__000009_9_D5.Velocity = 0.000000
	__Note__000009_9_D5.Info = ``

	__Player__000000_player.Name = `player`
	__Player__000000_player.Status = models.PLAYING

	// Setup of pointers
	// setup of Freqency instances pointers
	// setup of Note instances pointers
	__Note__000000_1_C5.Frequencies = append(__Note__000000_1_C5.Frequencies, __Freqency__000002_C5)
	__Note__000001_10_G4.Frequencies = append(__Note__000001_10_G4.Frequencies, __Freqency__000006_G4)
	__Note__000002_2_B4.Frequencies = append(__Note__000002_2_B4.Frequencies, __Freqency__000000_B4)
	__Note__000003_3_C5.Frequencies = append(__Note__000003_3_C5.Frequencies, __Freqency__000002_C5)
	__Note__000004_4_G4.Frequencies = append(__Note__000004_4_G4.Frequencies, __Freqency__000006_G4)
	__Note__000005_5_G4_.Frequencies = append(__Note__000005_5_G4_.Frequencies, __Freqency__000005_G_4)
	__Note__000006_6_C5.Frequencies = append(__Note__000006_6_C5.Frequencies, __Freqency__000002_C5)
	__Note__000007_7_B4.Frequencies = append(__Note__000007_7_B4.Frequencies, __Freqency__000000_B4)
	__Note__000008_8_C5.Frequencies = append(__Note__000008_8_C5.Frequencies, __Freqency__000002_C5)
	__Note__000009_9_D5.Frequencies = append(__Note__000009_9_D5.Frequencies, __Freqency__000003_D5)
	// setup of Player instances pointers
}
