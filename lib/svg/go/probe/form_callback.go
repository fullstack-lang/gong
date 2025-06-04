// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AnimateFormCallback(
	animate *models.Animate,
	probe *Probe,
	formGroup *table.FormGroup,
) (animateFormCallback *AnimateFormCallback) {
	animateFormCallback = new(AnimateFormCallback)
	animateFormCallback.probe = probe
	animateFormCallback.animate = animate
	animateFormCallback.formGroup = formGroup

	animateFormCallback.CreationMode = (animate == nil)

	return
}

type AnimateFormCallback struct {
	animate *models.Animate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (animateFormCallback *AnimateFormCallback) OnSave() {

	// log.Println("AnimateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	animateFormCallback.probe.formStage.Checkout()

	if animateFormCallback.animate == nil {
		animateFormCallback.animate = new(models.Animate).Stage(animateFormCallback.probe.stageOfInterest)
	}
	animate_ := animateFormCallback.animate
	_ = animate_

	for _, formDiv := range animateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(animate_.Name), formDiv)
		case "AttributeName":
			FormDivBasicFieldToField(&(animate_.AttributeName), formDiv)
		case "Values":
			FormDivBasicFieldToField(&(animate_.Values), formDiv)
		case "From":
			FormDivBasicFieldToField(&(animate_.From), formDiv)
		case "To":
			FormDivBasicFieldToField(&(animate_.To), formDiv)
		case "Dur":
			FormDivBasicFieldToField(&(animate_.Dur), formDiv)
		case "RepeatCount":
			FormDivBasicFieldToField(&(animate_.RepeatCount), formDiv)
		case "Circle:Animations":
			// WARNING : this form deals with the N-N association "Circle.Animations []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Circle
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Circle"
				rf.Fieldname = "Animations"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Circle)
					if !ok {
						log.Fatalln("Source of Circle.Animations []*Animate, is not an Circle instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animations, animate_)
					formerSource.Animations = slices.Delete(formerSource.Animations, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Circle
			for _circle := range *models.GetGongstructInstancesSet[models.Circle](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _circle.GetName() == newSourceName.GetName() {
					newSource = _circle // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Circle.Animations []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animations = append(newSource.Animations, animate_)
		case "Ellipse:Animates":
			// WARNING : this form deals with the N-N association "Ellipse.Animates []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Ellipse
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Ellipse"
				rf.Fieldname = "Animates"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Ellipse)
					if !ok {
						log.Fatalln("Source of Ellipse.Animates []*Animate, is not an Ellipse instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animates, animate_)
					formerSource.Animates = slices.Delete(formerSource.Animates, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Ellipse
			for _ellipse := range *models.GetGongstructInstancesSet[models.Ellipse](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _ellipse.GetName() == newSourceName.GetName() {
					newSource = _ellipse // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Ellipse.Animates []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animates = append(newSource.Animates, animate_)
		case "Line:Animates":
			// WARNING : this form deals with the N-N association "Line.Animates []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Line
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Line"
				rf.Fieldname = "Animates"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Line)
					if !ok {
						log.Fatalln("Source of Line.Animates []*Animate, is not an Line instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animates, animate_)
					formerSource.Animates = slices.Delete(formerSource.Animates, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Line
			for _line := range *models.GetGongstructInstancesSet[models.Line](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _line.GetName() == newSourceName.GetName() {
					newSource = _line // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Line.Animates []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animates = append(newSource.Animates, animate_)
		case "LinkAnchoredText:Animates":
			// WARNING : this form deals with the N-N association "LinkAnchoredText.Animates []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.LinkAnchoredText
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "LinkAnchoredText"
				rf.Fieldname = "Animates"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.LinkAnchoredText)
					if !ok {
						log.Fatalln("Source of LinkAnchoredText.Animates []*Animate, is not an LinkAnchoredText instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animates, animate_)
					formerSource.Animates = slices.Delete(formerSource.Animates, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.LinkAnchoredText
			for _linkanchoredtext := range *models.GetGongstructInstancesSet[models.LinkAnchoredText](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _linkanchoredtext.GetName() == newSourceName.GetName() {
					newSource = _linkanchoredtext // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of LinkAnchoredText.Animates []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animates = append(newSource.Animates, animate_)
		case "Path:Animates":
			// WARNING : this form deals with the N-N association "Path.Animates []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Path
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Path"
				rf.Fieldname = "Animates"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Path)
					if !ok {
						log.Fatalln("Source of Path.Animates []*Animate, is not an Path instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animates, animate_)
					formerSource.Animates = slices.Delete(formerSource.Animates, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Path
			for _path := range *models.GetGongstructInstancesSet[models.Path](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _path.GetName() == newSourceName.GetName() {
					newSource = _path // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Path.Animates []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animates = append(newSource.Animates, animate_)
		case "Polygone:Animates":
			// WARNING : this form deals with the N-N association "Polygone.Animates []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Polygone
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Polygone"
				rf.Fieldname = "Animates"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Polygone)
					if !ok {
						log.Fatalln("Source of Polygone.Animates []*Animate, is not an Polygone instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animates, animate_)
					formerSource.Animates = slices.Delete(formerSource.Animates, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Polygone
			for _polygone := range *models.GetGongstructInstancesSet[models.Polygone](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _polygone.GetName() == newSourceName.GetName() {
					newSource = _polygone // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Polygone.Animates []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animates = append(newSource.Animates, animate_)
		case "Polyline:Animates":
			// WARNING : this form deals with the N-N association "Polyline.Animates []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Polyline
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Polyline"
				rf.Fieldname = "Animates"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Polyline)
					if !ok {
						log.Fatalln("Source of Polyline.Animates []*Animate, is not an Polyline instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animates, animate_)
					formerSource.Animates = slices.Delete(formerSource.Animates, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Polyline
			for _polyline := range *models.GetGongstructInstancesSet[models.Polyline](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _polyline.GetName() == newSourceName.GetName() {
					newSource = _polyline // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Polyline.Animates []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animates = append(newSource.Animates, animate_)
		case "Rect:Animations":
			// WARNING : this form deals with the N-N association "Rect.Animations []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Rect
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Rect"
				rf.Fieldname = "Animations"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Rect)
					if !ok {
						log.Fatalln("Source of Rect.Animations []*Animate, is not an Rect instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animations, animate_)
					formerSource.Animations = slices.Delete(formerSource.Animations, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Rect
			for _rect := range *models.GetGongstructInstancesSet[models.Rect](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _rect.GetName() == newSourceName.GetName() {
					newSource = _rect // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Rect.Animations []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animations = append(newSource.Animations, animate_)
		case "RectAnchoredText:Animates":
			// WARNING : this form deals with the N-N association "RectAnchoredText.Animates []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.RectAnchoredText
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "RectAnchoredText"
				rf.Fieldname = "Animates"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.RectAnchoredText)
					if !ok {
						log.Fatalln("Source of RectAnchoredText.Animates []*Animate, is not an RectAnchoredText instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animates, animate_)
					formerSource.Animates = slices.Delete(formerSource.Animates, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.RectAnchoredText
			for _rectanchoredtext := range *models.GetGongstructInstancesSet[models.RectAnchoredText](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _rectanchoredtext.GetName() == newSourceName.GetName() {
					newSource = _rectanchoredtext // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of RectAnchoredText.Animates []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animates = append(newSource.Animates, animate_)
		case "Text:Animates":
			// WARNING : this form deals with the N-N association "Text.Animates []*Animate" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Animate). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Text
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Text"
				rf.Fieldname = "Animates"
				formerAssociationSource := models.GetReverseFieldOwner(
					animateFormCallback.probe.stageOfInterest,
					animate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Text)
					if !ok {
						log.Fatalln("Source of Text.Animates []*Animate, is not an Text instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Animates, animate_)
					formerSource.Animates = slices.Delete(formerSource.Animates, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Text
			for _text := range *models.GetGongstructInstancesSet[models.Text](animateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _text.GetName() == newSourceName.GetName() {
					newSource = _text // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Text.Animates []*Animate, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Animates = append(newSource.Animates, animate_)
		}
	}

	// manage the suppress operation
	if animateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		animate_.Unstage(animateFormCallback.probe.stageOfInterest)
	}

	animateFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Animate](
		animateFormCallback.probe,
	)
	animateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if animateFormCallback.CreationMode || animateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		animateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(animateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnimateFormCallback(
			nil,
			animateFormCallback.probe,
			newFormGroup,
		)
		animate := new(models.Animate)
		FillUpForm(animate, newFormGroup, animateFormCallback.probe)
		animateFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(animateFormCallback.probe)
}
func __gong__New__CircleFormCallback(
	circle *models.Circle,
	probe *Probe,
	formGroup *table.FormGroup,
) (circleFormCallback *CircleFormCallback) {
	circleFormCallback = new(CircleFormCallback)
	circleFormCallback.probe = probe
	circleFormCallback.circle = circle
	circleFormCallback.formGroup = formGroup

	circleFormCallback.CreationMode = (circle == nil)

	return
}

type CircleFormCallback struct {
	circle *models.Circle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (circleFormCallback *CircleFormCallback) OnSave() {

	// log.Println("CircleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circleFormCallback.probe.formStage.Checkout()

	if circleFormCallback.circle == nil {
		circleFormCallback.circle = new(models.Circle).Stage(circleFormCallback.probe.stageOfInterest)
	}
	circle_ := circleFormCallback.circle
	_ = circle_

	for _, formDiv := range circleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(circle_.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(circle_.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(circle_.CY), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(circle_.Radius), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(circle_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(circle_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(circle_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(circle_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(circle_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(circle_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(circle_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(circle_.Transform), formDiv)
		case "Animations":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](circleFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					circleFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			circle_.Animations = instanceSlice

		case "Layer:Circles":
			// WARNING : this form deals with the N-N association "Layer.Circles []*Circle" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Circle). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "Circles"
				formerAssociationSource := models.GetReverseFieldOwner(
					circleFormCallback.probe.stageOfInterest,
					circle_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.Circles []*Circle, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Circles, circle_)
					formerSource.Circles = slices.Delete(formerSource.Circles, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](circleFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.Circles []*Circle, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Circles = append(newSource.Circles, circle_)
		}
	}

	// manage the suppress operation
	if circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circle_.Unstage(circleFormCallback.probe.stageOfInterest)
	}

	circleFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Circle](
		circleFormCallback.probe,
	)
	circleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if circleFormCallback.CreationMode || circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(circleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CircleFormCallback(
			nil,
			circleFormCallback.probe,
			newFormGroup,
		)
		circle := new(models.Circle)
		FillUpForm(circle, newFormGroup, circleFormCallback.probe)
		circleFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(circleFormCallback.probe)
}
func __gong__New__EllipseFormCallback(
	ellipse *models.Ellipse,
	probe *Probe,
	formGroup *table.FormGroup,
) (ellipseFormCallback *EllipseFormCallback) {
	ellipseFormCallback = new(EllipseFormCallback)
	ellipseFormCallback.probe = probe
	ellipseFormCallback.ellipse = ellipse
	ellipseFormCallback.formGroup = formGroup

	ellipseFormCallback.CreationMode = (ellipse == nil)

	return
}

type EllipseFormCallback struct {
	ellipse *models.Ellipse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (ellipseFormCallback *EllipseFormCallback) OnSave() {

	// log.Println("EllipseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ellipseFormCallback.probe.formStage.Checkout()

	if ellipseFormCallback.ellipse == nil {
		ellipseFormCallback.ellipse = new(models.Ellipse).Stage(ellipseFormCallback.probe.stageOfInterest)
	}
	ellipse_ := ellipseFormCallback.ellipse
	_ = ellipse_

	for _, formDiv := range ellipseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(ellipse_.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(ellipse_.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(ellipse_.CY), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(ellipse_.RX), formDiv)
		case "RY":
			FormDivBasicFieldToField(&(ellipse_.RY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(ellipse_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(ellipse_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(ellipse_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(ellipse_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(ellipse_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(ellipse_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(ellipse_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(ellipse_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](ellipseFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					ellipseFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			ellipse_.Animates = instanceSlice

		case "Layer:Ellipses":
			// WARNING : this form deals with the N-N association "Layer.Ellipses []*Ellipse" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Ellipse). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "Ellipses"
				formerAssociationSource := models.GetReverseFieldOwner(
					ellipseFormCallback.probe.stageOfInterest,
					ellipse_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.Ellipses []*Ellipse, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Ellipses, ellipse_)
					formerSource.Ellipses = slices.Delete(formerSource.Ellipses, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](ellipseFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.Ellipses []*Ellipse, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Ellipses = append(newSource.Ellipses, ellipse_)
		}
	}

	// manage the suppress operation
	if ellipseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ellipse_.Unstage(ellipseFormCallback.probe.stageOfInterest)
	}

	ellipseFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Ellipse](
		ellipseFormCallback.probe,
	)
	ellipseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if ellipseFormCallback.CreationMode || ellipseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ellipseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(ellipseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EllipseFormCallback(
			nil,
			ellipseFormCallback.probe,
			newFormGroup,
		)
		ellipse := new(models.Ellipse)
		FillUpForm(ellipse, newFormGroup, ellipseFormCallback.probe)
		ellipseFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(ellipseFormCallback.probe)
}
func __gong__New__LayerFormCallback(
	layer *models.Layer,
	probe *Probe,
	formGroup *table.FormGroup,
) (layerFormCallback *LayerFormCallback) {
	layerFormCallback = new(LayerFormCallback)
	layerFormCallback.probe = probe
	layerFormCallback.layer = layer
	layerFormCallback.formGroup = formGroup

	layerFormCallback.CreationMode = (layer == nil)

	return
}

type LayerFormCallback struct {
	layer *models.Layer

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (layerFormCallback *LayerFormCallback) OnSave() {

	// log.Println("LayerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layerFormCallback.probe.formStage.Checkout()

	if layerFormCallback.layer == nil {
		layerFormCallback.layer = new(models.Layer).Stage(layerFormCallback.probe.stageOfInterest)
	}
	layer_ := layerFormCallback.layer
	_ = layer_

	for _, formDiv := range layerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(layer_.Name), formDiv)
		case "Rects":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Rect](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Rect, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Rect)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.Rects = instanceSlice

		case "Texts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Text](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Text, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Text)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.Texts = instanceSlice

		case "Circles":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Circle](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Circle, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Circle)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.Circles = instanceSlice

		case "Lines":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Line](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Line, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Line)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.Lines = instanceSlice

		case "Ellipses":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Ellipse](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Ellipse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Ellipse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.Ellipses = instanceSlice

		case "Polylines":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Polyline](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Polyline, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Polyline)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.Polylines = instanceSlice

		case "Polygones":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Polygone](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Polygone, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Polygone)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.Polygones = instanceSlice

		case "Paths":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Path](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Path, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Path)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.Paths = instanceSlice

		case "Links":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Link](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Link, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Link)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.Links = instanceSlice

		case "RectLinkLinks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RectLinkLink](layerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RectLinkLink, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RectLinkLink)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layer_.RectLinkLinks = instanceSlice

		case "SVG:Layers":
			// WARNING : this form deals with the N-N association "SVG.Layers []*Layer" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Layer). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.SVG
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "SVG"
				rf.Fieldname = "Layers"
				formerAssociationSource := models.GetReverseFieldOwner(
					layerFormCallback.probe.stageOfInterest,
					layer_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.SVG)
					if !ok {
						log.Fatalln("Source of SVG.Layers []*Layer, is not an SVG instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Layers, layer_)
					formerSource.Layers = slices.Delete(formerSource.Layers, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.SVG
			for _svg := range *models.GetGongstructInstancesSet[models.SVG](layerFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _svg.GetName() == newSourceName.GetName() {
					newSource = _svg // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of SVG.Layers []*Layer, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Layers = append(newSource.Layers, layer_)
		}
	}

	// manage the suppress operation
	if layerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layer_.Unstage(layerFormCallback.probe.stageOfInterest)
	}

	layerFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Layer](
		layerFormCallback.probe,
	)
	layerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if layerFormCallback.CreationMode || layerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(layerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LayerFormCallback(
			nil,
			layerFormCallback.probe,
			newFormGroup,
		)
		layer := new(models.Layer)
		FillUpForm(layer, newFormGroup, layerFormCallback.probe)
		layerFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(layerFormCallback.probe)
}
func __gong__New__LineFormCallback(
	line *models.Line,
	probe *Probe,
	formGroup *table.FormGroup,
) (lineFormCallback *LineFormCallback) {
	lineFormCallback = new(LineFormCallback)
	lineFormCallback.probe = probe
	lineFormCallback.line = line
	lineFormCallback.formGroup = formGroup

	lineFormCallback.CreationMode = (line == nil)

	return
}

type LineFormCallback struct {
	line *models.Line

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lineFormCallback *LineFormCallback) OnSave() {

	// log.Println("LineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lineFormCallback.probe.formStage.Checkout()

	if lineFormCallback.line == nil {
		lineFormCallback.line = new(models.Line).Stage(lineFormCallback.probe.stageOfInterest)
	}
	line_ := lineFormCallback.line
	_ = line_

	for _, formDiv := range lineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(line_.Name), formDiv)
		case "X1":
			FormDivBasicFieldToField(&(line_.X1), formDiv)
		case "Y1":
			FormDivBasicFieldToField(&(line_.Y1), formDiv)
		case "X2":
			FormDivBasicFieldToField(&(line_.X2), formDiv)
		case "Y2":
			FormDivBasicFieldToField(&(line_.Y2), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(line_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(line_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(line_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(line_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(line_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(line_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(line_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(line_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](lineFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					lineFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			line_.Animates = instanceSlice

		case "MouseClickX":
			FormDivBasicFieldToField(&(line_.MouseClickX), formDiv)
		case "MouseClickY":
			FormDivBasicFieldToField(&(line_.MouseClickY), formDiv)
		case "Layer:Lines":
			// WARNING : this form deals with the N-N association "Layer.Lines []*Line" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Line). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "Lines"
				formerAssociationSource := models.GetReverseFieldOwner(
					lineFormCallback.probe.stageOfInterest,
					line_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.Lines []*Line, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Lines, line_)
					formerSource.Lines = slices.Delete(formerSource.Lines, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](lineFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.Lines []*Line, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Lines = append(newSource.Lines, line_)
		}
	}

	// manage the suppress operation
	if lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		line_.Unstage(lineFormCallback.probe.stageOfInterest)
	}

	lineFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Line](
		lineFormCallback.probe,
	)
	lineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lineFormCallback.CreationMode || lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(lineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LineFormCallback(
			nil,
			lineFormCallback.probe,
			newFormGroup,
		)
		line := new(models.Line)
		FillUpForm(line, newFormGroup, lineFormCallback.probe)
		lineFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(lineFormCallback.probe)
}
func __gong__New__LinkFormCallback(
	link *models.Link,
	probe *Probe,
	formGroup *table.FormGroup,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.probe = probe
	linkFormCallback.link = link
	linkFormCallback.formGroup = formGroup

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	// log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.probe.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.probe.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	for _, formDiv := range linkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(link_.Type), formDiv)
		case "IsBezierCurve":
			FormDivBasicFieldToField(&(link_.IsBezierCurve), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(link_.Start), linkFormCallback.probe.stageOfInterest, formDiv)
		case "StartAnchorType":
			FormDivEnumStringFieldToField(&(link_.StartAnchorType), formDiv)
		case "End":
			FormDivSelectFieldToField(&(link_.End), linkFormCallback.probe.stageOfInterest, formDiv)
		case "EndAnchorType":
			FormDivEnumStringFieldToField(&(link_.EndAnchorType), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(link_.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(link_.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(link_.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(link_.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(link_.CornerOffsetRatio), formDiv)
		case "CornerRadius":
			FormDivBasicFieldToField(&(link_.CornerRadius), formDiv)
		case "HasEndArrow":
			FormDivBasicFieldToField(&(link_.HasEndArrow), formDiv)
		case "EndArrowSize":
			FormDivBasicFieldToField(&(link_.EndArrowSize), formDiv)
		case "HasStartArrow":
			FormDivBasicFieldToField(&(link_.HasStartArrow), formDiv)
		case "StartArrowSize":
			FormDivBasicFieldToField(&(link_.StartArrowSize), formDiv)
		case "TextAtArrowStart":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredText](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkAnchoredText, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkAnchoredText)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			link_.TextAtArrowStart = instanceSlice

		case "TextAtArrowEnd":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredText](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkAnchoredText, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkAnchoredText)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			link_.TextAtArrowEnd = instanceSlice

		case "ControlPoints":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Point](linkFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Point, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Point)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			link_.ControlPoints = instanceSlice

		case "Color":
			FormDivBasicFieldToField(&(link_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(link_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(link_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(link_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(link_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(link_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(link_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(link_.Transform), formDiv)
		case "Layer:Links":
			// WARNING : this form deals with the N-N association "Layer.Links []*Link" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Link). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "Links"
				formerAssociationSource := models.GetReverseFieldOwner(
					linkFormCallback.probe.stageOfInterest,
					link_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.Links []*Link, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Links, link_)
					formerSource.Links = slices.Delete(formerSource.Links, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](linkFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.Links []*Link, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Links = append(newSource.Links, link_)
		}
	}

	// manage the suppress operation
	if linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		link_.Unstage(linkFormCallback.probe.stageOfInterest)
	}

	linkFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Link](
		linkFormCallback.probe,
	)
	linkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode || linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(linkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			linkFormCallback.probe,
			newFormGroup,
		)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.probe)
		linkFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(linkFormCallback.probe)
}
func __gong__New__LinkAnchoredTextFormCallback(
	linkanchoredtext *models.LinkAnchoredText,
	probe *Probe,
	formGroup *table.FormGroup,
) (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) {
	linkanchoredtextFormCallback = new(LinkAnchoredTextFormCallback)
	linkanchoredtextFormCallback.probe = probe
	linkanchoredtextFormCallback.linkanchoredtext = linkanchoredtext
	linkanchoredtextFormCallback.formGroup = formGroup

	linkanchoredtextFormCallback.CreationMode = (linkanchoredtext == nil)

	return
}

type LinkAnchoredTextFormCallback struct {
	linkanchoredtext *models.LinkAnchoredText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) OnSave() {

	// log.Println("LinkAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkanchoredtextFormCallback.probe.formStage.Checkout()

	if linkanchoredtextFormCallback.linkanchoredtext == nil {
		linkanchoredtextFormCallback.linkanchoredtext = new(models.LinkAnchoredText).Stage(linkanchoredtextFormCallback.probe.stageOfInterest)
	}
	linkanchoredtext_ := linkanchoredtextFormCallback.linkanchoredtext
	_ = linkanchoredtext_

	for _, formDiv := range linkanchoredtextFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(linkanchoredtext_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(linkanchoredtext_.Content), formDiv)
		case "AutomaticLayout":
			FormDivBasicFieldToField(&(linkanchoredtext_.AutomaticLayout), formDiv)
		case "LinkAnchorType":
			FormDivEnumStringFieldToField(&(linkanchoredtext_.LinkAnchorType), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(linkanchoredtext_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(linkanchoredtext_.Y_Offset), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(linkanchoredtext_.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(linkanchoredtext_.FontSize), formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&(linkanchoredtext_.FontStyle), formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&(linkanchoredtext_.LetterSpacing), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(linkanchoredtext_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(linkanchoredtext_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(linkanchoredtext_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(linkanchoredtext_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](linkanchoredtextFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					linkanchoredtextFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			linkanchoredtext_.Animates = instanceSlice

		case "Link:TextAtArrowStart":
			// WARNING : this form deals with the N-N association "Link.TextAtArrowStart []*LinkAnchoredText" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of LinkAnchoredText). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Link
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Link"
				rf.Fieldname = "TextAtArrowStart"
				formerAssociationSource := models.GetReverseFieldOwner(
					linkanchoredtextFormCallback.probe.stageOfInterest,
					linkanchoredtext_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Link)
					if !ok {
						log.Fatalln("Source of Link.TextAtArrowStart []*LinkAnchoredText, is not an Link instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TextAtArrowStart, linkanchoredtext_)
					formerSource.TextAtArrowStart = slices.Delete(formerSource.TextAtArrowStart, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Link
			for _link := range *models.GetGongstructInstancesSet[models.Link](linkanchoredtextFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _link.GetName() == newSourceName.GetName() {
					newSource = _link // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Link.TextAtArrowStart []*LinkAnchoredText, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TextAtArrowStart = append(newSource.TextAtArrowStart, linkanchoredtext_)
		case "Link:TextAtArrowEnd":
			// WARNING : this form deals with the N-N association "Link.TextAtArrowEnd []*LinkAnchoredText" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of LinkAnchoredText). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Link
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Link"
				rf.Fieldname = "TextAtArrowEnd"
				formerAssociationSource := models.GetReverseFieldOwner(
					linkanchoredtextFormCallback.probe.stageOfInterest,
					linkanchoredtext_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Link)
					if !ok {
						log.Fatalln("Source of Link.TextAtArrowEnd []*LinkAnchoredText, is not an Link instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TextAtArrowEnd, linkanchoredtext_)
					formerSource.TextAtArrowEnd = slices.Delete(formerSource.TextAtArrowEnd, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Link
			for _link := range *models.GetGongstructInstancesSet[models.Link](linkanchoredtextFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _link.GetName() == newSourceName.GetName() {
					newSource = _link // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Link.TextAtArrowEnd []*LinkAnchoredText, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TextAtArrowEnd = append(newSource.TextAtArrowEnd, linkanchoredtext_)
		}
	}

	// manage the suppress operation
	if linkanchoredtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkanchoredtext_.Unstage(linkanchoredtextFormCallback.probe.stageOfInterest)
	}

	linkanchoredtextFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.LinkAnchoredText](
		linkanchoredtextFormCallback.probe,
	)
	linkanchoredtextFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if linkanchoredtextFormCallback.CreationMode || linkanchoredtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkanchoredtextFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(linkanchoredtextFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkAnchoredTextFormCallback(
			nil,
			linkanchoredtextFormCallback.probe,
			newFormGroup,
		)
		linkanchoredtext := new(models.LinkAnchoredText)
		FillUpForm(linkanchoredtext, newFormGroup, linkanchoredtextFormCallback.probe)
		linkanchoredtextFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(linkanchoredtextFormCallback.probe)
}
func __gong__New__PathFormCallback(
	path *models.Path,
	probe *Probe,
	formGroup *table.FormGroup,
) (pathFormCallback *PathFormCallback) {
	pathFormCallback = new(PathFormCallback)
	pathFormCallback.probe = probe
	pathFormCallback.path = path
	pathFormCallback.formGroup = formGroup

	pathFormCallback.CreationMode = (path == nil)

	return
}

type PathFormCallback struct {
	path *models.Path

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (pathFormCallback *PathFormCallback) OnSave() {

	// log.Println("PathFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pathFormCallback.probe.formStage.Checkout()

	if pathFormCallback.path == nil {
		pathFormCallback.path = new(models.Path).Stage(pathFormCallback.probe.stageOfInterest)
	}
	path_ := pathFormCallback.path
	_ = path_

	for _, formDiv := range pathFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(path_.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(path_.Definition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(path_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(path_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(path_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(path_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(path_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(path_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(path_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(path_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](pathFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					pathFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			path_.Animates = instanceSlice

		case "Layer:Paths":
			// WARNING : this form deals with the N-N association "Layer.Paths []*Path" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Path). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "Paths"
				formerAssociationSource := models.GetReverseFieldOwner(
					pathFormCallback.probe.stageOfInterest,
					path_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.Paths []*Path, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Paths, path_)
					formerSource.Paths = slices.Delete(formerSource.Paths, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](pathFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.Paths []*Path, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Paths = append(newSource.Paths, path_)
		}
	}

	// manage the suppress operation
	if pathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		path_.Unstage(pathFormCallback.probe.stageOfInterest)
	}

	pathFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Path](
		pathFormCallback.probe,
	)
	pathFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pathFormCallback.CreationMode || pathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pathFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(pathFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PathFormCallback(
			nil,
			pathFormCallback.probe,
			newFormGroup,
		)
		path := new(models.Path)
		FillUpForm(path, newFormGroup, pathFormCallback.probe)
		pathFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(pathFormCallback.probe)
}
func __gong__New__PointFormCallback(
	point *models.Point,
	probe *Probe,
	formGroup *table.FormGroup,
) (pointFormCallback *PointFormCallback) {
	pointFormCallback = new(PointFormCallback)
	pointFormCallback.probe = probe
	pointFormCallback.point = point
	pointFormCallback.formGroup = formGroup

	pointFormCallback.CreationMode = (point == nil)

	return
}

type PointFormCallback struct {
	point *models.Point

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (pointFormCallback *PointFormCallback) OnSave() {

	// log.Println("PointFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pointFormCallback.probe.formStage.Checkout()

	if pointFormCallback.point == nil {
		pointFormCallback.point = new(models.Point).Stage(pointFormCallback.probe.stageOfInterest)
	}
	point_ := pointFormCallback.point
	_ = point_

	for _, formDiv := range pointFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(point_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(point_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(point_.Y), formDiv)
		case "Link:ControlPoints":
			// WARNING : this form deals with the N-N association "Link.ControlPoints []*Point" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Point). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Link
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Link"
				rf.Fieldname = "ControlPoints"
				formerAssociationSource := models.GetReverseFieldOwner(
					pointFormCallback.probe.stageOfInterest,
					point_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Link)
					if !ok {
						log.Fatalln("Source of Link.ControlPoints []*Point, is not an Link instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ControlPoints, point_)
					formerSource.ControlPoints = slices.Delete(formerSource.ControlPoints, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Link
			for _link := range *models.GetGongstructInstancesSet[models.Link](pointFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _link.GetName() == newSourceName.GetName() {
					newSource = _link // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Link.ControlPoints []*Point, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ControlPoints = append(newSource.ControlPoints, point_)
		}
	}

	// manage the suppress operation
	if pointFormCallback.formGroup.HasSuppressButtonBeenPressed {
		point_.Unstage(pointFormCallback.probe.stageOfInterest)
	}

	pointFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Point](
		pointFormCallback.probe,
	)
	pointFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pointFormCallback.CreationMode || pointFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pointFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(pointFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PointFormCallback(
			nil,
			pointFormCallback.probe,
			newFormGroup,
		)
		point := new(models.Point)
		FillUpForm(point, newFormGroup, pointFormCallback.probe)
		pointFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(pointFormCallback.probe)
}
func __gong__New__PolygoneFormCallback(
	polygone *models.Polygone,
	probe *Probe,
	formGroup *table.FormGroup,
) (polygoneFormCallback *PolygoneFormCallback) {
	polygoneFormCallback = new(PolygoneFormCallback)
	polygoneFormCallback.probe = probe
	polygoneFormCallback.polygone = polygone
	polygoneFormCallback.formGroup = formGroup

	polygoneFormCallback.CreationMode = (polygone == nil)

	return
}

type PolygoneFormCallback struct {
	polygone *models.Polygone

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (polygoneFormCallback *PolygoneFormCallback) OnSave() {

	// log.Println("PolygoneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polygoneFormCallback.probe.formStage.Checkout()

	if polygoneFormCallback.polygone == nil {
		polygoneFormCallback.polygone = new(models.Polygone).Stage(polygoneFormCallback.probe.stageOfInterest)
	}
	polygone_ := polygoneFormCallback.polygone
	_ = polygone_

	for _, formDiv := range polygoneFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(polygone_.Name), formDiv)
		case "Points":
			FormDivBasicFieldToField(&(polygone_.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(polygone_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(polygone_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(polygone_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(polygone_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(polygone_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(polygone_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(polygone_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(polygone_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](polygoneFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					polygoneFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			polygone_.Animates = instanceSlice

		case "Layer:Polygones":
			// WARNING : this form deals with the N-N association "Layer.Polygones []*Polygone" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Polygone). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "Polygones"
				formerAssociationSource := models.GetReverseFieldOwner(
					polygoneFormCallback.probe.stageOfInterest,
					polygone_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.Polygones []*Polygone, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Polygones, polygone_)
					formerSource.Polygones = slices.Delete(formerSource.Polygones, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](polygoneFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.Polygones []*Polygone, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Polygones = append(newSource.Polygones, polygone_)
		}
	}

	// manage the suppress operation
	if polygoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		polygone_.Unstage(polygoneFormCallback.probe.stageOfInterest)
	}

	polygoneFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Polygone](
		polygoneFormCallback.probe,
	)
	polygoneFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if polygoneFormCallback.CreationMode || polygoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		polygoneFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(polygoneFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PolygoneFormCallback(
			nil,
			polygoneFormCallback.probe,
			newFormGroup,
		)
		polygone := new(models.Polygone)
		FillUpForm(polygone, newFormGroup, polygoneFormCallback.probe)
		polygoneFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(polygoneFormCallback.probe)
}
func __gong__New__PolylineFormCallback(
	polyline *models.Polyline,
	probe *Probe,
	formGroup *table.FormGroup,
) (polylineFormCallback *PolylineFormCallback) {
	polylineFormCallback = new(PolylineFormCallback)
	polylineFormCallback.probe = probe
	polylineFormCallback.polyline = polyline
	polylineFormCallback.formGroup = formGroup

	polylineFormCallback.CreationMode = (polyline == nil)

	return
}

type PolylineFormCallback struct {
	polyline *models.Polyline

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (polylineFormCallback *PolylineFormCallback) OnSave() {

	// log.Println("PolylineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polylineFormCallback.probe.formStage.Checkout()

	if polylineFormCallback.polyline == nil {
		polylineFormCallback.polyline = new(models.Polyline).Stage(polylineFormCallback.probe.stageOfInterest)
	}
	polyline_ := polylineFormCallback.polyline
	_ = polyline_

	for _, formDiv := range polylineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(polyline_.Name), formDiv)
		case "Points":
			FormDivBasicFieldToField(&(polyline_.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(polyline_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(polyline_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(polyline_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(polyline_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(polyline_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(polyline_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(polyline_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(polyline_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](polylineFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					polylineFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			polyline_.Animates = instanceSlice

		case "Layer:Polylines":
			// WARNING : this form deals with the N-N association "Layer.Polylines []*Polyline" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Polyline). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "Polylines"
				formerAssociationSource := models.GetReverseFieldOwner(
					polylineFormCallback.probe.stageOfInterest,
					polyline_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.Polylines []*Polyline, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Polylines, polyline_)
					formerSource.Polylines = slices.Delete(formerSource.Polylines, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](polylineFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.Polylines []*Polyline, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Polylines = append(newSource.Polylines, polyline_)
		}
	}

	// manage the suppress operation
	if polylineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		polyline_.Unstage(polylineFormCallback.probe.stageOfInterest)
	}

	polylineFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Polyline](
		polylineFormCallback.probe,
	)
	polylineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if polylineFormCallback.CreationMode || polylineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		polylineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(polylineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PolylineFormCallback(
			nil,
			polylineFormCallback.probe,
			newFormGroup,
		)
		polyline := new(models.Polyline)
		FillUpForm(polyline, newFormGroup, polylineFormCallback.probe)
		polylineFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(polylineFormCallback.probe)
}
func __gong__New__RectFormCallback(
	rect *models.Rect,
	probe *Probe,
	formGroup *table.FormGroup,
) (rectFormCallback *RectFormCallback) {
	rectFormCallback = new(RectFormCallback)
	rectFormCallback.probe = probe
	rectFormCallback.rect = rect
	rectFormCallback.formGroup = formGroup

	rectFormCallback.CreationMode = (rect == nil)

	return
}

type RectFormCallback struct {
	rect *models.Rect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rectFormCallback *RectFormCallback) OnSave() {

	// log.Println("RectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectFormCallback.probe.formStage.Checkout()

	if rectFormCallback.rect == nil {
		rectFormCallback.rect = new(models.Rect).Stage(rectFormCallback.probe.stageOfInterest)
	}
	rect_ := rectFormCallback.rect
	_ = rect_

	for _, formDiv := range rectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rect_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rect_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rect_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(rect_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(rect_.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(rect_.RX), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rect_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rect_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rect_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rect_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rect_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rect_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rect_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rect_.Transform), formDiv)
		case "Animations":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			rect_.Animations = instanceSlice

		case "IsSelectable":
			FormDivBasicFieldToField(&(rect_.IsSelectable), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(rect_.IsSelected), formDiv)
		case "CanHaveLeftHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveLeftHandle), formDiv)
		case "HasLeftHandle":
			FormDivBasicFieldToField(&(rect_.HasLeftHandle), formDiv)
		case "CanHaveRightHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveRightHandle), formDiv)
		case "HasRightHandle":
			FormDivBasicFieldToField(&(rect_.HasRightHandle), formDiv)
		case "CanHaveTopHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveTopHandle), formDiv)
		case "HasTopHandle":
			FormDivBasicFieldToField(&(rect_.HasTopHandle), formDiv)
		case "IsScalingProportionally":
			FormDivBasicFieldToField(&(rect_.IsScalingProportionally), formDiv)
		case "CanHaveBottomHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveBottomHandle), formDiv)
		case "HasBottomHandle":
			FormDivBasicFieldToField(&(rect_.HasBottomHandle), formDiv)
		case "CanMoveHorizontaly":
			FormDivBasicFieldToField(&(rect_.CanMoveHorizontaly), formDiv)
		case "CanMoveVerticaly":
			FormDivBasicFieldToField(&(rect_.CanMoveVerticaly), formDiv)
		case "RectAnchoredTexts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredText](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RectAnchoredText, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RectAnchoredText)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			rect_.RectAnchoredTexts = instanceSlice

		case "RectAnchoredRects":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredRect](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RectAnchoredRect, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RectAnchoredRect)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			rect_.RectAnchoredRects = instanceSlice

		case "RectAnchoredPaths":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredPath](rectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RectAnchoredPath, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RectAnchoredPath)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			rect_.RectAnchoredPaths = instanceSlice

		case "ChangeColorWhenHovered":
			FormDivBasicFieldToField(&(rect_.ChangeColorWhenHovered), formDiv)
		case "ColorWhenHovered":
			FormDivBasicFieldToField(&(rect_.ColorWhenHovered), formDiv)
		case "OriginalColor":
			FormDivBasicFieldToField(&(rect_.OriginalColor), formDiv)
		case "FillOpacityWhenHovered":
			FormDivBasicFieldToField(&(rect_.FillOpacityWhenHovered), formDiv)
		case "OriginalFillOpacity":
			FormDivBasicFieldToField(&(rect_.OriginalFillOpacity), formDiv)
		case "Layer:Rects":
			// WARNING : this form deals with the N-N association "Layer.Rects []*Rect" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Rect). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "Rects"
				formerAssociationSource := models.GetReverseFieldOwner(
					rectFormCallback.probe.stageOfInterest,
					rect_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.Rects []*Rect, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Rects, rect_)
					formerSource.Rects = slices.Delete(formerSource.Rects, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](rectFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.Rects []*Rect, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Rects = append(newSource.Rects, rect_)
		}
	}

	// manage the suppress operation
	if rectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rect_.Unstage(rectFormCallback.probe.stageOfInterest)
	}

	rectFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Rect](
		rectFormCallback.probe,
	)
	rectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rectFormCallback.CreationMode || rectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(rectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectFormCallback(
			nil,
			rectFormCallback.probe,
			newFormGroup,
		)
		rect := new(models.Rect)
		FillUpForm(rect, newFormGroup, rectFormCallback.probe)
		rectFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(rectFormCallback.probe)
}
func __gong__New__RectAnchoredPathFormCallback(
	rectanchoredpath *models.RectAnchoredPath,
	probe *Probe,
	formGroup *table.FormGroup,
) (rectanchoredpathFormCallback *RectAnchoredPathFormCallback) {
	rectanchoredpathFormCallback = new(RectAnchoredPathFormCallback)
	rectanchoredpathFormCallback.probe = probe
	rectanchoredpathFormCallback.rectanchoredpath = rectanchoredpath
	rectanchoredpathFormCallback.formGroup = formGroup

	rectanchoredpathFormCallback.CreationMode = (rectanchoredpath == nil)

	return
}

type RectAnchoredPathFormCallback struct {
	rectanchoredpath *models.RectAnchoredPath

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rectanchoredpathFormCallback *RectAnchoredPathFormCallback) OnSave() {

	// log.Println("RectAnchoredPathFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredpathFormCallback.probe.formStage.Checkout()

	if rectanchoredpathFormCallback.rectanchoredpath == nil {
		rectanchoredpathFormCallback.rectanchoredpath = new(models.RectAnchoredPath).Stage(rectanchoredpathFormCallback.probe.stageOfInterest)
	}
	rectanchoredpath_ := rectanchoredpathFormCallback.rectanchoredpath
	_ = rectanchoredpath_

	for _, formDiv := range rectanchoredpathFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredpath_.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(rectanchoredpath_.Definition), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredpath_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredpath_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredpath_.RectAnchorType), formDiv)
		case "ScalePropotionnally":
			FormDivBasicFieldToField(&(rectanchoredpath_.ScalePropotionnally), formDiv)
		case "AppliedScaling":
			FormDivBasicFieldToField(&(rectanchoredpath_.AppliedScaling), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredpath_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredpath_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredpath_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rectanchoredpath_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredpath_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredpath_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredpath_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredpath_.Transform), formDiv)
		case "Rect:RectAnchoredPaths":
			// WARNING : this form deals with the N-N association "Rect.RectAnchoredPaths []*RectAnchoredPath" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of RectAnchoredPath). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Rect
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Rect"
				rf.Fieldname = "RectAnchoredPaths"
				formerAssociationSource := models.GetReverseFieldOwner(
					rectanchoredpathFormCallback.probe.stageOfInterest,
					rectanchoredpath_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Rect)
					if !ok {
						log.Fatalln("Source of Rect.RectAnchoredPaths []*RectAnchoredPath, is not an Rect instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RectAnchoredPaths, rectanchoredpath_)
					formerSource.RectAnchoredPaths = slices.Delete(formerSource.RectAnchoredPaths, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Rect
			for _rect := range *models.GetGongstructInstancesSet[models.Rect](rectanchoredpathFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _rect.GetName() == newSourceName.GetName() {
					newSource = _rect // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Rect.RectAnchoredPaths []*RectAnchoredPath, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RectAnchoredPaths = append(newSource.RectAnchoredPaths, rectanchoredpath_)
		}
	}

	// manage the suppress operation
	if rectanchoredpathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredpath_.Unstage(rectanchoredpathFormCallback.probe.stageOfInterest)
	}

	rectanchoredpathFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.RectAnchoredPath](
		rectanchoredpathFormCallback.probe,
	)
	rectanchoredpathFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rectanchoredpathFormCallback.CreationMode || rectanchoredpathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredpathFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(rectanchoredpathFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectAnchoredPathFormCallback(
			nil,
			rectanchoredpathFormCallback.probe,
			newFormGroup,
		)
		rectanchoredpath := new(models.RectAnchoredPath)
		FillUpForm(rectanchoredpath, newFormGroup, rectanchoredpathFormCallback.probe)
		rectanchoredpathFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(rectanchoredpathFormCallback.probe)
}
func __gong__New__RectAnchoredRectFormCallback(
	rectanchoredrect *models.RectAnchoredRect,
	probe *Probe,
	formGroup *table.FormGroup,
) (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) {
	rectanchoredrectFormCallback = new(RectAnchoredRectFormCallback)
	rectanchoredrectFormCallback.probe = probe
	rectanchoredrectFormCallback.rectanchoredrect = rectanchoredrect
	rectanchoredrectFormCallback.formGroup = formGroup

	rectanchoredrectFormCallback.CreationMode = (rectanchoredrect == nil)

	return
}

type RectAnchoredRectFormCallback struct {
	rectanchoredrect *models.RectAnchoredRect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) OnSave() {

	// log.Println("RectAnchoredRectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredrectFormCallback.probe.formStage.Checkout()

	if rectanchoredrectFormCallback.rectanchoredrect == nil {
		rectanchoredrectFormCallback.rectanchoredrect = new(models.RectAnchoredRect).Stage(rectanchoredrectFormCallback.probe.stageOfInterest)
	}
	rectanchoredrect_ := rectanchoredrectFormCallback.rectanchoredrect
	_ = rectanchoredrect_

	for _, formDiv := range rectanchoredrectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredrect_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rectanchoredrect_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rectanchoredrect_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(rectanchoredrect_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(rectanchoredrect_.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(rectanchoredrect_.RX), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredrect_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredrect_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredrect_.RectAnchorType), formDiv)
		case "WidthFollowRect":
			FormDivBasicFieldToField(&(rectanchoredrect_.WidthFollowRect), formDiv)
		case "HeightFollowRect":
			FormDivBasicFieldToField(&(rectanchoredrect_.HeightFollowRect), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredrect_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredrect_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredrect_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredrect_.Transform), formDiv)
		case "Rect:RectAnchoredRects":
			// WARNING : this form deals with the N-N association "Rect.RectAnchoredRects []*RectAnchoredRect" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of RectAnchoredRect). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Rect
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Rect"
				rf.Fieldname = "RectAnchoredRects"
				formerAssociationSource := models.GetReverseFieldOwner(
					rectanchoredrectFormCallback.probe.stageOfInterest,
					rectanchoredrect_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Rect)
					if !ok {
						log.Fatalln("Source of Rect.RectAnchoredRects []*RectAnchoredRect, is not an Rect instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RectAnchoredRects, rectanchoredrect_)
					formerSource.RectAnchoredRects = slices.Delete(formerSource.RectAnchoredRects, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Rect
			for _rect := range *models.GetGongstructInstancesSet[models.Rect](rectanchoredrectFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _rect.GetName() == newSourceName.GetName() {
					newSource = _rect // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Rect.RectAnchoredRects []*RectAnchoredRect, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RectAnchoredRects = append(newSource.RectAnchoredRects, rectanchoredrect_)
		}
	}

	// manage the suppress operation
	if rectanchoredrectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredrect_.Unstage(rectanchoredrectFormCallback.probe.stageOfInterest)
	}

	rectanchoredrectFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.RectAnchoredRect](
		rectanchoredrectFormCallback.probe,
	)
	rectanchoredrectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rectanchoredrectFormCallback.CreationMode || rectanchoredrectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredrectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(rectanchoredrectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectAnchoredRectFormCallback(
			nil,
			rectanchoredrectFormCallback.probe,
			newFormGroup,
		)
		rectanchoredrect := new(models.RectAnchoredRect)
		FillUpForm(rectanchoredrect, newFormGroup, rectanchoredrectFormCallback.probe)
		rectanchoredrectFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(rectanchoredrectFormCallback.probe)
}
func __gong__New__RectAnchoredTextFormCallback(
	rectanchoredtext *models.RectAnchoredText,
	probe *Probe,
	formGroup *table.FormGroup,
) (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) {
	rectanchoredtextFormCallback = new(RectAnchoredTextFormCallback)
	rectanchoredtextFormCallback.probe = probe
	rectanchoredtextFormCallback.rectanchoredtext = rectanchoredtext
	rectanchoredtextFormCallback.formGroup = formGroup

	rectanchoredtextFormCallback.CreationMode = (rectanchoredtext == nil)

	return
}

type RectAnchoredTextFormCallback struct {
	rectanchoredtext *models.RectAnchoredText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) OnSave() {

	// log.Println("RectAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredtextFormCallback.probe.formStage.Checkout()

	if rectanchoredtextFormCallback.rectanchoredtext == nil {
		rectanchoredtextFormCallback.rectanchoredtext = new(models.RectAnchoredText).Stage(rectanchoredtextFormCallback.probe.stageOfInterest)
	}
	rectanchoredtext_ := rectanchoredtextFormCallback.rectanchoredtext
	_ = rectanchoredtext_

	for _, formDiv := range rectanchoredtextFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredtext_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(rectanchoredtext_.Content), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontSize), formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontStyle), formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&(rectanchoredtext_.LetterSpacing), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredtext_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredtext_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.RectAnchorType), formDiv)
		case "TextAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.TextAnchorType), formDiv)
		case "WritingMode":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.WritingMode), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredtext_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredtext_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredtext_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredtext_.Transform), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](rectanchoredtextFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rectanchoredtextFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			rectanchoredtext_.Animates = instanceSlice

		case "Rect:RectAnchoredTexts":
			// WARNING : this form deals with the N-N association "Rect.RectAnchoredTexts []*RectAnchoredText" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of RectAnchoredText). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Rect
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Rect"
				rf.Fieldname = "RectAnchoredTexts"
				formerAssociationSource := models.GetReverseFieldOwner(
					rectanchoredtextFormCallback.probe.stageOfInterest,
					rectanchoredtext_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Rect)
					if !ok {
						log.Fatalln("Source of Rect.RectAnchoredTexts []*RectAnchoredText, is not an Rect instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RectAnchoredTexts, rectanchoredtext_)
					formerSource.RectAnchoredTexts = slices.Delete(formerSource.RectAnchoredTexts, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Rect
			for _rect := range *models.GetGongstructInstancesSet[models.Rect](rectanchoredtextFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _rect.GetName() == newSourceName.GetName() {
					newSource = _rect // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Rect.RectAnchoredTexts []*RectAnchoredText, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RectAnchoredTexts = append(newSource.RectAnchoredTexts, rectanchoredtext_)
		}
	}

	// manage the suppress operation
	if rectanchoredtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredtext_.Unstage(rectanchoredtextFormCallback.probe.stageOfInterest)
	}

	rectanchoredtextFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.RectAnchoredText](
		rectanchoredtextFormCallback.probe,
	)
	rectanchoredtextFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rectanchoredtextFormCallback.CreationMode || rectanchoredtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectanchoredtextFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(rectanchoredtextFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectAnchoredTextFormCallback(
			nil,
			rectanchoredtextFormCallback.probe,
			newFormGroup,
		)
		rectanchoredtext := new(models.RectAnchoredText)
		FillUpForm(rectanchoredtext, newFormGroup, rectanchoredtextFormCallback.probe)
		rectanchoredtextFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(rectanchoredtextFormCallback.probe)
}
func __gong__New__RectLinkLinkFormCallback(
	rectlinklink *models.RectLinkLink,
	probe *Probe,
	formGroup *table.FormGroup,
) (rectlinklinkFormCallback *RectLinkLinkFormCallback) {
	rectlinklinkFormCallback = new(RectLinkLinkFormCallback)
	rectlinklinkFormCallback.probe = probe
	rectlinklinkFormCallback.rectlinklink = rectlinklink
	rectlinklinkFormCallback.formGroup = formGroup

	rectlinklinkFormCallback.CreationMode = (rectlinklink == nil)

	return
}

type RectLinkLinkFormCallback struct {
	rectlinklink *models.RectLinkLink

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rectlinklinkFormCallback *RectLinkLinkFormCallback) OnSave() {

	// log.Println("RectLinkLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectlinklinkFormCallback.probe.formStage.Checkout()

	if rectlinklinkFormCallback.rectlinklink == nil {
		rectlinklinkFormCallback.rectlinklink = new(models.RectLinkLink).Stage(rectlinklinkFormCallback.probe.stageOfInterest)
	}
	rectlinklink_ := rectlinklinkFormCallback.rectlinklink
	_ = rectlinklink_

	for _, formDiv := range rectlinklinkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectlinklink_.Name), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(rectlinklink_.Start), rectlinklinkFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(rectlinklink_.End), rectlinklinkFormCallback.probe.stageOfInterest, formDiv)
		case "TargetAnchorPosition":
			FormDivBasicFieldToField(&(rectlinklink_.TargetAnchorPosition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectlinklink_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectlinklink_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectlinklink_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectlinklink_.Transform), formDiv)
		case "Layer:RectLinkLinks":
			// WARNING : this form deals with the N-N association "Layer.RectLinkLinks []*RectLinkLink" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of RectLinkLink). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "RectLinkLinks"
				formerAssociationSource := models.GetReverseFieldOwner(
					rectlinklinkFormCallback.probe.stageOfInterest,
					rectlinklink_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.RectLinkLinks []*RectLinkLink, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RectLinkLinks, rectlinklink_)
					formerSource.RectLinkLinks = slices.Delete(formerSource.RectLinkLinks, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](rectlinklinkFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.RectLinkLinks []*RectLinkLink, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RectLinkLinks = append(newSource.RectLinkLinks, rectlinklink_)
		}
	}

	// manage the suppress operation
	if rectlinklinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectlinklink_.Unstage(rectlinklinkFormCallback.probe.stageOfInterest)
	}

	rectlinklinkFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.RectLinkLink](
		rectlinklinkFormCallback.probe,
	)
	rectlinklinkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rectlinklinkFormCallback.CreationMode || rectlinklinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rectlinklinkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(rectlinklinkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RectLinkLinkFormCallback(
			nil,
			rectlinklinkFormCallback.probe,
			newFormGroup,
		)
		rectlinklink := new(models.RectLinkLink)
		FillUpForm(rectlinklink, newFormGroup, rectlinklinkFormCallback.probe)
		rectlinklinkFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(rectlinklinkFormCallback.probe)
}
func __gong__New__SVGFormCallback(
	svg *models.SVG,
	probe *Probe,
	formGroup *table.FormGroup,
) (svgFormCallback *SVGFormCallback) {
	svgFormCallback = new(SVGFormCallback)
	svgFormCallback.probe = probe
	svgFormCallback.svg = svg
	svgFormCallback.formGroup = formGroup

	svgFormCallback.CreationMode = (svg == nil)

	return
}

type SVGFormCallback struct {
	svg *models.SVG

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (svgFormCallback *SVGFormCallback) OnSave() {

	// log.Println("SVGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgFormCallback.probe.formStage.Checkout()

	if svgFormCallback.svg == nil {
		svgFormCallback.svg = new(models.SVG).Stage(svgFormCallback.probe.stageOfInterest)
	}
	svg_ := svgFormCallback.svg
	_ = svg_

	for _, formDiv := range svgFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svg_.Name), formDiv)
		case "Layers":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Layer](svgFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Layer, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Layer)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					svgFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			svg_.Layers = instanceSlice

		case "DrawingState":
			FormDivEnumStringFieldToField(&(svg_.DrawingState), formDiv)
		case "StartRect":
			FormDivSelectFieldToField(&(svg_.StartRect), svgFormCallback.probe.stageOfInterest, formDiv)
		case "EndRect":
			FormDivSelectFieldToField(&(svg_.EndRect), svgFormCallback.probe.stageOfInterest, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&(svg_.IsEditable), formDiv)
		case "IsSVGFrontEndFileGenerated":
			FormDivBasicFieldToField(&(svg_.IsSVGFrontEndFileGenerated), formDiv)
		case "IsSVGBackEndFileGenerated":
			FormDivBasicFieldToField(&(svg_.IsSVGBackEndFileGenerated), formDiv)
		case "DefaultDirectoryForGeneratedImages":
			FormDivBasicFieldToField(&(svg_.DefaultDirectoryForGeneratedImages), formDiv)
		}
	}

	// manage the suppress operation
	if svgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svg_.Unstage(svgFormCallback.probe.stageOfInterest)
	}

	svgFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SVG](
		svgFormCallback.probe,
	)
	svgFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if svgFormCallback.CreationMode || svgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(svgFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SVGFormCallback(
			nil,
			svgFormCallback.probe,
			newFormGroup,
		)
		svg := new(models.SVG)
		FillUpForm(svg, newFormGroup, svgFormCallback.probe)
		svgFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(svgFormCallback.probe)
}
func __gong__New__SvgTextFormCallback(
	svgtext *models.SvgText,
	probe *Probe,
	formGroup *table.FormGroup,
) (svgtextFormCallback *SvgTextFormCallback) {
	svgtextFormCallback = new(SvgTextFormCallback)
	svgtextFormCallback.probe = probe
	svgtextFormCallback.svgtext = svgtext
	svgtextFormCallback.formGroup = formGroup

	svgtextFormCallback.CreationMode = (svgtext == nil)

	return
}

type SvgTextFormCallback struct {
	svgtext *models.SvgText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (svgtextFormCallback *SvgTextFormCallback) OnSave() {

	// log.Println("SvgTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgtextFormCallback.probe.formStage.Checkout()

	if svgtextFormCallback.svgtext == nil {
		svgtextFormCallback.svgtext = new(models.SvgText).Stage(svgtextFormCallback.probe.stageOfInterest)
	}
	svgtext_ := svgtextFormCallback.svgtext
	_ = svgtext_

	for _, formDiv := range svgtextFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svgtext_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(svgtext_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if svgtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgtext_.Unstage(svgtextFormCallback.probe.stageOfInterest)
	}

	svgtextFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SvgText](
		svgtextFormCallback.probe,
	)
	svgtextFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if svgtextFormCallback.CreationMode || svgtextFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgtextFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(svgtextFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SvgTextFormCallback(
			nil,
			svgtextFormCallback.probe,
			newFormGroup,
		)
		svgtext := new(models.SvgText)
		FillUpForm(svgtext, newFormGroup, svgtextFormCallback.probe)
		svgtextFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(svgtextFormCallback.probe)
}
func __gong__New__TextFormCallback(
	text *models.Text,
	probe *Probe,
	formGroup *table.FormGroup,
) (textFormCallback *TextFormCallback) {
	textFormCallback = new(TextFormCallback)
	textFormCallback.probe = probe
	textFormCallback.text = text
	textFormCallback.formGroup = formGroup

	textFormCallback.CreationMode = (text == nil)

	return
}

type TextFormCallback struct {
	text *models.Text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (textFormCallback *TextFormCallback) OnSave() {

	// log.Println("TextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	textFormCallback.probe.formStage.Checkout()

	if textFormCallback.text == nil {
		textFormCallback.text = new(models.Text).Stage(textFormCallback.probe.stageOfInterest)
	}
	text_ := textFormCallback.text
	_ = text_

	for _, formDiv := range textFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(text_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(text_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(text_.Y), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(text_.Content), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(text_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(text_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(text_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(text_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(text_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(text_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(text_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(text_.Transform), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(text_.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(text_.FontSize), formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&(text_.FontStyle), formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&(text_.LetterSpacing), formDiv)
		case "Animates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](textFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Animate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Animate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					textFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			text_.Animates = instanceSlice

		case "Layer:Texts":
			// WARNING : this form deals with the N-N association "Layer.Texts []*Text" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Text). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layer
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layer"
				rf.Fieldname = "Texts"
				formerAssociationSource := models.GetReverseFieldOwner(
					textFormCallback.probe.stageOfInterest,
					text_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layer)
					if !ok {
						log.Fatalln("Source of Layer.Texts []*Text, is not an Layer instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Texts, text_)
					formerSource.Texts = slices.Delete(formerSource.Texts, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layer
			for _layer := range *models.GetGongstructInstancesSet[models.Layer](textFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layer.GetName() == newSourceName.GetName() {
					newSource = _layer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layer.Texts []*Text, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Texts = append(newSource.Texts, text_)
		}
	}

	// manage the suppress operation
	if textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		text_.Unstage(textFormCallback.probe.stageOfInterest)
	}

	textFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Text](
		textFormCallback.probe,
	)
	textFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if textFormCallback.CreationMode || textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		textFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(textFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TextFormCallback(
			nil,
			textFormCallback.probe,
			newFormGroup,
		)
		text := new(models.Text)
		FillUpForm(text, newFormGroup, textFormCallback.probe)
		textFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(textFormCallback.probe)
}
