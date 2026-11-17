import { Component, Inject, Input, OnDestroy, OnInit, ChangeDetectorRef } from '@angular/core';
import { Subscription, forkJoin } from 'rxjs';

import * as form from '../../../../form/src/public-api'

import { FormGroup, FormBuilder, Validators, AbstractControl, ValidationErrors } from '@angular/forms';

import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatOptionModule } from '@angular/material/core';


import { TableSpecificComponent } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/table-specific/table-specific.component';
import { ConfirmationDialogComponent } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/dialog/dialog.component';

import { MatRadioModule } from '@angular/material/radio';
import { MatSliderModule } from '@angular/material/slider';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatDividerModule } from "@angular/material/divider";
import { MatExpansionModule } from '@angular/material/expansion';

import { TableDialogData } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/table-dialog-data';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatInputModule } from '@angular/material/input';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { decodeStringToIntArray_json, encodeIntArrayToString_json } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/association-storage';

import { TableExtraPathEnum } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/table/src/lib/TableExtraPathEnum';
import { TableExtraNameEnum } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/table/src/lib/TableExtraNameEnum';

export interface FormChunk {
  IsAccordionGroup: boolean;
  Panels: FormAccordionPanel[];
  NormalDivs: form.FormDiv[];
}

export interface FormAccordionPanel {
  Label: string;
  Divs: form.FormDiv[];
  IsExpanded?: boolean;
}

@Component({
  selector: 'lib-form-specific',
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatOptionModule,
    MatIconModule,
    MatButtonModule,
    MatInputModule,
    MatSelectModule,
    MatCheckboxModule,
    MatRadioModule,
    MatSliderModule,
    MatPaginatorModule,
    MatTooltipModule,
    MatDividerModule,
    MatExpansionModule
  ],
  templateUrl: './form-specific.html',
  styleUrl: './form-specific.css',
})
export class FormSpecific implements OnInit, OnDestroy {
  @Input() Name: string = ""

  private frontRepoSub: Subscription | undefined
  private debounceTimers: Map<number, ReturnType<typeof setTimeout>> = new Map()
  private liveUpdateSubs: Map<string, Subscription> = new Map()
  private readonly debounceDelayMs = 200

  public gongtableFrontRepo?: form.FrontRepo

  // for selection
  selectedFormGroup: form.FormGroup | undefined = undefined

  // generated form by getting info from the back
  angularFormGroup: FormGroup | undefined

  chunks: FormChunk[] = []

  // map from instance key (${TypeLabel}:::${Label}) to map of panelLabel -> isExpanded
  private expandedPanelsByInstance = new Map<string, Map<string, boolean>>()

  private getInstanceKey(): string {
    if (!this.selectedFormGroup) return ''
    const typeLabel = this.selectedFormGroup.TypeLabel || 'default'
    const label = this.selectedFormGroup.Label || this.selectedFormGroup.Name || 'default'
    return `${typeLabel}:::${label}`
  }

  onPanelExpandedChange(panelLabel: string, isExpanded: boolean): void {
    const key = this.getInstanceKey()
    if (!key) return
    let map = this.expandedPanelsByInstance.get(key)
    if (!map) {
      map = new Map<string, boolean>()
      this.expandedPanelsByInstance.set(key, map)
    }
    map.set(panelLabel, isExpanded)
  }

  private hasDivContent(div: form.FormDiv): boolean {
    return !!(
      div.IsADivider ||
      (div.FormFields && div.FormFields.length > 0) ||
      (div.CheckBoxs && div.CheckBoxs.length > 0) ||
      div.FormEditAssocButton ||
      div.FormSortAssocButton
    )
  }

  getFieldControlName(div: form.FormDiv, formField: form.FormField, index: number): string {
    return `field_${div.ID}_${formField.ID}_${index}`
  }

  getCheckBoxControlName(div: form.FormDiv, checkBox: form.CheckBox, index: number): string {
    return `cb_${div.ID}_${checkBox.ID}_${index}`
  }

  shouldShowHeader(): boolean {
    if (!this.selectedFormGroup) return false
    return !this.selectedFormGroup.HideSubmitButton ||
      !!this.selectedFormGroup.Label ||
      !!this.selectedFormGroup.TypeLabel ||
      !!this.selectedFormGroup.HasSuppressButton
  }

