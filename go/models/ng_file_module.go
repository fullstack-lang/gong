package models

const NgFileModule = `import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { {{TitlePkgName}}Module } from '{{pkgname}}'

// mandatory
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,

    HttpClientModule,
    {{TitlePkgName}}Module
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
`
