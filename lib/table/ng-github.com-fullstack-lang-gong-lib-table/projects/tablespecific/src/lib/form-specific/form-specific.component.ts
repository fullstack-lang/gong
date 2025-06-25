import { Component, Inject, Input, OnInit } from '@angular/core';
import { Subscription, forkJoin } from 'rxjs';

import * as table from '../../../../table/src/public-api'

import { FormGroup, FormBuilder, Validators, AbstractControl, ValidationErrors } from '@angular/forms';

import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatOptionModule } from '@angular/material/core';


import { TableSpecificComponent } from '../table-specific/table-specific.component';
import { ConfirmationDialogComponent } from '../dialog/dialog.component';

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatTooltipModule } from '@angular/material/tooltip';


import { TableDialogData } from '../table-dialog-data';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatInputModule } from '@angular/material/input';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { decodeStringToIntArray_json, encodeIntArrayToString_json } from '../association-storage';

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
    MatPaginatorModule,
    MatTooltipModule,
  ],
  templateUrl: './form-specific.component.html',
  styleUrl: './form-specific.component.css'
})
export class FormSpecificComponent {
  @Input() Name: string = ""

  // within the same stack, there can be multiple form. This one is the form to display
  @Input() FormName: string = ""

  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  private commutNbFromBackSubscription: Subscription = new Subscription
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date

  public gongtableFrontRepo?: table.FrontRepo

  // for selection
  selectedFormGroup: table.FormGroup | undefined = undefined

  // generated form by getting info from the back
  angularFormGroup: FormGroup | undefined

  currentFormEditAssocButton: table.FormEditAssocButton | undefined = undefined

  dialogRef: MatDialogRef<TableSpecificComponent, any> | undefined = undefined

  constructor(
    public dialog: MatDialog,
    private gongtableFrontRepoService: table.FrontRepoService,
    private gongtableCommitNbFromBackService: table.CommitNbFromBackService,
    private formBuilder: FormBuilder,

    private formFieldStringService: table.FormFieldStringService,
    private formFieldIntService: table.FormFieldIntService,
    private formFieldFloat64Service: table.FormFieldFloat64Service,
    private formFieldDateService: table.FormFieldDateService,
    private formFieldTimeService: table.FormFieldTimeService,
    private formFieldDateTimeService: table.FormFieldDateTimeService,
    private checkBoxService: table.CheckBoxService,
    private formFieldSelectService: table.FormFieldSelectService,
    private formEditAssocButtonService: table.FormEditAssocButtonService,
    private formSortAssocButtonService: table.FormSortAssocButtonService,

    private formGroupService: table.FormGroupService,

    public confirmationDialog: MatDialog,
  ) {

  }

  ngOnInit(): void {

    this.gongtableFrontRepoService.connectToWebSocket(this.Name).subscribe(
      gongtablesFrontRepo => {
        this.gongtableFrontRepo = gongtablesFrontRepo

        this.selectedFormGroup = undefined

        // refactorable version
        for (let form of this.gongtableFrontRepo.getFrontArray<table.FormGroup>(table.FormGroup.GONGSTRUCT_NAME)) {
          if (form.Name == this.FormName) {
            this.selectedFormGroup = form
          }
        }

        this.angularFormGroup = undefined
        if (this.selectedFormGroup == undefined) {
          return
        }

        let generatedFormGroupConfig: { [key: string]: [any, any] } = {};

        if (this.selectedFormGroup.FormDivs == undefined) {
          return
        }

        for (let formDiv of this.selectedFormGroup.FormDivs) {
          if (formDiv.FormFields) {

            for (let formField of formDiv.FormFields) {
              if (formField.FormFieldString) {
                generatedFormGroupConfig[formField.Name] = [formField.FormFieldString.Value, Validators.required]
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
                generatedFormGroupConfig[formField.Name] = [formField.FormFieldInt.Value.toString(), validators]
              }
              if (formField.FormFieldFloat64) {
                let validators = [Validators.required]

                if (formField.FormFieldFloat64.HasMinValidator) {
                  validators.push(Validators.min(formField.FormFieldFloat64.MinValue))
                }
                if (formField.FormFieldFloat64.HasMaxValidator) {
                  validators.push(Validators.max(formField.FormFieldFloat64.MaxValue))
                }
                generatedFormGroupConfig[formField.Name] = [formField.FormFieldFloat64.Value.toString(), validators]
              }

              if (formField.FormFieldDate) {
                let displayedString = formField.FormFieldDate.Value.toString().substring(0, 10)
                generatedFormGroupConfig[formField.Name] = [displayedString, Validators.required]
              }
              if (formField.FormFieldDateTime) {
                let displayedString = formField.FormFieldDateTime.Value.toString()
                generatedFormGroupConfig[formField.Name] = [displayedString, Validators.required]
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

                generatedFormGroupConfig[formField.Name] = [timeStr, Validators.required]
              }
              if (formField.FormFieldSelect) {
                formField.FormFieldSelect.Options.sort((a, b) => a.Name.localeCompare(b.Name))
                if (formField.FormFieldSelect.CanBeEmpty) {
                  generatedFormGroupConfig[formField.Name] = [formField.FormFieldSelect.Value?.Name, []]
                } else {
                  generatedFormGroupConfig[formField.Name] = [formField.FormFieldSelect.Value?.Name, Validators.required]
                }
              }
            }
          }

          if (formDiv.CheckBoxs) {
            for (let checkBox of formDiv.CheckBoxs) {
              generatedFormGroupConfig[checkBox.Name] = [checkBox.Value, Validators.required]
            }
          }

        }

        this.angularFormGroup = this.formBuilder.group(generatedFormGroupConfig)
      }
    )
  }