  isCompactDiv(div: form.FormDiv): boolean {
    if (div.FormFields && div.FormFields.length > 0) {
      return div.FormFields.every(f =>
        !!(f.FormFieldInt && f.FormFieldInt.IsSlider) ||
        !!(f.FormFieldFloat64 && f.FormFieldFloat64.IsSlider)
      )
    }
    return !!(div.CheckBoxs && div.CheckBoxs.length > 0 && !div.FormEditAssocButton && !div.FormSortAssocButton)
  }

  formatNumber(value: any): string {
    if (value === undefined || value === null || value === '') return '0'
    const num = +value
    if (Number.isNaN(num)) return `${value}`
    if (Number.isInteger(num)) return `${num}`
    return `${parseFloat(num.toFixed(4))}`
  }

  formatLabel = (value: number): string => {
    if (Math.abs(value) >= 1000) {
      return Math.round(value / 1000) + 'k'
    }
    return this.formatNumber(value)
  }

  getSliderDisplayValue(div: form.FormDiv, formField: form.FormField, index: number): string {
    const ctrlName = this.getFieldControlName(div, formField, index)
    const ctrl = this.angularFormGroup?.get(ctrlName)
    if (ctrl) {
      return this.formatNumber(ctrl.value)
    }
    if (formField.FormFieldInt) return `${formField.FormFieldInt.Value}`
    if (formField.FormFieldFloat64) return this.formatNumber(formField.FormFieldFloat64.Value)
    return ''
  }

  currentFormEditAssocButton: form.FormEditAssocButton | undefined = undefined

  dialogRef: MatDialogRef<TableSpecificComponent, any> | undefined = undefined

  constructor(
    public dialog: MatDialog,
    private frontRepoService: form.FrontRepoService,
    private formBuilder: FormBuilder,

    private formFieldStringService: form.FormFieldStringService,
    private formFieldIntService: form.FormFieldIntService,
    private formFieldFloat64Service: form.FormFieldFloat64Service,
    private formFieldDateService: form.FormFieldDateService,
    private formFieldTimeService: form.FormFieldTimeService,
    private formFieldDateTimeService: form.FormFieldDateTimeService,
    private checkBoxService: form.CheckBoxService,
    private formFieldSelectService: form.FormFieldSelectService,
    private formEditAssocButtonService: form.FormEditAssocButtonService,
    private formSortAssocButtonService: form.FormSortAssocButtonService,

    private formGroupService: form.FormGroupService,

    public confirmationDialog: MatDialog,
    private cdr: ChangeDetectorRef,
  ) {

  }

