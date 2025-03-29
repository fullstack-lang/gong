// generated code - do not edit

import { AsSplitAreaAPI } from './assplitarea-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { AsSplit } from './assplit'
import { Button } from './button'
import { Cursor } from './cursor'
import { Doc } from './doc'
import { Form } from './form'
import { Load } from './load'
import { Slider } from './slider'
import { Split } from './split'
import { Svg } from './svg'
import { Table } from './table'
import { Tone } from './tone'
import { Tree } from './tree'
import { Xlsx } from './xlsx'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AsSplitArea {

	static GONGSTRUCT_NAME = "AsSplitArea"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	ShowNameInHeader: boolean = false
	Size: number = 0
	IsAny: boolean = false
	HasDiv: boolean = false
	DivStyle: string = ""

	// insertion point for pointers and slices of pointers declarations
	AsSplit?: AsSplit

	Button?: Button

	Cursor?: Cursor

	Doc?: Doc

	Form?: Form

	Load?: Load

	Slider?: Slider

	Split?: Split

	Svg?: Svg

	Table?: Table

	Tone?: Tone

	Tree?: Tree

	Xlsx?: Xlsx

}

export function CopyAsSplitAreaToAsSplitAreaAPI(assplitarea: AsSplitArea, assplitareaAPI: AsSplitAreaAPI) {

	assplitareaAPI.CreatedAt = assplitarea.CreatedAt
	assplitareaAPI.DeletedAt = assplitarea.DeletedAt
	assplitareaAPI.ID = assplitarea.ID

	// insertion point for basic fields copy operations
	assplitareaAPI.Name = assplitarea.Name
	assplitareaAPI.ShowNameInHeader = assplitarea.ShowNameInHeader
	assplitareaAPI.Size = assplitarea.Size
	assplitareaAPI.IsAny = assplitarea.IsAny
	assplitareaAPI.HasDiv = assplitarea.HasDiv
	assplitareaAPI.DivStyle = assplitarea.DivStyle

	// insertion point for pointer fields encoding
	assplitareaAPI.AsSplitAreaPointersEncoding.AsSplitID.Valid = true
	if (assplitarea.AsSplit != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.AsSplitID.Int64 = assplitarea.AsSplit.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.AsSplitID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.ButtonID.Valid = true
	if (assplitarea.Button != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.ButtonID.Int64 = assplitarea.Button.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.ButtonID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.CursorID.Valid = true
	if (assplitarea.Cursor != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.CursorID.Int64 = assplitarea.Cursor.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.CursorID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.DocID.Valid = true
	if (assplitarea.Doc != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.DocID.Int64 = assplitarea.Doc.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.DocID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.FormID.Valid = true
	if (assplitarea.Form != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.FormID.Int64 = assplitarea.Form.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.FormID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.LoadID.Valid = true
	if (assplitarea.Load != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.LoadID.Int64 = assplitarea.Load.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.LoadID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.SliderID.Valid = true
	if (assplitarea.Slider != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.SliderID.Int64 = assplitarea.Slider.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.SliderID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.SplitID.Valid = true
	if (assplitarea.Split != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.SplitID.Int64 = assplitarea.Split.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.SplitID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.SvgID.Valid = true
	if (assplitarea.Svg != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.SvgID.Int64 = assplitarea.Svg.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.SvgID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.TableID.Valid = true
	if (assplitarea.Table != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.TableID.Int64 = assplitarea.Table.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.TableID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.ToneID.Valid = true
	if (assplitarea.Tone != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.ToneID.Int64 = assplitarea.Tone.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.ToneID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.TreeID.Valid = true
	if (assplitarea.Tree != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.TreeID.Int64 = assplitarea.Tree.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.TreeID.Int64 = 0 		
	}

	assplitareaAPI.AsSplitAreaPointersEncoding.XlsxID.Valid = true
	if (assplitarea.Xlsx != undefined) {
		assplitareaAPI.AsSplitAreaPointersEncoding.XlsxID.Int64 = assplitarea.Xlsx.ID  
	} else {
		assplitareaAPI.AsSplitAreaPointersEncoding.XlsxID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyAsSplitAreaAPIToAsSplitArea update basic, pointers and slice of pointers fields of assplitarea
// from respectively the basic fields and encoded fields of pointers and slices of pointers of assplitareaAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI: AsSplitAreaAPI, assplitarea: AsSplitArea, frontRepo: FrontRepo) {

	assplitarea.CreatedAt = assplitareaAPI.CreatedAt
	assplitarea.DeletedAt = assplitareaAPI.DeletedAt
	assplitarea.ID = assplitareaAPI.ID

	// insertion point for basic fields copy operations
	assplitarea.Name = assplitareaAPI.Name
	assplitarea.ShowNameInHeader = assplitareaAPI.ShowNameInHeader
	assplitarea.Size = assplitareaAPI.Size
	assplitarea.IsAny = assplitareaAPI.IsAny
	assplitarea.HasDiv = assplitareaAPI.HasDiv
	assplitarea.DivStyle = assplitareaAPI.DivStyle

	// insertion point for pointer fields encoding
	assplitarea.AsSplit = frontRepo.map_ID_AsSplit.get(assplitareaAPI.AsSplitAreaPointersEncoding.AsSplitID.Int64)
	assplitarea.Button = frontRepo.map_ID_Button.get(assplitareaAPI.AsSplitAreaPointersEncoding.ButtonID.Int64)
	assplitarea.Cursor = frontRepo.map_ID_Cursor.get(assplitareaAPI.AsSplitAreaPointersEncoding.CursorID.Int64)
	assplitarea.Doc = frontRepo.map_ID_Doc.get(assplitareaAPI.AsSplitAreaPointersEncoding.DocID.Int64)
	assplitarea.Form = frontRepo.map_ID_Form.get(assplitareaAPI.AsSplitAreaPointersEncoding.FormID.Int64)
	assplitarea.Load = frontRepo.map_ID_Load.get(assplitareaAPI.AsSplitAreaPointersEncoding.LoadID.Int64)
	assplitarea.Slider = frontRepo.map_ID_Slider.get(assplitareaAPI.AsSplitAreaPointersEncoding.SliderID.Int64)
	assplitarea.Split = frontRepo.map_ID_Split.get(assplitareaAPI.AsSplitAreaPointersEncoding.SplitID.Int64)
	assplitarea.Svg = frontRepo.map_ID_Svg.get(assplitareaAPI.AsSplitAreaPointersEncoding.SvgID.Int64)
	assplitarea.Table = frontRepo.map_ID_Table.get(assplitareaAPI.AsSplitAreaPointersEncoding.TableID.Int64)
	assplitarea.Tone = frontRepo.map_ID_Tone.get(assplitareaAPI.AsSplitAreaPointersEncoding.ToneID.Int64)
	assplitarea.Tree = frontRepo.map_ID_Tree.get(assplitareaAPI.AsSplitAreaPointersEncoding.TreeID.Int64)
	assplitarea.Xlsx = frontRepo.map_ID_Xlsx.get(assplitareaAPI.AsSplitAreaPointersEncoding.XlsxID.Int64)

	// insertion point for slice of pointers fields encoding
}
