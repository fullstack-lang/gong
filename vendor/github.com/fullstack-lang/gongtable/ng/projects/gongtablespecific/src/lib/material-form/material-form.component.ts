import { Component, Inject, Input, OnInit } from '@angular/core';
import { Subscription, forkJoin } from 'rxjs';

import * as gongtable from 'gongtable'

import { FormGroup, FormBuilder, Validators, AbstractControl, ValidationErrors } from '@angular/forms';

import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule } from '@angular/material/dialog';
import { NgIf } from '@angular/common';
import { TableDialogData } from '../table-dialog-data';
import { MaterialTableComponent } from '../material-table/material-table.component';

@Component({
  selector: 'lib-material-form',
  templateUrl: './material-form.component.html',
  styleUrls: ['./material-form.component.css']
})
export class MaterialFormComponent implements OnInit {

  @Input() DataStack: string = ""

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

  public gongtableFrontRepo?: gongtable.FrontRepo

  // for selection
  selectedFormGroup: gongtable.FormGroupDB | undefined = undefined;

  // angular stuff
  myFormSample: FormGroup | undefined

  // generated form by getting info from the back
  generatedForm: FormGroup | undefined

  constructor(
    public dialog: MatDialog,
    private gongtableFrontRepoService: gongtable.FrontRepoService,
    private gongtableCommitNbFromBackService: gongtable.CommitNbFromBackService,
    private formBuilder: FormBuilder,

    private formFieldStringService: gongtable.FormFieldStringService,
    private formFieldIntService: gongtable.FormFieldIntService,
    private formFieldFloat64Service: gongtable.FormFieldFloat64Service,
    private formFieldDateService: gongtable.FormFieldDateService,
    private formFieldTimeService: gongtable.FormFieldTimeService,
    private formFieldDateTimeService: gongtable.FormFieldDateTimeService,
    private checkBoxService: gongtable.CheckBoxService,
    private formFieldSelectService: gongtable.FormFieldSelectService,
    private formEditAssocButtonService: gongtable.FormEditAssocButtonService,
    private formSortAssocButtonService: gongtable.FormSortAssocButtonService,

    private formGroupService: gongtable.FormGroupService,
  ) {

  }

  ngOnInit(): void {
    this.startAutoRefresh(500); // Refresh every 500 ms (half second)

  }

  ngOnDestroy(): void {
    this.stopAutoRefresh();
  }


  stopAutoRefresh(): void {
    if (this.commutNbFromBackSubscription) {
      this.commutNbFromBackSubscription.unsubscribe();
    }
  }