  ngOnInit(): void {
    this.frontRepoSub?.unsubscribe()

    this.frontRepoSub = this.frontRepoService.connectToWebSocket(this.Name).subscribe(
      gongtablesFrontRepo => {
        this.gongtableFrontRepo = gongtablesFrontRepo

        this.selectedFormGroup = undefined

        for (let form_ of this.gongtableFrontRepo.getFrontArray<form.FormGroup>(form.FormGroup.GONGSTRUCT_NAME)) {
          this.selectedFormGroup = form_
        }

        this.angularFormGroup = undefined
        this.chunks = []
        if (this.selectedFormGroup == undefined) {
          this.cdr.markForCheck()
          return
        }

        if (this.selectedFormGroup.FormDivs) {
          const instanceKey = this.getInstanceKey();
          const expandedMap = this.expandedPanelsByInstance.get(instanceKey);

          let currentChunk: FormChunk | undefined = undefined;
          let currentPanel: FormAccordionPanel | undefined = undefined;

          for (let div of this.selectedFormGroup.FormDivs) {
            if (div.IsAStartAccordionGroup) {
              if (!currentChunk || !currentChunk.IsAccordionGroup) {
                currentChunk = { IsAccordionGroup: true, Panels: [], NormalDivs: [] };
                this.chunks.push(currentChunk);
              }
              const savedExpanded = expandedMap ? expandedMap.get(div.AccordionGroupName) : undefined;
              const isExpanded = savedExpanded !== undefined ? savedExpanded : (div.IsExpandedInitially || false);
              currentPanel = { Label: div.AccordionGroupName, Divs: [], IsExpanded: isExpanded };
              currentChunk.Panels.push(currentPanel);
            }

            if (div.IsAEndAccordionGroup) {
              if (this.hasDivContent(div)) {
                if (currentPanel) {
                  currentPanel.Divs.push(div);
                } else if (currentChunk && !currentChunk.IsAccordionGroup) {
                  currentChunk.NormalDivs.push(div);
                }
              }
              currentChunk = undefined;
              currentPanel = undefined;
              continue;
            }

            if (this.hasDivContent(div)) {
              if (currentPanel) {
                currentPanel.Divs.push(div);
              } else {
                if (!currentChunk || currentChunk.IsAccordionGroup) {
                  currentChunk = { IsAccordionGroup: false, Panels: [], NormalDivs: [] };
                  this.chunks.push(currentChunk);
                }
                currentChunk.NormalDivs.push(div);
              }
            }
          }
        }

        const generatedFormGroupConfig: { [key: string]: any } = {};

        if (this.selectedFormGroup.FormDivs) {
          for (let formDiv of this.selectedFormGroup.FormDivs) {
            if (formDiv.FormFields) {
              formDiv.FormFields.forEach((formField, index) => {
                const uniqueFormControlName = this.getFieldControlName(formDiv, formField, index);
                const isDisabled = !!formField.IsDisabled;

                if (formField.FormFieldString) {
                  generatedFormGroupConfig[uniqueFormControlName] = [{ value: formField.FormFieldString.Value, disabled: isDisabled }, Validators.required]
                }
                if (formField.FormFieldInt) {
                  let validators = [Validators.required]

                  if (formField.FormFieldInt.HasMinValidator) {
                    validators.push(Validators.min(formField.FormFieldInt.MinValue))
                  }
                  if (formField.FormFieldInt.HasMaxValidator) {
                    validators.push(Validators.max(formField.FormFieldInt.MaxValue))
                  }
                  validators.push(integerValidator)
                  const initVal = formField.FormFieldInt.IsSlider
                    ? formField.FormFieldInt.Value
                    : formField.FormFieldInt.Value.toString()
                  generatedFormGroupConfig[uniqueFormControlName] = [{ value: initVal, disabled: isDisabled }, validators]
                }
                if (formField.FormFieldFloat64) {
                  let validators = [Validators.required]

                  if (formField.FormFieldFloat64.HasMinValidator) {
                    validators.push(Validators.min(formField.FormFieldFloat64.MinValue))
                  }
                  if (formField.FormFieldFloat64.HasMaxValidator) {
                    validators.push(Validators.max(formField.FormFieldFloat64.MaxValue))
                  }
                  const initVal = formField.FormFieldFloat64.IsSlider
                    ? formField.FormFieldFloat64.Value
                    : formField.FormFieldFloat64.Value.toString()
                  generatedFormGroupConfig[uniqueFormControlName] = [{ value: initVal, disabled: isDisabled }, validators]
                }

                if (formField.FormFieldDate) {
                  let displayedString = formField.FormFieldDate.Value.toString().substring(0, 10)
                  generatedFormGroupConfig[uniqueFormControlName] = [{ value: displayedString, disabled: isDisabled }, Validators.required]
                }
                if (formField.FormFieldDateTime) {
                  let displayedString = formField.FormFieldDateTime.Value.toString()
                  generatedFormGroupConfig[uniqueFormControlName] = [{ value: displayedString, disabled: isDisabled }, Validators.required]
                }
                if (formField.FormFieldTime) {
                  let timeValue = new Date(formField.FormFieldTime.Value)
                  let hours = timeValue.getUTCHours();
                  let minutes = timeValue.getUTCMinutes();
                  let seconds = timeValue.getUTCSeconds();

                  let hoursStr = hours < 10 ? '0' + hours.toString() : hours.toString();
                  let minutesStr = minutes < 10 ? '0' + minutes.toString() : minutes.toString();
                  let secondsStr = seconds < 10 ? '0' + seconds.toString() : seconds.toString();

                  const timeStr = `${hoursStr}:${minutesStr}:${secondsStr}`

                  generatedFormGroupConfig[uniqueFormControlName] = [{ value: timeStr, disabled: isDisabled }, Validators.required]
                }
                if (formField.FormFieldSelect) {

                  if (formField.FormFieldSelect.PreserveInitialOrder == false) {
                    formField.FormFieldSelect.Options.sort((a, b) => a.Name.localeCompare(b.Name))
                  }

                  if (formField.FormFieldSelect.CanBeEmpty) {
                    generatedFormGroupConfig[uniqueFormControlName] = [{ value: formField.FormFieldSelect.Value?.Name, disabled: isDisabled }, []]
                  } else {
                    generatedFormGroupConfig[uniqueFormControlName] = [{ value: formField.FormFieldSelect.Value?.Name, disabled: isDisabled }, Validators.required]
                  }
                }
              })
            }

            if (formDiv.CheckBoxs) {
              formDiv.CheckBoxs.forEach((checkBox, index) => {
                const uniqueFormControlName = this.getCheckBoxControlName(formDiv, checkBox, index);
                generatedFormGroupConfig[uniqueFormControlName] = [{ value: checkBox.Value, disabled: !!checkBox.IsDisabled }, Validators.required]
              });
            }

          }
        }

        this.angularFormGroup = this.formBuilder.group(generatedFormGroupConfig)
        this.cdr.markForCheck()
      }
    )
  }

