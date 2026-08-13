import * as i0 from '@angular/core';
import { NgModule, DOCUMENT, Inject, Injectable } from '@angular/core';
import * as i1 from '@angular/common/http';
import { HttpParams, HttpHeaders } from '@angular/common/http';
import { BehaviorSubject, of, Observable, combineLatest, interval } from 'rxjs';
import { tap, catchError, shareReplay, switchMap } from 'rxjs/operators';
import * as i2 from '@angular/common';
import { DOCUMENT as DOCUMENT$1 } from '@angular/common';

class SplitModule {
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SplitModule, deps: [], target: i0.ɵɵFactoryTarget.NgModule });
    static ɵmod = i0.ɵɵngDeclareNgModule({ minVersion: "14.0.0", version: "20.3.26", ngImport: i0, type: SplitModule });
    static ɵinj = i0.ɵɵngDeclareInjector({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SplitModule });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SplitModule, decorators: [{
            type: NgModule,
            args: [{
                    declarations: [],
                    imports: [],
                    exports: [],
                    providers: [],
                }]
        }] });

// generated code - do not edit
class AsSplit {
    static GONGSTRUCT_NAME = "AsSplit";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Direction = "";
    IsSizeInPixel = false;
    IsWithCustomGutterSize = false;
    GutterSize = 0;
    // insertion point for pointers and slices of pointers declarations
    AsSplitAreas = [];
    CreatedAt;
    DeletedAt;
}
function CopyAsSplitToAsSplitAPI(assplit, assplitAPI) {
    assplitAPI.CreatedAt = assplit.CreatedAt;
    assplitAPI.DeletedAt = assplit.DeletedAt;
    assplitAPI.ID = assplit.ID;
    // insertion point for basic fields copy operations
    assplitAPI.Name = assplit.Name;
    assplitAPI.Direction = assplit.Direction;
    assplitAPI.IsSizeInPixel = assplit.IsSizeInPixel;
    assplitAPI.IsWithCustomGutterSize = assplit.IsWithCustomGutterSize;
    assplitAPI.GutterSize = assplit.GutterSize;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    assplitAPI.AsSplitPointersEncoding.AsSplitAreas = [];
    for (let _assplitarea of assplit.AsSplitAreas) {
        assplitAPI.AsSplitPointersEncoding.AsSplitAreas.push(_assplitarea.ID);
    }
}
// CopyAsSplitAPIToAsSplit update basic, pointers and slice of pointers fields of assplit
// from respectively the basic fields and encoded fields of pointers and slices of pointers of assplitAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyAsSplitAPIToAsSplit(assplitAPI, assplit, frontRepo) {
    assplit.CreatedAt = assplitAPI.CreatedAt;
    assplit.DeletedAt = assplitAPI.DeletedAt;
    assplit.ID = assplitAPI.ID;
    // insertion point for basic fields copy operations
    assplit.Name = assplitAPI.Name;
    assplit.Direction = assplitAPI.Direction;
    assplit.IsSizeInPixel = assplitAPI.IsSizeInPixel;
    assplit.IsWithCustomGutterSize = assplitAPI.IsWithCustomGutterSize;
    assplit.GutterSize = assplitAPI.GutterSize;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    if (!Array.isArray(assplitAPI.AsSplitPointersEncoding.AsSplitAreas)) {
        console.error('Rects is not an array:', assplitAPI.AsSplitPointersEncoding.AsSplitAreas);
        return;
    }
    assplit.AsSplitAreas = new Array();
    for (let _id of assplitAPI.AsSplitPointersEncoding.AsSplitAreas) {
        let _assplitarea = frontRepo.map_ID_AsSplitArea.get(_id);
        if (_assplitarea != undefined) {
            assplit.AsSplitAreas.push(_assplitarea);
        }
    }
}

