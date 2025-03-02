import { Component, OnInit } from '@angular/core';

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as table from '../../projects/table/src/public-api'

import { TableSpecificComponent } from '../../projects/tablespecific/src/lib/table-specific/table-specific.component'
import { FormSpecificComponent } from "../../projects/tablespecific/src/lib/form-specific/form-specific.component";


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatRadioModule,
    MatButtonModule,
    MatIconModule,
    AngularSplitModule,
    TableSpecificComponent,
    FormSpecificComponent
],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  table_view = table.TableName.TableDefaultName.toString()
  manualy_edited_table_probe = 'Manual Edited Table Stack Probe'

  form_view = table.FormGroupName.FormGroupDefaultName.toString()
  manualy_edited_form_probe = 'Manual Edited Form Probe'

  generated_table_probe_stack = 'Generated Table Stack Probe'

  view = this.form_view

  TableTestNameEnum = table.TableTestNameEnum

  views: string[] = [
    this.table_view,
    this.form_view,
    this.manualy_edited_table_probe,
    this.manualy_edited_form_probe,
    this.generated_table_probe_stack];

  FormName = "Form 1"

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "table"
  StackType = table.StackType

  TableExtraPathEnum = table.TableExtraPathEnum

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
