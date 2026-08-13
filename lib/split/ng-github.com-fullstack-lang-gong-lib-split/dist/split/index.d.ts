import * as i0 from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { Location } from '@angular/common';

declare class SplitModule {
    static ɵfac: i0.ɵɵFactoryDeclaration<SplitModule, never>;
    static ɵmod: i0.ɵɵNgModuleDeclaration<SplitModule, never, never, never>;
    static ɵinj: i0.ɵɵInjectorDeclaration<SplitModule>;
}

declare class AsSplitAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Direction: string;
    IsSizeInPixel: boolean;
    IsWithCustomGutterSize: boolean;
    GutterSize: number;
    AsSplitPointersEncoding: AsSplitPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class AsSplitPointersEncoding {
    AsSplitAreas: number[];
}

declare class NullInt64 {
    Int64: number;
    Valid: boolean;
}

declare class AsSplitAreaAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    ShowNameInHeader: boolean;
    Size: number;
    IsAny: boolean;
    HasDiv: boolean;
    DivStyle: string;
    AsSplitAreaPointersEncoding: AsSplitAreaPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class AsSplitAreaPointersEncoding {
    AsSplitID: NullInt64;
    ButtonID: NullInt64;
    CursorID: NullInt64;
    FormID: NullInt64;
    LoadID: NullInt64;
    MarkdownID: NullInt64;
    SliderID: NullInt64;
    SplitID: NullInt64;
    SvgID: NullInt64;
    TableID: NullInt64;
    ToneID: NullInt64;
    TreeID: NullInt64;
    ThreejsID: NullInt64;
    XlsxID: NullInt64;
}

declare class ButtonAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    ButtonPointersEncoding: ButtonPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class ButtonPointersEncoding {
}

