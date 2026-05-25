import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { CheckBoxAPI } from './checkbox-api'
import { CheckBox, CopyCheckBoxAPIToCheckBox } from './checkbox'
import { CheckBoxService } from './checkbox.service'

import { FormDivAPI } from './formdiv-api'
import { FormDiv, CopyFormDivAPIToFormDiv } from './formdiv'
import { FormDivService } from './formdiv.service'

import { FormEditAssocButtonAPI } from './formeditassocbutton-api'
import { FormEditAssocButton, CopyFormEditAssocButtonAPIToFormEditAssocButton } from './formeditassocbutton'
import { FormEditAssocButtonService } from './formeditassocbutton.service'

import { FormFieldAPI } from './formfield-api'
import { FormField, CopyFormFieldAPIToFormField } from './formfield'
import { FormFieldService } from './formfield.service'

import { FormFieldDateAPI } from './formfielddate-api'
import { FormFieldDate, CopyFormFieldDateAPIToFormFieldDate } from './formfielddate'
import { FormFieldDateService } from './formfielddate.service'

import { FormFieldDateTimeAPI } from './formfielddatetime-api'
import { FormFieldDateTime, CopyFormFieldDateTimeAPIToFormFieldDateTime } from './formfielddatetime'
import { FormFieldDateTimeService } from './formfielddatetime.service'

import { FormFieldFloat64API } from './formfieldfloat64-api'
import { FormFieldFloat64, CopyFormFieldFloat64APIToFormFieldFloat64 } from './formfieldfloat64'
import { FormFieldFloat64Service } from './formfieldfloat64.service'

import { FormFieldIntAPI } from './formfieldint-api'
import { FormFieldInt, CopyFormFieldIntAPIToFormFieldInt } from './formfieldint'
import { FormFieldIntService } from './formfieldint.service'

import { FormFieldSelectAPI } from './formfieldselect-api'
import { FormFieldSelect, CopyFormFieldSelectAPIToFormFieldSelect } from './formfieldselect'
import { FormFieldSelectService } from './formfieldselect.service'

import { FormFieldStringAPI } from './formfieldstring-api'
import { FormFieldString, CopyFormFieldStringAPIToFormFieldString } from './formfieldstring'
import { FormFieldStringService } from './formfieldstring.service'

import { FormFieldTimeAPI } from './formfieldtime-api'
import { FormFieldTime, CopyFormFieldTimeAPIToFormFieldTime } from './formfieldtime'
import { FormFieldTimeService } from './formfieldtime.service'

import { FormGroupAPI } from './formgroup-api'
import { FormGroup, CopyFormGroupAPIToFormGroup } from './formgroup'
import { FormGroupService } from './formgroup.service'

import { FormSortAssocButtonAPI } from './formsortassocbutton-api'
import { FormSortAssocButton, CopyFormSortAssocButtonAPIToFormSortAssocButton } from './formsortassocbutton'
import { FormSortAssocButtonService } from './formsortassocbutton.service'