// generated code - do not edit
class AsSplitArea {
    static GONGSTRUCT_NAME = "AsSplitArea";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    ShowNameInHeader = false;
    Size = 0;
    IsAny = false;
    HasDiv = false;
    DivStyle = "";
    // insertion point for pointers and slices of pointers declarations
    AsSplit;
    Button;
    Cursor;
    Form;
    Load;
    Markdown;
    Slider;
    Split;
    Svg;
    Table;
    Tone;
    Tree;
    Threejs;
    Xlsx;
    CreatedAt;
    DeletedAt;
}
function CopyAsSplitAreaToAsSplitAreaAPI(assplitarea, assplitareaAPI) {
    assplitareaAPI.CreatedAt = assplitarea.CreatedAt;
    assplitareaAPI.DeletedAt = assplitarea.DeletedAt;
    assplitareaAPI.ID = assplitarea.ID;
    // insertion point for basic fields copy operations
    assplitareaAPI.Name = assplitarea.Name;
    assplitareaAPI.ShowNameInHeader = assplitarea.ShowNameInHeader;
    assplitareaAPI.Size = assplitarea.Size;
    assplitareaAPI.IsAny = assplitarea.IsAny;
    assplitareaAPI.HasDiv = assplitarea.HasDiv;
    assplitareaAPI.DivStyle = assplitarea.DivStyle;
    // insertion point for pointer fields encoding
    assplitareaAPI.AsSplitAreaPointersEncoding.AsSplitID.Valid = true;
    if (assplitarea.AsSplit != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.AsSplitID.Int64 = assplitarea.AsSplit.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.AsSplitID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.ButtonID.Valid = true;
    if (assplitarea.Button != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.ButtonID.Int64 = assplitarea.Button.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.ButtonID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.CursorID.Valid = true;
    if (assplitarea.Cursor != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.CursorID.Int64 = assplitarea.Cursor.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.CursorID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.FormID.Valid = true;
    if (assplitarea.Form != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.FormID.Int64 = assplitarea.Form.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.FormID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.LoadID.Valid = true;
    if (assplitarea.Load != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.LoadID.Int64 = assplitarea.Load.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.LoadID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.MarkdownID.Valid = true;
    if (assplitarea.Markdown != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.MarkdownID.Int64 = assplitarea.Markdown.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.MarkdownID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.SliderID.Valid = true;
    if (assplitarea.Slider != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.SliderID.Int64 = assplitarea.Slider.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.SliderID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.SplitID.Valid = true;
    if (assplitarea.Split != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.SplitID.Int64 = assplitarea.Split.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.SplitID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.SvgID.Valid = true;
    if (assplitarea.Svg != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.SvgID.Int64 = assplitarea.Svg.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.SvgID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.TableID.Valid = true;
    if (assplitarea.Table != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.TableID.Int64 = assplitarea.Table.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.TableID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.ToneID.Valid = true;
    if (assplitarea.Tone != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.ToneID.Int64 = assplitarea.Tone.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.ToneID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.TreeID.Valid = true;
    if (assplitarea.Tree != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.TreeID.Int64 = assplitarea.Tree.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.TreeID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.ThreejsID.Valid = true;
    if (assplitarea.Threejs != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.ThreejsID.Int64 = assplitarea.Threejs.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.ThreejsID.Int64 = 0;
    }
    assplitareaAPI.AsSplitAreaPointersEncoding.XlsxID.Valid = true;
    if (assplitarea.Xlsx != undefined) {
        assplitareaAPI.AsSplitAreaPointersEncoding.XlsxID.Int64 = assplitarea.Xlsx.ID;
    }
    else {
        assplitareaAPI.AsSplitAreaPointersEncoding.XlsxID.Int64 = 0;
    }
    // insertion point for slice of pointers fields encoding
}
// CopyAsSplitAreaAPIToAsSplitArea update basic, pointers and slice of pointers fields of assplitarea
// from respectively the basic fields and encoded fields of pointers and slices of pointers of assplitareaAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI, assplitarea, frontRepo) {
    assplitarea.CreatedAt = assplitareaAPI.CreatedAt;
    assplitarea.DeletedAt = assplitareaAPI.DeletedAt;
    assplitarea.ID = assplitareaAPI.ID;
    // insertion point for basic fields copy operations
    assplitarea.Name = assplitareaAPI.Name;
    assplitarea.ShowNameInHeader = assplitareaAPI.ShowNameInHeader;
    assplitarea.Size = assplitareaAPI.Size;
    assplitarea.IsAny = assplitareaAPI.IsAny;
    assplitarea.HasDiv = assplitareaAPI.HasDiv;
    assplitarea.DivStyle = assplitareaAPI.DivStyle;
    // insertion point for pointer fields encoding
    assplitarea.AsSplit = frontRepo.map_ID_AsSplit.get(assplitareaAPI.AsSplitAreaPointersEncoding.AsSplitID.Int64);
    assplitarea.Button = frontRepo.map_ID_Button.get(assplitareaAPI.AsSplitAreaPointersEncoding.ButtonID.Int64);
    assplitarea.Cursor = frontRepo.map_ID_Cursor.get(assplitareaAPI.AsSplitAreaPointersEncoding.CursorID.Int64);
    assplitarea.Form = frontRepo.map_ID_Form.get(assplitareaAPI.AsSplitAreaPointersEncoding.FormID.Int64);
    assplitarea.Load = frontRepo.map_ID_Load.get(assplitareaAPI.AsSplitAreaPointersEncoding.LoadID.Int64);
    assplitarea.Markdown = frontRepo.map_ID_Markdown.get(assplitareaAPI.AsSplitAreaPointersEncoding.MarkdownID.Int64);
    assplitarea.Slider = frontRepo.map_ID_Slider.get(assplitareaAPI.AsSplitAreaPointersEncoding.SliderID.Int64);
    assplitarea.Split = frontRepo.map_ID_Split.get(assplitareaAPI.AsSplitAreaPointersEncoding.SplitID.Int64);
    assplitarea.Svg = frontRepo.map_ID_Svg.get(assplitareaAPI.AsSplitAreaPointersEncoding.SvgID.Int64);
    assplitarea.Table = frontRepo.map_ID_Table.get(assplitareaAPI.AsSplitAreaPointersEncoding.TableID.Int64);
    assplitarea.Tone = frontRepo.map_ID_Tone.get(assplitareaAPI.AsSplitAreaPointersEncoding.ToneID.Int64);
    assplitarea.Tree = frontRepo.map_ID_Tree.get(assplitareaAPI.AsSplitAreaPointersEncoding.TreeID.Int64);
    assplitarea.Threejs = frontRepo.map_ID_Threejs.get(assplitareaAPI.AsSplitAreaPointersEncoding.ThreejsID.Int64);
    assplitarea.Xlsx = frontRepo.map_ID_Xlsx.get(assplitareaAPI.AsSplitAreaPointersEncoding.XlsxID.Int64);
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Button {
    static GONGSTRUCT_NAME = "Button";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyButtonToButtonAPI(button, buttonAPI) {
    buttonAPI.CreatedAt = button.CreatedAt;
    buttonAPI.DeletedAt = button.DeletedAt;
    buttonAPI.ID = button.ID;
    // insertion point for basic fields copy operations
    buttonAPI.Name = button.Name;
    buttonAPI.StackName = button.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyButtonAPIToButton update basic, pointers and slice of pointers fields of button
// from respectively the basic fields and encoded fields of pointers and slices of pointers of buttonAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyButtonAPIToButton(buttonAPI, button, frontRepo) {
    button.CreatedAt = buttonAPI.CreatedAt;
    button.DeletedAt = buttonAPI.DeletedAt;
    button.ID = buttonAPI.ID;
    // insertion point for basic fields copy operations
    button.Name = buttonAPI.Name;
    button.StackName = buttonAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Cursor {
    static GONGSTRUCT_NAME = "Cursor";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    Style = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyCursorToCursorAPI(cursor, cursorAPI) {
    cursorAPI.CreatedAt = cursor.CreatedAt;
    cursorAPI.DeletedAt = cursor.DeletedAt;
    cursorAPI.ID = cursor.ID;
    // insertion point for basic fields copy operations
    cursorAPI.Name = cursor.Name;
    cursorAPI.StackName = cursor.StackName;
    cursorAPI.Style = cursor.Style;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyCursorAPIToCursor update basic, pointers and slice of pointers fields of cursor
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cursorAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyCursorAPIToCursor(cursorAPI, cursor, frontRepo) {
    cursor.CreatedAt = cursorAPI.CreatedAt;
    cursor.DeletedAt = cursorAPI.DeletedAt;
    cursor.ID = cursorAPI.ID;
    // insertion point for basic fields copy operations
    cursor.Name = cursorAPI.Name;
    cursor.StackName = cursorAPI.StackName;
    cursor.Style = cursorAPI.Style;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class FavIcon {
    static GONGSTRUCT_NAME = "FavIcon";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    SVG = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyFavIconToFavIconAPI(favicon, faviconAPI) {
    faviconAPI.CreatedAt = favicon.CreatedAt;
    faviconAPI.DeletedAt = favicon.DeletedAt;
    faviconAPI.ID = favicon.ID;
    // insertion point for basic fields copy operations
    faviconAPI.Name = favicon.Name;
    faviconAPI.SVG = favicon.SVG;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyFavIconAPIToFavIcon update basic, pointers and slice of pointers fields of favicon
// from respectively the basic fields and encoded fields of pointers and slices of pointers of faviconAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyFavIconAPIToFavIcon(faviconAPI, favicon, frontRepo) {
    favicon.CreatedAt = faviconAPI.CreatedAt;
    favicon.DeletedAt = faviconAPI.DeletedAt;
    favicon.ID = faviconAPI.ID;
    // insertion point for basic fields copy operations
    favicon.Name = faviconAPI.Name;
    favicon.SVG = faviconAPI.SVG;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Form {
    static GONGSTRUCT_NAME = "Form";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyFormToFormAPI(form, formAPI) {
    formAPI.CreatedAt = form.CreatedAt;
    formAPI.DeletedAt = form.DeletedAt;
    formAPI.ID = form.ID;
    // insertion point for basic fields copy operations
    formAPI.Name = form.Name;
    formAPI.StackName = form.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyFormAPIToForm update basic, pointers and slice of pointers fields of form
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyFormAPIToForm(formAPI, form, frontRepo) {
    form.CreatedAt = formAPI.CreatedAt;
    form.DeletedAt = formAPI.DeletedAt;
    form.ID = formAPI.ID;
    // insertion point for basic fields copy operations
    form.Name = formAPI.Name;
    form.StackName = formAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Load {
    static GONGSTRUCT_NAME = "Load";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyLoadToLoadAPI(load, loadAPI) {
    loadAPI.CreatedAt = load.CreatedAt;
    loadAPI.DeletedAt = load.DeletedAt;
    loadAPI.ID = load.ID;
    // insertion point for basic fields copy operations
    loadAPI.Name = load.Name;
    loadAPI.StackName = load.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyLoadAPIToLoad update basic, pointers and slice of pointers fields of load
// from respectively the basic fields and encoded fields of pointers and slices of pointers of loadAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyLoadAPIToLoad(loadAPI, load, frontRepo) {
    load.CreatedAt = loadAPI.CreatedAt;
    load.DeletedAt = loadAPI.DeletedAt;
    load.ID = loadAPI.ID;
    // insertion point for basic fields copy operations
    load.Name = loadAPI.Name;
    load.StackName = loadAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class LogoOnTheLeft {
    static GONGSTRUCT_NAME = "LogoOnTheLeft";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Width = 0;
    Height = 0;
    SVG = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyLogoOnTheLeftToLogoOnTheLeftAPI(logoontheleft, logoontheleftAPI) {
    logoontheleftAPI.CreatedAt = logoontheleft.CreatedAt;
    logoontheleftAPI.DeletedAt = logoontheleft.DeletedAt;
    logoontheleftAPI.ID = logoontheleft.ID;
    // insertion point for basic fields copy operations
    logoontheleftAPI.Name = logoontheleft.Name;
    logoontheleftAPI.Width = logoontheleft.Width;
    logoontheleftAPI.Height = logoontheleft.Height;
    logoontheleftAPI.SVG = logoontheleft.SVG;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyLogoOnTheLeftAPIToLogoOnTheLeft update basic, pointers and slice of pointers fields of logoontheleft
// from respectively the basic fields and encoded fields of pointers and slices of pointers of logoontheleftAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyLogoOnTheLeftAPIToLogoOnTheLeft(logoontheleftAPI, logoontheleft, frontRepo) {
    logoontheleft.CreatedAt = logoontheleftAPI.CreatedAt;
    logoontheleft.DeletedAt = logoontheleftAPI.DeletedAt;
    logoontheleft.ID = logoontheleftAPI.ID;
    // insertion point for basic fields copy operations
    logoontheleft.Name = logoontheleftAPI.Name;
    logoontheleft.Width = logoontheleftAPI.Width;
    logoontheleft.Height = logoontheleftAPI.Height;
    logoontheleft.SVG = logoontheleftAPI.SVG;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class LogoOnTheRight {
    static GONGSTRUCT_NAME = "LogoOnTheRight";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Width = 0;
    Height = 0;
    SVG = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyLogoOnTheRightToLogoOnTheRightAPI(logoontheright, logoontherightAPI) {
    logoontherightAPI.CreatedAt = logoontheright.CreatedAt;
    logoontherightAPI.DeletedAt = logoontheright.DeletedAt;
    logoontherightAPI.ID = logoontheright.ID;
    // insertion point for basic fields copy operations
    logoontherightAPI.Name = logoontheright.Name;
    logoontherightAPI.Width = logoontheright.Width;
    logoontherightAPI.Height = logoontheright.Height;
    logoontherightAPI.SVG = logoontheright.SVG;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyLogoOnTheRightAPIToLogoOnTheRight update basic, pointers and slice of pointers fields of logoontheright
// from respectively the basic fields and encoded fields of pointers and slices of pointers of logoontherightAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyLogoOnTheRightAPIToLogoOnTheRight(logoontherightAPI, logoontheright, frontRepo) {
    logoontheright.CreatedAt = logoontherightAPI.CreatedAt;
    logoontheright.DeletedAt = logoontherightAPI.DeletedAt;
    logoontheright.ID = logoontherightAPI.ID;
    // insertion point for basic fields copy operations
    logoontheright.Name = logoontherightAPI.Name;
    logoontheright.Width = logoontherightAPI.Width;
    logoontheright.Height = logoontherightAPI.Height;
    logoontheright.SVG = logoontherightAPI.SVG;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Markdown {
    static GONGSTRUCT_NAME = "Markdown";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyMarkdownToMarkdownAPI(markdown, markdownAPI) {
    markdownAPI.CreatedAt = markdown.CreatedAt;
    markdownAPI.DeletedAt = markdown.DeletedAt;
    markdownAPI.ID = markdown.ID;
    // insertion point for basic fields copy operations
    markdownAPI.Name = markdown.Name;
    markdownAPI.StackName = markdown.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyMarkdownAPIToMarkdown update basic, pointers and slice of pointers fields of markdown
// from respectively the basic fields and encoded fields of pointers and slices of pointers of markdownAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyMarkdownAPIToMarkdown(markdownAPI, markdown, frontRepo) {
    markdown.CreatedAt = markdownAPI.CreatedAt;
    markdown.DeletedAt = markdownAPI.DeletedAt;
    markdown.ID = markdownAPI.ID;
    // insertion point for basic fields copy operations
    markdown.Name = markdownAPI.Name;
    markdown.StackName = markdownAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Slider {
    static GONGSTRUCT_NAME = "Slider";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopySliderToSliderAPI(slider, sliderAPI) {
    sliderAPI.CreatedAt = slider.CreatedAt;
    sliderAPI.DeletedAt = slider.DeletedAt;
    sliderAPI.ID = slider.ID;
    // insertion point for basic fields copy operations
    sliderAPI.Name = slider.Name;
    sliderAPI.StackName = slider.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopySliderAPIToSlider update basic, pointers and slice of pointers fields of slider
// from respectively the basic fields and encoded fields of pointers and slices of pointers of sliderAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopySliderAPIToSlider(sliderAPI, slider, frontRepo) {
    slider.CreatedAt = sliderAPI.CreatedAt;
    slider.DeletedAt = sliderAPI.DeletedAt;
    slider.ID = sliderAPI.ID;
    // insertion point for basic fields copy operations
    slider.Name = sliderAPI.Name;
    slider.StackName = sliderAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Split {
    static GONGSTRUCT_NAME = "Split";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopySplitToSplitAPI(split, splitAPI) {
    splitAPI.CreatedAt = split.CreatedAt;
    splitAPI.DeletedAt = split.DeletedAt;
    splitAPI.ID = split.ID;
    // insertion point for basic fields copy operations
    splitAPI.Name = split.Name;
    splitAPI.StackName = split.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopySplitAPIToSplit update basic, pointers and slice of pointers fields of split
// from respectively the basic fields and encoded fields of pointers and slices of pointers of splitAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopySplitAPIToSplit(splitAPI, split, frontRepo) {
    split.CreatedAt = splitAPI.CreatedAt;
    split.DeletedAt = splitAPI.DeletedAt;
    split.ID = splitAPI.ID;
    // insertion point for basic fields copy operations
    split.Name = splitAPI.Name;
    split.StackName = splitAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Svg {
    static GONGSTRUCT_NAME = "Svg";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    Style = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopySvgToSvgAPI(svg, svgAPI) {
    svgAPI.CreatedAt = svg.CreatedAt;
    svgAPI.DeletedAt = svg.DeletedAt;
    svgAPI.ID = svg.ID;
    // insertion point for basic fields copy operations
    svgAPI.Name = svg.Name;
    svgAPI.StackName = svg.StackName;
    svgAPI.Style = svg.Style;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopySvgAPIToSvg update basic, pointers and slice of pointers fields of svg
// from respectively the basic fields and encoded fields of pointers and slices of pointers of svgAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopySvgAPIToSvg(svgAPI, svg, frontRepo) {
    svg.CreatedAt = svgAPI.CreatedAt;
    svg.DeletedAt = svgAPI.DeletedAt;
    svg.ID = svgAPI.ID;
    // insertion point for basic fields copy operations
    svg.Name = svgAPI.Name;
    svg.StackName = svgAPI.StackName;
    svg.Style = svgAPI.Style;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Table {
    static GONGSTRUCT_NAME = "Table";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyTableToTableAPI(table, tableAPI) {
    tableAPI.CreatedAt = table.CreatedAt;
    tableAPI.DeletedAt = table.DeletedAt;
    tableAPI.ID = table.ID;
    // insertion point for basic fields copy operations
    tableAPI.Name = table.Name;
    tableAPI.StackName = table.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyTableAPIToTable update basic, pointers and slice of pointers fields of table
// from respectively the basic fields and encoded fields of pointers and slices of pointers of tableAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyTableAPIToTable(tableAPI, table, frontRepo) {
    table.CreatedAt = tableAPI.CreatedAt;
    table.DeletedAt = tableAPI.DeletedAt;
    table.ID = tableAPI.ID;
    // insertion point for basic fields copy operations
    table.Name = tableAPI.Name;
    table.StackName = tableAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Threejs {
    static GONGSTRUCT_NAME = "Threejs";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyThreejsToThreejsAPI(threejs, threejsAPI) {
    threejsAPI.CreatedAt = threejs.CreatedAt;
    threejsAPI.DeletedAt = threejs.DeletedAt;
    threejsAPI.ID = threejs.ID;
    // insertion point for basic fields copy operations
    threejsAPI.Name = threejs.Name;
    threejsAPI.StackName = threejs.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyThreejsAPIToThreejs update basic, pointers and slice of pointers fields of threejs
// from respectively the basic fields and encoded fields of pointers and slices of pointers of threejsAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyThreejsAPIToThreejs(threejsAPI, threejs, frontRepo) {
    threejs.CreatedAt = threejsAPI.CreatedAt;
    threejs.DeletedAt = threejsAPI.DeletedAt;
    threejs.ID = threejsAPI.ID;
    // insertion point for basic fields copy operations
    threejs.Name = threejsAPI.Name;
    threejs.StackName = threejsAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Title {
    static GONGSTRUCT_NAME = "Title";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyTitleToTitleAPI(title, titleAPI) {
    titleAPI.CreatedAt = title.CreatedAt;
    titleAPI.DeletedAt = title.DeletedAt;
    titleAPI.ID = title.ID;
    // insertion point for basic fields copy operations
    titleAPI.Name = title.Name;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyTitleAPIToTitle update basic, pointers and slice of pointers fields of title
// from respectively the basic fields and encoded fields of pointers and slices of pointers of titleAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyTitleAPIToTitle(titleAPI, title, frontRepo) {
    title.CreatedAt = titleAPI.CreatedAt;
    title.DeletedAt = titleAPI.DeletedAt;
    title.ID = titleAPI.ID;
    // insertion point for basic fields copy operations
    title.Name = titleAPI.Name;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Tone {
    static GONGSTRUCT_NAME = "Tone";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyToneToToneAPI(tone, toneAPI) {
    toneAPI.CreatedAt = tone.CreatedAt;
    toneAPI.DeletedAt = tone.DeletedAt;
    toneAPI.ID = tone.ID;
    // insertion point for basic fields copy operations
    toneAPI.Name = tone.Name;
    toneAPI.StackName = tone.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyToneAPIToTone update basic, pointers and slice of pointers fields of tone
// from respectively the basic fields and encoded fields of pointers and slices of pointers of toneAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyToneAPIToTone(toneAPI, tone, frontRepo) {
    tone.CreatedAt = toneAPI.CreatedAt;
    tone.DeletedAt = toneAPI.DeletedAt;
    tone.ID = toneAPI.ID;
    // insertion point for basic fields copy operations
    tone.Name = toneAPI.Name;
    tone.StackName = toneAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class Tree {
    static GONGSTRUCT_NAME = "Tree";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyTreeToTreeAPI(tree, treeAPI) {
    treeAPI.CreatedAt = tree.CreatedAt;
    treeAPI.DeletedAt = tree.DeletedAt;
    treeAPI.ID = tree.ID;
    // insertion point for basic fields copy operations
    treeAPI.Name = tree.Name;
    treeAPI.StackName = tree.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyTreeAPIToTree update basic, pointers and slice of pointers fields of tree
// from respectively the basic fields and encoded fields of pointers and slices of pointers of treeAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyTreeAPIToTree(treeAPI, tree, frontRepo) {
    tree.CreatedAt = treeAPI.CreatedAt;
    tree.DeletedAt = treeAPI.DeletedAt;
    tree.ID = treeAPI.ID;
    // insertion point for basic fields copy operations
    tree.Name = treeAPI.Name;
    tree.StackName = treeAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class View {
    static GONGSTRUCT_NAME = "View";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    ShowViewName = false;
    IsSelectedView = false;
    Direction = "";
    IsSecondaryView = false;
    IsSizeInPixel = false;
    IsWithCustomGutterSize = false;
    GutterSize = 0;
    // insertion point for pointers and slices of pointers declarations
    RootAsSplitAreas = [];
    CreatedAt;
    DeletedAt;
}
function CopyViewToViewAPI(view, viewAPI) {
    viewAPI.CreatedAt = view.CreatedAt;
    viewAPI.DeletedAt = view.DeletedAt;
    viewAPI.ID = view.ID;
    // insertion point for basic fields copy operations
    viewAPI.Name = view.Name;
    viewAPI.ShowViewName = view.ShowViewName;
    viewAPI.IsSelectedView = view.IsSelectedView;
    viewAPI.Direction = view.Direction;
    viewAPI.IsSecondaryView = view.IsSecondaryView;
    viewAPI.IsSizeInPixel = view.IsSizeInPixel;
    viewAPI.IsWithCustomGutterSize = view.IsWithCustomGutterSize;
    viewAPI.GutterSize = view.GutterSize;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    viewAPI.ViewPointersEncoding.RootAsSplitAreas = [];
    for (let _assplitarea of view.RootAsSplitAreas) {
        viewAPI.ViewPointersEncoding.RootAsSplitAreas.push(_assplitarea.ID);
    }
}
// CopyViewAPIToView update basic, pointers and slice of pointers fields of view
// from respectively the basic fields and encoded fields of pointers and slices of pointers of viewAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyViewAPIToView(viewAPI, view, frontRepo) {
    view.CreatedAt = viewAPI.CreatedAt;
    view.DeletedAt = viewAPI.DeletedAt;
    view.ID = viewAPI.ID;
    // insertion point for basic fields copy operations
    view.Name = viewAPI.Name;
    view.ShowViewName = viewAPI.ShowViewName;
    view.IsSelectedView = viewAPI.IsSelectedView;
    view.Direction = viewAPI.Direction;
    view.IsSecondaryView = viewAPI.IsSecondaryView;
    view.IsSizeInPixel = viewAPI.IsSizeInPixel;
    view.IsWithCustomGutterSize = viewAPI.IsWithCustomGutterSize;
    view.GutterSize = viewAPI.GutterSize;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
    if (!Array.isArray(viewAPI.ViewPointersEncoding.RootAsSplitAreas)) {
        console.error('Rects is not an array:', viewAPI.ViewPointersEncoding.RootAsSplitAreas);
        return;
    }
    view.RootAsSplitAreas = new Array();
    for (let _id of viewAPI.ViewPointersEncoding.RootAsSplitAreas) {
        let _assplitarea = frontRepo.map_ID_AsSplitArea.get(_id);
        if (_assplitarea != undefined) {
            view.RootAsSplitAreas.push(_assplitarea);
        }
    }
}

// generated code - do not edit
class Xlsx {
    static GONGSTRUCT_NAME = "Xlsx";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for pointers and slices of pointers declarations
    CreatedAt;
    DeletedAt;
}
function CopyXlsxToXlsxAPI(xlsx, xlsxAPI) {
    xlsxAPI.CreatedAt = xlsx.CreatedAt;
    xlsxAPI.DeletedAt = xlsx.DeletedAt;
    xlsxAPI.ID = xlsx.ID;
    // insertion point for basic fields copy operations
    xlsxAPI.Name = xlsx.Name;
    xlsxAPI.StackName = xlsx.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}
// CopyXlsxAPIToXlsx update basic, pointers and slice of pointers fields of xlsx
// from respectively the basic fields and encoded fields of pointers and slices of pointers of xlsxAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
function CopyXlsxAPIToXlsx(xlsxAPI, xlsx, frontRepo) {
    xlsx.CreatedAt = xlsxAPI.CreatedAt;
    xlsx.DeletedAt = xlsxAPI.DeletedAt;
    xlsx.ID = xlsxAPI.ID;
    // insertion point for basic fields copy operations
    xlsx.Name = xlsxAPI.Name;
    xlsx.StackName = xlsxAPI.StackName;
    // insertion point for pointer fields encoding
    // insertion point for slice of pointers fields encoding
}

// generated code - do not edit
class BackRepoData {
    // insertion point for declarations
    AsSplitAPIs = new Array();
    AsSplitAreaAPIs = new Array();
    ButtonAPIs = new Array();
    CursorAPIs = new Array();
    FavIconAPIs = new Array();
    FormAPIs = new Array();
    LoadAPIs = new Array();
    LogoOnTheLeftAPIs = new Array();
    LogoOnTheRightAPIs = new Array();
    MarkdownAPIs = new Array();
    SliderAPIs = new Array();
    SplitAPIs = new Array();
    SvgAPIs = new Array();
    TableAPIs = new Array();
    ThreejsAPIs = new Array();
    TitleAPIs = new Array();
    ToneAPIs = new Array();
    TreeAPIs = new Array();
    ViewAPIs = new Array();
    XlsxAPIs = new Array();
    // index of the web socket for this stack type (unique among all stack instances)
    GONG__Index;
    constructor(data) {
        // insertion point for copies
        this.AsSplitAPIs = data?.AsSplitAPIs || [];
        this.AsSplitAreaAPIs = data?.AsSplitAreaAPIs || [];
        this.ButtonAPIs = data?.ButtonAPIs || [];
        this.CursorAPIs = data?.CursorAPIs || [];
        this.FavIconAPIs = data?.FavIconAPIs || [];
        this.FormAPIs = data?.FormAPIs || [];
        this.LoadAPIs = data?.LoadAPIs || [];
        this.LogoOnTheLeftAPIs = data?.LogoOnTheLeftAPIs || [];
        this.LogoOnTheRightAPIs = data?.LogoOnTheRightAPIs || [];
        this.MarkdownAPIs = data?.MarkdownAPIs || [];
        this.SliderAPIs = data?.SliderAPIs || [];
        this.SplitAPIs = data?.SplitAPIs || [];
        this.SvgAPIs = data?.SvgAPIs || [];
        this.TableAPIs = data?.TableAPIs || [];
        this.ThreejsAPIs = data?.ThreejsAPIs || [];
        this.TitleAPIs = data?.TitleAPIs || [];
        this.ToneAPIs = data?.ToneAPIs || [];
        this.TreeAPIs = data?.TreeAPIs || [];
        this.ViewAPIs = data?.ViewAPIs || [];
        this.XlsxAPIs = data?.XlsxAPIs || [];
        this.GONG__Index = data?.GONG__Index ?? -1; // Assign Index here
    }
}

class AsSplitAPI {
    static GONGSTRUCT_NAME = "AsSplit";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Direction = "";
    IsSizeInPixel = false;
    IsWithCustomGutterSize = false;
    GutterSize = 0;
    // insertion point for other decls
    AsSplitPointersEncoding = new AsSplitPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class AsSplitPointersEncoding {
    // insertion point for pointers and slices of pointers encoding fields
    AsSplitAreas = [];
}

// generated code, do not edit
// generated by ng_file_service_ts
class AsSplitService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    AsSplitServiceChanged = new BehaviorSubject("");
    assplitsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.assplitsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/assplits';
    }
    /** GET assplits from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getAsSplits(Name, frontRepo);
    }
    getAsSplits(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.assplitsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getAsSplits', [])));
    }
    /** GET assplit by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getAsSplit(id, Name, frontRepo);
    }
    getAsSplit(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.assplitsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched assplit id=${id}`)),
        catchError(this.handleError(`getAsSplit id=${id}`)));
    }
    // postFront copy assplit to a version with encoded pointers and post to the back
    postFront(assplit, Name) {
        let assplitAPI = new AsSplitAPI;
        CopyAsSplitToAsSplitAPI(assplit, assplitAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.assplitsUrl, assplitAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postAsSplit')));
    }
    /** POST: add a new assplit to the server */
    post(assplitdb, Name, frontRepo) {
        return this.postAsSplit(assplitdb, Name, frontRepo);
    }
    postAsSplit(assplitdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.assplitsUrl, assplitdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted assplitdb id=${assplitdb.ID}`)
        }), catchError(this.handleError('postAsSplit')));
    }
    /** DELETE: delete the assplitdb from the server */
    delete(assplitdb, Name) {
        return this.deleteAsSplit(assplitdb, Name);
    }
    deleteAsSplit(assplitdb, Name) {
        const id = typeof assplitdb === 'number' ? assplitdb : assplitdb.ID;
        const url = `${this.assplitsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted assplitdb id=${id}`)), catchError(this.handleError('deleteAsSplit')));
    }
    // updateFront copy assplit to a version with encoded pointers and update to the back
    updateFront(assplit, Name) {
        let assplitAPI = new AsSplitAPI;
        CopyAsSplitToAsSplitAPI(assplit, assplitAPI);
        const id = typeof assplitAPI === 'number' ? assplitAPI : assplitAPI.ID;
        const url = `${this.assplitsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, assplitAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateAsSplit')));
    }
    /** PUT: update the assplitdb on the server */
    update(assplitdb, Name, frontRepo) {
        return this.updateAsSplit(assplitdb, Name, frontRepo);
    }
    updateAsSplit(assplitdb, Name, frontRepo) {
        const id = typeof assplitdb === 'number' ? assplitdb : assplitdb.ID;
        const url = `${this.assplitsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, assplitdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated assplitdb id=${assplitdb.ID}`)
        }), catchError(this.handleError('updateAsSplit')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in AsSplitService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("AsSplitService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: AsSplitService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: AsSplitService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: AsSplitService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// define the type of nullable Int64 in order to support back pointers IDs
class NullInt64 {
    Int64 = 0;
    Valid = false;
}

// usefull for managing pointer ID values that can be nullable
class AsSplitAreaAPI {
    static GONGSTRUCT_NAME = "AsSplitArea";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    ShowNameInHeader = false;
    Size = 0;
    IsAny = false;
    HasDiv = false;
    DivStyle = "";
    // insertion point for other decls
    AsSplitAreaPointersEncoding = new AsSplitAreaPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class AsSplitAreaPointersEncoding {
    // insertion point for pointers and slices of pointers encoding fields
    AsSplitID = new NullInt64; // if pointer is null, AsSplit.ID = 0
    ButtonID = new NullInt64; // if pointer is null, Button.ID = 0
    CursorID = new NullInt64; // if pointer is null, Cursor.ID = 0
    FormID = new NullInt64; // if pointer is null, Form.ID = 0
    LoadID = new NullInt64; // if pointer is null, Load.ID = 0
    MarkdownID = new NullInt64; // if pointer is null, Markdown.ID = 0
    SliderID = new NullInt64; // if pointer is null, Slider.ID = 0
    SplitID = new NullInt64; // if pointer is null, Split.ID = 0
    SvgID = new NullInt64; // if pointer is null, Svg.ID = 0
    TableID = new NullInt64; // if pointer is null, Table.ID = 0
    ToneID = new NullInt64; // if pointer is null, Tone.ID = 0
    TreeID = new NullInt64; // if pointer is null, Tree.ID = 0
    ThreejsID = new NullInt64; // if pointer is null, Threejs.ID = 0
    XlsxID = new NullInt64; // if pointer is null, Xlsx.ID = 0
}

// generated code, do not edit
// generated by ng_file_service_ts
class AsSplitAreaService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    AsSplitAreaServiceChanged = new BehaviorSubject("");
    assplitareasUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.assplitareasUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/assplitareas';
    }
    /** GET assplitareas from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getAsSplitAreas(Name, frontRepo);
    }
    getAsSplitAreas(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.assplitareasUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getAsSplitAreas', [])));
    }
    /** GET assplitarea by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getAsSplitArea(id, Name, frontRepo);
    }
    getAsSplitArea(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.assplitareasUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched assplitarea id=${id}`)),
        catchError(this.handleError(`getAsSplitArea id=${id}`)));
    }
    // postFront copy assplitarea to a version with encoded pointers and post to the back
    postFront(assplitarea, Name) {
        let assplitareaAPI = new AsSplitAreaAPI;
        CopyAsSplitAreaToAsSplitAreaAPI(assplitarea, assplitareaAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.assplitareasUrl, assplitareaAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postAsSplitArea')));
    }
    /** POST: add a new assplitarea to the server */
    post(assplitareadb, Name, frontRepo) {
        return this.postAsSplitArea(assplitareadb, Name, frontRepo);
    }
    postAsSplitArea(assplitareadb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.assplitareasUrl, assplitareadb, httpOptions).pipe(tap(_ => {
            // this.log(`posted assplitareadb id=${assplitareadb.ID}`)
        }), catchError(this.handleError('postAsSplitArea')));
    }
    /** DELETE: delete the assplitareadb from the server */
    delete(assplitareadb, Name) {
        return this.deleteAsSplitArea(assplitareadb, Name);
    }
    deleteAsSplitArea(assplitareadb, Name) {
        const id = typeof assplitareadb === 'number' ? assplitareadb : assplitareadb.ID;
        const url = `${this.assplitareasUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted assplitareadb id=${id}`)), catchError(this.handleError('deleteAsSplitArea')));
    }
    // updateFront copy assplitarea to a version with encoded pointers and update to the back
    updateFront(assplitarea, Name) {
        let assplitareaAPI = new AsSplitAreaAPI;
        CopyAsSplitAreaToAsSplitAreaAPI(assplitarea, assplitareaAPI);
        const id = typeof assplitareaAPI === 'number' ? assplitareaAPI : assplitareaAPI.ID;
        const url = `${this.assplitareasUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, assplitareaAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateAsSplitArea')));
    }
    /** PUT: update the assplitareadb on the server */
    update(assplitareadb, Name, frontRepo) {
        return this.updateAsSplitArea(assplitareadb, Name, frontRepo);
    }
    updateAsSplitArea(assplitareadb, Name, frontRepo) {
        const id = typeof assplitareadb === 'number' ? assplitareadb : assplitareadb.ID;
        const url = `${this.assplitareasUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, assplitareadb, httpOptions).pipe(tap(_ => {
            // this.log(`updated assplitareadb id=${assplitareadb.ID}`)
        }), catchError(this.handleError('updateAsSplitArea')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in AsSplitAreaService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("AsSplitAreaService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: AsSplitAreaService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: AsSplitAreaService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: AsSplitAreaService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class ButtonAPI {
    static GONGSTRUCT_NAME = "Button";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    ButtonPointersEncoding = new ButtonPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class ButtonPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class ButtonService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    ButtonServiceChanged = new BehaviorSubject("");
    buttonsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.buttonsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/buttons';
    }
    /** GET buttons from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getButtons(Name, frontRepo);
    }
    getButtons(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.buttonsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getButtons', [])));
    }
    /** GET button by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getButton(id, Name, frontRepo);
    }
    getButton(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.buttonsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched button id=${id}`)),
        catchError(this.handleError(`getButton id=${id}`)));
    }
    // postFront copy button to a version with encoded pointers and post to the back
    postFront(button, Name) {
        let buttonAPI = new ButtonAPI;
        CopyButtonToButtonAPI(button, buttonAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.buttonsUrl, buttonAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postButton')));
    }
    /** POST: add a new button to the server */
    post(buttondb, Name, frontRepo) {
        return this.postButton(buttondb, Name, frontRepo);
    }
    postButton(buttondb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.buttonsUrl, buttondb, httpOptions).pipe(tap(_ => {
            // this.log(`posted buttondb id=${buttondb.ID}`)
        }), catchError(this.handleError('postButton')));
    }
    /** DELETE: delete the buttondb from the server */
    delete(buttondb, Name) {
        return this.deleteButton(buttondb, Name);
    }
    deleteButton(buttondb, Name) {
        const id = typeof buttondb === 'number' ? buttondb : buttondb.ID;
        const url = `${this.buttonsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted buttondb id=${id}`)), catchError(this.handleError('deleteButton')));
    }
    // updateFront copy button to a version with encoded pointers and update to the back
    updateFront(button, Name) {
        let buttonAPI = new ButtonAPI;
        CopyButtonToButtonAPI(button, buttonAPI);
        const id = typeof buttonAPI === 'number' ? buttonAPI : buttonAPI.ID;
        const url = `${this.buttonsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, buttonAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateButton')));
    }
    /** PUT: update the buttondb on the server */
    update(buttondb, Name, frontRepo) {
        return this.updateButton(buttondb, Name, frontRepo);
    }
    updateButton(buttondb, Name, frontRepo) {
        const id = typeof buttondb === 'number' ? buttondb : buttondb.ID;
        const url = `${this.buttonsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, buttondb, httpOptions).pipe(tap(_ => {
            // this.log(`updated buttondb id=${buttondb.ID}`)
        }), catchError(this.handleError('updateButton')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in ButtonService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("ButtonService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ButtonService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ButtonService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ButtonService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class CursorAPI {
    static GONGSTRUCT_NAME = "Cursor";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    Style = "";
    // insertion point for other decls
    CursorPointersEncoding = new CursorPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class CursorPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class CursorService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    CursorServiceChanged = new BehaviorSubject("");
    cursorsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.cursorsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/cursors';
    }
    /** GET cursors from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getCursors(Name, frontRepo);
    }
    getCursors(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.cursorsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getCursors', [])));
    }
    /** GET cursor by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getCursor(id, Name, frontRepo);
    }
    getCursor(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.cursorsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched cursor id=${id}`)),
        catchError(this.handleError(`getCursor id=${id}`)));
    }
    // postFront copy cursor to a version with encoded pointers and post to the back
    postFront(cursor, Name) {
        let cursorAPI = new CursorAPI;
        CopyCursorToCursorAPI(cursor, cursorAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.cursorsUrl, cursorAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postCursor')));
    }
    /** POST: add a new cursor to the server */
    post(cursordb, Name, frontRepo) {
        return this.postCursor(cursordb, Name, frontRepo);
    }
    postCursor(cursordb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.cursorsUrl, cursordb, httpOptions).pipe(tap(_ => {
            // this.log(`posted cursordb id=${cursordb.ID}`)
        }), catchError(this.handleError('postCursor')));
    }
    /** DELETE: delete the cursordb from the server */
    delete(cursordb, Name) {
        return this.deleteCursor(cursordb, Name);
    }
    deleteCursor(cursordb, Name) {
        const id = typeof cursordb === 'number' ? cursordb : cursordb.ID;
        const url = `${this.cursorsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted cursordb id=${id}`)), catchError(this.handleError('deleteCursor')));
    }
    // updateFront copy cursor to a version with encoded pointers and update to the back
    updateFront(cursor, Name) {
        let cursorAPI = new CursorAPI;
        CopyCursorToCursorAPI(cursor, cursorAPI);
        const id = typeof cursorAPI === 'number' ? cursorAPI : cursorAPI.ID;
        const url = `${this.cursorsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, cursorAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateCursor')));
    }
    /** PUT: update the cursordb on the server */
    update(cursordb, Name, frontRepo) {
        return this.updateCursor(cursordb, Name, frontRepo);
    }
    updateCursor(cursordb, Name, frontRepo) {
        const id = typeof cursordb === 'number' ? cursordb : cursordb.ID;
        const url = `${this.cursorsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, cursordb, httpOptions).pipe(tap(_ => {
            // this.log(`updated cursordb id=${cursordb.ID}`)
        }), catchError(this.handleError('updateCursor')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in CursorService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("CursorService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: CursorService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: CursorService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: CursorService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class FavIconAPI {
    static GONGSTRUCT_NAME = "FavIcon";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    SVG = "";
    // insertion point for other decls
    FavIconPointersEncoding = new FavIconPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class FavIconPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class FavIconService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    FavIconServiceChanged = new BehaviorSubject("");
    faviconsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.faviconsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/favicons';
    }
    /** GET favicons from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getFavIcons(Name, frontRepo);
    }
    getFavIcons(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.faviconsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getFavIcons', [])));
    }
    /** GET favicon by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getFavIcon(id, Name, frontRepo);
    }
    getFavIcon(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.faviconsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched favicon id=${id}`)),
        catchError(this.handleError(`getFavIcon id=${id}`)));
    }
    // postFront copy favicon to a version with encoded pointers and post to the back
    postFront(favicon, Name) {
        let faviconAPI = new FavIconAPI;
        CopyFavIconToFavIconAPI(favicon, faviconAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.faviconsUrl, faviconAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postFavIcon')));
    }
    /** POST: add a new favicon to the server */
    post(favicondb, Name, frontRepo) {
        return this.postFavIcon(favicondb, Name, frontRepo);
    }
    postFavIcon(favicondb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.faviconsUrl, favicondb, httpOptions).pipe(tap(_ => {
            // this.log(`posted favicondb id=${favicondb.ID}`)
        }), catchError(this.handleError('postFavIcon')));
    }
    /** DELETE: delete the favicondb from the server */
    delete(favicondb, Name) {
        return this.deleteFavIcon(favicondb, Name);
    }
    deleteFavIcon(favicondb, Name) {
        const id = typeof favicondb === 'number' ? favicondb : favicondb.ID;
        const url = `${this.faviconsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted favicondb id=${id}`)), catchError(this.handleError('deleteFavIcon')));
    }
    // updateFront copy favicon to a version with encoded pointers and update to the back
    updateFront(favicon, Name) {
        let faviconAPI = new FavIconAPI;
        CopyFavIconToFavIconAPI(favicon, faviconAPI);
        const id = typeof faviconAPI === 'number' ? faviconAPI : faviconAPI.ID;
        const url = `${this.faviconsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, faviconAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateFavIcon')));
    }
    /** PUT: update the favicondb on the server */
    update(favicondb, Name, frontRepo) {
        return this.updateFavIcon(favicondb, Name, frontRepo);
    }
    updateFavIcon(favicondb, Name, frontRepo) {
        const id = typeof favicondb === 'number' ? favicondb : favicondb.ID;
        const url = `${this.faviconsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, favicondb, httpOptions).pipe(tap(_ => {
            // this.log(`updated favicondb id=${favicondb.ID}`)
        }), catchError(this.handleError('updateFavIcon')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in FavIconService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("FavIconService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: FavIconService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: FavIconService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: FavIconService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class FormAPI {
    static GONGSTRUCT_NAME = "Form";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    FormPointersEncoding = new FormPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class FormPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class FormService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    FormServiceChanged = new BehaviorSubject("");
    formsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.formsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/forms';
    }
    /** GET forms from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getForms(Name, frontRepo);
    }
    getForms(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.formsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getForms', [])));
    }
    /** GET form by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getForm(id, Name, frontRepo);
    }
    getForm(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.formsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched form id=${id}`)),
        catchError(this.handleError(`getForm id=${id}`)));
    }
    // postFront copy form to a version with encoded pointers and post to the back
    postFront(form, Name) {
        let formAPI = new FormAPI;
        CopyFormToFormAPI(form, formAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.formsUrl, formAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postForm')));
    }
    /** POST: add a new form to the server */
    post(formdb, Name, frontRepo) {
        return this.postForm(formdb, Name, frontRepo);
    }
    postForm(formdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.formsUrl, formdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted formdb id=${formdb.ID}`)
        }), catchError(this.handleError('postForm')));
    }
    /** DELETE: delete the formdb from the server */
    delete(formdb, Name) {
        return this.deleteForm(formdb, Name);
    }
    deleteForm(formdb, Name) {
        const id = typeof formdb === 'number' ? formdb : formdb.ID;
        const url = `${this.formsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted formdb id=${id}`)), catchError(this.handleError('deleteForm')));
    }
    // updateFront copy form to a version with encoded pointers and update to the back
    updateFront(form, Name) {
        let formAPI = new FormAPI;
        CopyFormToFormAPI(form, formAPI);
        const id = typeof formAPI === 'number' ? formAPI : formAPI.ID;
        const url = `${this.formsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, formAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateForm')));
    }
    /** PUT: update the formdb on the server */
    update(formdb, Name, frontRepo) {
        return this.updateForm(formdb, Name, frontRepo);
    }
    updateForm(formdb, Name, frontRepo) {
        const id = typeof formdb === 'number' ? formdb : formdb.ID;
        const url = `${this.formsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, formdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated formdb id=${formdb.ID}`)
        }), catchError(this.handleError('updateForm')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in FormService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("FormService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: FormService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: FormService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: FormService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class LoadAPI {
    static GONGSTRUCT_NAME = "Load";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    LoadPointersEncoding = new LoadPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class LoadPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class LoadService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    LoadServiceChanged = new BehaviorSubject("");
    loadsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.loadsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/loads';
    }
    /** GET loads from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getLoads(Name, frontRepo);
    }
    getLoads(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.loadsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getLoads', [])));
    }
    /** GET load by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getLoad(id, Name, frontRepo);
    }
    getLoad(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.loadsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched load id=${id}`)),
        catchError(this.handleError(`getLoad id=${id}`)));
    }
    // postFront copy load to a version with encoded pointers and post to the back
    postFront(load, Name) {
        let loadAPI = new LoadAPI;
        CopyLoadToLoadAPI(load, loadAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.loadsUrl, loadAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postLoad')));
    }
    /** POST: add a new load to the server */
    post(loaddb, Name, frontRepo) {
        return this.postLoad(loaddb, Name, frontRepo);
    }
    postLoad(loaddb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.loadsUrl, loaddb, httpOptions).pipe(tap(_ => {
            // this.log(`posted loaddb id=${loaddb.ID}`)
        }), catchError(this.handleError('postLoad')));
    }
    /** DELETE: delete the loaddb from the server */
    delete(loaddb, Name) {
        return this.deleteLoad(loaddb, Name);
    }
    deleteLoad(loaddb, Name) {
        const id = typeof loaddb === 'number' ? loaddb : loaddb.ID;
        const url = `${this.loadsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted loaddb id=${id}`)), catchError(this.handleError('deleteLoad')));
    }
    // updateFront copy load to a version with encoded pointers and update to the back
    updateFront(load, Name) {
        let loadAPI = new LoadAPI;
        CopyLoadToLoadAPI(load, loadAPI);
        const id = typeof loadAPI === 'number' ? loadAPI : loadAPI.ID;
        const url = `${this.loadsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, loadAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateLoad')));
    }
    /** PUT: update the loaddb on the server */
    update(loaddb, Name, frontRepo) {
        return this.updateLoad(loaddb, Name, frontRepo);
    }
    updateLoad(loaddb, Name, frontRepo) {
        const id = typeof loaddb === 'number' ? loaddb : loaddb.ID;
        const url = `${this.loadsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, loaddb, httpOptions).pipe(tap(_ => {
            // this.log(`updated loaddb id=${loaddb.ID}`)
        }), catchError(this.handleError('updateLoad')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in LoadService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("LoadService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: LoadService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: LoadService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: LoadService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class LogoOnTheLeftAPI {
    static GONGSTRUCT_NAME = "LogoOnTheLeft";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Width = 0;
    Height = 0;
    SVG = "";
    // insertion point for other decls
    LogoOnTheLeftPointersEncoding = new LogoOnTheLeftPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class LogoOnTheLeftPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class LogoOnTheLeftService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    LogoOnTheLeftServiceChanged = new BehaviorSubject("");
    logoontheleftsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.logoontheleftsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/logoonthelefts';
    }
    /** GET logoonthelefts from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getLogoOnTheLefts(Name, frontRepo);
    }
    getLogoOnTheLefts(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.logoontheleftsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getLogoOnTheLefts', [])));
    }
    /** GET logoontheleft by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getLogoOnTheLeft(id, Name, frontRepo);
    }
    getLogoOnTheLeft(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.logoontheleftsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched logoontheleft id=${id}`)),
        catchError(this.handleError(`getLogoOnTheLeft id=${id}`)));
    }
    // postFront copy logoontheleft to a version with encoded pointers and post to the back
    postFront(logoontheleft, Name) {
        let logoontheleftAPI = new LogoOnTheLeftAPI;
        CopyLogoOnTheLeftToLogoOnTheLeftAPI(logoontheleft, logoontheleftAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.logoontheleftsUrl, logoontheleftAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postLogoOnTheLeft')));
    }
    /** POST: add a new logoontheleft to the server */
    post(logoontheleftdb, Name, frontRepo) {
        return this.postLogoOnTheLeft(logoontheleftdb, Name, frontRepo);
    }
    postLogoOnTheLeft(logoontheleftdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.logoontheleftsUrl, logoontheleftdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted logoontheleftdb id=${logoontheleftdb.ID}`)
        }), catchError(this.handleError('postLogoOnTheLeft')));
    }
    /** DELETE: delete the logoontheleftdb from the server */
    delete(logoontheleftdb, Name) {
        return this.deleteLogoOnTheLeft(logoontheleftdb, Name);
    }
    deleteLogoOnTheLeft(logoontheleftdb, Name) {
        const id = typeof logoontheleftdb === 'number' ? logoontheleftdb : logoontheleftdb.ID;
        const url = `${this.logoontheleftsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted logoontheleftdb id=${id}`)), catchError(this.handleError('deleteLogoOnTheLeft')));
    }
    // updateFront copy logoontheleft to a version with encoded pointers and update to the back
    updateFront(logoontheleft, Name) {
        let logoontheleftAPI = new LogoOnTheLeftAPI;
        CopyLogoOnTheLeftToLogoOnTheLeftAPI(logoontheleft, logoontheleftAPI);
        const id = typeof logoontheleftAPI === 'number' ? logoontheleftAPI : logoontheleftAPI.ID;
        const url = `${this.logoontheleftsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, logoontheleftAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateLogoOnTheLeft')));
    }
    /** PUT: update the logoontheleftdb on the server */
    update(logoontheleftdb, Name, frontRepo) {
        return this.updateLogoOnTheLeft(logoontheleftdb, Name, frontRepo);
    }
    updateLogoOnTheLeft(logoontheleftdb, Name, frontRepo) {
        const id = typeof logoontheleftdb === 'number' ? logoontheleftdb : logoontheleftdb.ID;
        const url = `${this.logoontheleftsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, logoontheleftdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated logoontheleftdb id=${logoontheleftdb.ID}`)
        }), catchError(this.handleError('updateLogoOnTheLeft')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in LogoOnTheLeftService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("LogoOnTheLeftService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: LogoOnTheLeftService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: LogoOnTheLeftService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: LogoOnTheLeftService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class LogoOnTheRightAPI {
    static GONGSTRUCT_NAME = "LogoOnTheRight";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    Width = 0;
    Height = 0;
    SVG = "";
    // insertion point for other decls
    LogoOnTheRightPointersEncoding = new LogoOnTheRightPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class LogoOnTheRightPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class LogoOnTheRightService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    LogoOnTheRightServiceChanged = new BehaviorSubject("");
    logoontherightsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.logoontherightsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/logoontherights';
    }
    /** GET logoontherights from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getLogoOnTheRights(Name, frontRepo);
    }
    getLogoOnTheRights(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.logoontherightsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getLogoOnTheRights', [])));
    }
    /** GET logoontheright by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getLogoOnTheRight(id, Name, frontRepo);
    }
    getLogoOnTheRight(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.logoontherightsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched logoontheright id=${id}`)),
        catchError(this.handleError(`getLogoOnTheRight id=${id}`)));
    }
    // postFront copy logoontheright to a version with encoded pointers and post to the back
    postFront(logoontheright, Name) {
        let logoontherightAPI = new LogoOnTheRightAPI;
        CopyLogoOnTheRightToLogoOnTheRightAPI(logoontheright, logoontherightAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.logoontherightsUrl, logoontherightAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postLogoOnTheRight')));
    }
    /** POST: add a new logoontheright to the server */
    post(logoontherightdb, Name, frontRepo) {
        return this.postLogoOnTheRight(logoontherightdb, Name, frontRepo);
    }
    postLogoOnTheRight(logoontherightdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.logoontherightsUrl, logoontherightdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted logoontherightdb id=${logoontherightdb.ID}`)
        }), catchError(this.handleError('postLogoOnTheRight')));
    }
    /** DELETE: delete the logoontherightdb from the server */
    delete(logoontherightdb, Name) {
        return this.deleteLogoOnTheRight(logoontherightdb, Name);
    }
    deleteLogoOnTheRight(logoontherightdb, Name) {
        const id = typeof logoontherightdb === 'number' ? logoontherightdb : logoontherightdb.ID;
        const url = `${this.logoontherightsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted logoontherightdb id=${id}`)), catchError(this.handleError('deleteLogoOnTheRight')));
    }
    // updateFront copy logoontheright to a version with encoded pointers and update to the back
    updateFront(logoontheright, Name) {
        let logoontherightAPI = new LogoOnTheRightAPI;
        CopyLogoOnTheRightToLogoOnTheRightAPI(logoontheright, logoontherightAPI);
        const id = typeof logoontherightAPI === 'number' ? logoontherightAPI : logoontherightAPI.ID;
        const url = `${this.logoontherightsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, logoontherightAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateLogoOnTheRight')));
    }
    /** PUT: update the logoontherightdb on the server */
    update(logoontherightdb, Name, frontRepo) {
        return this.updateLogoOnTheRight(logoontherightdb, Name, frontRepo);
    }
    updateLogoOnTheRight(logoontherightdb, Name, frontRepo) {
        const id = typeof logoontherightdb === 'number' ? logoontherightdb : logoontherightdb.ID;
        const url = `${this.logoontherightsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, logoontherightdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated logoontherightdb id=${logoontherightdb.ID}`)
        }), catchError(this.handleError('updateLogoOnTheRight')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in LogoOnTheRightService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("LogoOnTheRightService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: LogoOnTheRightService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: LogoOnTheRightService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: LogoOnTheRightService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class MarkdownAPI {
    static GONGSTRUCT_NAME = "Markdown";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    MarkdownPointersEncoding = new MarkdownPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class MarkdownPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class MarkdownService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    MarkdownServiceChanged = new BehaviorSubject("");
    markdownsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.markdownsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/markdowns';
    }
    /** GET markdowns from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getMarkdowns(Name, frontRepo);
    }
    getMarkdowns(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.markdownsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getMarkdowns', [])));
    }
    /** GET markdown by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getMarkdown(id, Name, frontRepo);
    }
    getMarkdown(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.markdownsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched markdown id=${id}`)),
        catchError(this.handleError(`getMarkdown id=${id}`)));
    }
    // postFront copy markdown to a version with encoded pointers and post to the back
    postFront(markdown, Name) {
        let markdownAPI = new MarkdownAPI;
        CopyMarkdownToMarkdownAPI(markdown, markdownAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.markdownsUrl, markdownAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postMarkdown')));
    }
    /** POST: add a new markdown to the server */
    post(markdowndb, Name, frontRepo) {
        return this.postMarkdown(markdowndb, Name, frontRepo);
    }
    postMarkdown(markdowndb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.markdownsUrl, markdowndb, httpOptions).pipe(tap(_ => {
            // this.log(`posted markdowndb id=${markdowndb.ID}`)
        }), catchError(this.handleError('postMarkdown')));
    }
    /** DELETE: delete the markdowndb from the server */
    delete(markdowndb, Name) {
        return this.deleteMarkdown(markdowndb, Name);
    }
    deleteMarkdown(markdowndb, Name) {
        const id = typeof markdowndb === 'number' ? markdowndb : markdowndb.ID;
        const url = `${this.markdownsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted markdowndb id=${id}`)), catchError(this.handleError('deleteMarkdown')));
    }
    // updateFront copy markdown to a version with encoded pointers and update to the back
    updateFront(markdown, Name) {
        let markdownAPI = new MarkdownAPI;
        CopyMarkdownToMarkdownAPI(markdown, markdownAPI);
        const id = typeof markdownAPI === 'number' ? markdownAPI : markdownAPI.ID;
        const url = `${this.markdownsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, markdownAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateMarkdown')));
    }
    /** PUT: update the markdowndb on the server */
    update(markdowndb, Name, frontRepo) {
        return this.updateMarkdown(markdowndb, Name, frontRepo);
    }
    updateMarkdown(markdowndb, Name, frontRepo) {
        const id = typeof markdowndb === 'number' ? markdowndb : markdowndb.ID;
        const url = `${this.markdownsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, markdowndb, httpOptions).pipe(tap(_ => {
            // this.log(`updated markdowndb id=${markdowndb.ID}`)
        }), catchError(this.handleError('updateMarkdown')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in MarkdownService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("MarkdownService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: MarkdownService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: MarkdownService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: MarkdownService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class SliderAPI {
    static GONGSTRUCT_NAME = "Slider";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    SliderPointersEncoding = new SliderPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class SliderPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class SliderService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    SliderServiceChanged = new BehaviorSubject("");
    slidersUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.slidersUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/sliders';
    }
    /** GET sliders from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getSliders(Name, frontRepo);
    }
    getSliders(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.slidersUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getSliders', [])));
    }
    /** GET slider by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getSlider(id, Name, frontRepo);
    }
    getSlider(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.slidersUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched slider id=${id}`)),
        catchError(this.handleError(`getSlider id=${id}`)));
    }
    // postFront copy slider to a version with encoded pointers and post to the back
    postFront(slider, Name) {
        let sliderAPI = new SliderAPI;
        CopySliderToSliderAPI(slider, sliderAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.slidersUrl, sliderAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postSlider')));
    }
    /** POST: add a new slider to the server */
    post(sliderdb, Name, frontRepo) {
        return this.postSlider(sliderdb, Name, frontRepo);
    }
    postSlider(sliderdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.slidersUrl, sliderdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted sliderdb id=${sliderdb.ID}`)
        }), catchError(this.handleError('postSlider')));
    }
    /** DELETE: delete the sliderdb from the server */
    delete(sliderdb, Name) {
        return this.deleteSlider(sliderdb, Name);
    }
    deleteSlider(sliderdb, Name) {
        const id = typeof sliderdb === 'number' ? sliderdb : sliderdb.ID;
        const url = `${this.slidersUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted sliderdb id=${id}`)), catchError(this.handleError('deleteSlider')));
    }
    // updateFront copy slider to a version with encoded pointers and update to the back
    updateFront(slider, Name) {
        let sliderAPI = new SliderAPI;
        CopySliderToSliderAPI(slider, sliderAPI);
        const id = typeof sliderAPI === 'number' ? sliderAPI : sliderAPI.ID;
        const url = `${this.slidersUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, sliderAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateSlider')));
    }
    /** PUT: update the sliderdb on the server */
    update(sliderdb, Name, frontRepo) {
        return this.updateSlider(sliderdb, Name, frontRepo);
    }
    updateSlider(sliderdb, Name, frontRepo) {
        const id = typeof sliderdb === 'number' ? sliderdb : sliderdb.ID;
        const url = `${this.slidersUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, sliderdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated sliderdb id=${sliderdb.ID}`)
        }), catchError(this.handleError('updateSlider')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in SliderService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("SliderService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SliderService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SliderService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SliderService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class SplitAPI {
    static GONGSTRUCT_NAME = "Split";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    SplitPointersEncoding = new SplitPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class SplitPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class SplitService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    SplitServiceChanged = new BehaviorSubject("");
    splitsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.splitsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/splits';
    }
    /** GET splits from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getSplits(Name, frontRepo);
    }
    getSplits(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.splitsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getSplits', [])));
    }
    /** GET split by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getSplit(id, Name, frontRepo);
    }
    getSplit(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.splitsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched split id=${id}`)),
        catchError(this.handleError(`getSplit id=${id}`)));
    }
    // postFront copy split to a version with encoded pointers and post to the back
    postFront(split, Name) {
        let splitAPI = new SplitAPI;
        CopySplitToSplitAPI(split, splitAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.splitsUrl, splitAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postSplit')));
    }
    /** POST: add a new split to the server */
    post(splitdb, Name, frontRepo) {
        return this.postSplit(splitdb, Name, frontRepo);
    }
    postSplit(splitdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.splitsUrl, splitdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted splitdb id=${splitdb.ID}`)
        }), catchError(this.handleError('postSplit')));
    }
    /** DELETE: delete the splitdb from the server */
    delete(splitdb, Name) {
        return this.deleteSplit(splitdb, Name);
    }
    deleteSplit(splitdb, Name) {
        const id = typeof splitdb === 'number' ? splitdb : splitdb.ID;
        const url = `${this.splitsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted splitdb id=${id}`)), catchError(this.handleError('deleteSplit')));
    }
    // updateFront copy split to a version with encoded pointers and update to the back
    updateFront(split, Name) {
        let splitAPI = new SplitAPI;
        CopySplitToSplitAPI(split, splitAPI);
        const id = typeof splitAPI === 'number' ? splitAPI : splitAPI.ID;
        const url = `${this.splitsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, splitAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateSplit')));
    }
    /** PUT: update the splitdb on the server */
    update(splitdb, Name, frontRepo) {
        return this.updateSplit(splitdb, Name, frontRepo);
    }
    updateSplit(splitdb, Name, frontRepo) {
        const id = typeof splitdb === 'number' ? splitdb : splitdb.ID;
        const url = `${this.splitsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, splitdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated splitdb id=${splitdb.ID}`)
        }), catchError(this.handleError('updateSplit')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in SplitService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("SplitService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SplitService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SplitService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SplitService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class SvgAPI {
    static GONGSTRUCT_NAME = "Svg";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    Style = "";
    // insertion point for other decls
    SvgPointersEncoding = new SvgPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class SvgPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class SvgService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    SvgServiceChanged = new BehaviorSubject("");
    svgsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.svgsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/svgs';
    }
    /** GET svgs from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getSvgs(Name, frontRepo);
    }
    getSvgs(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.svgsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getSvgs', [])));
    }
    /** GET svg by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getSvg(id, Name, frontRepo);
    }
    getSvg(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.svgsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched svg id=${id}`)),
        catchError(this.handleError(`getSvg id=${id}`)));
    }
    // postFront copy svg to a version with encoded pointers and post to the back
    postFront(svg, Name) {
        let svgAPI = new SvgAPI;
        CopySvgToSvgAPI(svg, svgAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.svgsUrl, svgAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postSvg')));
    }
    /** POST: add a new svg to the server */
    post(svgdb, Name, frontRepo) {
        return this.postSvg(svgdb, Name, frontRepo);
    }
    postSvg(svgdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.svgsUrl, svgdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted svgdb id=${svgdb.ID}`)
        }), catchError(this.handleError('postSvg')));
    }
    /** DELETE: delete the svgdb from the server */
    delete(svgdb, Name) {
        return this.deleteSvg(svgdb, Name);
    }
    deleteSvg(svgdb, Name) {
        const id = typeof svgdb === 'number' ? svgdb : svgdb.ID;
        const url = `${this.svgsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted svgdb id=${id}`)), catchError(this.handleError('deleteSvg')));
    }
    // updateFront copy svg to a version with encoded pointers and update to the back
    updateFront(svg, Name) {
        let svgAPI = new SvgAPI;
        CopySvgToSvgAPI(svg, svgAPI);
        const id = typeof svgAPI === 'number' ? svgAPI : svgAPI.ID;
        const url = `${this.svgsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, svgAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateSvg')));
    }
    /** PUT: update the svgdb on the server */
    update(svgdb, Name, frontRepo) {
        return this.updateSvg(svgdb, Name, frontRepo);
    }
    updateSvg(svgdb, Name, frontRepo) {
        const id = typeof svgdb === 'number' ? svgdb : svgdb.ID;
        const url = `${this.svgsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, svgdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated svgdb id=${svgdb.ID}`)
        }), catchError(this.handleError('updateSvg')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in SvgService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("SvgService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SvgService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SvgService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: SvgService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class TableAPI {
    static GONGSTRUCT_NAME = "Table";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    TablePointersEncoding = new TablePointersEncoding;
    CreatedAt;
    DeletedAt;
}
class TablePointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class TableService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    TableServiceChanged = new BehaviorSubject("");
    tablesUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.tablesUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/tables';
    }
    /** GET tables from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getTables(Name, frontRepo);
    }
    getTables(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.tablesUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getTables', [])));
    }
    /** GET table by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getTable(id, Name, frontRepo);
    }
    getTable(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.tablesUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched table id=${id}`)),
        catchError(this.handleError(`getTable id=${id}`)));
    }
    // postFront copy table to a version with encoded pointers and post to the back
    postFront(table, Name) {
        let tableAPI = new TableAPI;
        CopyTableToTableAPI(table, tableAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.tablesUrl, tableAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postTable')));
    }
    /** POST: add a new table to the server */
    post(tabledb, Name, frontRepo) {
        return this.postTable(tabledb, Name, frontRepo);
    }
    postTable(tabledb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.tablesUrl, tabledb, httpOptions).pipe(tap(_ => {
            // this.log(`posted tabledb id=${tabledb.ID}`)
        }), catchError(this.handleError('postTable')));
    }
    /** DELETE: delete the tabledb from the server */
    delete(tabledb, Name) {
        return this.deleteTable(tabledb, Name);
    }
    deleteTable(tabledb, Name) {
        const id = typeof tabledb === 'number' ? tabledb : tabledb.ID;
        const url = `${this.tablesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted tabledb id=${id}`)), catchError(this.handleError('deleteTable')));
    }
    // updateFront copy table to a version with encoded pointers and update to the back
    updateFront(table, Name) {
        let tableAPI = new TableAPI;
        CopyTableToTableAPI(table, tableAPI);
        const id = typeof tableAPI === 'number' ? tableAPI : tableAPI.ID;
        const url = `${this.tablesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, tableAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateTable')));
    }
    /** PUT: update the tabledb on the server */
    update(tabledb, Name, frontRepo) {
        return this.updateTable(tabledb, Name, frontRepo);
    }
    updateTable(tabledb, Name, frontRepo) {
        const id = typeof tabledb === 'number' ? tabledb : tabledb.ID;
        const url = `${this.tablesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, tabledb, httpOptions).pipe(tap(_ => {
            // this.log(`updated tabledb id=${tabledb.ID}`)
        }), catchError(this.handleError('updateTable')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in TableService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("TableService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: TableService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: TableService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: TableService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class ThreejsAPI {
    static GONGSTRUCT_NAME = "Threejs";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    ThreejsPointersEncoding = new ThreejsPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class ThreejsPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class ThreejsService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    ThreejsServiceChanged = new BehaviorSubject("");
    threejssUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.threejssUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/threejss';
    }
    /** GET threejss from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getThreejss(Name, frontRepo);
    }
    getThreejss(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.threejssUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getThreejss', [])));
    }
    /** GET threejs by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getThreejs(id, Name, frontRepo);
    }
    getThreejs(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.threejssUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched threejs id=${id}`)),
        catchError(this.handleError(`getThreejs id=${id}`)));
    }
    // postFront copy threejs to a version with encoded pointers and post to the back
    postFront(threejs, Name) {
        let threejsAPI = new ThreejsAPI;
        CopyThreejsToThreejsAPI(threejs, threejsAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.threejssUrl, threejsAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postThreejs')));
    }
    /** POST: add a new threejs to the server */
    post(threejsdb, Name, frontRepo) {
        return this.postThreejs(threejsdb, Name, frontRepo);
    }
    postThreejs(threejsdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.threejssUrl, threejsdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted threejsdb id=${threejsdb.ID}`)
        }), catchError(this.handleError('postThreejs')));
    }
    /** DELETE: delete the threejsdb from the server */
    delete(threejsdb, Name) {
        return this.deleteThreejs(threejsdb, Name);
    }
    deleteThreejs(threejsdb, Name) {
        const id = typeof threejsdb === 'number' ? threejsdb : threejsdb.ID;
        const url = `${this.threejssUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted threejsdb id=${id}`)), catchError(this.handleError('deleteThreejs')));
    }
    // updateFront copy threejs to a version with encoded pointers and update to the back
    updateFront(threejs, Name) {
        let threejsAPI = new ThreejsAPI;
        CopyThreejsToThreejsAPI(threejs, threejsAPI);
        const id = typeof threejsAPI === 'number' ? threejsAPI : threejsAPI.ID;
        const url = `${this.threejssUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, threejsAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateThreejs')));
    }
    /** PUT: update the threejsdb on the server */
    update(threejsdb, Name, frontRepo) {
        return this.updateThreejs(threejsdb, Name, frontRepo);
    }
    updateThreejs(threejsdb, Name, frontRepo) {
        const id = typeof threejsdb === 'number' ? threejsdb : threejsdb.ID;
        const url = `${this.threejssUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, threejsdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated threejsdb id=${threejsdb.ID}`)
        }), catchError(this.handleError('updateThreejs')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in ThreejsService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("ThreejsService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ThreejsService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ThreejsService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ThreejsService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class TitleAPI {
    static GONGSTRUCT_NAME = "Title";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    // insertion point for other decls
    TitlePointersEncoding = new TitlePointersEncoding;
    CreatedAt;
    DeletedAt;
}
class TitlePointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class TitleService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    TitleServiceChanged = new BehaviorSubject("");
    titlesUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.titlesUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/titles';
    }
    /** GET titles from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getTitles(Name, frontRepo);
    }
    getTitles(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.titlesUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getTitles', [])));
    }
    /** GET title by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getTitle(id, Name, frontRepo);
    }
    getTitle(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.titlesUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched title id=${id}`)),
        catchError(this.handleError(`getTitle id=${id}`)));
    }
    // postFront copy title to a version with encoded pointers and post to the back
    postFront(title, Name) {
        let titleAPI = new TitleAPI;
        CopyTitleToTitleAPI(title, titleAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.titlesUrl, titleAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postTitle')));
    }
    /** POST: add a new title to the server */
    post(titledb, Name, frontRepo) {
        return this.postTitle(titledb, Name, frontRepo);
    }
    postTitle(titledb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.titlesUrl, titledb, httpOptions).pipe(tap(_ => {
            // this.log(`posted titledb id=${titledb.ID}`)
        }), catchError(this.handleError('postTitle')));
    }
    /** DELETE: delete the titledb from the server */
    delete(titledb, Name) {
        return this.deleteTitle(titledb, Name);
    }
    deleteTitle(titledb, Name) {
        const id = typeof titledb === 'number' ? titledb : titledb.ID;
        const url = `${this.titlesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted titledb id=${id}`)), catchError(this.handleError('deleteTitle')));
    }
    // updateFront copy title to a version with encoded pointers and update to the back
    updateFront(title, Name) {
        let titleAPI = new TitleAPI;
        CopyTitleToTitleAPI(title, titleAPI);
        const id = typeof titleAPI === 'number' ? titleAPI : titleAPI.ID;
        const url = `${this.titlesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, titleAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateTitle')));
    }
    /** PUT: update the titledb on the server */
    update(titledb, Name, frontRepo) {
        return this.updateTitle(titledb, Name, frontRepo);
    }
    updateTitle(titledb, Name, frontRepo) {
        const id = typeof titledb === 'number' ? titledb : titledb.ID;
        const url = `${this.titlesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, titledb, httpOptions).pipe(tap(_ => {
            // this.log(`updated titledb id=${titledb.ID}`)
        }), catchError(this.handleError('updateTitle')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in TitleService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("TitleService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: TitleService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: TitleService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: TitleService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class ToneAPI {
    static GONGSTRUCT_NAME = "Tone";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    TonePointersEncoding = new TonePointersEncoding;
    CreatedAt;
    DeletedAt;
}
class TonePointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class ToneService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    ToneServiceChanged = new BehaviorSubject("");
    tonesUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.tonesUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/tones';
    }
    /** GET tones from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getTones(Name, frontRepo);
    }
    getTones(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.tonesUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getTones', [])));
    }
    /** GET tone by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getTone(id, Name, frontRepo);
    }
    getTone(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.tonesUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched tone id=${id}`)),
        catchError(this.handleError(`getTone id=${id}`)));
    }
    // postFront copy tone to a version with encoded pointers and post to the back
    postFront(tone, Name) {
        let toneAPI = new ToneAPI;
        CopyToneToToneAPI(tone, toneAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.tonesUrl, toneAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postTone')));
    }
    /** POST: add a new tone to the server */
    post(tonedb, Name, frontRepo) {
        return this.postTone(tonedb, Name, frontRepo);
    }
    postTone(tonedb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.tonesUrl, tonedb, httpOptions).pipe(tap(_ => {
            // this.log(`posted tonedb id=${tonedb.ID}`)
        }), catchError(this.handleError('postTone')));
    }
    /** DELETE: delete the tonedb from the server */
    delete(tonedb, Name) {
        return this.deleteTone(tonedb, Name);
    }
    deleteTone(tonedb, Name) {
        const id = typeof tonedb === 'number' ? tonedb : tonedb.ID;
        const url = `${this.tonesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted tonedb id=${id}`)), catchError(this.handleError('deleteTone')));
    }
    // updateFront copy tone to a version with encoded pointers and update to the back
    updateFront(tone, Name) {
        let toneAPI = new ToneAPI;
        CopyToneToToneAPI(tone, toneAPI);
        const id = typeof toneAPI === 'number' ? toneAPI : toneAPI.ID;
        const url = `${this.tonesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, toneAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateTone')));
    }
    /** PUT: update the tonedb on the server */
    update(tonedb, Name, frontRepo) {
        return this.updateTone(tonedb, Name, frontRepo);
    }
    updateTone(tonedb, Name, frontRepo) {
        const id = typeof tonedb === 'number' ? tonedb : tonedb.ID;
        const url = `${this.tonesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, tonedb, httpOptions).pipe(tap(_ => {
            // this.log(`updated tonedb id=${tonedb.ID}`)
        }), catchError(this.handleError('updateTone')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in ToneService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("ToneService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ToneService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ToneService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ToneService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class TreeAPI {
    static GONGSTRUCT_NAME = "Tree";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    TreePointersEncoding = new TreePointersEncoding;
    CreatedAt;
    DeletedAt;
}
class TreePointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class TreeService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    TreeServiceChanged = new BehaviorSubject("");
    treesUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.treesUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/trees';
    }
    /** GET trees from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getTrees(Name, frontRepo);
    }
    getTrees(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.treesUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getTrees', [])));
    }
    /** GET tree by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getTree(id, Name, frontRepo);
    }
    getTree(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.treesUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched tree id=${id}`)),
        catchError(this.handleError(`getTree id=${id}`)));
    }
    // postFront copy tree to a version with encoded pointers and post to the back
    postFront(tree, Name) {
        let treeAPI = new TreeAPI;
        CopyTreeToTreeAPI(tree, treeAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.treesUrl, treeAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postTree')));
    }
    /** POST: add a new tree to the server */
    post(treedb, Name, frontRepo) {
        return this.postTree(treedb, Name, frontRepo);
    }
    postTree(treedb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.treesUrl, treedb, httpOptions).pipe(tap(_ => {
            // this.log(`posted treedb id=${treedb.ID}`)
        }), catchError(this.handleError('postTree')));
    }
    /** DELETE: delete the treedb from the server */
    delete(treedb, Name) {
        return this.deleteTree(treedb, Name);
    }
    deleteTree(treedb, Name) {
        const id = typeof treedb === 'number' ? treedb : treedb.ID;
        const url = `${this.treesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted treedb id=${id}`)), catchError(this.handleError('deleteTree')));
    }
    // updateFront copy tree to a version with encoded pointers and update to the back
    updateFront(tree, Name) {
        let treeAPI = new TreeAPI;
        CopyTreeToTreeAPI(tree, treeAPI);
        const id = typeof treeAPI === 'number' ? treeAPI : treeAPI.ID;
        const url = `${this.treesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, treeAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateTree')));
    }
    /** PUT: update the treedb on the server */
    update(treedb, Name, frontRepo) {
        return this.updateTree(treedb, Name, frontRepo);
    }
    updateTree(treedb, Name, frontRepo) {
        const id = typeof treedb === 'number' ? treedb : treedb.ID;
        const url = `${this.treesUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, treedb, httpOptions).pipe(tap(_ => {
            // this.log(`updated treedb id=${treedb.ID}`)
        }), catchError(this.handleError('updateTree')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in TreeService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("TreeService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: TreeService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: TreeService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: TreeService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

class ViewAPI {
    static GONGSTRUCT_NAME = "View";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    ShowViewName = false;
    IsSelectedView = false;
    Direction = "";
    IsSecondaryView = false;
    IsSizeInPixel = false;
    IsWithCustomGutterSize = false;
    GutterSize = 0;
    // insertion point for other decls
    ViewPointersEncoding = new ViewPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class ViewPointersEncoding {
    // insertion point for pointers and slices of pointers encoding fields
    RootAsSplitAreas = [];
}

// generated code, do not edit
// generated by ng_file_service_ts
class ViewService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    ViewServiceChanged = new BehaviorSubject("");
    viewsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.viewsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/views';
    }
    /** GET views from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getViews(Name, frontRepo);
    }
    getViews(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.viewsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getViews', [])));
    }
    /** GET view by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getView(id, Name, frontRepo);
    }
    getView(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.viewsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched view id=${id}`)),
        catchError(this.handleError(`getView id=${id}`)));
    }
    // postFront copy view to a version with encoded pointers and post to the back
    postFront(view, Name) {
        let viewAPI = new ViewAPI;
        CopyViewToViewAPI(view, viewAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.viewsUrl, viewAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postView')));
    }
    /** POST: add a new view to the server */
    post(viewdb, Name, frontRepo) {
        return this.postView(viewdb, Name, frontRepo);
    }
    postView(viewdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.viewsUrl, viewdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted viewdb id=${viewdb.ID}`)
        }), catchError(this.handleError('postView')));
    }
    /** DELETE: delete the viewdb from the server */
    delete(viewdb, Name) {
        return this.deleteView(viewdb, Name);
    }
    deleteView(viewdb, Name) {
        const id = typeof viewdb === 'number' ? viewdb : viewdb.ID;
        const url = `${this.viewsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted viewdb id=${id}`)), catchError(this.handleError('deleteView')));
    }
    // updateFront copy view to a version with encoded pointers and update to the back
    updateFront(view, Name) {
        let viewAPI = new ViewAPI;
        CopyViewToViewAPI(view, viewAPI);
        const id = typeof viewAPI === 'number' ? viewAPI : viewAPI.ID;
        const url = `${this.viewsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, viewAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateView')));
    }
    /** PUT: update the viewdb on the server */
    update(viewdb, Name, frontRepo) {
        return this.updateView(viewdb, Name, frontRepo);
    }
    updateView(viewdb, Name, frontRepo) {
        const id = typeof viewdb === 'number' ? viewdb : viewdb.ID;
        const url = `${this.viewsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, viewdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated viewdb id=${viewdb.ID}`)
        }), catchError(this.handleError('updateView')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in ViewService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("ViewService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ViewService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ViewService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: ViewService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// insertion point for imports
class XlsxAPI {
    static GONGSTRUCT_NAME = "Xlsx";
    ID = 0;
    // insertion point for basic fields declarations
    Name = "";
    StackName = "";
    // insertion point for other decls
    XlsxPointersEncoding = new XlsxPointersEncoding;
    CreatedAt;
    DeletedAt;
}
class XlsxPointersEncoding {
}

// generated code, do not edit
// generated by ng_file_service_ts
// insertion point for imports
class XlsxService {
    http;
    document;
    // Kamar Raïmo: Adding a way to communicate between components that share information
    // so that they are notified of a change.
    XlsxServiceChanged = new BehaviorSubject("");
    xlsxsUrl;
    constructor(http, document) {
        this.http = http;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.xlsxsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/xlsxs';
    }
    /** GET xlsxs from the server */
    // gets is more robust to refactoring
    gets(Name, frontRepo) {
        return this.getXlsxs(Name, frontRepo);
    }
    getXlsxs(Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        return this.http.get(this.xlsxsUrl, { params: params })
            .pipe(tap(), catchError(this.handleError('getXlsxs', [])));
    }
    /** GET xlsx by id. Will 404 if id not found */
    // more robust API to refactoring
    get(id, Name, frontRepo) {
        return this.getXlsx(id, Name, frontRepo);
    }
    getXlsx(id, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        const url = `${this.xlsxsUrl}/${id}`;
        return this.http.get(url, { params: params }).pipe(
        // tap(_ => this.log(`fetched xlsx id=${id}`)),
        catchError(this.handleError(`getXlsx id=${id}`)));
    }
    // postFront copy xlsx to a version with encoded pointers and post to the back
    postFront(xlsx, Name) {
        let xlsxAPI = new XlsxAPI;
        CopyXlsxToXlsxAPI(xlsx, xlsxAPI);
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.xlsxsUrl, xlsxAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('postXlsx')));
    }
    /** POST: add a new xlsx to the server */
    post(xlsxdb, Name, frontRepo) {
        return this.postXlsx(xlsxdb, Name, frontRepo);
    }
    postXlsx(xlsxdb, Name, frontRepo) {
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.post(this.xlsxsUrl, xlsxdb, httpOptions).pipe(tap(_ => {
            // this.log(`posted xlsxdb id=${xlsxdb.ID}`)
        }), catchError(this.handleError('postXlsx')));
    }
    /** DELETE: delete the xlsxdb from the server */
    delete(xlsxdb, Name) {
        return this.deleteXlsx(xlsxdb, Name);
    }
    deleteXlsx(xlsxdb, Name) {
        const id = typeof xlsxdb === 'number' ? xlsxdb : xlsxdb.ID;
        const url = `${this.xlsxsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.delete(url, httpOptions).pipe(tap(_ => this.log(`deleted xlsxdb id=${id}`)), catchError(this.handleError('deleteXlsx')));
    }
    // updateFront copy xlsx to a version with encoded pointers and update to the back
    updateFront(xlsx, Name) {
        let xlsxAPI = new XlsxAPI;
        CopyXlsxToXlsxAPI(xlsx, xlsxAPI);
        const id = typeof xlsxAPI === 'number' ? xlsxAPI : xlsxAPI.ID;
        const url = `${this.xlsxsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, xlsxAPI, httpOptions).pipe(tap(_ => {
        }), catchError(this.handleError('updateXlsx')));
    }
    /** PUT: update the xlsxdb on the server */
    update(xlsxdb, Name, frontRepo) {
        return this.updateXlsx(xlsxdb, Name, frontRepo);
    }
    updateXlsx(xlsxdb, Name, frontRepo) {
        const id = typeof xlsxdb === 'number' ? xlsxdb : xlsxdb.ID;
        const url = `${this.xlsxsUrl}/${id}`;
        let params = new HttpParams().set("Name", Name);
        let httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            params: params
        };
        return this.http.put(url, xlsxdb, httpOptions).pipe(tap(_ => {
            // this.log(`updated xlsxdb id=${xlsxdb.ID}`)
        }), catchError(this.handleError('updateXlsx')));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in XlsxService', result) {
        return (error) => {
            // TODO: send the error to remote logging
            console.error("XlsxService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: XlsxService, deps: [{ token: i1.HttpClient }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: XlsxService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: XlsxService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

const StackType = "github.com/fullstack-lang/gong/lib/split/go/models";
// FrontRepo stores all instances in a front repository (design pattern repository)
class FrontRepo {
    array_AsSplits = new Array(); // array of front instances
    map_ID_AsSplit = new Map(); // map of front instances
    array_AsSplitAreas = new Array(); // array of front instances
    map_ID_AsSplitArea = new Map(); // map of front instances
    array_Buttons = new Array(); // array of front instances
    map_ID_Button = new Map(); // map of front instances
    array_Cursors = new Array(); // array of front instances
    map_ID_Cursor = new Map(); // map of front instances
    array_FavIcons = new Array(); // array of front instances
    map_ID_FavIcon = new Map(); // map of front instances
    array_Forms = new Array(); // array of front instances
    map_ID_Form = new Map(); // map of front instances
    array_Loads = new Array(); // array of front instances
    map_ID_Load = new Map(); // map of front instances
    array_LogoOnTheLefts = new Array(); // array of front instances
    map_ID_LogoOnTheLeft = new Map(); // map of front instances
    array_LogoOnTheRights = new Array(); // array of front instances
    map_ID_LogoOnTheRight = new Map(); // map of front instances
    array_Markdowns = new Array(); // array of front instances
    map_ID_Markdown = new Map(); // map of front instances
    array_Sliders = new Array(); // array of front instances
    map_ID_Slider = new Map(); // map of front instances
    array_Splits = new Array(); // array of front instances
    map_ID_Split = new Map(); // map of front instances
    array_Svgs = new Array(); // array of front instances
    map_ID_Svg = new Map(); // map of front instances
    array_Tables = new Array(); // array of front instances
    map_ID_Table = new Map(); // map of front instances
    array_Threejss = new Array(); // array of front instances
    map_ID_Threejs = new Map(); // map of front instances
    array_Titles = new Array(); // array of front instances
    map_ID_Title = new Map(); // map of front instances
    array_Tones = new Array(); // array of front instances
    map_ID_Tone = new Map(); // map of front instances
    array_Trees = new Array(); // array of front instances
    map_ID_Tree = new Map(); // map of front instances
    array_Views = new Array(); // array of front instances
    map_ID_View = new Map(); // map of front instances
    array_Xlsxs = new Array(); // array of front instances
    map_ID_Xlsx = new Map(); // map of front instances
    GONG__Index = -1;
    // getFrontArray allows for a get function that is robust to refactoring of the named struct name
    // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
    // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
    getFrontArray(gongStructName) {
        switch (gongStructName) {
            // insertion point
            case 'AsSplit':
                return this.array_AsSplits;
            case 'AsSplitArea':
                return this.array_AsSplitAreas;
            case 'Button':
                return this.array_Buttons;
            case 'Cursor':
                return this.array_Cursors;
            case 'FavIcon':
                return this.array_FavIcons;
            case 'Form':
                return this.array_Forms;
            case 'Load':
                return this.array_Loads;
            case 'LogoOnTheLeft':
                return this.array_LogoOnTheLefts;
            case 'LogoOnTheRight':
                return this.array_LogoOnTheRights;
            case 'Markdown':
                return this.array_Markdowns;
            case 'Slider':
                return this.array_Sliders;
            case 'Split':
                return this.array_Splits;
            case 'Svg':
                return this.array_Svgs;
            case 'Table':
                return this.array_Tables;
            case 'Threejs':
                return this.array_Threejss;
            case 'Title':
                return this.array_Titles;
            case 'Tone':
                return this.array_Tones;
            case 'Tree':
                return this.array_Trees;
            case 'View':
                return this.array_Views;
            case 'Xlsx':
                return this.array_Xlsxs;
            default:
                throw new Error("Type not recognized");
        }
    }
    getFrontMap(gongStructName) {
        switch (gongStructName) {
            // insertion point
            case 'AsSplit':
                return this.map_ID_AsSplit;
            case 'AsSplitArea':
                return this.map_ID_AsSplitArea;
            case 'Button':
                return this.map_ID_Button;
            case 'Cursor':
                return this.map_ID_Cursor;
            case 'FavIcon':
                return this.map_ID_FavIcon;
            case 'Form':
                return this.map_ID_Form;
            case 'Load':
                return this.map_ID_Load;
            case 'LogoOnTheLeft':
                return this.map_ID_LogoOnTheLeft;
            case 'LogoOnTheRight':
                return this.map_ID_LogoOnTheRight;
            case 'Markdown':
                return this.map_ID_Markdown;
            case 'Slider':
                return this.map_ID_Slider;
            case 'Split':
                return this.map_ID_Split;
            case 'Svg':
                return this.map_ID_Svg;
            case 'Table':
                return this.map_ID_Table;
            case 'Threejs':
                return this.map_ID_Threejs;
            case 'Title':
                return this.map_ID_Title;
            case 'Tone':
                return this.map_ID_Tone;
            case 'Tree':
                return this.map_ID_Tree;
            case 'View':
                return this.map_ID_View;
            case 'Xlsx':
                return this.map_ID_Xlsx;
            default:
                throw new Error("Type not recognized");
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
class DialogData {
    ID = 0; // ID of the calling instance
    // the reverse pointer is the name of the generated field on the destination
    // struct of the ONE-MANY association
    ReversePointer = ""; // field of {{Structname}} that serve as reverse pointer
    OrderingMode = false; // if true, this is for ordering items
    // there are different selection mode : ONE_MANY or MANY_MANY
    SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE;
    // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
    //
    // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
    // 
    // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
    // at the end of the ONE-MANY association
    SourceStruct = ""; // The "Aclass"
    SourceField = ""; // the "AnarrayofbUse"
    IntermediateStruct = ""; // the "AclassBclassUse" 
    IntermediateStructField = ""; // the "Bclass" as field
    NextAssociationStruct = ""; // the "Bclass"
    Name = "";
}
var SelectionMode;
(function (SelectionMode) {
    SelectionMode["ONE_MANY_ASSOCIATION_MODE"] = "ONE_MANY_ASSOCIATION_MODE";
    SelectionMode["MANY_MANY_ASSOCIATION_MODE"] = "MANY_MANY_ASSOCIATION_MODE";
})(SelectionMode || (SelectionMode = {}));
//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
class FrontRepoService {
    http;
    assplitService;
    assplitareaService;
    buttonService;
    cursorService;
    faviconService;
    formService;
    loadService;
    logoontheleftService;
    logoontherightService;
    markdownService;
    sliderService;
    splitService;
    svgService;
    tableService;
    threejsService;
    titleService;
    toneService;
    treeService;
    viewService;
    xlsxService;
    Name = "";
    httpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    };
    //
    // Store of all instances of the stack
    //
    frontRepo = new (FrontRepo);
    // Manage open WebSocket connections
    webSocketConnections = new Map();
    constructor(http, // insertion point sub template 
    assplitService, assplitareaService, buttonService, cursorService, faviconService, formService, loadService, logoontheleftService, logoontherightService, markdownService, sliderService, splitService, svgService, tableService, threejsService, titleService, toneService, treeService, viewService, xlsxService) {
        this.http = http;
        this.assplitService = assplitService;
        this.assplitareaService = assplitareaService;
        this.buttonService = buttonService;
        this.cursorService = cursorService;
        this.faviconService = faviconService;
        this.formService = formService;
        this.loadService = loadService;
        this.logoontheleftService = logoontheleftService;
        this.logoontherightService = logoontherightService;
        this.markdownService = markdownService;
        this.sliderService = sliderService;
        this.splitService = splitService;
        this.svgService = svgService;
        this.tableService = tableService;
        this.threejsService = threejsService;
        this.titleService = titleService;
        this.toneService = toneService;
        this.treeService = treeService;
        this.viewService = viewService;
        this.xlsxService = xlsxService;
    }
    // postService provides a post function for each struct name
    postService(structName, instanceToBePosted) {
        let service = this[structName.toLowerCase() + "Service" + "Service"];
        let servicePostFunction = service[("post" + structName)];
        servicePostFunction(instanceToBePosted).subscribe(instance => {
            let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged")];
            behaviorSubject.next("post");
        });
    }
    // deleteService provides a delete function for each struct name
    deleteService(structName, instanceToBeDeleted) {
        let service = this[structName.toLowerCase() + "Service"];
        let serviceDeleteFunction = service["delete" + structName];
        serviceDeleteFunction(instanceToBeDeleted).subscribe(instance => {
            let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged")];
            behaviorSubject.next("delete");
        });
    }
    // typing of observable can be messy in typescript. Therefore, one force the type
    observableFrontRepo;
    //
    // pull performs a GET on all struct of the stack and redeem association pointers 
    //
    // This is an observable. Therefore, the control flow forks with
    // - pull() return immediatly the observable
    // - the observable observer, if it subscribe, is called when all GET calls are performs
    pull(Name = "") {
        this.Name = Name;
        this.observableFrontRepo = [
            of(null), // see above for justification
            // insertion point sub template
            this.assplitService.getAsSplits(this.Name, this.frontRepo),
            this.assplitareaService.getAsSplitAreas(this.Name, this.frontRepo),
            this.buttonService.getButtons(this.Name, this.frontRepo),
            this.cursorService.getCursors(this.Name, this.frontRepo),
            this.faviconService.getFavIcons(this.Name, this.frontRepo),
            this.formService.getForms(this.Name, this.frontRepo),
            this.loadService.getLoads(this.Name, this.frontRepo),
            this.logoontheleftService.getLogoOnTheLefts(this.Name, this.frontRepo),
            this.logoontherightService.getLogoOnTheRights(this.Name, this.frontRepo),
            this.markdownService.getMarkdowns(this.Name, this.frontRepo),
            this.sliderService.getSliders(this.Name, this.frontRepo),
            this.splitService.getSplits(this.Name, this.frontRepo),
            this.svgService.getSvgs(this.Name, this.frontRepo),
            this.tableService.getTables(this.Name, this.frontRepo),
            this.threejsService.getThreejss(this.Name, this.frontRepo),
            this.titleService.getTitles(this.Name, this.frontRepo),
            this.toneService.getTones(this.Name, this.frontRepo),
            this.treeService.getTrees(this.Name, this.frontRepo),
            this.viewService.getViews(this.Name, this.frontRepo),
            this.xlsxService.getXlsxs(this.Name, this.frontRepo),
        ];
        return new Observable((observer) => {
            combineLatest(this.observableFrontRepo).subscribe(([___of_null, // see above for the explanation about of
            // insertion point sub template for declarations 
            assplits_, assplitareas_, buttons_, cursors_, favicons_, forms_, loads_, logoonthelefts_, logoontherights_, markdowns_, sliders_, splits_, svgs_, tables_, threejss_, titles_, tones_, trees_, views_, xlsxs_,]) => {
                let _this = this;
                // Typing can be messy with many items. Therefore, type casting is necessary here
                // insertion point sub template for type casting 
                var assplits;
                assplits = assplits_;
                var assplitareas;
                assplitareas = assplitareas_;
                var buttons;
                buttons = buttons_;
                var cursors;
                cursors = cursors_;
                var favicons;
                favicons = favicons_;
                var forms;
                forms = forms_;
                var loads;
                loads = loads_;
                var logoonthelefts;
                logoonthelefts = logoonthelefts_;
                var logoontherights;
                logoontherights = logoontherights_;
                var markdowns;
                markdowns = markdowns_;
                var sliders;
                sliders = sliders_;
                var splits;
                splits = splits_;
                var svgs;
                svgs = svgs_;
                var tables;
                tables = tables_;
                var threejss;
                threejss = threejss_;
                var titles;
                titles = titles_;
                var tones;
                tones = tones_;
                var trees;
                trees = trees_;
                var views;
                views = views_;
                var xlsxs;
                xlsxs = xlsxs_;
                // 
                // First Step: init map of instances
                // insertion point sub template for init 
                // init the arrays
                this.frontRepo.array_AsSplits = [];
                this.frontRepo.map_ID_AsSplit.clear();
                assplits.forEach(assplitAPI => {
                    let assplit = new AsSplit;
                    this.frontRepo.array_AsSplits.push(assplit);
                    this.frontRepo.map_ID_AsSplit.set(assplitAPI.ID, assplit);
                });
                // init the arrays
                this.frontRepo.array_AsSplitAreas = [];
                this.frontRepo.map_ID_AsSplitArea.clear();
                assplitareas.forEach(assplitareaAPI => {
                    let assplitarea = new AsSplitArea;
                    this.frontRepo.array_AsSplitAreas.push(assplitarea);
                    this.frontRepo.map_ID_AsSplitArea.set(assplitareaAPI.ID, assplitarea);
                });
                // init the arrays
                this.frontRepo.array_Buttons = [];
                this.frontRepo.map_ID_Button.clear();
                buttons.forEach(buttonAPI => {
                    let button = new Button;
                    this.frontRepo.array_Buttons.push(button);
                    this.frontRepo.map_ID_Button.set(buttonAPI.ID, button);
                });
                // init the arrays
                this.frontRepo.array_Cursors = [];
                this.frontRepo.map_ID_Cursor.clear();
                cursors.forEach(cursorAPI => {
                    let cursor = new Cursor;
                    this.frontRepo.array_Cursors.push(cursor);
                    this.frontRepo.map_ID_Cursor.set(cursorAPI.ID, cursor);
                });
                // init the arrays
                this.frontRepo.array_FavIcons = [];
                this.frontRepo.map_ID_FavIcon.clear();
                favicons.forEach(faviconAPI => {
                    let favicon = new FavIcon;
                    this.frontRepo.array_FavIcons.push(favicon);
                    this.frontRepo.map_ID_FavIcon.set(faviconAPI.ID, favicon);
                });
                // init the arrays
                this.frontRepo.array_Forms = [];
                this.frontRepo.map_ID_Form.clear();
                forms.forEach(formAPI => {
                    let form = new Form;
                    this.frontRepo.array_Forms.push(form);
                    this.frontRepo.map_ID_Form.set(formAPI.ID, form);
                });
                // init the arrays
                this.frontRepo.array_Loads = [];
                this.frontRepo.map_ID_Load.clear();
                loads.forEach(loadAPI => {
                    let load = new Load;
                    this.frontRepo.array_Loads.push(load);
                    this.frontRepo.map_ID_Load.set(loadAPI.ID, load);
                });
                // init the arrays
                this.frontRepo.array_LogoOnTheLefts = [];
                this.frontRepo.map_ID_LogoOnTheLeft.clear();
                logoonthelefts.forEach(logoontheleftAPI => {
                    let logoontheleft = new LogoOnTheLeft;
                    this.frontRepo.array_LogoOnTheLefts.push(logoontheleft);
                    this.frontRepo.map_ID_LogoOnTheLeft.set(logoontheleftAPI.ID, logoontheleft);
                });
                // init the arrays
                this.frontRepo.array_LogoOnTheRights = [];
                this.frontRepo.map_ID_LogoOnTheRight.clear();
                logoontherights.forEach(logoontherightAPI => {
                    let logoontheright = new LogoOnTheRight;
                    this.frontRepo.array_LogoOnTheRights.push(logoontheright);
                    this.frontRepo.map_ID_LogoOnTheRight.set(logoontherightAPI.ID, logoontheright);
                });
                // init the arrays
                this.frontRepo.array_Markdowns = [];
                this.frontRepo.map_ID_Markdown.clear();
                markdowns.forEach(markdownAPI => {
                    let markdown = new Markdown;
                    this.frontRepo.array_Markdowns.push(markdown);
                    this.frontRepo.map_ID_Markdown.set(markdownAPI.ID, markdown);
                });
                // init the arrays
                this.frontRepo.array_Sliders = [];
                this.frontRepo.map_ID_Slider.clear();
                sliders.forEach(sliderAPI => {
                    let slider = new Slider;
                    this.frontRepo.array_Sliders.push(slider);
                    this.frontRepo.map_ID_Slider.set(sliderAPI.ID, slider);
                });
                // init the arrays
                this.frontRepo.array_Splits = [];
                this.frontRepo.map_ID_Split.clear();
                splits.forEach(splitAPI => {
                    let split = new Split;
                    this.frontRepo.array_Splits.push(split);
                    this.frontRepo.map_ID_Split.set(splitAPI.ID, split);
                });
                // init the arrays
                this.frontRepo.array_Svgs = [];
                this.frontRepo.map_ID_Svg.clear();
                svgs.forEach(svgAPI => {
                    let svg = new Svg;
                    this.frontRepo.array_Svgs.push(svg);
                    this.frontRepo.map_ID_Svg.set(svgAPI.ID, svg);
                });
                // init the arrays
                this.frontRepo.array_Tables = [];
                this.frontRepo.map_ID_Table.clear();
                tables.forEach(tableAPI => {
                    let table = new Table;
                    this.frontRepo.array_Tables.push(table);
                    this.frontRepo.map_ID_Table.set(tableAPI.ID, table);
                });
                // init the arrays
                this.frontRepo.array_Threejss = [];
                this.frontRepo.map_ID_Threejs.clear();
                threejss.forEach(threejsAPI => {
                    let threejs = new Threejs;
                    this.frontRepo.array_Threejss.push(threejs);
                    this.frontRepo.map_ID_Threejs.set(threejsAPI.ID, threejs);
                });
                // init the arrays
                this.frontRepo.array_Titles = [];
                this.frontRepo.map_ID_Title.clear();
                titles.forEach(titleAPI => {
                    let title = new Title;
                    this.frontRepo.array_Titles.push(title);
                    this.frontRepo.map_ID_Title.set(titleAPI.ID, title);
                });
                // init the arrays
                this.frontRepo.array_Tones = [];
                this.frontRepo.map_ID_Tone.clear();
                tones.forEach(toneAPI => {
                    let tone = new Tone;
                    this.frontRepo.array_Tones.push(tone);
                    this.frontRepo.map_ID_Tone.set(toneAPI.ID, tone);
                });
                // init the arrays
                this.frontRepo.array_Trees = [];
                this.frontRepo.map_ID_Tree.clear();
                trees.forEach(treeAPI => {
                    let tree = new Tree;
                    this.frontRepo.array_Trees.push(tree);
                    this.frontRepo.map_ID_Tree.set(treeAPI.ID, tree);
                });
                // init the arrays
                this.frontRepo.array_Views = [];
                this.frontRepo.map_ID_View.clear();
                views.forEach(viewAPI => {
                    let view = new View;
                    this.frontRepo.array_Views.push(view);
                    this.frontRepo.map_ID_View.set(viewAPI.ID, view);
                });
                // init the arrays
                this.frontRepo.array_Xlsxs = [];
                this.frontRepo.map_ID_Xlsx.clear();
                xlsxs.forEach(xlsxAPI => {
                    let xlsx = new Xlsx;
                    this.frontRepo.array_Xlsxs.push(xlsx);
                    this.frontRepo.map_ID_Xlsx.set(xlsxAPI.ID, xlsx);
                });
                // 
                // Second Step: reddeem front objects
                // insertion point sub template for redeem 
                // fill up front objects
                assplits.forEach(assplitAPI => {
                    let assplit = this.frontRepo.map_ID_AsSplit.get(assplitAPI.ID);
                    CopyAsSplitAPIToAsSplit(assplitAPI, assplit, this.frontRepo);
                });
                // fill up front objects
                assplitareas.forEach(assplitareaAPI => {
                    let assplitarea = this.frontRepo.map_ID_AsSplitArea.get(assplitareaAPI.ID);
                    CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI, assplitarea, this.frontRepo);
                });
                // fill up front objects
                buttons.forEach(buttonAPI => {
                    let button = this.frontRepo.map_ID_Button.get(buttonAPI.ID);
                    CopyButtonAPIToButton(buttonAPI, button, this.frontRepo);
                });
                // fill up front objects
                cursors.forEach(cursorAPI => {
                    let cursor = this.frontRepo.map_ID_Cursor.get(cursorAPI.ID);
                    CopyCursorAPIToCursor(cursorAPI, cursor, this.frontRepo);
                });
                // fill up front objects
                favicons.forEach(faviconAPI => {
                    let favicon = this.frontRepo.map_ID_FavIcon.get(faviconAPI.ID);
                    CopyFavIconAPIToFavIcon(faviconAPI, favicon, this.frontRepo);
                });
                // fill up front objects
                forms.forEach(formAPI => {
                    let form = this.frontRepo.map_ID_Form.get(formAPI.ID);
                    CopyFormAPIToForm(formAPI, form, this.frontRepo);
                });
                // fill up front objects
                loads.forEach(loadAPI => {
                    let load = this.frontRepo.map_ID_Load.get(loadAPI.ID);
                    CopyLoadAPIToLoad(loadAPI, load, this.frontRepo);
                });
                // fill up front objects
                logoonthelefts.forEach(logoontheleftAPI => {
                    let logoontheleft = this.frontRepo.map_ID_LogoOnTheLeft.get(logoontheleftAPI.ID);
                    CopyLogoOnTheLeftAPIToLogoOnTheLeft(logoontheleftAPI, logoontheleft, this.frontRepo);
                });
                // fill up front objects
                logoontherights.forEach(logoontherightAPI => {
                    let logoontheright = this.frontRepo.map_ID_LogoOnTheRight.get(logoontherightAPI.ID);
                    CopyLogoOnTheRightAPIToLogoOnTheRight(logoontherightAPI, logoontheright, this.frontRepo);
                });
                // fill up front objects
                markdowns.forEach(markdownAPI => {
                    let markdown = this.frontRepo.map_ID_Markdown.get(markdownAPI.ID);
                    CopyMarkdownAPIToMarkdown(markdownAPI, markdown, this.frontRepo);
                });
                // fill up front objects
                sliders.forEach(sliderAPI => {
                    let slider = this.frontRepo.map_ID_Slider.get(sliderAPI.ID);
                    CopySliderAPIToSlider(sliderAPI, slider, this.frontRepo);
                });
                // fill up front objects
                splits.forEach(splitAPI => {
                    let split = this.frontRepo.map_ID_Split.get(splitAPI.ID);
                    CopySplitAPIToSplit(splitAPI, split, this.frontRepo);
                });
                // fill up front objects
                svgs.forEach(svgAPI => {
                    let svg = this.frontRepo.map_ID_Svg.get(svgAPI.ID);
                    CopySvgAPIToSvg(svgAPI, svg, this.frontRepo);
                });
                // fill up front objects
                tables.forEach(tableAPI => {
                    let table = this.frontRepo.map_ID_Table.get(tableAPI.ID);
                    CopyTableAPIToTable(tableAPI, table, this.frontRepo);
                });
                // fill up front objects
                threejss.forEach(threejsAPI => {
                    let threejs = this.frontRepo.map_ID_Threejs.get(threejsAPI.ID);
                    CopyThreejsAPIToThreejs(threejsAPI, threejs, this.frontRepo);
                });
                // fill up front objects
                titles.forEach(titleAPI => {
                    let title = this.frontRepo.map_ID_Title.get(titleAPI.ID);
                    CopyTitleAPIToTitle(titleAPI, title, this.frontRepo);
                });
                // fill up front objects
                tones.forEach(toneAPI => {
                    let tone = this.frontRepo.map_ID_Tone.get(toneAPI.ID);
                    CopyToneAPIToTone(toneAPI, tone, this.frontRepo);
                });
                // fill up front objects
                trees.forEach(treeAPI => {
                    let tree = this.frontRepo.map_ID_Tree.get(treeAPI.ID);
                    CopyTreeAPIToTree(treeAPI, tree, this.frontRepo);
                });
                // fill up front objects
                views.forEach(viewAPI => {
                    let view = this.frontRepo.map_ID_View.get(viewAPI.ID);
                    CopyViewAPIToView(viewAPI, view, this.frontRepo);
                });
                // fill up front objects
                xlsxs.forEach(xlsxAPI => {
                    let xlsx = this.frontRepo.map_ID_Xlsx.get(xlsxAPI.ID);
                    CopyXlsxAPIToXlsx(xlsxAPI, xlsx, this.frontRepo);
                });
                // hand over control flow to observer
                observer.next(this.frontRepo);
            });
        });
    }
    connectToWebSocket(Name) {
        // console.log("github.com/fullstack-lang/gong/lib/split/go; connectToWebSocket: started", Name)
        // Check if a connection for this name already exists
        if (this.webSocketConnections.has(Name)) {
            // console.log("github.com/fullstack-lang/gong/lib/split/go; connectToWebSocket: returning existing connection")
            return this.webSocketConnections.get(Name);
        }
        //
        // Create a new connection
        //
        let host = window.location.host;
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        if (host === 'localhost:4200') {
            host = 'localhost:8080';
        }
        // Construct the base path using the dynamic host and protocol
        // The API path remains the same.
        let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/split/go/v1/ws/stage`;
        let params = new HttpParams().set("Name", Name);
        let paramString = params.toString();
        let url = `${basePath}?${paramString}`;
        const newConnection$ = new Observable(observer => {
            // console.log("github.com/fullstack-lang/gong/lib/split/go; connectToWebSocket: new Observable created")
            let socket;
            const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null;
            const processData = (dataString) => {
                // console.log("github.com/fullstack-lang/gong/lib/split/go; connectToWebSocket: processData called")
                const backRepoData = new BackRepoData(JSON.parse(dataString));
                let frontRepo = new (FrontRepo)();
                frontRepo.GONG__Index = backRepoData.GONG__Index;
                // 
                // First Step: init map of instances
                // insertion point sub template for init 
                // init the arrays
                frontRepo.array_AsSplits = [];
                frontRepo.map_ID_AsSplit.clear();
                backRepoData.AsSplitAPIs.forEach(assplitAPI => {
                    let assplit = new AsSplit;
                    frontRepo.array_AsSplits.push(assplit);
                    frontRepo.map_ID_AsSplit.set(assplitAPI.ID, assplit);
                });
                // init the arrays
                frontRepo.array_AsSplitAreas = [];
                frontRepo.map_ID_AsSplitArea.clear();
                backRepoData.AsSplitAreaAPIs.forEach(assplitareaAPI => {
                    let assplitarea = new AsSplitArea;
                    frontRepo.array_AsSplitAreas.push(assplitarea);
                    frontRepo.map_ID_AsSplitArea.set(assplitareaAPI.ID, assplitarea);
                });
                // init the arrays
                frontRepo.array_Buttons = [];
                frontRepo.map_ID_Button.clear();
                backRepoData.ButtonAPIs.forEach(buttonAPI => {
                    let button = new Button;
                    frontRepo.array_Buttons.push(button);
                    frontRepo.map_ID_Button.set(buttonAPI.ID, button);
                });
                // init the arrays
                frontRepo.array_Cursors = [];
                frontRepo.map_ID_Cursor.clear();
                backRepoData.CursorAPIs.forEach(cursorAPI => {
                    let cursor = new Cursor;
                    frontRepo.array_Cursors.push(cursor);
                    frontRepo.map_ID_Cursor.set(cursorAPI.ID, cursor);
                });
                // init the arrays
                frontRepo.array_FavIcons = [];
                frontRepo.map_ID_FavIcon.clear();
                backRepoData.FavIconAPIs.forEach(faviconAPI => {
                    let favicon = new FavIcon;
                    frontRepo.array_FavIcons.push(favicon);
                    frontRepo.map_ID_FavIcon.set(faviconAPI.ID, favicon);
                });
                // init the arrays
                frontRepo.array_Forms = [];
                frontRepo.map_ID_Form.clear();
                backRepoData.FormAPIs.forEach(formAPI => {
                    let form = new Form;
                    frontRepo.array_Forms.push(form);
                    frontRepo.map_ID_Form.set(formAPI.ID, form);
                });
                // init the arrays
                frontRepo.array_Loads = [];
                frontRepo.map_ID_Load.clear();
                backRepoData.LoadAPIs.forEach(loadAPI => {
                    let load = new Load;
                    frontRepo.array_Loads.push(load);
                    frontRepo.map_ID_Load.set(loadAPI.ID, load);
                });
                // init the arrays
                frontRepo.array_LogoOnTheLefts = [];
                frontRepo.map_ID_LogoOnTheLeft.clear();
                backRepoData.LogoOnTheLeftAPIs.forEach(logoontheleftAPI => {
                    let logoontheleft = new LogoOnTheLeft;
                    frontRepo.array_LogoOnTheLefts.push(logoontheleft);
                    frontRepo.map_ID_LogoOnTheLeft.set(logoontheleftAPI.ID, logoontheleft);
                });
                // init the arrays
                frontRepo.array_LogoOnTheRights = [];
                frontRepo.map_ID_LogoOnTheRight.clear();
                backRepoData.LogoOnTheRightAPIs.forEach(logoontherightAPI => {
                    let logoontheright = new LogoOnTheRight;
                    frontRepo.array_LogoOnTheRights.push(logoontheright);
                    frontRepo.map_ID_LogoOnTheRight.set(logoontherightAPI.ID, logoontheright);
                });
                // init the arrays
                frontRepo.array_Markdowns = [];
                frontRepo.map_ID_Markdown.clear();
                backRepoData.MarkdownAPIs.forEach(markdownAPI => {
                    let markdown = new Markdown;
                    frontRepo.array_Markdowns.push(markdown);
                    frontRepo.map_ID_Markdown.set(markdownAPI.ID, markdown);
                });
                // init the arrays
                frontRepo.array_Sliders = [];
                frontRepo.map_ID_Slider.clear();
                backRepoData.SliderAPIs.forEach(sliderAPI => {
                    let slider = new Slider;
                    frontRepo.array_Sliders.push(slider);
                    frontRepo.map_ID_Slider.set(sliderAPI.ID, slider);
                });
                // init the arrays
                frontRepo.array_Splits = [];
                frontRepo.map_ID_Split.clear();
                backRepoData.SplitAPIs.forEach(splitAPI => {
                    let split = new Split;
                    frontRepo.array_Splits.push(split);
                    frontRepo.map_ID_Split.set(splitAPI.ID, split);
                });
                // init the arrays
                frontRepo.array_Svgs = [];
                frontRepo.map_ID_Svg.clear();
                backRepoData.SvgAPIs.forEach(svgAPI => {
                    let svg = new Svg;
                    frontRepo.array_Svgs.push(svg);
                    frontRepo.map_ID_Svg.set(svgAPI.ID, svg);
                });
                // init the arrays
                frontRepo.array_Tables = [];
                frontRepo.map_ID_Table.clear();
                backRepoData.TableAPIs.forEach(tableAPI => {
                    let table = new Table;
                    frontRepo.array_Tables.push(table);
                    frontRepo.map_ID_Table.set(tableAPI.ID, table);
                });
                // init the arrays
                frontRepo.array_Threejss = [];
                frontRepo.map_ID_Threejs.clear();
                backRepoData.ThreejsAPIs.forEach(threejsAPI => {
                    let threejs = new Threejs;
                    frontRepo.array_Threejss.push(threejs);
                    frontRepo.map_ID_Threejs.set(threejsAPI.ID, threejs);
                });
                // init the arrays
                frontRepo.array_Titles = [];
                frontRepo.map_ID_Title.clear();
                backRepoData.TitleAPIs.forEach(titleAPI => {
                    let title = new Title;
                    frontRepo.array_Titles.push(title);
                    frontRepo.map_ID_Title.set(titleAPI.ID, title);
                });
                // init the arrays
                frontRepo.array_Tones = [];
                frontRepo.map_ID_Tone.clear();
                backRepoData.ToneAPIs.forEach(toneAPI => {
                    let tone = new Tone;
                    frontRepo.array_Tones.push(tone);
                    frontRepo.map_ID_Tone.set(toneAPI.ID, tone);
                });
                // init the arrays
                frontRepo.array_Trees = [];
                frontRepo.map_ID_Tree.clear();
                backRepoData.TreeAPIs.forEach(treeAPI => {
                    let tree = new Tree;
                    frontRepo.array_Trees.push(tree);
                    frontRepo.map_ID_Tree.set(treeAPI.ID, tree);
                });
                // init the arrays
                frontRepo.array_Views = [];
                frontRepo.map_ID_View.clear();
                backRepoData.ViewAPIs.forEach(viewAPI => {
                    let view = new View;
                    frontRepo.array_Views.push(view);
                    frontRepo.map_ID_View.set(viewAPI.ID, view);
                });
                // init the arrays
                frontRepo.array_Xlsxs = [];
                frontRepo.map_ID_Xlsx.clear();
                backRepoData.XlsxAPIs.forEach(xlsxAPI => {
                    let xlsx = new Xlsx;
                    frontRepo.array_Xlsxs.push(xlsx);
                    frontRepo.map_ID_Xlsx.set(xlsxAPI.ID, xlsx);
                });
                // 
                // Second Step: reddeem front objects
                // insertion point sub template for redeem 
                // fill up front objects
                backRepoData.AsSplitAPIs.forEach(assplitAPI => {
                    let assplit = frontRepo.map_ID_AsSplit.get(assplitAPI.ID);
                    CopyAsSplitAPIToAsSplit(assplitAPI, assplit, frontRepo);
                });
                // fill up front objects
                backRepoData.AsSplitAreaAPIs.forEach(assplitareaAPI => {
                    let assplitarea = frontRepo.map_ID_AsSplitArea.get(assplitareaAPI.ID);
                    CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI, assplitarea, frontRepo);
                });
                // fill up front objects
                backRepoData.ButtonAPIs.forEach(buttonAPI => {
                    let button = frontRepo.map_ID_Button.get(buttonAPI.ID);
                    CopyButtonAPIToButton(buttonAPI, button, frontRepo);
                });
                // fill up front objects
                backRepoData.CursorAPIs.forEach(cursorAPI => {
                    let cursor = frontRepo.map_ID_Cursor.get(cursorAPI.ID);
                    CopyCursorAPIToCursor(cursorAPI, cursor, frontRepo);
                });
                // fill up front objects
                backRepoData.FavIconAPIs.forEach(faviconAPI => {
                    let favicon = frontRepo.map_ID_FavIcon.get(faviconAPI.ID);
                    CopyFavIconAPIToFavIcon(faviconAPI, favicon, frontRepo);
                });
                // fill up front objects
                backRepoData.FormAPIs.forEach(formAPI => {
                    let form = frontRepo.map_ID_Form.get(formAPI.ID);
                    CopyFormAPIToForm(formAPI, form, frontRepo);
                });
                // fill up front objects
                backRepoData.LoadAPIs.forEach(loadAPI => {
                    let load = frontRepo.map_ID_Load.get(loadAPI.ID);
                    CopyLoadAPIToLoad(loadAPI, load, frontRepo);
                });
                // fill up front objects
                backRepoData.LogoOnTheLeftAPIs.forEach(logoontheleftAPI => {
                    let logoontheleft = frontRepo.map_ID_LogoOnTheLeft.get(logoontheleftAPI.ID);
                    CopyLogoOnTheLeftAPIToLogoOnTheLeft(logoontheleftAPI, logoontheleft, frontRepo);
                });
                // fill up front objects
                backRepoData.LogoOnTheRightAPIs.forEach(logoontherightAPI => {
                    let logoontheright = frontRepo.map_ID_LogoOnTheRight.get(logoontherightAPI.ID);
                    CopyLogoOnTheRightAPIToLogoOnTheRight(logoontherightAPI, logoontheright, frontRepo);
                });
                // fill up front objects
                backRepoData.MarkdownAPIs.forEach(markdownAPI => {
                    let markdown = frontRepo.map_ID_Markdown.get(markdownAPI.ID);
                    CopyMarkdownAPIToMarkdown(markdownAPI, markdown, frontRepo);
                });
                // fill up front objects
                backRepoData.SliderAPIs.forEach(sliderAPI => {
                    let slider = frontRepo.map_ID_Slider.get(sliderAPI.ID);
                    CopySliderAPIToSlider(sliderAPI, slider, frontRepo);
                });
                // fill up front objects
                backRepoData.SplitAPIs.forEach(splitAPI => {
                    let split = frontRepo.map_ID_Split.get(splitAPI.ID);
                    CopySplitAPIToSplit(splitAPI, split, frontRepo);
                });
                // fill up front objects
                backRepoData.SvgAPIs.forEach(svgAPI => {
                    let svg = frontRepo.map_ID_Svg.get(svgAPI.ID);
                    CopySvgAPIToSvg(svgAPI, svg, frontRepo);
                });
                // fill up front objects
                backRepoData.TableAPIs.forEach(tableAPI => {
                    let table = frontRepo.map_ID_Table.get(tableAPI.ID);
                    CopyTableAPIToTable(tableAPI, table, frontRepo);
                });
                // fill up front objects
                backRepoData.ThreejsAPIs.forEach(threejsAPI => {
                    let threejs = frontRepo.map_ID_Threejs.get(threejsAPI.ID);
                    CopyThreejsAPIToThreejs(threejsAPI, threejs, frontRepo);
                });
                // fill up front objects
                backRepoData.TitleAPIs.forEach(titleAPI => {
                    let title = frontRepo.map_ID_Title.get(titleAPI.ID);
                    CopyTitleAPIToTitle(titleAPI, title, frontRepo);
                });
                // fill up front objects
                backRepoData.ToneAPIs.forEach(toneAPI => {
                    let tone = frontRepo.map_ID_Tone.get(toneAPI.ID);
                    CopyToneAPIToTone(toneAPI, tone, frontRepo);
                });
                // fill up front objects
                backRepoData.TreeAPIs.forEach(treeAPI => {
                    let tree = frontRepo.map_ID_Tree.get(treeAPI.ID);
                    CopyTreeAPIToTree(treeAPI, tree, frontRepo);
                });
                // fill up front objects
                backRepoData.ViewAPIs.forEach(viewAPI => {
                    let view = frontRepo.map_ID_View.get(viewAPI.ID);
                    CopyViewAPIToView(viewAPI, view, frontRepo);
                });
                // fill up front objects
                backRepoData.XlsxAPIs.forEach(xlsxAPI => {
                    let xlsx = frontRepo.map_ID_Xlsx.get(xlsxAPI.ID);
                    CopyXlsxAPIToXlsx(xlsxAPI, xlsx, frontRepo);
                });
                observer.next(frontRepo);
            };
            // 3. Connection Loop
            const attemptConnection = (retries) => {
                // console.log("github.com/fullstack-lang/gong/lib/split/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)
                // A. WASM OFFLINE MODE (Check if Go is ready)
                if (window.openWasmSocket) {
                    // console.log("github.com/fullstack-lang/gong/lib/split/go; attemptConnection: openWasmSocket exists, calling it");
                    window.openWasmSocket("github.com/fullstack-lang/gong/lib/split/go", Name, processData);
                    return;
                }
                // B. WAITING FOR WASM
                if (isOfflineMode && retries > 0) {
                    // console.log("github.com/fullstack-lang/gong/lib/split/go; attemptConnection: WAITING FOR WASM. Retries left:", retries)
                    setTimeout(() => attemptConnection(retries - 1), 100);
                    return;
                }
                // C. STANDARD SERVER MODE
                if (!isOfflineMode) {
                    // console.log("github.com/fullstack-lang/gong/lib/split/go; attemptConnection: STANDARD SERVER MODE. url =", url)
                    socket = new WebSocket(url);
                    socket.onopen = (event) => {
                        // console.log("github.com/fullstack-lang/gong/lib/split/go; WebSocket: onopen", event)
                    };
                    socket.onmessage = event => {
                        // console.log("github.com/fullstack-lang/gong/lib/split/go; WebSocket: onmessage")
                        processData(event.data);
                    };
                    socket.onerror = event => {
                        console.error("github.com/fullstack-lang/gong/lib/split/go WebSocket: onerror", event);
                        observer.error(event);
                    };
                    socket.onclose = (event) => {
                        // console.log("github.com/fullstack-lang/gong/lib/split/go; WebSocket: onclose", event)
                        observer.complete();
                    };
                }
                else {
                    console.error("github.com/fullstack-lang/gong/lib/split/go, attemptConnection: Offline mode detected, but WASM backend failed to load.");
                    observer.error("Offline mode detected, but WASM backend failed to load.");
                }
            };
            attemptConnection(50);
            // Teardown logic: Called when the last subscriber unsubscribes.
            return () => {
                this.webSocketConnections.delete(Name); // Remove from cache
                if (socket) {
                    socket.close();
                }
            };
        }).pipe(
        // This is the key:
        // - shareReplay makes this a "multicast" observable, sharing the single WebSocket among subscribers.
        // - { bufferSize: 1, refCount: true } means:
        //   - bufferSize: 1 => new subscribers get the last emitted value immediately.
        //   - refCount: true => the connection starts with the first subscriber and stops with the last.
        shareReplay({ bufferSize: 1, refCount: true }));
        // Store the new connection observable in the map
        this.webSocketConnections.set(Name, newConnection$);
        return newConnection$;
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: FrontRepoService, deps: [{ token: i1.HttpClient }, { token: AsSplitService }, { token: AsSplitAreaService }, { token: ButtonService }, { token: CursorService }, { token: FavIconService }, { token: FormService }, { token: LoadService }, { token: LogoOnTheLeftService }, { token: LogoOnTheRightService }, { token: MarkdownService }, { token: SliderService }, { token: SplitService }, { token: SvgService }, { token: TableService }, { token: ThreejsService }, { token: TitleService }, { token: ToneService }, { token: TreeService }, { token: ViewService }, { token: XlsxService }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: FrontRepoService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: FrontRepoService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: AsSplitService }, { type: AsSplitAreaService }, { type: ButtonService }, { type: CursorService }, { type: FavIconService }, { type: FormService }, { type: LoadService }, { type: LogoOnTheLeftService }, { type: LogoOnTheRightService }, { type: MarkdownService }, { type: SliderService }, { type: SplitService }, { type: SvgService }, { type: TableService }, { type: ThreejsService }, { type: TitleService }, { type: ToneService }, { type: TreeService }, { type: ViewService }, { type: XlsxService }] });
// insertion point for get unique ID per struct 
function getAsSplitUniqueID(id) {
    return 31 * id;
}
function getAsSplitAreaUniqueID(id) {
    return 37 * id;
}
function getButtonUniqueID(id) {
    return 41 * id;
}
function getCursorUniqueID(id) {
    return 43 * id;
}
function getFavIconUniqueID(id) {
    return 47 * id;
}
function getFormUniqueID(id) {
    return 53 * id;
}
function getLoadUniqueID(id) {
    return 59 * id;
}
function getLogoOnTheLeftUniqueID(id) {
    return 61 * id;
}
function getLogoOnTheRightUniqueID(id) {
    return 67 * id;
}
function getMarkdownUniqueID(id) {
    return 71 * id;
}
function getSliderUniqueID(id) {
    return 73 * id;
}
function getSplitUniqueID(id) {
    return 79 * id;
}
function getSvgUniqueID(id) {
    return 83 * id;
}
function getTableUniqueID(id) {
    return 89 * id;
}
function getThreejsUniqueID(id) {
    return 97 * id;
}
function getTitleUniqueID(id) {
    return 101 * id;
}
function getToneUniqueID(id) {
    return 103 * id;
}
function getTreeUniqueID(id) {
    return 107 * id;
}
function getViewUniqueID(id) {
    return 109 * id;
}
function getXlsxUniqueID(id) {
    return 113 * id;
}

// generated file, do not edit
class CommitNbFromBackService {
    http;
    location;
    document;
    httpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    };
    commitNbFromBackUrl;
    constructor(http, location, document) {
        this.http = http;
        this.location = location;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.commitNbFromBackUrl = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/commitfrombacknb';
    }
    getCommitNbFromBack(intervalMs, Name = "") {
        let params = new HttpParams().set("Name", Name);
        return interval(intervalMs).pipe(switchMap(() => this.http.get(this.commitNbFromBackUrl, { params: params }).pipe(catchError(error => {
            // Handle the error here, e.g. log it, show a notification, etc.
            console.error('Error fetching commit number:', error);
            // Return a default value, a new Observable, or rethrow the error
            return of(0); // Here, we return 0 as a default value
        }))));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation in CommitNbFromBackService', result) {
        return (error) => {
            // TODO: send the error to remote logging infrastructure
            console.error("in CommitNbFromBackService" + error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log('${operation} failed: ${error.message}');
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
        console.log(message);
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: CommitNbFromBackService, deps: [{ token: i1.HttpClient }, { token: i2.Location }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: CommitNbFromBackService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: CommitNbFromBackService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: i2.Location }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

// generated file, do not edit
class PushFromFrontNbService {
    http;
    location;
    document;
    httpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    };
    pushFromFrontNbURL;
    constructor(http, location, document) {
        this.http = http;
        this.location = location;
        this.document = document;
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin;
        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080");
        // compute path to the service
        this.pushFromFrontNbURL = origin + '/api/github.com/fullstack-lang/gong/lib/split/go/v1/pushfromfrontnb';
    }
    // observable of the commit nb getter
    getPushFromFrontNb() {
        return this.http.get(this.pushFromFrontNbURL)
            .pipe(tap(_ => this.log('fetched commit nb')), catchError(this.handleError('getPushFromFrontNb', -1)));
    }
    getPushNbFromFront(intervalMs, Name = "") {
        let params = new HttpParams().set("Name", Name);
        return interval(intervalMs).pipe(switchMap(() => this.http.get(this.pushFromFrontNbURL, { params: params }).pipe(catchError(error => {
            // Handle the error here, e.g. log it, show a notification, etc.
            console.error('Error fetching commit number:', error);
            // Return a default value, a new Observable, or rethrow the error
            return of(0); // Here, we return 0 as a default value
        }))));
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    handleError(operation = 'operation', result) {
        return (error) => {
            // TODO: send the error to remote logging infrastructure
            console.error(error); // log to console instead
            // TODO: better job of transforming error for user consumption
            this.log('${operation} failed: ${error.message}');
            // Let the app keep running by returning an empty result.
            return of(result);
        };
    }
    log(message) {
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: PushFromFrontNbService, deps: [{ token: i1.HttpClient }, { token: i2.Location }, { token: DOCUMENT }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: PushFromFrontNbService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: PushFromFrontNbService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: i1.HttpClient }, { type: i2.Location }, { type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT]
                }] }] });

class WebSocketService {
    document;
    constructor(document) {
        this.document = document;
    }
    connect(stackPath) {
        return new Observable(subscriber => {
            // Determine if we are running in the offline HTML file
            const isOfflineMode = this.document.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null;
            const attemptConnection = (retries) => {
                // 1. WASM OFFLINE MODE (Check if Go is ready)
                if (window.openWasmSocket) {
                    const onMessageFromGo = (message) => {
                        subscriber.next(JSON.parse(message));
                    };
                    window.openWasmSocket(stackPath, onMessageFromGo);
                    return; // Successfully connected to WASM
                }
                // 2. STILL WAITING FOR WASM
                if (isOfflineMode && retries > 0) {
                    // If we are offline, WASM might just be decoding. Wait 100ms and check again.
                    setTimeout(() => attemptConnection(retries - 1), 100);
                    return;
                }
                // 3. STANDARD SERVER MODE (Fallback)
                if (!isOfflineMode) {
                    let protocol = this.document.location.protocol === 'https:' ? 'wss://' : 'ws://';
                    let port = this.document.location.port ? ':' + this.document.location.port : '';
                    let host = this.document.location.hostname;
                    const wsUrl = `${protocol}${host}${port}/api/github.com/fullstack-lang/gong/test/test4/go/v1/ws/stage?Name=${stackPath}`;
                    const ws = new WebSocket(wsUrl);
                    ws.onmessage = (event) => subscriber.next(JSON.parse(event.data));
                    ws.onerror = (error) => subscriber.error(error);
                    ws.onclose = () => subscriber.complete();
                }
                else {
                    subscriber.error("Offline mode detected, but WASM backend failed to load.");
                }
            };
            // Start the connection loop, allowing up to 5 seconds (50 * 100ms) for WASM to boot
            attemptConnection(50);
        });
    }
    static ɵfac = i0.ɵɵngDeclareFactory({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: WebSocketService, deps: [{ token: DOCUMENT$1 }], target: i0.ɵɵFactoryTarget.Injectable });
    static ɵprov = i0.ɵɵngDeclareInjectable({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: WebSocketService, providedIn: 'root' });
}
i0.ɵɵngDeclareClassMetadata({ minVersion: "12.0.0", version: "20.3.26", ngImport: i0, type: WebSocketService, decorators: [{
            type: Injectable,
            args: [{
                    providedIn: 'root'
                }]
        }], ctorParameters: () => [{ type: Document, decorators: [{
                    type: Inject,
                    args: [DOCUMENT$1]
                }] }] });

// generated from ng_file_enum.ts.go
var Direction;
(function (Direction) {
    // insertion point	
    Direction["Vertical"] = "vertical";
    Direction["Horizontal"] = "horizontal";
})(Direction || (Direction = {}));
const DirectionList = [
    { value: Direction.Vertical, viewValue: "vertical" },
    { value: Direction.Horizontal, viewValue: "horizontal" },
];

// generated from ng_file_public_api_ts.go
/*
* Public API Surface of split
*/

/**
 * Generated bundle index. Do not edit.
 */

export { AsSplit, AsSplitAPI, AsSplitArea, AsSplitAreaAPI, AsSplitAreaPointersEncoding, AsSplitAreaService, AsSplitPointersEncoding, AsSplitService, BackRepoData, Button, ButtonAPI, ButtonPointersEncoding, ButtonService, CommitNbFromBackService, CopyAsSplitAPIToAsSplit, CopyAsSplitAreaAPIToAsSplitArea, CopyAsSplitAreaToAsSplitAreaAPI, CopyAsSplitToAsSplitAPI, CopyButtonAPIToButton, CopyButtonToButtonAPI, CopyCursorAPIToCursor, CopyCursorToCursorAPI, CopyFavIconAPIToFavIcon, CopyFavIconToFavIconAPI, CopyFormAPIToForm, CopyFormToFormAPI, CopyLoadAPIToLoad, CopyLoadToLoadAPI, CopyLogoOnTheLeftAPIToLogoOnTheLeft, CopyLogoOnTheLeftToLogoOnTheLeftAPI, CopyLogoOnTheRightAPIToLogoOnTheRight, CopyLogoOnTheRightToLogoOnTheRightAPI, CopyMarkdownAPIToMarkdown, CopyMarkdownToMarkdownAPI, CopySliderAPIToSlider, CopySliderToSliderAPI, CopySplitAPIToSplit, CopySplitToSplitAPI, CopySvgAPIToSvg, CopySvgToSvgAPI, CopyTableAPIToTable, CopyTableToTableAPI, CopyThreejsAPIToThreejs, CopyThreejsToThreejsAPI, CopyTitleAPIToTitle, CopyTitleToTitleAPI, CopyToneAPIToTone, CopyToneToToneAPI, CopyTreeAPIToTree, CopyTreeToTreeAPI, CopyViewAPIToView, CopyViewToViewAPI, CopyXlsxAPIToXlsx, CopyXlsxToXlsxAPI, Cursor, CursorAPI, CursorPointersEncoding, CursorService, DialogData, Direction, DirectionList, FavIcon, FavIconAPI, FavIconPointersEncoding, FavIconService, Form, FormAPI, FormPointersEncoding, FormService, FrontRepo, FrontRepoService, Load, LoadAPI, LoadPointersEncoding, LoadService, LogoOnTheLeft, LogoOnTheLeftAPI, LogoOnTheLeftPointersEncoding, LogoOnTheLeftService, LogoOnTheRight, LogoOnTheRightAPI, LogoOnTheRightPointersEncoding, LogoOnTheRightService, Markdown, MarkdownAPI, MarkdownPointersEncoding, MarkdownService, NullInt64, PushFromFrontNbService, SelectionMode, Slider, SliderAPI, SliderPointersEncoding, SliderService, Split, SplitAPI, SplitModule, SplitPointersEncoding, SplitService, StackType, Svg, SvgAPI, SvgPointersEncoding, SvgService, Table, TableAPI, TablePointersEncoding, TableService, Threejs, ThreejsAPI, ThreejsPointersEncoding, ThreejsService, Title, TitleAPI, TitlePointersEncoding, TitleService, Tone, ToneAPI, TonePointersEncoding, ToneService, Tree, TreeAPI, TreePointersEncoding, TreeService, View, ViewAPI, ViewPointersEncoding, ViewService, WebSocketService, Xlsx, XlsxAPI, XlsxPointersEncoding, XlsxService, getAsSplitAreaUniqueID, getAsSplitUniqueID, getButtonUniqueID, getCursorUniqueID, getFavIconUniqueID, getFormUniqueID, getLoadUniqueID, getLogoOnTheLeftUniqueID, getLogoOnTheRightUniqueID, getMarkdownUniqueID, getSliderUniqueID, getSplitUniqueID, getSvgUniqueID, getTableUniqueID, getThreejsUniqueID, getTitleUniqueID, getToneUniqueID, getTreeUniqueID, getViewUniqueID, getXlsxUniqueID };
//# sourceMappingURL=split.mjs.map