  ngOnDestroy(): void {
    this.frontRepoSub?.unsubscribe()
    this.debounceTimers.forEach(timer => clearTimeout(timer))
    this.debounceTimers.clear()
    this.liveUpdateSubs.forEach(sub => sub.unsubscribe())
    this.liveUpdateSubs.clear()
  }

  private extractSliderEventValue($event: Event, div: form.FormDiv, formField: form.FormField, index: number): number {
    const target = $event.target as HTMLInputElement | null
    const ctrlName = this.getFieldControlName(div, formField, index)
    let numVal: number
    if (target && target.value !== undefined && target.value !== '') {
      numVal = +target.value
    } else {
      numVal = +(this.angularFormGroup?.get(ctrlName)?.value ?? 0)
    }
    this.angularFormGroup?.get(ctrlName)?.setValue(numVal, { emitEvent: false })
    this.cdr.markForCheck()
    return numVal
  }

  private isSliderLiveUpdate(formField: form.FormField): boolean {
    return !!(
      (formField.FormFieldInt && formField.FormFieldInt.IsLiveUpdate) ||
      (formField.FormFieldFloat64 && formField.FormFieldFloat64.IsLiveUpdate)
    )
  }

  onSliderInput($event: Event, formField: form.FormField, div: form.FormDiv, index: number): void {
    const numVal = this.extractSliderEventValue($event, div, formField, index)
    if (!this.isSliderLiveUpdate(formField)) {
      return
    }

    const fieldId = formField.ID
    const existingTimer = this.debounceTimers.get(fieldId)
    if (existingTimer) {
      clearTimeout(existingTimer)
    }

    const timer = setTimeout(() => {
      this.debounceTimers.delete(fieldId)
      this.pushSliderUpdate(formField, numVal)
    }, this.debounceDelayMs)

    this.debounceTimers.set(fieldId, timer)
  }

  onSliderChange($event: Event, formField: form.FormField, div: form.FormDiv, index: number): void {
    const numVal = this.extractSliderEventValue($event, div, formField, index)
    if (!this.isSliderLiveUpdate(formField)) {
      return
    }

    const fieldId = formField.ID
    const existingTimer = this.debounceTimers.get(fieldId)
    if (existingTimer) {
      clearTimeout(existingTimer)
      this.debounceTimers.delete(fieldId)
    }
    this.pushSliderUpdate(formField, numVal)
  }

  private pushSliderUpdate(formField: form.FormField, numVal: number): void {
    if (formField.FormFieldInt) {
      const ffInt = formField.FormFieldInt
      const intVal = Math.round(numVal)
      ffInt.Value = intVal
      const subKey = `int_${ffInt.ID}`
      this.liveUpdateSubs.get(subKey)?.unsubscribe()
      const sub = this.formFieldIntService.updateFront(ffInt, this.Name).subscribe()
      this.liveUpdateSubs.set(subKey, sub)
    } else if (formField.FormFieldFloat64) {
      const ffFloat = formField.FormFieldFloat64
      ffFloat.Value = numVal
      const subKey = `float_${ffFloat.ID}`
      this.liveUpdateSubs.get(subKey)?.unsubscribe()
      const sub = this.formFieldFloat64Service.updateFront(ffFloat, this.Name).subscribe()
      this.liveUpdateSubs.set(subKey, sub)
    }
  }