import { OptionAPI } from './option-api'
import { Option, CopyOptionAPIToOption } from './option'
import { OptionService } from './option.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/form/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_CheckBoxs = new Array<CheckBox>() // array of front instances
	map_ID_CheckBox = new Map<number, CheckBox>() // map of front instances

	array_FormDivs = new Array<FormDiv>() // array of front instances
	map_ID_FormDiv = new Map<number, FormDiv>() // map of front instances

	array_FormEditAssocButtons = new Array<FormEditAssocButton>() // array of front instances
	map_ID_FormEditAssocButton = new Map<number, FormEditAssocButton>() // map of front instances

	array_FormFields = new Array<FormField>() // array of front instances
	map_ID_FormField = new Map<number, FormField>() // map of front instances

	array_FormFieldDates = new Array<FormFieldDate>() // array of front instances
	map_ID_FormFieldDate = new Map<number, FormFieldDate>() // map of front instances

	array_FormFieldDateTimes = new Array<FormFieldDateTime>() // array of front instances
	map_ID_FormFieldDateTime = new Map<number, FormFieldDateTime>() // map of front instances

	array_FormFieldFloat64s = new Array<FormFieldFloat64>() // array of front instances
	map_ID_FormFieldFloat64 = new Map<number, FormFieldFloat64>() // map of front instances

	array_FormFieldInts = new Array<FormFieldInt>() // array of front instances
	map_ID_FormFieldInt = new Map<number, FormFieldInt>() // map of front instances

	array_FormFieldSelects = new Array<FormFieldSelect>() // array of front instances
	map_ID_FormFieldSelect = new Map<number, FormFieldSelect>() // map of front instances

	array_FormFieldStrings = new Array<FormFieldString>() // array of front instances
	map_ID_FormFieldString = new Map<number, FormFieldString>() // map of front instances

	array_FormFieldTimes = new Array<FormFieldTime>() // array of front instances
	map_ID_FormFieldTime = new Map<number, FormFieldTime>() // map of front instances

	array_FormGroups = new Array<FormGroup>() // array of front instances
	map_ID_FormGroup = new Map<number, FormGroup>() // map of front instances

	array_FormSortAssocButtons = new Array<FormSortAssocButton>() // array of front instances
	map_ID_FormSortAssocButton = new Map<number, FormSortAssocButton>() // map of front instances

	array_Options = new Array<Option>() // array of front instances
	map_ID_Option = new Map<number, Option>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'CheckBox':
				return this.array_CheckBoxs as unknown as Array<Type>
			case 'FormDiv':
				return this.array_FormDivs as unknown as Array<Type>
			case 'FormEditAssocButton':
				return this.array_FormEditAssocButtons as unknown as Array<Type>
			case 'FormField':
				return this.array_FormFields as unknown as Array<Type>
			case 'FormFieldDate':
				return this.array_FormFieldDates as unknown as Array<Type>
			case 'FormFieldDateTime':
				return this.array_FormFieldDateTimes as unknown as Array<Type>
			case 'FormFieldFloat64':
				return this.array_FormFieldFloat64s as unknown as Array<Type>
			case 'FormFieldInt':
				return this.array_FormFieldInts as unknown as Array<Type>
			case 'FormFieldSelect':
				return this.array_FormFieldSelects as unknown as Array<Type>
			case 'FormFieldString':
				return this.array_FormFieldStrings as unknown as Array<Type>
			case 'FormFieldTime':
				return this.array_FormFieldTimes as unknown as Array<Type>
			case 'FormGroup':
				return this.array_FormGroups as unknown as Array<Type>
			case 'FormSortAssocButton':
				return this.array_FormSortAssocButtons as unknown as Array<Type>
			case 'Option':
				return this.array_Options as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'CheckBox':
				return this.map_ID_CheckBox as unknown as Map<number, Type>
			case 'FormDiv':
				return this.map_ID_FormDiv as unknown as Map<number, Type>
			case 'FormEditAssocButton':
				return this.map_ID_FormEditAssocButton as unknown as Map<number, Type>
			case 'FormField':
				return this.map_ID_FormField as unknown as Map<number, Type>
			case 'FormFieldDate':
				return this.map_ID_FormFieldDate as unknown as Map<number, Type>
			case 'FormFieldDateTime':
				return this.map_ID_FormFieldDateTime as unknown as Map<number, Type>
			case 'FormFieldFloat64':
				return this.map_ID_FormFieldFloat64 as unknown as Map<number, Type>
			case 'FormFieldInt':
				return this.map_ID_FormFieldInt as unknown as Map<number, Type>
			case 'FormFieldSelect':
				return this.map_ID_FormFieldSelect as unknown as Map<number, Type>
			case 'FormFieldString':
				return this.map_ID_FormFieldString as unknown as Map<number, Type>
			case 'FormFieldTime':
				return this.map_ID_FormFieldTime as unknown as Map<number, Type>
			case 'FormGroup':
				return this.map_ID_FormGroup as unknown as Map<number, Type>
			case 'FormSortAssocButton':
				return this.map_ID_FormSortAssocButton as unknown as Map<number, Type>
			case 'Option':
				return this.map_ID_Option as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized")
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	Name: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	Name: string = ""

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	}

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	// Manage open WebSocket connections
	private webSocketConnections = new Map<string, Observable<FrontRepo>>()


	constructor(
		private http: HttpClient, // insertion point sub template 
		private checkboxService: CheckBoxService,
		private formdivService: FormDivService,
		private formeditassocbuttonService: FormEditAssocButtonService,
		private formfieldService: FormFieldService,
		private formfielddateService: FormFieldDateService,
		private formfielddatetimeService: FormFieldDateTimeService,
		private formfieldfloat64Service: FormFieldFloat64Service,
		private formfieldintService: FormFieldIntService,
		private formfieldselectService: FormFieldSelectService,
		private formfieldstringService: FormFieldStringService,
		private formfieldtimeService: FormFieldTimeService,
		private formgroupService: FormGroupService,
		private formsortassocbuttonService: FormSortAssocButtonService,
		private optionService: OptionService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		)
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		)
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo!: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<CheckBoxAPI[]>,
		Observable<FormDivAPI[]>,
		Observable<FormEditAssocButtonAPI[]>,
		Observable<FormFieldAPI[]>,
		Observable<FormFieldDateAPI[]>,
		Observable<FormFieldDateTimeAPI[]>,
		Observable<FormFieldFloat64API[]>,
		Observable<FormFieldIntAPI[]>,
		Observable<FormFieldSelectAPI[]>,
		Observable<FormFieldStringAPI[]>,
		Observable<FormFieldTimeAPI[]>,
		Observable<FormGroupAPI[]>,
		Observable<FormSortAssocButtonAPI[]>,
		Observable<OptionAPI[]>,
	]

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(Name: string = ""): Observable<FrontRepo> {

		this.Name = Name

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.checkboxService.getCheckBoxs(this.Name, this.frontRepo),
			this.formdivService.getFormDivs(this.Name, this.frontRepo),
			this.formeditassocbuttonService.getFormEditAssocButtons(this.Name, this.frontRepo),
			this.formfieldService.getFormFields(this.Name, this.frontRepo),
			this.formfielddateService.getFormFieldDates(this.Name, this.frontRepo),
			this.formfielddatetimeService.getFormFieldDateTimes(this.Name, this.frontRepo),
			this.formfieldfloat64Service.getFormFieldFloat64s(this.Name, this.frontRepo),
			this.formfieldintService.getFormFieldInts(this.Name, this.frontRepo),
			this.formfieldselectService.getFormFieldSelects(this.Name, this.frontRepo),
			this.formfieldstringService.getFormFieldStrings(this.Name, this.frontRepo),
			this.formfieldtimeService.getFormFieldTimes(this.Name, this.frontRepo),
			this.formgroupService.getFormGroups(this.Name, this.frontRepo),
			this.formsortassocbuttonService.getFormSortAssocButtons(this.Name, this.frontRepo),
			this.optionService.getOptions(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						checkboxs_,
						formdivs_,
						formeditassocbuttons_,
						formfields_,
						formfielddates_,
						formfielddatetimes_,
						formfieldfloat64s_,
						formfieldints_,
						formfieldselects_,
						formfieldstrings_,
						formfieldtimes_,
						formgroups_,
						formsortassocbuttons_,
						options_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var checkboxs: CheckBoxAPI[]
						checkboxs = checkboxs_ as CheckBoxAPI[]
						var formdivs: FormDivAPI[]
						formdivs = formdivs_ as FormDivAPI[]
						var formeditassocbuttons: FormEditAssocButtonAPI[]
						formeditassocbuttons = formeditassocbuttons_ as FormEditAssocButtonAPI[]
						var formfields: FormFieldAPI[]
						formfields = formfields_ as FormFieldAPI[]
						var formfielddates: FormFieldDateAPI[]
						formfielddates = formfielddates_ as FormFieldDateAPI[]
						var formfielddatetimes: FormFieldDateTimeAPI[]
						formfielddatetimes = formfielddatetimes_ as FormFieldDateTimeAPI[]
						var formfieldfloat64s: FormFieldFloat64API[]
						formfieldfloat64s = formfieldfloat64s_ as FormFieldFloat64API[]
						var formfieldints: FormFieldIntAPI[]
						formfieldints = formfieldints_ as FormFieldIntAPI[]
						var formfieldselects: FormFieldSelectAPI[]
						formfieldselects = formfieldselects_ as FormFieldSelectAPI[]
						var formfieldstrings: FormFieldStringAPI[]
						formfieldstrings = formfieldstrings_ as FormFieldStringAPI[]
						var formfieldtimes: FormFieldTimeAPI[]
						formfieldtimes = formfieldtimes_ as FormFieldTimeAPI[]
						var formgroups: FormGroupAPI[]
						formgroups = formgroups_ as FormGroupAPI[]
						var formsortassocbuttons: FormSortAssocButtonAPI[]
						formsortassocbuttons = formsortassocbuttons_ as FormSortAssocButtonAPI[]
						var options: OptionAPI[]
						options = options_ as OptionAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_CheckBoxs = []
						this.frontRepo.map_ID_CheckBox.clear()

						checkboxs.forEach(
							checkboxAPI => {
								let checkbox = new CheckBox
								this.frontRepo.array_CheckBoxs.push(checkbox)
								this.frontRepo.map_ID_CheckBox.set(checkboxAPI.ID, checkbox)
							}
						)

						// init the arrays
						this.frontRepo.array_FormDivs = []
						this.frontRepo.map_ID_FormDiv.clear()

						formdivs.forEach(
							formdivAPI => {
								let formdiv = new FormDiv
								this.frontRepo.array_FormDivs.push(formdiv)
								this.frontRepo.map_ID_FormDiv.set(formdivAPI.ID, formdiv)
							}
						)

						// init the arrays
						this.frontRepo.array_FormEditAssocButtons = []
						this.frontRepo.map_ID_FormEditAssocButton.clear()

						formeditassocbuttons.forEach(
							formeditassocbuttonAPI => {
								let formeditassocbutton = new FormEditAssocButton
								this.frontRepo.array_FormEditAssocButtons.push(formeditassocbutton)
								this.frontRepo.map_ID_FormEditAssocButton.set(formeditassocbuttonAPI.ID, formeditassocbutton)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFields = []
						this.frontRepo.map_ID_FormField.clear()

						formfields.forEach(
							formfieldAPI => {
								let formfield = new FormField
								this.frontRepo.array_FormFields.push(formfield)
								this.frontRepo.map_ID_FormField.set(formfieldAPI.ID, formfield)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldDates = []
						this.frontRepo.map_ID_FormFieldDate.clear()

						formfielddates.forEach(
							formfielddateAPI => {
								let formfielddate = new FormFieldDate
								this.frontRepo.array_FormFieldDates.push(formfielddate)
								this.frontRepo.map_ID_FormFieldDate.set(formfielddateAPI.ID, formfielddate)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldDateTimes = []
						this.frontRepo.map_ID_FormFieldDateTime.clear()

						formfielddatetimes.forEach(
							formfielddatetimeAPI => {
								let formfielddatetime = new FormFieldDateTime
								this.frontRepo.array_FormFieldDateTimes.push(formfielddatetime)
								this.frontRepo.map_ID_FormFieldDateTime.set(formfielddatetimeAPI.ID, formfielddatetime)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldFloat64s = []
						this.frontRepo.map_ID_FormFieldFloat64.clear()

						formfieldfloat64s.forEach(
							formfieldfloat64API => {
								let formfieldfloat64 = new FormFieldFloat64
								this.frontRepo.array_FormFieldFloat64s.push(formfieldfloat64)
								this.frontRepo.map_ID_FormFieldFloat64.set(formfieldfloat64API.ID, formfieldfloat64)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldInts = []
						this.frontRepo.map_ID_FormFieldInt.clear()

						formfieldints.forEach(
							formfieldintAPI => {
								let formfieldint = new FormFieldInt
								this.frontRepo.array_FormFieldInts.push(formfieldint)
								this.frontRepo.map_ID_FormFieldInt.set(formfieldintAPI.ID, formfieldint)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldSelects = []
						this.frontRepo.map_ID_FormFieldSelect.clear()

						formfieldselects.forEach(
							formfieldselectAPI => {
								let formfieldselect = new FormFieldSelect
								this.frontRepo.array_FormFieldSelects.push(formfieldselect)
								this.frontRepo.map_ID_FormFieldSelect.set(formfieldselectAPI.ID, formfieldselect)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldStrings = []
						this.frontRepo.map_ID_FormFieldString.clear()

						formfieldstrings.forEach(
							formfieldstringAPI => {
								let formfieldstring = new FormFieldString
								this.frontRepo.array_FormFieldStrings.push(formfieldstring)
								this.frontRepo.map_ID_FormFieldString.set(formfieldstringAPI.ID, formfieldstring)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldTimes = []
						this.frontRepo.map_ID_FormFieldTime.clear()

						formfieldtimes.forEach(
							formfieldtimeAPI => {
								let formfieldtime = new FormFieldTime
								this.frontRepo.array_FormFieldTimes.push(formfieldtime)
								this.frontRepo.map_ID_FormFieldTime.set(formfieldtimeAPI.ID, formfieldtime)
							}
						)

						// init the arrays
						this.frontRepo.array_FormGroups = []
						this.frontRepo.map_ID_FormGroup.clear()

						formgroups.forEach(
							formgroupAPI => {
								let formgroup = new FormGroup
								this.frontRepo.array_FormGroups.push(formgroup)
								this.frontRepo.map_ID_FormGroup.set(formgroupAPI.ID, formgroup)
							}
						)

						// init the arrays
						this.frontRepo.array_FormSortAssocButtons = []
						this.frontRepo.map_ID_FormSortAssocButton.clear()

						formsortassocbuttons.forEach(
							formsortassocbuttonAPI => {
								let formsortassocbutton = new FormSortAssocButton
								this.frontRepo.array_FormSortAssocButtons.push(formsortassocbutton)
								this.frontRepo.map_ID_FormSortAssocButton.set(formsortassocbuttonAPI.ID, formsortassocbutton)
							}
						)

						// init the arrays
						this.frontRepo.array_Options = []
						this.frontRepo.map_ID_Option.clear()

						options.forEach(
							optionAPI => {
								let option = new Option
								this.frontRepo.array_Options.push(option)
								this.frontRepo.map_ID_Option.set(optionAPI.ID, option)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						checkboxs.forEach(
							checkboxAPI => {
								let checkbox = this.frontRepo.map_ID_CheckBox.get(checkboxAPI.ID)
								CopyCheckBoxAPIToCheckBox(checkboxAPI, checkbox!, this.frontRepo)
							}
						)

						// fill up front objects
						formdivs.forEach(
							formdivAPI => {
								let formdiv = this.frontRepo.map_ID_FormDiv.get(formdivAPI.ID)
								CopyFormDivAPIToFormDiv(formdivAPI, formdiv!, this.frontRepo)
							}
						)

						// fill up front objects
						formeditassocbuttons.forEach(
							formeditassocbuttonAPI => {
								let formeditassocbutton = this.frontRepo.map_ID_FormEditAssocButton.get(formeditassocbuttonAPI.ID)
								CopyFormEditAssocButtonAPIToFormEditAssocButton(formeditassocbuttonAPI, formeditassocbutton!, this.frontRepo)
							}
						)

						// fill up front objects
						formfields.forEach(
							formfieldAPI => {
								let formfield = this.frontRepo.map_ID_FormField.get(formfieldAPI.ID)
								CopyFormFieldAPIToFormField(formfieldAPI, formfield!, this.frontRepo)
							}
						)

						// fill up front objects
						formfielddates.forEach(
							formfielddateAPI => {
								let formfielddate = this.frontRepo.map_ID_FormFieldDate.get(formfielddateAPI.ID)
								CopyFormFieldDateAPIToFormFieldDate(formfielddateAPI, formfielddate!, this.frontRepo)
							}
						)

						// fill up front objects
						formfielddatetimes.forEach(
							formfielddatetimeAPI => {
								let formfielddatetime = this.frontRepo.map_ID_FormFieldDateTime.get(formfielddatetimeAPI.ID)
								CopyFormFieldDateTimeAPIToFormFieldDateTime(formfielddatetimeAPI, formfielddatetime!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldfloat64s.forEach(
							formfieldfloat64API => {
								let formfieldfloat64 = this.frontRepo.map_ID_FormFieldFloat64.get(formfieldfloat64API.ID)
								CopyFormFieldFloat64APIToFormFieldFloat64(formfieldfloat64API, formfieldfloat64!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldints.forEach(
							formfieldintAPI => {
								let formfieldint = this.frontRepo.map_ID_FormFieldInt.get(formfieldintAPI.ID)
								CopyFormFieldIntAPIToFormFieldInt(formfieldintAPI, formfieldint!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldselects.forEach(
							formfieldselectAPI => {
								let formfieldselect = this.frontRepo.map_ID_FormFieldSelect.get(formfieldselectAPI.ID)
								CopyFormFieldSelectAPIToFormFieldSelect(formfieldselectAPI, formfieldselect!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldstrings.forEach(
							formfieldstringAPI => {
								let formfieldstring = this.frontRepo.map_ID_FormFieldString.get(formfieldstringAPI.ID)
								CopyFormFieldStringAPIToFormFieldString(formfieldstringAPI, formfieldstring!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldtimes.forEach(
							formfieldtimeAPI => {
								let formfieldtime = this.frontRepo.map_ID_FormFieldTime.get(formfieldtimeAPI.ID)
								CopyFormFieldTimeAPIToFormFieldTime(formfieldtimeAPI, formfieldtime!, this.frontRepo)
							}
						)

						// fill up front objects
						formgroups.forEach(
							formgroupAPI => {
								let formgroup = this.frontRepo.map_ID_FormGroup.get(formgroupAPI.ID)
								CopyFormGroupAPIToFormGroup(formgroupAPI, formgroup!, this.frontRepo)
							}
						)

						// fill up front objects
						formsortassocbuttons.forEach(
							formsortassocbuttonAPI => {
								let formsortassocbutton = this.frontRepo.map_ID_FormSortAssocButton.get(formsortassocbuttonAPI.ID)
								CopyFormSortAssocButtonAPIToFormSortAssocButton(formsortassocbuttonAPI, formsortassocbutton!, this.frontRepo)
							}
						)

						// fill up front objects
						options.forEach(
							optionAPI => {
								let option = this.frontRepo.map_ID_Option.get(optionAPI.ID)
								CopyOptionAPIToOption(optionAPI, option!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(Name: string): Observable<FrontRepo> {

		console.log("github.com/fullstack-lang/gong/lib/form/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			console.log("github.com/fullstack-lang/gong/lib/form/go; connectToWebSocket: returning existing connection")
			return this.webSocketConnections.get(Name)!
		}

		//
		// Create a new connection
		//
		let host = window.location.host
		const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'

		if (host === 'localhost:4200') {
			host = 'localhost:8080'
		}

		// Construct the base path using the dynamic host and protocol
		// The API path remains the same.
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/form/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			console.log("github.com/fullstack-lang/gong/lib/form/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:'

			const processData = (dataString: string) => {
				console.log("github.com/fullstack-lang/gong/lib/form/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_CheckBoxs = []
				frontRepo.map_ID_CheckBox.clear()

				backRepoData.CheckBoxAPIs.forEach(
					checkboxAPI => {
						let checkbox = new CheckBox
						frontRepo.array_CheckBoxs.push(checkbox)
						frontRepo.map_ID_CheckBox.set(checkboxAPI.ID, checkbox)
					}
				)

				// init the arrays
				frontRepo.array_FormDivs = []
				frontRepo.map_ID_FormDiv.clear()

				backRepoData.FormDivAPIs.forEach(
					formdivAPI => {
						let formdiv = new FormDiv
						frontRepo.array_FormDivs.push(formdiv)
						frontRepo.map_ID_FormDiv.set(formdivAPI.ID, formdiv)
					}
				)

				// init the arrays
				frontRepo.array_FormEditAssocButtons = []
				frontRepo.map_ID_FormEditAssocButton.clear()

				backRepoData.FormEditAssocButtonAPIs.forEach(
					formeditassocbuttonAPI => {
						let formeditassocbutton = new FormEditAssocButton
						frontRepo.array_FormEditAssocButtons.push(formeditassocbutton)
						frontRepo.map_ID_FormEditAssocButton.set(formeditassocbuttonAPI.ID, formeditassocbutton)
					}
				)

				// init the arrays
				frontRepo.array_FormFields = []
				frontRepo.map_ID_FormField.clear()

				backRepoData.FormFieldAPIs.forEach(
					formfieldAPI => {
						let formfield = new FormField
						frontRepo.array_FormFields.push(formfield)
						frontRepo.map_ID_FormField.set(formfieldAPI.ID, formfield)
					}
				)

				// init the arrays
				frontRepo.array_FormFieldDates = []
				frontRepo.map_ID_FormFieldDate.clear()

				backRepoData.FormFieldDateAPIs.forEach(
					formfielddateAPI => {
						let formfielddate = new FormFieldDate
						frontRepo.array_FormFieldDates.push(formfielddate)
						frontRepo.map_ID_FormFieldDate.set(formfielddateAPI.ID, formfielddate)
					}
				)

				// init the arrays
				frontRepo.array_FormFieldDateTimes = []
				frontRepo.map_ID_FormFieldDateTime.clear()

				backRepoData.FormFieldDateTimeAPIs.forEach(
					formfielddatetimeAPI => {
						let formfielddatetime = new FormFieldDateTime
						frontRepo.array_FormFieldDateTimes.push(formfielddatetime)
						frontRepo.map_ID_FormFieldDateTime.set(formfielddatetimeAPI.ID, formfielddatetime)
					}
				)

				// init the arrays
				frontRepo.array_FormFieldFloat64s = []
				frontRepo.map_ID_FormFieldFloat64.clear()

				backRepoData.FormFieldFloat64APIs.forEach(
					formfieldfloat64API => {
						let formfieldfloat64 = new FormFieldFloat64
						frontRepo.array_FormFieldFloat64s.push(formfieldfloat64)
						frontRepo.map_ID_FormFieldFloat64.set(formfieldfloat64API.ID, formfieldfloat64)
					}
				)

				// init the arrays
				frontRepo.array_FormFieldInts = []
				frontRepo.map_ID_FormFieldInt.clear()

				backRepoData.FormFieldIntAPIs.forEach(
					formfieldintAPI => {
						let formfieldint = new FormFieldInt
						frontRepo.array_FormFieldInts.push(formfieldint)
						frontRepo.map_ID_FormFieldInt.set(formfieldintAPI.ID, formfieldint)
					}
				)

				// init the arrays
				frontRepo.array_FormFieldSelects = []
				frontRepo.map_ID_FormFieldSelect.clear()

				backRepoData.FormFieldSelectAPIs.forEach(
					formfieldselectAPI => {
						let formfieldselect = new FormFieldSelect
						frontRepo.array_FormFieldSelects.push(formfieldselect)
						frontRepo.map_ID_FormFieldSelect.set(formfieldselectAPI.ID, formfieldselect)
					}
				)

				// init the arrays
				frontRepo.array_FormFieldStrings = []
				frontRepo.map_ID_FormFieldString.clear()

				backRepoData.FormFieldStringAPIs.forEach(
					formfieldstringAPI => {
						let formfieldstring = new FormFieldString
						frontRepo.array_FormFieldStrings.push(formfieldstring)
						frontRepo.map_ID_FormFieldString.set(formfieldstringAPI.ID, formfieldstring)
					}
				)

				// init the arrays
				frontRepo.array_FormFieldTimes = []
				frontRepo.map_ID_FormFieldTime.clear()

				backRepoData.FormFieldTimeAPIs.forEach(
					formfieldtimeAPI => {
						let formfieldtime = new FormFieldTime
						frontRepo.array_FormFieldTimes.push(formfieldtime)
						frontRepo.map_ID_FormFieldTime.set(formfieldtimeAPI.ID, formfieldtime)
					}
				)

				// init the arrays
				frontRepo.array_FormGroups = []
				frontRepo.map_ID_FormGroup.clear()

				backRepoData.FormGroupAPIs.forEach(
					formgroupAPI => {
						let formgroup = new FormGroup
						frontRepo.array_FormGroups.push(formgroup)
						frontRepo.map_ID_FormGroup.set(formgroupAPI.ID, formgroup)
					}
				)

				// init the arrays
				frontRepo.array_FormSortAssocButtons = []
				frontRepo.map_ID_FormSortAssocButton.clear()

				backRepoData.FormSortAssocButtonAPIs.forEach(
					formsortassocbuttonAPI => {
						let formsortassocbutton = new FormSortAssocButton
						frontRepo.array_FormSortAssocButtons.push(formsortassocbutton)
						frontRepo.map_ID_FormSortAssocButton.set(formsortassocbuttonAPI.ID, formsortassocbutton)
					}
				)

				// init the arrays
				frontRepo.array_Options = []
				frontRepo.map_ID_Option.clear()

				backRepoData.OptionAPIs.forEach(
					optionAPI => {
						let option = new Option
						frontRepo.array_Options.push(option)
						frontRepo.map_ID_Option.set(optionAPI.ID, option)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.CheckBoxAPIs.forEach(
					checkboxAPI => {
						let checkbox = frontRepo.map_ID_CheckBox.get(checkboxAPI.ID)
						CopyCheckBoxAPIToCheckBox(checkboxAPI, checkbox!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormDivAPIs.forEach(
					formdivAPI => {
						let formdiv = frontRepo.map_ID_FormDiv.get(formdivAPI.ID)
						CopyFormDivAPIToFormDiv(formdivAPI, formdiv!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormEditAssocButtonAPIs.forEach(
					formeditassocbuttonAPI => {
						let formeditassocbutton = frontRepo.map_ID_FormEditAssocButton.get(formeditassocbuttonAPI.ID)
						CopyFormEditAssocButtonAPIToFormEditAssocButton(formeditassocbuttonAPI, formeditassocbutton!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldAPIs.forEach(
					formfieldAPI => {
						let formfield = frontRepo.map_ID_FormField.get(formfieldAPI.ID)
						CopyFormFieldAPIToFormField(formfieldAPI, formfield!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldDateAPIs.forEach(
					formfielddateAPI => {
						let formfielddate = frontRepo.map_ID_FormFieldDate.get(formfielddateAPI.ID)
						CopyFormFieldDateAPIToFormFieldDate(formfielddateAPI, formfielddate!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldDateTimeAPIs.forEach(
					formfielddatetimeAPI => {
						let formfielddatetime = frontRepo.map_ID_FormFieldDateTime.get(formfielddatetimeAPI.ID)
						CopyFormFieldDateTimeAPIToFormFieldDateTime(formfielddatetimeAPI, formfielddatetime!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldFloat64APIs.forEach(
					formfieldfloat64API => {
						let formfieldfloat64 = frontRepo.map_ID_FormFieldFloat64.get(formfieldfloat64API.ID)
						CopyFormFieldFloat64APIToFormFieldFloat64(formfieldfloat64API, formfieldfloat64!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldIntAPIs.forEach(
					formfieldintAPI => {
						let formfieldint = frontRepo.map_ID_FormFieldInt.get(formfieldintAPI.ID)
						CopyFormFieldIntAPIToFormFieldInt(formfieldintAPI, formfieldint!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldSelectAPIs.forEach(
					formfieldselectAPI => {
						let formfieldselect = frontRepo.map_ID_FormFieldSelect.get(formfieldselectAPI.ID)
						CopyFormFieldSelectAPIToFormFieldSelect(formfieldselectAPI, formfieldselect!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldStringAPIs.forEach(
					formfieldstringAPI => {
						let formfieldstring = frontRepo.map_ID_FormFieldString.get(formfieldstringAPI.ID)
						CopyFormFieldStringAPIToFormFieldString(formfieldstringAPI, formfieldstring!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldTimeAPIs.forEach(
					formfieldtimeAPI => {
						let formfieldtime = frontRepo.map_ID_FormFieldTime.get(formfieldtimeAPI.ID)
						CopyFormFieldTimeAPIToFormFieldTime(formfieldtimeAPI, formfieldtime!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormGroupAPIs.forEach(
					formgroupAPI => {
						let formgroup = frontRepo.map_ID_FormGroup.get(formgroupAPI.ID)
						CopyFormGroupAPIToFormGroup(formgroupAPI, formgroup!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormSortAssocButtonAPIs.forEach(
					formsortassocbuttonAPI => {
						let formsortassocbutton = frontRepo.map_ID_FormSortAssocButton.get(formsortassocbuttonAPI.ID)
						CopyFormSortAssocButtonAPIToFormSortAssocButton(formsortassocbuttonAPI, formsortassocbutton!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.OptionAPIs.forEach(
					optionAPI => {
						let option = frontRepo.map_ID_Option.get(optionAPI.ID)
						CopyOptionAPIToOption(optionAPI, option!, frontRepo)
					}
				)


				observer.next(frontRepo)
			}

			// 3. Connection Loop
			const attemptConnection = (retries: number): void => {
				console.log("github.com/fullstack-lang/gong/lib/form/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)

				// A. WASM OFFLINE MODE (Check if Go is ready)
				if ((window as any).openWasmSocket) {
					console.log("github.com/fullstack-lang/gong/lib/form/go; attemptConnection: openWasmSocket exists, calling it");
					(window as any).openWasmSocket(Name, processData);
					return;
				}

				// B. WAITING FOR WASM
				if (isOfflineMode && retries > 0) {
					console.log("github.com/fullstack-lang/gong/lib/form/go; attemptConnection: WAITING FOR WASM. Retries left:", retries)
					setTimeout(() => attemptConnection(retries - 1), 100);
					return;
				}

				// C. STANDARD SERVER MODE
				if (!isOfflineMode) {
					console.log("github.com/fullstack-lang/gong/lib/form/go; attemptConnection: STANDARD SERVER MODE. url =", url)
					socket = new WebSocket(url)
					socket.onopen = (event) => {
						console.log("github.com/fullstack-lang/gong/lib/form/go; WebSocket: onopen", event)
					}
					socket.onmessage = event => {
						console.log("github.com/fullstack-lang/gong/lib/form/go; WebSocket: onmessage")
						processData(event.data)
					}
					socket.onerror = event => {
						console.error("github.com/fullstack-lang/gong/lib/form/go WebSocket: onerror", event)
						observer.error(event)
					}
					socket.onclose = (event) => {
						console.log("github.com/fullstack-lang/gong/lib/form/go; WebSocket: onclose", event)
						observer.complete()
					}
				} else {
					console.error("github.com/fullstack-lang/gong/lib/form/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.");
				}
			};

			attemptConnection(50);

			// Teardown logic: Called when the last subscriber unsubscribes.
			return () => {
				this.webSocketConnections.delete(Name) // Remove from cache
				if (socket) {
					socket.close()
				}
			}
		}).pipe(
			// This is the key:
			// - shareReplay makes this a "multicast" observable, sharing the single WebSocket among subscribers.
			// - { bufferSize: 1, refCount: true } means:
			//   - bufferSize: 1 => new subscribers get the last emitted value immediately.
			//   - refCount: true => the connection starts with the first subscriber and stops with the last.
			shareReplay({ bufferSize: 1, refCount: true })
		)

		// Store the new connection observable in the map
		this.webSocketConnections.set(Name, newConnection$)
		return newConnection$
	}
}

// insertion point for get unique ID per struct 
export function getCheckBoxUniqueID(id: number): number {
	return 31 * id
}
export function getFormDivUniqueID(id: number): number {
	return 37 * id
}
export function getFormEditAssocButtonUniqueID(id: number): number {
	return 41 * id
}
export function getFormFieldUniqueID(id: number): number {
	return 43 * id
}
export function getFormFieldDateUniqueID(id: number): number {
	return 47 * id
}
export function getFormFieldDateTimeUniqueID(id: number): number {
	return 53 * id
}
export function getFormFieldFloat64UniqueID(id: number): number {
	return 59 * id
}
export function getFormFieldIntUniqueID(id: number): number {
	return 61 * id
}
export function getFormFieldSelectUniqueID(id: number): number {
	return 67 * id
}
export function getFormFieldStringUniqueID(id: number): number {
	return 71 * id
}
export function getFormFieldTimeUniqueID(id: number): number {
	return 73 * id
}
export function getFormGroupUniqueID(id: number): number {
	return 79 * id
}
export function getFormSortAssocButtonUniqueID(id: number): number {
	return 83 * id
}
export function getOptionUniqueID(id: number): number {
	return 89 * id
}