  startAutoRefresh(intervalMs: number): void {
    this.commutNbFromBackSubscription = this.gongtableCommitNbFromBackService
      .getCommitNbFromBack(intervalMs, this.DataStack)
      .subscribe((commitNbFromBack: number) => {
        // console.log("OutletComponent, last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)

        if (this.lastCommitNbFromBack < commitNbFromBack) {
          const d = new Date()
          console.log("OutletComponent:", this.DataStack, d.toLocaleTimeString() + `.${d.getMilliseconds()}` +
            ", last commit increased nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
          this.lastCommitNbFromBack = commitNbFromBack
          this.refresh()
        }
      }
      )
  }

  refresh(): void {

    this.gongtableFrontRepoService.pull(this.DataStack).subscribe(
      gongtablesFrontRepo => {
        this.gongtableFrontRepo = gongtablesFrontRepo

        this.selectedFormGroup = undefined

        // refactorable version
        for (let form of this.gongtableFrontRepo.getArray<gongtable.FormGroupDB>(gongtable.FormGroupDB.GONGSTRUCT_NAME)) {
          if (form.Name == this.FormName) {
            this.selectedFormGroup = form
          }
        }

        this.generatedForm = undefined
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

        this.generatedForm = this.formBuilder.group(generatedFormGroupConfig)
      }
    )
  }

  submitForm() {

    const promises = []


    if (this.generatedForm == undefined) {
      return
    }
    console.log(this.generatedForm.valueChanges)

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
            let newValue = this.generatedForm.value[formField.Name]

            if (newValue != formFieldString.Value) {

              formFieldString.Value = newValue
              promises.push(this.formFieldStringService.updateFormFieldString(formFieldString, this.DataStack, this.gongtableFrontRepoService.frontRepo))
            }
          }
          if (formField.FormFieldInt) {
            let formFieldInt = formField.FormFieldInt
            let newValue: number = +this.generatedForm.value[formField.Name]

            if (newValue != formFieldInt.Value) {

              formFieldInt.Value = newValue
              promises.push(this.formFieldIntService.updateFormFieldInt(formFieldInt, this.DataStack, this.gongtableFrontRepoService.frontRepo))
            }
          }
          if (formField.FormFieldFloat64) {
            let formFieldFlFormFieldFloat64 = formField.FormFieldFloat64
            let newValue: number = +this.generatedForm.value[formField.Name]

            if (newValue != formFieldFlFormFieldFloat64.Value) {

              formFieldFlFormFieldFloat64.Value = newValue
              promises.push(this.formFieldFloat64Service.updateFormFieldFloat64(formFieldFlFormFieldFloat64, this.DataStack, this.gongtableFrontRepoService.frontRepo))
            }
          }
          if (formField.FormFieldDate) {
            // Assume formField is already defined
            let formFieldDate = formField.FormFieldDate

            let formFieldValue = this.generatedForm.value[formField.Name];

            // 1. Convert to a UTC formatted string and then to a Date object
            let dateObj = new Date(formFieldValue);
            let formattedDate = dateObj.toISOString();
            let dateObject = new Date(formattedDate);

            console.log(dateObject);
            console.log(formFieldDate.Value);

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
              promises.push(this.formFieldDateService.updateFormFieldDate(formFieldDate, this.DataStack, this.gongtableFrontRepoService.frontRepo))
            }

          }
          if (formField.FormFieldTime) {
            let formFieldTime = formField.FormFieldTime

            const [hours, minutes, seconds] = this.generatedForm.value[formField.Name].split(':').map(Number);
            const date = new Date("1970-01-01")
            date.setUTCHours(hours, minutes, seconds);
            // console.log("date for time", date.toUTCString())
            // console.log("date for backend time", new Date(formFieldTime.Value).toUTCString())

            if (date.getTime() != new Date(formFieldTime.Value).getTime()) {
              formFieldTime.Value = date
              promises.push(this.formFieldTimeService.updateFormFieldTime(formFieldTime, this.DataStack, this.gongtableFrontRepoService.frontRepo))
            }
          }
          if (formField.FormFieldDateTime) {
            let formFieldDateTime = formField.FormFieldDateTime

            let newValue = this.generatedForm.value[formField.Name]

            if (newValue != formFieldDateTime.Value) {
              formFieldDateTime.Value = newValue
              promises.push(this.formFieldDateTimeService.updateFormFieldDateTime(formFieldDateTime, this.DataStack, this.gongtableFrontRepoService.frontRepo))
            }
          }
          if (formField.FormFieldSelect) {
            let newValue = this.generatedForm.value[formField.Name]
            let formFieldSelect = formField.FormFieldSelect

            if (newValue != formFieldSelect.Value?.Name) {
              formFieldSelect.Value = newValue

              if (!formFieldSelect.CanBeEmpty && formFieldSelect.Options == undefined) {
                return
              }

              if (formFieldSelect.CanBeEmpty && formFieldSelect.Value == undefined) {
                formFieldSelect.FormFieldSelectPointersEncoding.ValueID.Int64 = 0
              }

              if (formFieldSelect.Options) {
                for (let option of formFieldSelect.Options) {
                  if (option.Name == newValue) {
                    formFieldSelect.Value = option
                    formFieldSelect.FormFieldSelectPointersEncoding.ValueID.Int64 = option.ID
                  }
                }
              }
              promises.push(this.formFieldSelectService.updateFormFieldSelect(formFieldSelect, this.DataStack, this.gongtableFrontRepoService.frontRepo))
            }
          }
        }
      }
      if (formDiv.CheckBoxs) {
        for (let checkBox of formDiv.CheckBoxs) {
          let newValue = this.generatedForm.value[checkBox.Name] as boolean
          if (newValue != checkBox.Value) {
            checkBox.Value = newValue
            promises.push(this.checkBoxService.updateCheckBox(checkBox, this.DataStack, this.gongtableFrontRepoService.frontRepo))
          }
        }
      }
    }

    // wait till all promises are completed to update the form group itself
    forkJoin(promises).subscribe(
      () => {
        this.formGroupService.updateFormGroup(this.selectedFormGroup!, this.DataStack, this.gongtableFrontRepoService.frontRepo).subscribe(
          () => {

            // a refresh is necessary to redeem all associations
            // this.refresh()
          }
        )
      }
    )

    if (promises.length == 0) {
      this.formGroupService.updateFormGroup(this.selectedFormGroup!, this.DataStack, this.gongtableFrontRepoService.frontRepo).subscribe(
        () => {
          // a refresh is necessary to redeem all associations
          // this.refresh()
        }
      )
    }

  }

  openTableAssociation(fieldName: string) {

    console.log("openTableAssociation: ", fieldName)

    if (this.generatedForm == undefined) {
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

          this.formEditAssocButtonService.updateFormEditAssocButton(formDiv.FormEditAssocButton, this.DataStack, this.gongtableFrontRepoService.frontRepo).subscribe(
            () => {
              console.log("assoc button updated")

              // when the association button is pressed
              this.dialog.open(MaterialTableComponent, {
                data: {
                  DataStack: this.DataStack + gongtable.TableExtraPathEnum.StackNamePostFixForTableForAssociation,
                  TableName: gongtable.TableExtraNameEnum.TableSelectExtraName
                },
              });
            }
          )
        }
      }
    }
  }

  openTableSort(fieldName: string) {

    console.log("openTableSort: ", fieldName)

    if (this.generatedForm == undefined) {
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

          this.formSortAssocButtonService.updateFormSortAssocButton(formDiv.FormSortAssocButton, this.DataStack, this.gongtableFrontRepoService.frontRepo).subscribe(
            () => {
              console.log("sort button updated")

              // when the association button is pressed
              this.dialog.open(MaterialTableComponent, {
                data: {
                  DataStack: this.DataStack + gongtable.TableExtraPathEnum.StackNamePostFixForTableForAssociationSorting,
                  TableName: gongtable.TableExtraNameEnum.TableSortExtraName
                },
              });
            }
          )
        }
      }
    }
  }

  getDynamicStyles(formField: gongtable.FormFieldDB): { [key: string]: any } {
    const styles: { [key: string]: any } = {} // Explicitly define the type here   
    if (formField) {
      if (formField.HasBespokeWidth) {
        styles['width.px'] = formField.BespokeWidthPx
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
