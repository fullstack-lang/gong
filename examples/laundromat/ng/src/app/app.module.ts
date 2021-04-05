import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { HttpClient } from '@angular/common/http';
import { Routes, RouterModule } from '@angular/router';

import { HttpClientModule } from '@angular/common/http';


import { MatRadioModule } from '@angular/material/radio';

import { FormsModule } from '@angular/forms';


// split
import { AngularSplitModule } from 'angular-split';


import { GongsimcontrolModule } from 'gongsimcontrol'
import { GongsimModule } from 'gongsim'


import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { LaundromatModule } from 'laundromat'

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,

    BrowserModule,
    HttpClientModule,

    MatRadioModule,

    FormsModule,

    // gongsim stack
    GongsimcontrolModule,
    GongsimModule,

    // gongdoc stack
    GongdocModule,
    GongdocdiagramsModule,

    LaundromatModule,

    RouterModule,

    AngularSplitModule.forRoot()
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