  onCheckBoxChange(checkBox: form.CheckBox, div: form.FormDiv, index: number): void {
    if (!checkBox.IsLiveUpdate) {
      return
    }
    const ctrlName = this.getCheckBoxControlName(div, checkBox, index)
    const newVal = !!this.angularFormGroup?.get(ctrlName)?.value
    if (newVal !== checkBox.Value) {
      checkBox.Value = newVal
      const subKey = `cb_${checkBox.ID}`
      this.liveUpdateSubs.get(subKey)?.unsubscribe()
      const sub = this.checkBoxService.updateFront(checkBox, this.Name).subscribe()
      this.liveUpdateSubs.set(subKey, sub)
    }
  }

  onSelectChange(formField: form.FormField, div: form.FormDiv, index: number): void {
    const sel = formField.FormFieldSelect
    if (!sel || !sel.IsLiveUpdate) {
      return
    }
    const ctrlName = this.getFieldControlName(div, formField, index)
    const newValueFromForm = this.angularFormGroup?.get(ctrlName)?.value
    if (newValueFromForm !== sel.Value?.Name) {
      let matchedOption: form.Option | undefined = undefined
      if (sel.Options) {
        for (let option of sel.Options) {
          if (option.Name === newValueFromForm) {
            matchedOption = option
            break
          }
        }
      }
      sel.Value = matchedOption
      const subKey = `sel_${sel.ID}`
      this.liveUpdateSubs.get(subKey)?.unsubscribe()
      const sub = this.formFieldSelectService.updateFront(sel, this.Name).subscribe()
      this.liveUpdateSubs.set(subKey, sub)
    }
  }

  submitForm() {

    const promises = []


    if (this.angularFormGroup == undefined) {
      return
    }

    if (this.selectedFormGroup == undefined) {
      return
    }

    if (this.selectedFormGroup.FormDivs == undefined) {
      return
    }

    const rawFormValues = this.angularFormGroup.getRawValue()

    for (let formDiv of this.selectedFormGroup.FormDivs) {
      if (formDiv.FormFields) {
        formDiv.FormFields.forEach((formField, index) => {
          const uniqueFormControlName = this.getFieldControlName(formDiv, formField, index);
          const newValueFromForm = rawFormValues[uniqueFormControlName];

          if (formField.FormFieldString) {
            let formFieldString = formField.FormFieldString
            if (newValueFromForm != formFieldString.Value) {

              formFieldString.Value = newValueFromForm
              promises.push(this.formFieldStringService.updateFront(formFieldString, this.Name))
            }
          }
          if (formField.FormFieldInt) {
            let formFieldInt = formField.FormFieldInt
            let newValue: number = +newValueFromForm

            if (newValue != formFieldInt.Value) {

              formFieldInt.Value = newValue
              promises.push(this.formFieldIntService.updateFront(formFieldInt, this.Name))
            }
          }
          if (formField.FormFieldFloat64) {
            let formFieldFlFormFieldFloat64 = formField.FormFieldFloat64
            let newValue: number = +newValueFromForm

            if (newValue != formFieldFlFormFieldFloat64.Value) {

              formFieldFlFormFieldFloat64.Value = newValue
              promises.push(this.formFieldFloat64Service.updateFront(formFieldFlFormFieldFloat64, this.Name))
            }
          }
          if (formField.FormFieldDate) {
            let formFieldDate = formField.FormFieldDate

            let formFieldValue = newValueFromForm;

            let dateObj = new Date(formFieldValue);
            let formattedDate = dateObj.toISOString();
            let dateObject = new Date(formattedDate);

            let inputDate = new Date(formFieldValue);
            let comparisonDate = new Date(formFieldDate.Value);

            let isSameDay = (date1: Date, date2: Date) => {
              return date1.getUTCFullYear() === date2.getUTCFullYear() &&
                date1.getUTCMonth() === date2.getUTCMonth() &&
                date1.getUTCDate() === date2.getUTCDate();
            }

            if (!isSameDay(inputDate, comparisonDate)) {
              formFieldDate.Value = dateObject;
              promises.push(this.formFieldDateService.updateFront(formFieldDate, this.Name))
            }

          }
          if (formField.FormFieldTime) {
            let formFieldTime = formField.FormFieldTime

            const [hours, minutes, seconds] = newValueFromForm.split(':').map(Number);
            const date = new Date("1970-01-01")
            date.setUTCHours(hours, minutes, seconds);

            if (date.getTime() != new Date(formFieldTime.Value).getTime()) {
              formFieldTime.Value = date
              promises.push(this.formFieldTimeService.updateFront(formFieldTime, this.Name))
            }
          }
          if (formField.FormFieldDateTime) {
            let formFieldDateTime = formField.FormFieldDateTime

            if (newValueFromForm != formFieldDateTime.Value) {
              formFieldDateTime.Value = newValueFromForm
              promises.push(this.formFieldDateTimeService.updateFront(formFieldDateTime, this.Name))
            }
          }
          if (formField.FormFieldSelect) {
            let formFieldSelect = formField.FormFieldSelect

            if (newValueFromForm != formFieldSelect.Value?.Name) {
              formFieldSelect.Value = newValueFromForm

              if (!formFieldSelect.CanBeEmpty && formFieldSelect.Options == undefined) {
                return
              }

              if (formFieldSelect.Options) {
                for (let option of formFieldSelect.Options) {
                  if (option.Name == newValueFromForm) {
                    formFieldSelect.Value = option
                  }
                }
              }
              promises.push(this.formFieldSelectService.updateFront(formFieldSelect, this.Name))
              if (formDiv.FormEditAssocButton) {
                formDiv.FormEditAssocButton.HasChanged = false
              }
            }
          }
        })
      }
      if (formDiv.CheckBoxs) {
        formDiv.CheckBoxs.forEach((checkBox, index) => {
          const uniqueFormControlName = this.getCheckBoxControlName(formDiv, checkBox, index);
          const newValue = rawFormValues[uniqueFormControlName] as boolean;
          if (newValue != checkBox.Value) {
            checkBox.Value = newValue
            promises.push(this.checkBoxService.updateFront(checkBox, this.Name))
          }
        });
      }
      if (formDiv.FormEditAssocButton && formDiv.FormEditAssocButton.HasChanged) {
        formDiv.FormEditAssocButton.IsForSavePurpose = true
        promises.push(this.formEditAssocButtonService.updateFront(formDiv.FormEditAssocButton, this.Name))
      }
    }

    // wait till all promises are completed to update the form group itself
    forkJoin(promises).subscribe(
      () => {
        this.formGroupService.updateFront(this.selectedFormGroup!, this.Name).subscribe(
          () => {
          }
        )
      }
    )

    if (promises.length == 0) {
      this.formGroupService.updateFront(this.selectedFormGroup!, this.Name).subscribe(
        () => {
          console.log("Form refreshed without updates")
        }
      )
    }

  }

