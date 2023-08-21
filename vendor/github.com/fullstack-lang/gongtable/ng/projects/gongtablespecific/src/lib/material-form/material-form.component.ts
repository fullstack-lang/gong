import { Component, Inject, Input, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';

import * as gongtable from 'gongtable'

import { FormGroup, FormBuilder, Validators } from '@angular/forms';

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
    private formFieldDateService: gongtable.FormFieldDateService,
    private formFieldTimeService: gongtable.FormFieldTimeService,
    private formFieldDateTimeService: gongtable.FormFieldDateTimeService,
    private checkBoxService: gongtable.CheckBoxService,
    private formFieldSelectService: gongtable.FormFieldSelectService,
    private formEditAssocButtonService: gongtable.FormEditAssocButtonService,
    private formSortAssocButtonService: gongtable.FormSortAssocButtonService,

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

        for (let form of this.gongtableFrontRepo.FormGroups_array) {
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

        this.selectedFormGroup.FormDivs.sort((t1, t2) => {
          if (t1.FormGroup_FormDivsDBID_Index && t2.FormGroup_FormDivsDBID_Index) {
            if (t1.FormGroup_FormDivsDBID_Index.Int64 > t2.FormGroup_FormDivsDBID_Index.Int64) {
              return 1;
            }
            if (t1.FormGroup_FormDivsDBID_Index.Int64 < t2.FormGroup_FormDivsDBID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        })

        for (let formDiv of this.selectedFormGroup.FormDivs) {
          if (formDiv.FormFields) {
            formDiv.FormFields.sort((t1, t2) => {
              if (t1.FormDiv_FormFieldsDBID_Index && t2.FormDiv_FormFieldsDBID_Index) {
                if (t1.FormDiv_FormFieldsDBID_Index.Int64 > t2.FormDiv_FormFieldsDBID_Index.Int64) {
                  return 1;
                }
                if (t1.FormDiv_FormFieldsDBID_Index.Int64 < t2.FormDiv_FormFieldsDBID_Index.Int64) {
                  return -1;
                }
              }
              return 0;
            })

            for (let formField of formDiv.FormFields) {
              if (formField.FormFieldString) {
                generatedFormGroupConfig[formField.Name] = [formField.FormFieldString.Value, Validators.required]
              }
              if (formField.FormFieldInt) {
                generatedFormGroupConfig[formField.Name] = [formField.FormFieldInt.Value.toString(), Validators.required]
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
                generatedFormGroupConfig[formField.Name] = [formField.FormFieldSelect.Value?.Name, Validators.required]
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
    if (this.generatedForm) {
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
                this.formFieldStringService.updateFormFieldString(formFieldString, this.DataStack).subscribe(
                  () => {
                    console.log("String Form Field updated")
                  }
                )
              }
            }
            if (formField.FormFieldInt) {
              let formFieldInt = formField.FormFieldInt
              let newValue: number = +this.generatedForm.value[formField.Name]

              if (newValue != formFieldInt.Value) {
                formFieldInt.Value = newValue
                this.formFieldIntService.updateFormFieldInt(formFieldInt, this.DataStack).subscribe(
                  () => {
                    console.log("Int Form Field updated")
                  }
                )
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
                this.formFieldDateService.updateFormFieldDate(formFieldDate, this.DataStack).subscribe(() => {
                  console.log("Date Form Field updated");
                });
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
                this.formFieldTimeService.updateFormFieldTime(formFieldTime, this.DataStack).subscribe(
                  () => {
                    console.log("Time Form Field updated")
                  }
                )
              }
            }
            if (formField.FormFieldDateTime) {
              let formFieldDateTime = formField.FormFieldDateTime

              let newValue = this.generatedForm.value[formField.Name]

              if (newValue != formFieldDateTime.Value) {
                formFieldDateTime.Value = newValue
                this.formFieldDateTimeService.updateFormFieldDateTime(formFieldDateTime, this.DataStack).subscribe(
                  () => {
                    console.log("Date Time Form Field updated")
                  }
                )
              }
            }
            if (formField.FormFieldSelect) {
              let newValue = this.generatedForm.value[formField.Name]
              let formFieldSelect = formField.FormFieldSelect

              if (newValue != formFieldSelect.Value?.Name) {
                formFieldSelect.Value = newValue

                if (formFieldSelect.Options == undefined) {
                  return
                }

                for (let option of formFieldSelect.Options) {
                  if (option.Name == newValue) {
                    formFieldSelect.Value = option
                    formFieldSelect.ValueID.Int64 = option.ID
                  }
                }

                let options = formFieldSelect.Options

                this.formFieldSelectService.updateFormFieldSelect(formFieldSelect, this.DataStack).subscribe(
                  () => {
                    console.log("Select Form Field updated")
                    formFieldSelect.Options = options
                  }
                )
              }
            }
          }
        }
        if (formDiv.CheckBoxs) {
          for (let checkBox of formDiv.CheckBoxs) {
            let newValue = this.generatedForm.value[checkBox.Name] as boolean
            if (newValue != checkBox.Value) {
              checkBox.Value = newValue
              this.checkBoxService.updateCheckBox(checkBox, this.DataStack).subscribe(
                () => {
                  console.log("Boolean Field updated")
                }
              )
            }

          }
        }
      }
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

          this.formEditAssocButtonService.updateFormEditAssocButton(formDiv.FormEditAssocButton, this.DataStack).subscribe(
            () => {
              console.log("assoc button updated")

              // when the association button is pressed
              this.dialog.open(MaterialTableComponent, {
                data: {
                  DataStack: this.DataStack + gongtable.TableExtraPathEnum.TableSelectExtra,
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

          this.formSortAssocButtonService.updateFormSortAssocButton(formDiv.FormSortAssocButton, this.DataStack).subscribe(
            () => {
              console.log("sort button updated")

              // when the association button is pressed
              this.dialog.open(MaterialTableComponent, {
                data: {
                  DataStack: this.DataStack + gongtable.TableExtraPathEnum.TableSortExtra,
                  TableName: gongtable.TableExtraNameEnum.TableSortExtraName
                },
              });
            }
          )
        }
      }
    }
  }
}