declare class Button {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyButtonToButtonAPI(button: Button, buttonAPI: ButtonAPI): void;
declare function CopyButtonAPIToButton(buttonAPI: ButtonAPI, button: Button, frontRepo: FrontRepo): void;

declare class CursorAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    Style: string;
    CursorPointersEncoding: CursorPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class CursorPointersEncoding {
}

declare class Cursor {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    Style: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyCursorToCursorAPI(cursor: Cursor, cursorAPI: CursorAPI): void;
declare function CopyCursorAPIToCursor(cursorAPI: CursorAPI, cursor: Cursor, frontRepo: FrontRepo): void;

declare class FormAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    FormPointersEncoding: FormPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class FormPointersEncoding {
}

declare class Form {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyFormToFormAPI(form: Form, formAPI: FormAPI): void;
declare function CopyFormAPIToForm(formAPI: FormAPI, form: Form, frontRepo: FrontRepo): void;

declare class LoadAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    LoadPointersEncoding: LoadPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class LoadPointersEncoding {
}

declare class Load {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyLoadToLoadAPI(load: Load, loadAPI: LoadAPI): void;
declare function CopyLoadAPIToLoad(loadAPI: LoadAPI, load: Load, frontRepo: FrontRepo): void;

declare class MarkdownAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    MarkdownPointersEncoding: MarkdownPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class MarkdownPointersEncoding {
}

declare class Markdown {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyMarkdownToMarkdownAPI(markdown: Markdown, markdownAPI: MarkdownAPI): void;
declare function CopyMarkdownAPIToMarkdown(markdownAPI: MarkdownAPI, markdown: Markdown, frontRepo: FrontRepo): void;

declare class SliderAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    SliderPointersEncoding: SliderPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class SliderPointersEncoding {
}

declare class Slider {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopySliderToSliderAPI(slider: Slider, sliderAPI: SliderAPI): void;
declare function CopySliderAPIToSlider(sliderAPI: SliderAPI, slider: Slider, frontRepo: FrontRepo): void;

declare class SplitAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    SplitPointersEncoding: SplitPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class SplitPointersEncoding {
}

declare class Split {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopySplitToSplitAPI(split: Split, splitAPI: SplitAPI): void;
declare function CopySplitAPIToSplit(splitAPI: SplitAPI, split: Split, frontRepo: FrontRepo): void;

declare class SvgAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    Style: string;
    SvgPointersEncoding: SvgPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class SvgPointersEncoding {
}

declare class Svg {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    Style: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopySvgToSvgAPI(svg: Svg, svgAPI: SvgAPI): void;
declare function CopySvgAPIToSvg(svgAPI: SvgAPI, svg: Svg, frontRepo: FrontRepo): void;

declare class TableAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    TablePointersEncoding: TablePointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class TablePointersEncoding {
}

declare class Table {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyTableToTableAPI(table: Table, tableAPI: TableAPI): void;
declare function CopyTableAPIToTable(tableAPI: TableAPI, table: Table, frontRepo: FrontRepo): void;

declare class ToneAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    TonePointersEncoding: TonePointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class TonePointersEncoding {
}

declare class Tone {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyToneToToneAPI(tone: Tone, toneAPI: ToneAPI): void;
declare function CopyToneAPIToTone(toneAPI: ToneAPI, tone: Tone, frontRepo: FrontRepo): void;

declare class TreeAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    TreePointersEncoding: TreePointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class TreePointersEncoding {
}

declare class Tree {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyTreeToTreeAPI(tree: Tree, treeAPI: TreeAPI): void;
declare function CopyTreeAPIToTree(treeAPI: TreeAPI, tree: Tree, frontRepo: FrontRepo): void;

declare class ThreejsAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    ThreejsPointersEncoding: ThreejsPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class ThreejsPointersEncoding {
}

declare class Threejs {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyThreejsToThreejsAPI(threejs: Threejs, threejsAPI: ThreejsAPI): void;
declare function CopyThreejsAPIToThreejs(threejsAPI: ThreejsAPI, threejs: Threejs, frontRepo: FrontRepo): void;

declare class XlsxAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    XlsxPointersEncoding: XlsxPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class XlsxPointersEncoding {
}

declare class Xlsx {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    StackName: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyXlsxToXlsxAPI(xlsx: Xlsx, xlsxAPI: XlsxAPI): void;
declare function CopyXlsxAPIToXlsx(xlsxAPI: XlsxAPI, xlsx: Xlsx, frontRepo: FrontRepo): void;

declare class AsSplitArea {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    ShowNameInHeader: boolean;
    Size: number;
    IsAny: boolean;
    HasDiv: boolean;
    DivStyle: string;
    AsSplit?: AsSplit;
    Button?: Button;
    Cursor?: Cursor;
    Form?: Form;
    Load?: Load;
    Markdown?: Markdown;
    Slider?: Slider;
    Split?: Split;
    Svg?: Svg;
    Table?: Table;
    Tone?: Tone;
    Tree?: Tree;
    Threejs?: Threejs;
    Xlsx?: Xlsx;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyAsSplitAreaToAsSplitAreaAPI(assplitarea: AsSplitArea, assplitareaAPI: AsSplitAreaAPI): void;
declare function CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI: AsSplitAreaAPI, assplitarea: AsSplitArea, frontRepo: FrontRepo): void;

declare class AsSplit {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Direction: string;
    IsSizeInPixel: boolean;
    IsWithCustomGutterSize: boolean;
    GutterSize: number;
    AsSplitAreas: Array<AsSplitArea>;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyAsSplitToAsSplitAPI(assplit: AsSplit, assplitAPI: AsSplitAPI): void;
declare function CopyAsSplitAPIToAsSplit(assplitAPI: AsSplitAPI, assplit: AsSplit, frontRepo: FrontRepo): void;

declare class AsSplitService {
    private http;
    private document;
    AsSplitServiceChanged: BehaviorSubject<string>;
    private assplitsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET assplits from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<AsSplitAPI[]>;
    getAsSplits(Name: string, frontRepo: FrontRepo): Observable<AsSplitAPI[]>;
    /** GET assplit by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<AsSplitAPI>;
    getAsSplit(id: number, Name: string, frontRepo: FrontRepo): Observable<AsSplitAPI>;
    postFront(assplit: AsSplit, Name: string): Observable<AsSplitAPI>;
    /** POST: add a new assplit to the server */
    post(assplitdb: AsSplitAPI, Name: string, frontRepo: FrontRepo): Observable<AsSplitAPI>;
    postAsSplit(assplitdb: AsSplitAPI, Name: string, frontRepo: FrontRepo): Observable<AsSplitAPI>;
    /** DELETE: delete the assplitdb from the server */
    delete(assplitdb: AsSplitAPI | number, Name: string): Observable<AsSplitAPI>;
    deleteAsSplit(assplitdb: AsSplitAPI | number, Name: string): Observable<AsSplitAPI>;
    updateFront(assplit: AsSplit, Name: string): Observable<AsSplitAPI>;
    /** PUT: update the assplitdb on the server */
    update(assplitdb: AsSplitAPI, Name: string, frontRepo: FrontRepo): Observable<AsSplitAPI>;
    updateAsSplit(assplitdb: AsSplitAPI, Name: string, frontRepo: FrontRepo): Observable<AsSplitAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<AsSplitService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<AsSplitService>;
}

declare class AsSplitAreaService {
    private http;
    private document;
    AsSplitAreaServiceChanged: BehaviorSubject<string>;
    private assplitareasUrl;
    constructor(http: HttpClient, document: Document);
    /** GET assplitareas from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<AsSplitAreaAPI[]>;
    getAsSplitAreas(Name: string, frontRepo: FrontRepo): Observable<AsSplitAreaAPI[]>;
    /** GET assplitarea by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<AsSplitAreaAPI>;
    getAsSplitArea(id: number, Name: string, frontRepo: FrontRepo): Observable<AsSplitAreaAPI>;
    postFront(assplitarea: AsSplitArea, Name: string): Observable<AsSplitAreaAPI>;
    /** POST: add a new assplitarea to the server */
    post(assplitareadb: AsSplitAreaAPI, Name: string, frontRepo: FrontRepo): Observable<AsSplitAreaAPI>;
    postAsSplitArea(assplitareadb: AsSplitAreaAPI, Name: string, frontRepo: FrontRepo): Observable<AsSplitAreaAPI>;
    /** DELETE: delete the assplitareadb from the server */
    delete(assplitareadb: AsSplitAreaAPI | number, Name: string): Observable<AsSplitAreaAPI>;
    deleteAsSplitArea(assplitareadb: AsSplitAreaAPI | number, Name: string): Observable<AsSplitAreaAPI>;
    updateFront(assplitarea: AsSplitArea, Name: string): Observable<AsSplitAreaAPI>;
    /** PUT: update the assplitareadb on the server */
    update(assplitareadb: AsSplitAreaAPI, Name: string, frontRepo: FrontRepo): Observable<AsSplitAreaAPI>;
    updateAsSplitArea(assplitareadb: AsSplitAreaAPI, Name: string, frontRepo: FrontRepo): Observable<AsSplitAreaAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<AsSplitAreaService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<AsSplitAreaService>;
}

declare class ButtonService {
    private http;
    private document;
    ButtonServiceChanged: BehaviorSubject<string>;
    private buttonsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET buttons from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<ButtonAPI[]>;
    getButtons(Name: string, frontRepo: FrontRepo): Observable<ButtonAPI[]>;
    /** GET button by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<ButtonAPI>;
    getButton(id: number, Name: string, frontRepo: FrontRepo): Observable<ButtonAPI>;
    postFront(button: Button, Name: string): Observable<ButtonAPI>;
    /** POST: add a new button to the server */
    post(buttondb: ButtonAPI, Name: string, frontRepo: FrontRepo): Observable<ButtonAPI>;
    postButton(buttondb: ButtonAPI, Name: string, frontRepo: FrontRepo): Observable<ButtonAPI>;
    /** DELETE: delete the buttondb from the server */
    delete(buttondb: ButtonAPI | number, Name: string): Observable<ButtonAPI>;
    deleteButton(buttondb: ButtonAPI | number, Name: string): Observable<ButtonAPI>;
    updateFront(button: Button, Name: string): Observable<ButtonAPI>;
    /** PUT: update the buttondb on the server */
    update(buttondb: ButtonAPI, Name: string, frontRepo: FrontRepo): Observable<ButtonAPI>;
    updateButton(buttondb: ButtonAPI, Name: string, frontRepo: FrontRepo): Observable<ButtonAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<ButtonService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<ButtonService>;
}

declare class CursorService {
    private http;
    private document;
    CursorServiceChanged: BehaviorSubject<string>;
    private cursorsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET cursors from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<CursorAPI[]>;
    getCursors(Name: string, frontRepo: FrontRepo): Observable<CursorAPI[]>;
    /** GET cursor by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<CursorAPI>;
    getCursor(id: number, Name: string, frontRepo: FrontRepo): Observable<CursorAPI>;
    postFront(cursor: Cursor, Name: string): Observable<CursorAPI>;
    /** POST: add a new cursor to the server */
    post(cursordb: CursorAPI, Name: string, frontRepo: FrontRepo): Observable<CursorAPI>;
    postCursor(cursordb: CursorAPI, Name: string, frontRepo: FrontRepo): Observable<CursorAPI>;
    /** DELETE: delete the cursordb from the server */
    delete(cursordb: CursorAPI | number, Name: string): Observable<CursorAPI>;
    deleteCursor(cursordb: CursorAPI | number, Name: string): Observable<CursorAPI>;
    updateFront(cursor: Cursor, Name: string): Observable<CursorAPI>;
    /** PUT: update the cursordb on the server */
    update(cursordb: CursorAPI, Name: string, frontRepo: FrontRepo): Observable<CursorAPI>;
    updateCursor(cursordb: CursorAPI, Name: string, frontRepo: FrontRepo): Observable<CursorAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<CursorService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<CursorService>;
}

declare class FavIconAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    SVG: string;
    FavIconPointersEncoding: FavIconPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class FavIconPointersEncoding {
}

declare class FavIcon {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    SVG: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyFavIconToFavIconAPI(favicon: FavIcon, faviconAPI: FavIconAPI): void;
declare function CopyFavIconAPIToFavIcon(faviconAPI: FavIconAPI, favicon: FavIcon, frontRepo: FrontRepo): void;

declare class FavIconService {
    private http;
    private document;
    FavIconServiceChanged: BehaviorSubject<string>;
    private faviconsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET favicons from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<FavIconAPI[]>;
    getFavIcons(Name: string, frontRepo: FrontRepo): Observable<FavIconAPI[]>;
    /** GET favicon by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<FavIconAPI>;
    getFavIcon(id: number, Name: string, frontRepo: FrontRepo): Observable<FavIconAPI>;
    postFront(favicon: FavIcon, Name: string): Observable<FavIconAPI>;
    /** POST: add a new favicon to the server */
    post(favicondb: FavIconAPI, Name: string, frontRepo: FrontRepo): Observable<FavIconAPI>;
    postFavIcon(favicondb: FavIconAPI, Name: string, frontRepo: FrontRepo): Observable<FavIconAPI>;
    /** DELETE: delete the favicondb from the server */
    delete(favicondb: FavIconAPI | number, Name: string): Observable<FavIconAPI>;
    deleteFavIcon(favicondb: FavIconAPI | number, Name: string): Observable<FavIconAPI>;
    updateFront(favicon: FavIcon, Name: string): Observable<FavIconAPI>;
    /** PUT: update the favicondb on the server */
    update(favicondb: FavIconAPI, Name: string, frontRepo: FrontRepo): Observable<FavIconAPI>;
    updateFavIcon(favicondb: FavIconAPI, Name: string, frontRepo: FrontRepo): Observable<FavIconAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<FavIconService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<FavIconService>;
}

declare class FormService {
    private http;
    private document;
    FormServiceChanged: BehaviorSubject<string>;
    private formsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET forms from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<FormAPI[]>;
    getForms(Name: string, frontRepo: FrontRepo): Observable<FormAPI[]>;
    /** GET form by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<FormAPI>;
    getForm(id: number, Name: string, frontRepo: FrontRepo): Observable<FormAPI>;
    postFront(form: Form, Name: string): Observable<FormAPI>;
    /** POST: add a new form to the server */
    post(formdb: FormAPI, Name: string, frontRepo: FrontRepo): Observable<FormAPI>;
    postForm(formdb: FormAPI, Name: string, frontRepo: FrontRepo): Observable<FormAPI>;
    /** DELETE: delete the formdb from the server */
    delete(formdb: FormAPI | number, Name: string): Observable<FormAPI>;
    deleteForm(formdb: FormAPI | number, Name: string): Observable<FormAPI>;
    updateFront(form: Form, Name: string): Observable<FormAPI>;
    /** PUT: update the formdb on the server */
    update(formdb: FormAPI, Name: string, frontRepo: FrontRepo): Observable<FormAPI>;
    updateForm(formdb: FormAPI, Name: string, frontRepo: FrontRepo): Observable<FormAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<FormService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<FormService>;
}

declare class LoadService {
    private http;
    private document;
    LoadServiceChanged: BehaviorSubject<string>;
    private loadsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET loads from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<LoadAPI[]>;
    getLoads(Name: string, frontRepo: FrontRepo): Observable<LoadAPI[]>;
    /** GET load by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<LoadAPI>;
    getLoad(id: number, Name: string, frontRepo: FrontRepo): Observable<LoadAPI>;
    postFront(load: Load, Name: string): Observable<LoadAPI>;
    /** POST: add a new load to the server */
    post(loaddb: LoadAPI, Name: string, frontRepo: FrontRepo): Observable<LoadAPI>;
    postLoad(loaddb: LoadAPI, Name: string, frontRepo: FrontRepo): Observable<LoadAPI>;
    /** DELETE: delete the loaddb from the server */
    delete(loaddb: LoadAPI | number, Name: string): Observable<LoadAPI>;
    deleteLoad(loaddb: LoadAPI | number, Name: string): Observable<LoadAPI>;
    updateFront(load: Load, Name: string): Observable<LoadAPI>;
    /** PUT: update the loaddb on the server */
    update(loaddb: LoadAPI, Name: string, frontRepo: FrontRepo): Observable<LoadAPI>;
    updateLoad(loaddb: LoadAPI, Name: string, frontRepo: FrontRepo): Observable<LoadAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<LoadService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<LoadService>;
}

declare class LogoOnTheLeftAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Width: number;
    Height: number;
    SVG: string;
    LogoOnTheLeftPointersEncoding: LogoOnTheLeftPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class LogoOnTheLeftPointersEncoding {
}

declare class LogoOnTheLeft {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Width: number;
    Height: number;
    SVG: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyLogoOnTheLeftToLogoOnTheLeftAPI(logoontheleft: LogoOnTheLeft, logoontheleftAPI: LogoOnTheLeftAPI): void;
declare function CopyLogoOnTheLeftAPIToLogoOnTheLeft(logoontheleftAPI: LogoOnTheLeftAPI, logoontheleft: LogoOnTheLeft, frontRepo: FrontRepo): void;

declare class LogoOnTheLeftService {
    private http;
    private document;
    LogoOnTheLeftServiceChanged: BehaviorSubject<string>;
    private logoontheleftsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET logoonthelefts from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<LogoOnTheLeftAPI[]>;
    getLogoOnTheLefts(Name: string, frontRepo: FrontRepo): Observable<LogoOnTheLeftAPI[]>;
    /** GET logoontheleft by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheLeftAPI>;
    getLogoOnTheLeft(id: number, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheLeftAPI>;
    postFront(logoontheleft: LogoOnTheLeft, Name: string): Observable<LogoOnTheLeftAPI>;
    /** POST: add a new logoontheleft to the server */
    post(logoontheleftdb: LogoOnTheLeftAPI, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheLeftAPI>;
    postLogoOnTheLeft(logoontheleftdb: LogoOnTheLeftAPI, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheLeftAPI>;
    /** DELETE: delete the logoontheleftdb from the server */
    delete(logoontheleftdb: LogoOnTheLeftAPI | number, Name: string): Observable<LogoOnTheLeftAPI>;
    deleteLogoOnTheLeft(logoontheleftdb: LogoOnTheLeftAPI | number, Name: string): Observable<LogoOnTheLeftAPI>;
    updateFront(logoontheleft: LogoOnTheLeft, Name: string): Observable<LogoOnTheLeftAPI>;
    /** PUT: update the logoontheleftdb on the server */
    update(logoontheleftdb: LogoOnTheLeftAPI, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheLeftAPI>;
    updateLogoOnTheLeft(logoontheleftdb: LogoOnTheLeftAPI, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheLeftAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<LogoOnTheLeftService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<LogoOnTheLeftService>;
}

declare class LogoOnTheRightAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Width: number;
    Height: number;
    SVG: string;
    LogoOnTheRightPointersEncoding: LogoOnTheRightPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class LogoOnTheRightPointersEncoding {
}

declare class LogoOnTheRight {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    Width: number;
    Height: number;
    SVG: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyLogoOnTheRightToLogoOnTheRightAPI(logoontheright: LogoOnTheRight, logoontherightAPI: LogoOnTheRightAPI): void;
declare function CopyLogoOnTheRightAPIToLogoOnTheRight(logoontherightAPI: LogoOnTheRightAPI, logoontheright: LogoOnTheRight, frontRepo: FrontRepo): void;

declare class LogoOnTheRightService {
    private http;
    private document;
    LogoOnTheRightServiceChanged: BehaviorSubject<string>;
    private logoontherightsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET logoontherights from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<LogoOnTheRightAPI[]>;
    getLogoOnTheRights(Name: string, frontRepo: FrontRepo): Observable<LogoOnTheRightAPI[]>;
    /** GET logoontheright by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheRightAPI>;
    getLogoOnTheRight(id: number, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheRightAPI>;
    postFront(logoontheright: LogoOnTheRight, Name: string): Observable<LogoOnTheRightAPI>;
    /** POST: add a new logoontheright to the server */
    post(logoontherightdb: LogoOnTheRightAPI, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheRightAPI>;
    postLogoOnTheRight(logoontherightdb: LogoOnTheRightAPI, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheRightAPI>;
    /** DELETE: delete the logoontherightdb from the server */
    delete(logoontherightdb: LogoOnTheRightAPI | number, Name: string): Observable<LogoOnTheRightAPI>;
    deleteLogoOnTheRight(logoontherightdb: LogoOnTheRightAPI | number, Name: string): Observable<LogoOnTheRightAPI>;
    updateFront(logoontheright: LogoOnTheRight, Name: string): Observable<LogoOnTheRightAPI>;
    /** PUT: update the logoontherightdb on the server */
    update(logoontherightdb: LogoOnTheRightAPI, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheRightAPI>;
    updateLogoOnTheRight(logoontherightdb: LogoOnTheRightAPI, Name: string, frontRepo: FrontRepo): Observable<LogoOnTheRightAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<LogoOnTheRightService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<LogoOnTheRightService>;
}

declare class MarkdownService {
    private http;
    private document;
    MarkdownServiceChanged: BehaviorSubject<string>;
    private markdownsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET markdowns from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<MarkdownAPI[]>;
    getMarkdowns(Name: string, frontRepo: FrontRepo): Observable<MarkdownAPI[]>;
    /** GET markdown by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<MarkdownAPI>;
    getMarkdown(id: number, Name: string, frontRepo: FrontRepo): Observable<MarkdownAPI>;
    postFront(markdown: Markdown, Name: string): Observable<MarkdownAPI>;
    /** POST: add a new markdown to the server */
    post(markdowndb: MarkdownAPI, Name: string, frontRepo: FrontRepo): Observable<MarkdownAPI>;
    postMarkdown(markdowndb: MarkdownAPI, Name: string, frontRepo: FrontRepo): Observable<MarkdownAPI>;
    /** DELETE: delete the markdowndb from the server */
    delete(markdowndb: MarkdownAPI | number, Name: string): Observable<MarkdownAPI>;
    deleteMarkdown(markdowndb: MarkdownAPI | number, Name: string): Observable<MarkdownAPI>;
    updateFront(markdown: Markdown, Name: string): Observable<MarkdownAPI>;
    /** PUT: update the markdowndb on the server */
    update(markdowndb: MarkdownAPI, Name: string, frontRepo: FrontRepo): Observable<MarkdownAPI>;
    updateMarkdown(markdowndb: MarkdownAPI, Name: string, frontRepo: FrontRepo): Observable<MarkdownAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<MarkdownService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<MarkdownService>;
}

declare class SliderService {
    private http;
    private document;
    SliderServiceChanged: BehaviorSubject<string>;
    private slidersUrl;
    constructor(http: HttpClient, document: Document);
    /** GET sliders from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<SliderAPI[]>;
    getSliders(Name: string, frontRepo: FrontRepo): Observable<SliderAPI[]>;
    /** GET slider by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<SliderAPI>;
    getSlider(id: number, Name: string, frontRepo: FrontRepo): Observable<SliderAPI>;
    postFront(slider: Slider, Name: string): Observable<SliderAPI>;
    /** POST: add a new slider to the server */
    post(sliderdb: SliderAPI, Name: string, frontRepo: FrontRepo): Observable<SliderAPI>;
    postSlider(sliderdb: SliderAPI, Name: string, frontRepo: FrontRepo): Observable<SliderAPI>;
    /** DELETE: delete the sliderdb from the server */
    delete(sliderdb: SliderAPI | number, Name: string): Observable<SliderAPI>;
    deleteSlider(sliderdb: SliderAPI | number, Name: string): Observable<SliderAPI>;
    updateFront(slider: Slider, Name: string): Observable<SliderAPI>;
    /** PUT: update the sliderdb on the server */
    update(sliderdb: SliderAPI, Name: string, frontRepo: FrontRepo): Observable<SliderAPI>;
    updateSlider(sliderdb: SliderAPI, Name: string, frontRepo: FrontRepo): Observable<SliderAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<SliderService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<SliderService>;
}

declare class SplitService {
    private http;
    private document;
    SplitServiceChanged: BehaviorSubject<string>;
    private splitsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET splits from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<SplitAPI[]>;
    getSplits(Name: string, frontRepo: FrontRepo): Observable<SplitAPI[]>;
    /** GET split by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<SplitAPI>;
    getSplit(id: number, Name: string, frontRepo: FrontRepo): Observable<SplitAPI>;
    postFront(split: Split, Name: string): Observable<SplitAPI>;
    /** POST: add a new split to the server */
    post(splitdb: SplitAPI, Name: string, frontRepo: FrontRepo): Observable<SplitAPI>;
    postSplit(splitdb: SplitAPI, Name: string, frontRepo: FrontRepo): Observable<SplitAPI>;
    /** DELETE: delete the splitdb from the server */
    delete(splitdb: SplitAPI | number, Name: string): Observable<SplitAPI>;
    deleteSplit(splitdb: SplitAPI | number, Name: string): Observable<SplitAPI>;
    updateFront(split: Split, Name: string): Observable<SplitAPI>;
    /** PUT: update the splitdb on the server */
    update(splitdb: SplitAPI, Name: string, frontRepo: FrontRepo): Observable<SplitAPI>;
    updateSplit(splitdb: SplitAPI, Name: string, frontRepo: FrontRepo): Observable<SplitAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<SplitService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<SplitService>;
}

declare class SvgService {
    private http;
    private document;
    SvgServiceChanged: BehaviorSubject<string>;
    private svgsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET svgs from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<SvgAPI[]>;
    getSvgs(Name: string, frontRepo: FrontRepo): Observable<SvgAPI[]>;
    /** GET svg by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<SvgAPI>;
    getSvg(id: number, Name: string, frontRepo: FrontRepo): Observable<SvgAPI>;
    postFront(svg: Svg, Name: string): Observable<SvgAPI>;
    /** POST: add a new svg to the server */
    post(svgdb: SvgAPI, Name: string, frontRepo: FrontRepo): Observable<SvgAPI>;
    postSvg(svgdb: SvgAPI, Name: string, frontRepo: FrontRepo): Observable<SvgAPI>;
    /** DELETE: delete the svgdb from the server */
    delete(svgdb: SvgAPI | number, Name: string): Observable<SvgAPI>;
    deleteSvg(svgdb: SvgAPI | number, Name: string): Observable<SvgAPI>;
    updateFront(svg: Svg, Name: string): Observable<SvgAPI>;
    /** PUT: update the svgdb on the server */
    update(svgdb: SvgAPI, Name: string, frontRepo: FrontRepo): Observable<SvgAPI>;
    updateSvg(svgdb: SvgAPI, Name: string, frontRepo: FrontRepo): Observable<SvgAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<SvgService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<SvgService>;
}

declare class TableService {
    private http;
    private document;
    TableServiceChanged: BehaviorSubject<string>;
    private tablesUrl;
    constructor(http: HttpClient, document: Document);
    /** GET tables from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<TableAPI[]>;
    getTables(Name: string, frontRepo: FrontRepo): Observable<TableAPI[]>;
    /** GET table by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<TableAPI>;
    getTable(id: number, Name: string, frontRepo: FrontRepo): Observable<TableAPI>;
    postFront(table: Table, Name: string): Observable<TableAPI>;
    /** POST: add a new table to the server */
    post(tabledb: TableAPI, Name: string, frontRepo: FrontRepo): Observable<TableAPI>;
    postTable(tabledb: TableAPI, Name: string, frontRepo: FrontRepo): Observable<TableAPI>;
    /** DELETE: delete the tabledb from the server */
    delete(tabledb: TableAPI | number, Name: string): Observable<TableAPI>;
    deleteTable(tabledb: TableAPI | number, Name: string): Observable<TableAPI>;
    updateFront(table: Table, Name: string): Observable<TableAPI>;
    /** PUT: update the tabledb on the server */
    update(tabledb: TableAPI, Name: string, frontRepo: FrontRepo): Observable<TableAPI>;
    updateTable(tabledb: TableAPI, Name: string, frontRepo: FrontRepo): Observable<TableAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<TableService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<TableService>;
}

declare class ThreejsService {
    private http;
    private document;
    ThreejsServiceChanged: BehaviorSubject<string>;
    private threejssUrl;
    constructor(http: HttpClient, document: Document);
    /** GET threejss from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<ThreejsAPI[]>;
    getThreejss(Name: string, frontRepo: FrontRepo): Observable<ThreejsAPI[]>;
    /** GET threejs by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<ThreejsAPI>;
    getThreejs(id: number, Name: string, frontRepo: FrontRepo): Observable<ThreejsAPI>;
    postFront(threejs: Threejs, Name: string): Observable<ThreejsAPI>;
    /** POST: add a new threejs to the server */
    post(threejsdb: ThreejsAPI, Name: string, frontRepo: FrontRepo): Observable<ThreejsAPI>;
    postThreejs(threejsdb: ThreejsAPI, Name: string, frontRepo: FrontRepo): Observable<ThreejsAPI>;
    /** DELETE: delete the threejsdb from the server */
    delete(threejsdb: ThreejsAPI | number, Name: string): Observable<ThreejsAPI>;
    deleteThreejs(threejsdb: ThreejsAPI | number, Name: string): Observable<ThreejsAPI>;
    updateFront(threejs: Threejs, Name: string): Observable<ThreejsAPI>;
    /** PUT: update the threejsdb on the server */
    update(threejsdb: ThreejsAPI, Name: string, frontRepo: FrontRepo): Observable<ThreejsAPI>;
    updateThreejs(threejsdb: ThreejsAPI, Name: string, frontRepo: FrontRepo): Observable<ThreejsAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<ThreejsService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<ThreejsService>;
}

declare class TitleAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    TitlePointersEncoding: TitlePointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class TitlePointersEncoding {
}

declare class Title {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyTitleToTitleAPI(title: Title, titleAPI: TitleAPI): void;
declare function CopyTitleAPIToTitle(titleAPI: TitleAPI, title: Title, frontRepo: FrontRepo): void;

declare class TitleService {
    private http;
    private document;
    TitleServiceChanged: BehaviorSubject<string>;
    private titlesUrl;
    constructor(http: HttpClient, document: Document);
    /** GET titles from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<TitleAPI[]>;
    getTitles(Name: string, frontRepo: FrontRepo): Observable<TitleAPI[]>;
    /** GET title by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<TitleAPI>;
    getTitle(id: number, Name: string, frontRepo: FrontRepo): Observable<TitleAPI>;
    postFront(title: Title, Name: string): Observable<TitleAPI>;
    /** POST: add a new title to the server */
    post(titledb: TitleAPI, Name: string, frontRepo: FrontRepo): Observable<TitleAPI>;
    postTitle(titledb: TitleAPI, Name: string, frontRepo: FrontRepo): Observable<TitleAPI>;
    /** DELETE: delete the titledb from the server */
    delete(titledb: TitleAPI | number, Name: string): Observable<TitleAPI>;
    deleteTitle(titledb: TitleAPI | number, Name: string): Observable<TitleAPI>;
    updateFront(title: Title, Name: string): Observable<TitleAPI>;
    /** PUT: update the titledb on the server */
    update(titledb: TitleAPI, Name: string, frontRepo: FrontRepo): Observable<TitleAPI>;
    updateTitle(titledb: TitleAPI, Name: string, frontRepo: FrontRepo): Observable<TitleAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<TitleService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<TitleService>;
}

declare class ToneService {
    private http;
    private document;
    ToneServiceChanged: BehaviorSubject<string>;
    private tonesUrl;
    constructor(http: HttpClient, document: Document);
    /** GET tones from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<ToneAPI[]>;
    getTones(Name: string, frontRepo: FrontRepo): Observable<ToneAPI[]>;
    /** GET tone by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<ToneAPI>;
    getTone(id: number, Name: string, frontRepo: FrontRepo): Observable<ToneAPI>;
    postFront(tone: Tone, Name: string): Observable<ToneAPI>;
    /** POST: add a new tone to the server */
    post(tonedb: ToneAPI, Name: string, frontRepo: FrontRepo): Observable<ToneAPI>;
    postTone(tonedb: ToneAPI, Name: string, frontRepo: FrontRepo): Observable<ToneAPI>;
    /** DELETE: delete the tonedb from the server */
    delete(tonedb: ToneAPI | number, Name: string): Observable<ToneAPI>;
    deleteTone(tonedb: ToneAPI | number, Name: string): Observable<ToneAPI>;
    updateFront(tone: Tone, Name: string): Observable<ToneAPI>;
    /** PUT: update the tonedb on the server */
    update(tonedb: ToneAPI, Name: string, frontRepo: FrontRepo): Observable<ToneAPI>;
    updateTone(tonedb: ToneAPI, Name: string, frontRepo: FrontRepo): Observable<ToneAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<ToneService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<ToneService>;
}

declare class TreeService {
    private http;
    private document;
    TreeServiceChanged: BehaviorSubject<string>;
    private treesUrl;
    constructor(http: HttpClient, document: Document);
    /** GET trees from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<TreeAPI[]>;
    getTrees(Name: string, frontRepo: FrontRepo): Observable<TreeAPI[]>;
    /** GET tree by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<TreeAPI>;
    getTree(id: number, Name: string, frontRepo: FrontRepo): Observable<TreeAPI>;
    postFront(tree: Tree, Name: string): Observable<TreeAPI>;
    /** POST: add a new tree to the server */
    post(treedb: TreeAPI, Name: string, frontRepo: FrontRepo): Observable<TreeAPI>;
    postTree(treedb: TreeAPI, Name: string, frontRepo: FrontRepo): Observable<TreeAPI>;
    /** DELETE: delete the treedb from the server */
    delete(treedb: TreeAPI | number, Name: string): Observable<TreeAPI>;
    deleteTree(treedb: TreeAPI | number, Name: string): Observable<TreeAPI>;
    updateFront(tree: Tree, Name: string): Observable<TreeAPI>;
    /** PUT: update the treedb on the server */
    update(treedb: TreeAPI, Name: string, frontRepo: FrontRepo): Observable<TreeAPI>;
    updateTree(treedb: TreeAPI, Name: string, frontRepo: FrontRepo): Observable<TreeAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<TreeService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<TreeService>;
}

declare class ViewAPI {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    ShowViewName: boolean;
    IsSelectedView: boolean;
    Direction: string;
    IsSecondaryView: boolean;
    IsSizeInPixel: boolean;
    IsWithCustomGutterSize: boolean;
    GutterSize: number;
    ViewPointersEncoding: ViewPointersEncoding;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare class ViewPointersEncoding {
    RootAsSplitAreas: number[];
}

declare class View {
    static GONGSTRUCT_NAME: string;
    ID: number;
    Name: string;
    ShowViewName: boolean;
    IsSelectedView: boolean;
    Direction: string;
    IsSecondaryView: boolean;
    IsSizeInPixel: boolean;
    IsWithCustomGutterSize: boolean;
    GutterSize: number;
    RootAsSplitAreas: Array<AsSplitArea>;
    CreatedAt?: string;
    DeletedAt?: string;
}
declare function CopyViewToViewAPI(view: View, viewAPI: ViewAPI): void;
declare function CopyViewAPIToView(viewAPI: ViewAPI, view: View, frontRepo: FrontRepo): void;

declare class ViewService {
    private http;
    private document;
    ViewServiceChanged: BehaviorSubject<string>;
    private viewsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET views from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<ViewAPI[]>;
    getViews(Name: string, frontRepo: FrontRepo): Observable<ViewAPI[]>;
    /** GET view by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<ViewAPI>;
    getView(id: number, Name: string, frontRepo: FrontRepo): Observable<ViewAPI>;
    postFront(view: View, Name: string): Observable<ViewAPI>;
    /** POST: add a new view to the server */
    post(viewdb: ViewAPI, Name: string, frontRepo: FrontRepo): Observable<ViewAPI>;
    postView(viewdb: ViewAPI, Name: string, frontRepo: FrontRepo): Observable<ViewAPI>;
    /** DELETE: delete the viewdb from the server */
    delete(viewdb: ViewAPI | number, Name: string): Observable<ViewAPI>;
    deleteView(viewdb: ViewAPI | number, Name: string): Observable<ViewAPI>;
    updateFront(view: View, Name: string): Observable<ViewAPI>;
    /** PUT: update the viewdb on the server */
    update(viewdb: ViewAPI, Name: string, frontRepo: FrontRepo): Observable<ViewAPI>;
    updateView(viewdb: ViewAPI, Name: string, frontRepo: FrontRepo): Observable<ViewAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<ViewService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<ViewService>;
}

declare class XlsxService {
    private http;
    private document;
    XlsxServiceChanged: BehaviorSubject<string>;
    private xlsxsUrl;
    constructor(http: HttpClient, document: Document);
    /** GET xlsxs from the server */
    gets(Name: string, frontRepo: FrontRepo): Observable<XlsxAPI[]>;
    getXlsxs(Name: string, frontRepo: FrontRepo): Observable<XlsxAPI[]>;
    /** GET xlsx by id. Will 404 if id not found */
    get(id: number, Name: string, frontRepo: FrontRepo): Observable<XlsxAPI>;
    getXlsx(id: number, Name: string, frontRepo: FrontRepo): Observable<XlsxAPI>;
    postFront(xlsx: Xlsx, Name: string): Observable<XlsxAPI>;
    /** POST: add a new xlsx to the server */
    post(xlsxdb: XlsxAPI, Name: string, frontRepo: FrontRepo): Observable<XlsxAPI>;
    postXlsx(xlsxdb: XlsxAPI, Name: string, frontRepo: FrontRepo): Observable<XlsxAPI>;
    /** DELETE: delete the xlsxdb from the server */
    delete(xlsxdb: XlsxAPI | number, Name: string): Observable<XlsxAPI>;
    deleteXlsx(xlsxdb: XlsxAPI | number, Name: string): Observable<XlsxAPI>;
    updateFront(xlsx: Xlsx, Name: string): Observable<XlsxAPI>;
    /** PUT: update the xlsxdb on the server */
    update(xlsxdb: XlsxAPI, Name: string, frontRepo: FrontRepo): Observable<XlsxAPI>;
    updateXlsx(xlsxdb: XlsxAPI, Name: string, frontRepo: FrontRepo): Observable<XlsxAPI>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<XlsxService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<XlsxService>;
}

declare const StackType = "github.com/fullstack-lang/gong/lib/split/go/models";
declare class FrontRepo {
    array_AsSplits: AsSplit[];
    map_ID_AsSplit: Map<number, AsSplit>;
    array_AsSplitAreas: AsSplitArea[];
    map_ID_AsSplitArea: Map<number, AsSplitArea>;
    array_Buttons: Button[];
    map_ID_Button: Map<number, Button>;
    array_Cursors: Cursor[];
    map_ID_Cursor: Map<number, Cursor>;
    array_FavIcons: FavIcon[];
    map_ID_FavIcon: Map<number, FavIcon>;
    array_Forms: Form[];
    map_ID_Form: Map<number, Form>;
    array_Loads: Load[];
    map_ID_Load: Map<number, Load>;
    array_LogoOnTheLefts: LogoOnTheLeft[];
    map_ID_LogoOnTheLeft: Map<number, LogoOnTheLeft>;
    array_LogoOnTheRights: LogoOnTheRight[];
    map_ID_LogoOnTheRight: Map<number, LogoOnTheRight>;
    array_Markdowns: Markdown[];
    map_ID_Markdown: Map<number, Markdown>;
    array_Sliders: Slider[];
    map_ID_Slider: Map<number, Slider>;
    array_Splits: Split[];
    map_ID_Split: Map<number, Split>;
    array_Svgs: Svg[];
    map_ID_Svg: Map<number, Svg>;
    array_Tables: Table[];
    map_ID_Table: Map<number, Table>;
    array_Threejss: Threejs[];
    map_ID_Threejs: Map<number, Threejs>;
    array_Titles: Title[];
    map_ID_Title: Map<number, Title>;
    array_Tones: Tone[];
    map_ID_Tone: Map<number, Tone>;
    array_Trees: Tree[];
    map_ID_Tree: Map<number, Tree>;
    array_Views: View[];
    map_ID_View: Map<number, View>;
    array_Xlsxs: Xlsx[];
    map_ID_Xlsx: Map<number, Xlsx>;
    GONG__Index: number;
    getFrontArray<Type>(gongStructName: string): Array<Type>;
    getFrontMap<Type>(gongStructName: string): Map<number, Type>;
}
declare class DialogData {
    ID: number;
    ReversePointer: string;
    OrderingMode: boolean;
    SelectionMode: SelectionMode;
    SourceStruct: string;
    SourceField: string;
    IntermediateStruct: string;
    IntermediateStructField: string;
    NextAssociationStruct: string;
    Name: string;
}
declare enum SelectionMode {
    ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
    MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE"
}
declare class FrontRepoService {
    private http;
    private assplitService;
    private assplitareaService;
    private buttonService;
    private cursorService;
    private faviconService;
    private formService;
    private loadService;
    private logoontheleftService;
    private logoontherightService;
    private markdownService;
    private sliderService;
    private splitService;
    private svgService;
    private tableService;
    private threejsService;
    private titleService;
    private toneService;
    private treeService;
    private viewService;
    private xlsxService;
    Name: string;
    httpOptions: {
        headers: HttpHeaders;
    };
    frontRepo: FrontRepo;
    private webSocketConnections;
    constructor(http: HttpClient, // insertion point sub template 
    assplitService: AsSplitService, assplitareaService: AsSplitAreaService, buttonService: ButtonService, cursorService: CursorService, faviconService: FavIconService, formService: FormService, loadService: LoadService, logoontheleftService: LogoOnTheLeftService, logoontherightService: LogoOnTheRightService, markdownService: MarkdownService, sliderService: SliderService, splitService: SplitService, svgService: SvgService, tableService: TableService, threejsService: ThreejsService, titleService: TitleService, toneService: ToneService, treeService: TreeService, viewService: ViewService, xlsxService: XlsxService);
    postService(structName: string, instanceToBePosted: any): void;
    deleteService(structName: string, instanceToBeDeleted: any): void;
    observableFrontRepo: [
        Observable<null>,
        Observable<AsSplitAPI[]>,
        Observable<AsSplitAreaAPI[]>,
        Observable<ButtonAPI[]>,
        Observable<CursorAPI[]>,
        Observable<FavIconAPI[]>,
        Observable<FormAPI[]>,
        Observable<LoadAPI[]>,
        Observable<LogoOnTheLeftAPI[]>,
        Observable<LogoOnTheRightAPI[]>,
        Observable<MarkdownAPI[]>,
        Observable<SliderAPI[]>,
        Observable<SplitAPI[]>,
        Observable<SvgAPI[]>,
        Observable<TableAPI[]>,
        Observable<ThreejsAPI[]>,
        Observable<TitleAPI[]>,
        Observable<ToneAPI[]>,
        Observable<TreeAPI[]>,
        Observable<ViewAPI[]>,
        Observable<XlsxAPI[]>
    ];
    pull(Name?: string): Observable<FrontRepo>;
    connectToWebSocket(Name: string): Observable<FrontRepo>;
    static ɵfac: i0.ɵɵFactoryDeclaration<FrontRepoService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<FrontRepoService>;
}
declare function getAsSplitUniqueID(id: number): number;
declare function getAsSplitAreaUniqueID(id: number): number;
declare function getButtonUniqueID(id: number): number;
declare function getCursorUniqueID(id: number): number;
declare function getFavIconUniqueID(id: number): number;
declare function getFormUniqueID(id: number): number;
declare function getLoadUniqueID(id: number): number;
declare function getLogoOnTheLeftUniqueID(id: number): number;
declare function getLogoOnTheRightUniqueID(id: number): number;
declare function getMarkdownUniqueID(id: number): number;
declare function getSliderUniqueID(id: number): number;
declare function getSplitUniqueID(id: number): number;
declare function getSvgUniqueID(id: number): number;
declare function getTableUniqueID(id: number): number;
declare function getThreejsUniqueID(id: number): number;
declare function getTitleUniqueID(id: number): number;
declare function getToneUniqueID(id: number): number;
declare function getTreeUniqueID(id: number): number;
declare function getViewUniqueID(id: number): number;
declare function getXlsxUniqueID(id: number): number;

declare class CommitNbFromBackService {
    private http;
    private location;
    private document;
    httpOptions: {
        headers: HttpHeaders;
    };
    private commitNbFromBackUrl;
    constructor(http: HttpClient, location: Location, document: Document);
    getCommitNbFromBack(intervalMs: number, Name?: string): Observable<number>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<CommitNbFromBackService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<CommitNbFromBackService>;
}

declare class PushFromFrontNbService {
    private http;
    private location;
    private document;
    httpOptions: {
        headers: HttpHeaders;
    };
    private pushFromFrontNbURL;
    constructor(http: HttpClient, location: Location, document: Document);
    getPushFromFrontNb(): Observable<number>;
    getPushNbFromFront(intervalMs: number, Name?: string): Observable<number>;
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    private handleError;
    private log;
    static ɵfac: i0.ɵɵFactoryDeclaration<PushFromFrontNbService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<PushFromFrontNbService>;
}

declare class BackRepoData {
    AsSplitAPIs: AsSplitAPI[];
    AsSplitAreaAPIs: AsSplitAreaAPI[];
    ButtonAPIs: ButtonAPI[];
    CursorAPIs: CursorAPI[];
    FavIconAPIs: FavIconAPI[];
    FormAPIs: FormAPI[];
    LoadAPIs: LoadAPI[];
    LogoOnTheLeftAPIs: LogoOnTheLeftAPI[];
    LogoOnTheRightAPIs: LogoOnTheRightAPI[];
    MarkdownAPIs: MarkdownAPI[];
    SliderAPIs: SliderAPI[];
    SplitAPIs: SplitAPI[];
    SvgAPIs: SvgAPI[];
    TableAPIs: TableAPI[];
    ThreejsAPIs: ThreejsAPI[];
    TitleAPIs: TitleAPI[];
    ToneAPIs: ToneAPI[];
    TreeAPIs: TreeAPI[];
    ViewAPIs: ViewAPI[];
    XlsxAPIs: XlsxAPI[];
    GONG__Index: number;
    constructor(data?: Partial<BackRepoData>);
}

declare class WebSocketService {
    private document;
    constructor(document: Document);
    connect(stackPath: string): Observable<any>;
    static ɵfac: i0.ɵɵFactoryDeclaration<WebSocketService, never>;
    static ɵprov: i0.ɵɵInjectableDeclaration<WebSocketService>;
}

declare enum Direction {
    Vertical = "vertical",
    Horizontal = "horizontal"
}
interface DirectionSelect {
    value: string;
    viewValue: string;
}
declare const DirectionList: DirectionSelect[];

export { AsSplit, AsSplitAPI, AsSplitArea, AsSplitAreaAPI, AsSplitAreaPointersEncoding, AsSplitAreaService, AsSplitPointersEncoding, AsSplitService, BackRepoData, Button, ButtonAPI, ButtonPointersEncoding, ButtonService, CommitNbFromBackService, CopyAsSplitAPIToAsSplit, CopyAsSplitAreaAPIToAsSplitArea, CopyAsSplitAreaToAsSplitAreaAPI, CopyAsSplitToAsSplitAPI, CopyButtonAPIToButton, CopyButtonToButtonAPI, CopyCursorAPIToCursor, CopyCursorToCursorAPI, CopyFavIconAPIToFavIcon, CopyFavIconToFavIconAPI, CopyFormAPIToForm, CopyFormToFormAPI, CopyLoadAPIToLoad, CopyLoadToLoadAPI, CopyLogoOnTheLeftAPIToLogoOnTheLeft, CopyLogoOnTheLeftToLogoOnTheLeftAPI, CopyLogoOnTheRightAPIToLogoOnTheRight, CopyLogoOnTheRightToLogoOnTheRightAPI, CopyMarkdownAPIToMarkdown, CopyMarkdownToMarkdownAPI, CopySliderAPIToSlider, CopySliderToSliderAPI, CopySplitAPIToSplit, CopySplitToSplitAPI, CopySvgAPIToSvg, CopySvgToSvgAPI, CopyTableAPIToTable, CopyTableToTableAPI, CopyThreejsAPIToThreejs, CopyThreejsToThreejsAPI, CopyTitleAPIToTitle, CopyTitleToTitleAPI, CopyToneAPIToTone, CopyToneToToneAPI, CopyTreeAPIToTree, CopyTreeToTreeAPI, CopyViewAPIToView, CopyViewToViewAPI, CopyXlsxAPIToXlsx, CopyXlsxToXlsxAPI, Cursor, CursorAPI, CursorPointersEncoding, CursorService, DialogData, Direction, DirectionList, FavIcon, FavIconAPI, FavIconPointersEncoding, FavIconService, Form, FormAPI, FormPointersEncoding, FormService, FrontRepo, FrontRepoService, Load, LoadAPI, LoadPointersEncoding, LoadService, LogoOnTheLeft, LogoOnTheLeftAPI, LogoOnTheLeftPointersEncoding, LogoOnTheLeftService, LogoOnTheRight, LogoOnTheRightAPI, LogoOnTheRightPointersEncoding, LogoOnTheRightService, Markdown, MarkdownAPI, MarkdownPointersEncoding, MarkdownService, NullInt64, PushFromFrontNbService, SelectionMode, Slider, SliderAPI, SliderPointersEncoding, SliderService, Split, SplitAPI, SplitModule, SplitPointersEncoding, SplitService, StackType, Svg, SvgAPI, SvgPointersEncoding, SvgService, Table, TableAPI, TablePointersEncoding, TableService, Threejs, ThreejsAPI, ThreejsPointersEncoding, ThreejsService, Title, TitleAPI, TitlePointersEncoding, TitleService, Tone, ToneAPI, TonePointersEncoding, ToneService, Tree, TreeAPI, TreePointersEncoding, TreeService, View, ViewAPI, ViewPointersEncoding, ViewService, WebSocketService, Xlsx, XlsxAPI, XlsxPointersEncoding, XlsxService, getAsSplitAreaUniqueID, getAsSplitUniqueID, getButtonUniqueID, getCursorUniqueID, getFavIconUniqueID, getFormUniqueID, getLoadUniqueID, getLogoOnTheLeftUniqueID, getLogoOnTheRightUniqueID, getMarkdownUniqueID, getSliderUniqueID, getSplitUniqueID, getSvgUniqueID, getTableUniqueID, getThreejsUniqueID, getTitleUniqueID, getToneUniqueID, getTreeUniqueID, getViewUniqueID, getXlsxUniqueID };
export type { DirectionSelect };