  openTableAssociation(fieldName: string) {

    console.log("openTableAssociation: ", fieldName)

    if (this.angularFormGroup == undefined) {
      return
    }

    if (this.selectedFormGroup == undefined) {
      return
    }

    if (this.selectedFormGroup.FormDivs == undefined) {
      return
    }

    for (let formDiv of this.selectedFormGroup.FormDivs) {
      if (formDiv.FormEditAssocButton) {
        if (formDiv.FormEditAssocButton.Name == fieldName) {

          this.currentFormEditAssocButton = formDiv.FormEditAssocButton

          formDiv.FormEditAssocButton.IsForSavePurpose = false
          this.formEditAssocButtonService.updateFront(formDiv.FormEditAssocButton, this.Name).subscribe(
            () => {
              console.log("assoc button updated")

              this.dialogRef = this.dialog.open(TableSpecificComponent, {
                maxWidth: '90vw',
                data: {
                  Name: this.Name + TableExtraPathEnum.StackNamePostFixForTableForAssociation,
                  AssociationStorage: this.currentFormEditAssocButton?.AssociationStorage
                },
              });

              this.dialogRef.afterClosed().subscribe(result => {
                console.log('The dialog was closed');

                if (result) {
                  let selectedIDs = new Array<number>();
                  for (let row of result) {
                    selectedIDs.push(row.ID - 1);
                  }

                  if (this.currentFormEditAssocButton) {
                    let currentAssociationIDs = decodeStringToIntArray_json(this.currentFormEditAssocButton.AssociationStorage);

                    let newOrderedIDs = new Array<number>();

                    for (let existingID of currentAssociationIDs) {
                      if (selectedIDs.includes(existingID)) {
                        newOrderedIDs.push(existingID);
                      }
                    }

                    for (let selectedID of selectedIDs) {
                      if (!currentAssociationIDs.includes(selectedID)) {
                        newOrderedIDs.push(selectedID);
                      }
                    }

                    this.currentFormEditAssocButton.AssociationStorage = encodeIntArrayToString_json(newOrderedIDs);
                    this.currentFormEditAssocButton.HasChanged = true;
                    console.log('Result:', this.currentFormEditAssocButton.AssociationStorage);

                    if (formDiv.FormFields && formDiv.FormFields.length > 0 && formDiv.FormFields[0].FormFieldSelect) {
                      const formField = formDiv.FormFields[0]
                      const uniqueFormControlName = this.getFieldControlName(formDiv, formField, 0)
                      const select = formField.FormFieldSelect
                      if (select) {
                        if (newOrderedIDs.length > 0) {
                          const lastID = newOrderedIDs[newOrderedIDs.length - 1]
                          if (select.Options && lastID >= 0 && lastID < select.Options.length) {
                            const option = select.Options[lastID]
                            select.Value = option
                            this.angularFormGroup?.get(uniqueFormControlName)?.setValue(option.Name)
                          }
                        } else {
                          select.Value = undefined
                          this.angularFormGroup?.get(uniqueFormControlName)?.setValue(null)
                        }
                      }
                    }
                  }
                }
              });
            }
          )
        }
      }
    }
  }

