package models

const AppModuleImport = `
import { HttpClientModule } from '@angular/common/http';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatCardModule } from '@angular/material/card'
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatRadioModule } from '@angular/material/radio';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

import { FormsModule } from '@angular/forms';

// to split the screen
import { AngularSplitModule } from 'angular-split';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { GongModule } from 'gong'
`
const AppModuleImport2 = `

AppRoutingModule,

HttpClientModule,

MatSliderModule,
MatSelectModule,
MatFormFieldModule,
MatInputModule,
MatDatepickerModule,
MatTableModule,
MatCheckboxModule,
MatButtonModule,
MatIconModule,
MatToolbarModule,
MatCardModule,
MatTooltipModule,
MatRadioModule,
MatSlideToggleModule,
FormsModule,

AngularSplitModule,

// gongsim stack
GongsimcontrolModule,
GongsimModule,

// gongdoc stack
GongdocModule,
GongdocdiagramsModule,

// gongleaflet stack
GongleafletModule,
GongleafletspecificModule,

//
GongModule,
`
