// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *AsSplitFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AsSplit", true)
			} else {
				FillUpFormFromGongstruct(onSave.assplit, probe)
			}
		case *AsSplitAreaFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AsSplitArea", true)
			} else {
				FillUpFormFromGongstruct(onSave.assplitarea, probe)
			}
		case *ButtonFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Button", true)
			} else {
				FillUpFormFromGongstruct(onSave.button, probe)
			}
		case *CursorFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Cursor", true)
			} else {
				FillUpFormFromGongstruct(onSave.cursor, probe)
			}
		case *FavIconFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "FavIcon", true)
			} else {
				FillUpFormFromGongstruct(onSave.favicon, probe)
			}
		case *FormFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Form", true)
			} else {
				FillUpFormFromGongstruct(onSave.form, probe)
			}
		case *LoadFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Load", true)
			} else {
				FillUpFormFromGongstruct(onSave.load, probe)
			}
		case *LogoOnTheLeftFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "LogoOnTheLeft", true)
			} else {
				FillUpFormFromGongstruct(onSave.logoontheleft, probe)
			}
		case *LogoOnTheRightFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "LogoOnTheRight", true)
			} else {
				FillUpFormFromGongstruct(onSave.logoontheright, probe)
			}
		case *MarkdownFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Markdown", true)
			} else {
				FillUpFormFromGongstruct(onSave.markdown, probe)
			}
		case *SliderFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Slider", true)
			} else {
				FillUpFormFromGongstruct(onSave.slider, probe)
			}
		case *SplitFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Split", true)
			} else {
				FillUpFormFromGongstruct(onSave.split, probe)
			}
		case *SvgFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Svg", true)
			} else {
				FillUpFormFromGongstruct(onSave.svg, probe)
			}
		case *TableFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Table", true)
			} else {
				FillUpFormFromGongstruct(onSave.table, probe)
			}
		case *TitleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Title", true)
			} else {
				FillUpFormFromGongstruct(onSave.title, probe)
			}
		case *ToneFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tone", true)
			} else {
				FillUpFormFromGongstruct(onSave.tone, probe)
			}
		case *TreeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tree", true)
			} else {
				FillUpFormFromGongstruct(onSave.tree, probe)
			}
		case *ViewFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "View", true)
			} else {
				FillUpFormFromGongstruct(onSave.view, probe)
			}
		case *XlsxFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Xlsx", true)
			} else {
				FillUpFormFromGongstruct(onSave.xlsx, probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "AsSplit":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AsSplit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AsSplitFormCallback(
			nil,
			probe,
			formGroup,
		)
		assplit := new(models.AsSplit)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(assplit, formGroup, probe)
	case "AsSplitArea":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AsSplitArea Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AsSplitAreaFormCallback(
			nil,
			probe,
			formGroup,
		)
		assplitarea := new(models.AsSplitArea)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(assplitarea, formGroup, probe)
	case "Button":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Button Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ButtonFormCallback(
			nil,
			probe,
			formGroup,
		)
		button := new(models.Button)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(button, formGroup, probe)
	case "Cursor":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Cursor Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CursorFormCallback(
			nil,
			probe,
			formGroup,
		)
		cursor := new(models.Cursor)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cursor, formGroup, probe)
	case "FavIcon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FavIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FavIconFormCallback(
			nil,
			probe,
			formGroup,
		)
		favicon := new(models.FavIcon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(favicon, formGroup, probe)
	case "Form":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Form Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFormCallback(
			nil,
			probe,
			formGroup,
		)
		form := new(models.Form)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(form, formGroup, probe)
	case "Load":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Load Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LoadFormCallback(
			nil,
			probe,
			formGroup,
		)
		load := new(models.Load)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(load, formGroup, probe)
	case "LogoOnTheLeft":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "LogoOnTheLeft Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LogoOnTheLeftFormCallback(
			nil,
			probe,
			formGroup,
		)
		logoontheleft := new(models.LogoOnTheLeft)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(logoontheleft, formGroup, probe)
	case "LogoOnTheRight":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "LogoOnTheRight Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LogoOnTheRightFormCallback(
			nil,
			probe,
			formGroup,
		)
		logoontheright := new(models.LogoOnTheRight)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(logoontheright, formGroup, probe)
	case "Markdown":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Markdown Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MarkdownFormCallback(
			nil,
			probe,
			formGroup,
		)
		markdown := new(models.Markdown)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(markdown, formGroup, probe)
	case "Slider":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Slider Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SliderFormCallback(
			nil,
			probe,
			formGroup,
		)
		slider := new(models.Slider)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(slider, formGroup, probe)
	case "Split":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Split Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SplitFormCallback(
			nil,
			probe,
			formGroup,
		)
		split := new(models.Split)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(split, formGroup, probe)
	case "Svg":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Svg Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SvgFormCallback(
			nil,
			probe,
			formGroup,
		)
		svg := new(models.Svg)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(svg, formGroup, probe)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Table Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableFormCallback(
			nil,
			probe,
			formGroup,
		)
		table := new(models.Table)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(table, formGroup, probe)
	case "Title":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Title Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TitleFormCallback(
			nil,
			probe,
			formGroup,
		)
		title := new(models.Title)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(title, formGroup, probe)
	case "Tone":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tone Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ToneFormCallback(
			nil,
			probe,
			formGroup,
		)
		tone := new(models.Tone)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tone, formGroup, probe)
	case "Tree":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tree Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TreeFormCallback(
			nil,
			probe,
			formGroup,
		)
		tree := new(models.Tree)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tree, formGroup, probe)
	case "View":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "View Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ViewFormCallback(
			nil,
			probe,
			formGroup,
		)
		view := new(models.View)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(view, formGroup, probe)
	case "Xlsx":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xlsx Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XlsxFormCallback(
			nil,
			probe,
			formGroup,
		)
		xlsx := new(models.Xlsx)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xlsx, formGroup, probe)
	}
	formStage.Commit()
}