  submitForm() {

    const promises = []


    if (this.angularFormGroup == undefined) {
      return
    }
    console.log(this.angularFormGroup.valueChanges)

    if (this.selectedFormGroup == undefined) {
      return
    }

    if (this.selectedFormGroup.FormDivs == undefined) {
      return
    }

    for (let formDiv of this.selectedFormGroup.FormDivs) {
      if (formDiv.FormFields) {
        for (let formField of formDiv.FormFields) {
          if (formField.FormFieldString) {
            let formFieldString = formField.FormFieldString
            let newValue = this.angularFormGroup.value[formField.Name]

            if (newValue != formFieldString.Value) {

              formFieldString.Value = newValue
              promises.push(this.formFieldStringService.updateFront(formFieldString, this.Name))
            }
          }
          if (formField.FormFieldInt) {
            let formFieldInt = formField.FormFieldInt
            let newValue: number = +this.angularFormGroup.value[formField.Name]

            if (newValue != formFieldInt.Value) {

              formFieldInt.Value = newValue
              promises.push(this.formFieldIntService.updateFront(formFieldInt, this.Name))
            }
          }
          if (formField.FormFieldFloat64) {
            let formFieldFlFormFieldFloat64 = formField.FormFieldFloat64
            let newValue: number = +this.angularFormGroup.value[formField.Name]

            if (newValue != formFieldFlFormFieldFloat64.Value) {

              formFieldFlFormFieldFloat64.Value = newValue
              promises.push(this.formFieldFloat64Service.updateFront(formFieldFlFormFieldFloat64, this.Name))
            }
          }
          if (formField.FormFieldDate) {
            // Assume formField is already defined
            let formFieldDate = formField.FormFieldDate

            let formFieldValue = this.angularFormGroup.value[formField.Name];

            // 1. Convert to a UTC formatted string and then to a Date object
            let dateObj = new Date(formFieldValue);
            let formattedDate = dateObj.toISOString();
            let dateObject = new Date(formattedDate);

            console.log("input date", dateObject);
            console.log("date before", formFieldDate.Value);

            // 2. Check if two dates are on the same day
            let inputDate = new Date(formFieldValue);
            let comparisonDate = new Date(formFieldDate.Value);

            let isSameDay = (date1: Date, date2: Date) => {
              return date1.getUTCFullYear() === date2.getUTCFullYear() &&
                date1.getUTCMonth() === date2.getUTCMonth() &&
                date1.getUTCDate() === date2.getUTCDate();
            }

            console.log(isSameDay(inputDate, comparisonDate));

            if (!isSameDay(inputDate, comparisonDate)) {
              formFieldDate.Value = dateObject;
              promises.push(this.formFieldDateService.updateFront(formFieldDate, this.Name))
            }

          }
          if (formField.FormFieldTime) {
            let formFieldTime = formField.FormFieldTime

            const [hours, minutes, seconds] = this.angularFormGroup.value[formField.Name].split(':').map(Number);
            const date = new Date("1970-01-01")
            date.setUTCHours(hours, minutes, seconds);
            // console.log("date for time", date.toUTCString())
            // console.log("date for backend time", new Date(formFieldTime.Value).toUTCString())

            if (date.getTime() != new Date(formFieldTime.Value).getTime()) {
              formFieldTime.Value = date
              promises.push(this.formFieldTimeService.updateFront(formFieldTime, this.Name))
            }
          }
          if (formField.FormFieldDateTime) {
            let formFieldDateTime = formField.FormFieldDateTime

            let newValue = this.angularFormGroup.value[formField.Name]

            if (newValue != formFieldDateTime.Value) {
              formFieldDateTime.Value = newValue
              promises.push(this.formFieldDateTimeService.updateFront(formFieldDateTime, this.Name))
            }
          }
          if (formField.FormFieldSelect) {
            let newValue = this.angularFormGroup.value[formField.Name]
            let formFieldSelect = formField.FormFieldSelect

            if (newValue != formFieldSelect.Value?.Name) {
              formFieldSelect.Value = newValue

              if (!formFieldSelect.CanBeEmpty && formFieldSelect.Options == undefined) {
                return
              }

              if (formFieldSelect.Options) {
                for (let option of formFieldSelect.Options) {
                  if (option.Name == newValue) {
                    formFieldSelect.Value = option
                  }
                }
              }
              promises.push(this.formFieldSelectService.updateFront(formFieldSelect, this.Name))
            }
          }
        }
      }
      if (formDiv.CheckBoxs) {
        for (let checkBox of formDiv.CheckBoxs) {
          let newValue = this.angularFormGroup.value[checkBox.Name] as boolean
          if (newValue != checkBox.Value) {
            checkBox.Value = newValue
            promises.push(this.checkBoxService.updateFront(checkBox, this.Name))
          }
        }
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
            console.log("Form refreshed", promises.length)
            // a refresh is necessary to redeem all associations
            // this.refresh()
          }
        )
      }
    )

    if (promises.length == 0) {
      this.formGroupService.updateFront(this.selectedFormGroup!, this.Name).subscribe(
        () => {
          console.log("Form refreshed without updates")

          // a refresh is necessary to redeem all associations
          // this.refresh()
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

          // This update will have the back generates the table stack for selecting the associated fields
          // this table stack front "-table-pick" will be fill with the storage for having
          // the rows selected
          formDiv.FormEditAssocButton.IsForSavePurpose = false
          this.formEditAssocButtonService.updateFront(formDiv.FormEditAssocButton, this.Name).subscribe(
            () => {
              console.log("assoc button updated")

              // when the association button is pressed
              this.dialogRef = this.dialog.open(TableSpecificComponent, {
                data: {
                  Name: this.Name + table.TableExtraPathEnum.StackNamePostFixForTableForAssociation,
                  TableName: table.TableExtraNameEnum.TableSelectExtraName,
                  AssociationStorage: this.currentFormEditAssocButton?.AssociationStorage
                },
              });

              this.dialogRef.afterClosed().subscribe(result => {
                console.log('The dialog was closed');

                if (result) {
                  // Extract selected IDs from result
                  let selectedIDs = new Array<number>();
                  for (let row of result) {
                    selectedIDs.push(row.Cells[0].CellInt.Value);
                  }

                  if (this.currentFormEditAssocButton) {
                    // Get current association storage as array
                    let currentAssociationIDs = decodeStringToIntArray_json(this.currentFormEditAssocButton.AssociationStorage);

                    // Create new ordered array
                    let newOrderedIDs = new Array<number>();

                    // First, preserve the original order for existing items that are still selected
                    for (let existingID of currentAssociationIDs) {
                      if (selectedIDs.includes(existingID)) {
                        newOrderedIDs.push(existingID);
                      }
                    }

                    // Then, append any new items that weren't in the original association storage
                    for (let selectedID of selectedIDs) {
                      if (!currentAssociationIDs.includes(selectedID)) {
                        newOrderedIDs.push(selectedID);
                      }
                    }

                    // Update the association storage with the new ordered array
                    this.currentFormEditAssocButton.AssociationStorage = encodeIntArrayToString_json(newOrderedIDs);
                    this.currentFormEditAssocButton.HasChanged = true;
                    console.log('Result:', this.currentFormEditAssocButton.AssociationStorage);
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

              // when the association button is pressed
              this.dialogRef = this.dialog.open(TableSpecificComponent, {
                data: {
                  Name: this.Name + table.TableExtraPathEnum.StackNamePostFixForTableForAssociationSorting,
                  TableName: table.TableExtraNameEnum.TableSortExtraName,
                  AssociationStorage: this.currentFormEditAssocButton?.AssociationStorage
                },
              })

              this.dialogRef.afterClosed().subscribe(result => {
                console.log('The sort dialog was closed');

                // Get the updated AssociationStorage from the dialog data
                const updatedAssociationStorage = this.dialogRef?.componentInstance?.tableDialogData?.AssociationStorage;

                if (updatedAssociationStorage && this.currentFormEditAssocButton) {
                  // Update the FormEditAssocButton with the new association storage
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

    // the update of the form will be called later
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
        // #651 do not generate an update of the group !
        // this.formGroupService.updateFront(
        //   this.selectedFormGroup, this.Name).subscribe(
        //     () => {

        //     }
        //   )
      }
    });
  }

  getDynamicStyles(formField: table.FormField): { [key: string]: any } {
    const styles: { [key: string]: any } = {} // Explicitly define the type here   
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