  openTableSort(fieldName: string) {
    console.log("openTableSort: ", fieldName)

    if (this.angularFormGroup == undefined) {
      return
    }

    if (this.selectedFormGroup == undefined) {
      return
    }

    if (this.selectedFormGroup.FormDivs == undefined) {
      return
    }

    for (let formDiv of this.selectedFormGroup.FormDivs) {
      if (formDiv.FormSortAssocButton) {
        if (formDiv.FormSortAssocButton.Name == fieldName) {

          this.currentFormEditAssocButton = formDiv.FormSortAssocButton.FormEditAssocButton

          this.formSortAssocButtonService.updateFront(formDiv.FormSortAssocButton, this.Name).subscribe(
            () => {
              console.log("sort button updated")

              this.dialogRef = this.dialog.open(TableSpecificComponent, {
                maxWidth: '90vw',
                data: {
                  Name: this.Name + TableExtraPathEnum.StackNamePostFixForTableForAssociationSorting,
                  TableName: TableExtraNameEnum.TableSortExtraName,
                  AssociationStorage: this.currentFormEditAssocButton?.AssociationStorage
                },
              })

              this.dialogRef.afterClosed().subscribe(result => {
                console.log('The sort dialog was closed');

                const updatedAssociationStorage = this.dialogRef?.componentInstance?.tableDialogData?.AssociationStorage;

                if (updatedAssociationStorage && this.currentFormEditAssocButton) {
                  this.currentFormEditAssocButton.AssociationStorage = updatedAssociationStorage;
                  this.currentFormEditAssocButton.HasChanged = true;
                  console.log('Updated AssociationStorage after sorting:', this.currentFormEditAssocButton.AssociationStorage);
                }
              })
            }
          )
        }
      }
    }
  }

  onSuppress() {
    if (this.selectedFormGroup == undefined) {
      return
    }
    this.selectedFormGroup.HasSuppressButtonBeenPressed = true
  }

  openDialog(): void {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      width: '250px',
      data: { message: 'Are you sure you want to delete this item?' }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        if (this.selectedFormGroup == undefined) {
          return
        }
        this.selectedFormGroup.HasSuppressButtonBeenPressed = true
        this.formGroupService.updateFront(
          this.selectedFormGroup, this.Name).subscribe(
            () => {

            }
          )
      }
    });
  }

  getDynamicStyles(formField: form.FormField): { [key: string]: any } {
    const styles: { [key: string]: any } = {}
    if (formField) {
      if (formField.HasBespokeWidth) {
        styles['width.px'] = formField.BespokeWidthPx
      }
      if (formField.HasBespokeHeight) {
        styles['height.px'] = formField.BespokeHeightPx
      }
    }
    return styles
  }

}

export function integerValidator(control: AbstractControl): ValidationErrors | null {
  const value = control.value
  if (value === null || value === '') return null
  if (!Number.isInteger(+value)) {
    return { 'integer': 'Input must be an integer' }
  }
  return null
}
