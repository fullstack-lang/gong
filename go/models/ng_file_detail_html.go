package models

const NgDetailTemplateHTML = `<div *ngIf="{{structname}}" class="details">
    <h1 class="details__title">{{structname}}</h1>
    <!-- insertion point for fields specific code -->{{` + string(rune(NgDetailHtmlInsertionPerStructFields)) + `}}
    <div class="details__save">
        <button mat-raised-button color="primary" (click)="save()">
			Save {{structname}}
		</button>
    </div>
</div>`

type NgDetailHtmlInsertionPoint int

const (
	NgDetailHtmlInsertionPerStructFields NgDetailHtmlInsertionPoint = iota
	NgDetailHtmlInsertionsNb
)

//
// Sub Templates
//
type NgDetailHtmlSubTemplate int

const (
	NgDetailHtmlEnum NgDetailHtmlSubTemplate = iota
	NgDetailHtmlBasicField
	NgDetailHtmlTimeField
	NgDetailHtmlBool
	NgDetailHtmlTimeDuration
	NgDetailPointerToStructHtmlFormField
	NgDetailSliceOfPointerToStructHtml
	NgDetailSliceOfPointerToStructReverseHtml
)

var NgDetailHtmlSubTemplateCode map[NgDetailHtmlSubTemplate]string = map[NgDetailHtmlSubTemplate]string{
	NgDetailHtmlEnum: `
    <!-- -->
    <div class="details__item">
        <mat-form-field appearance="fill">
            <mat-label>{{FieldName}}</mat-label>
            <mat-select [(ngModel)]="{{structname}}.{{FieldName}}">
                <mat-option *ngFor="let enum of {{EnumName}}List" [value]="enum.value">
                    {{enum.viewValue}}
                </mat-option>
            </mat-select>
        </mat-form-field>
    </div>`,

	NgDetailHtmlBasicField: `
    <!-- -->
    <div class="details__item">
        <form>
            <mat-form-field>
                <mat-label>{{FieldName}}</mat-label>
                <input {{TypeInput}} matInput [(ngModel)]="{{structname}}.{{FieldName}}">
            </mat-form-field>
        </form>
    </div>`,

	NgDetailHtmlTimeField: `
    <!-- -->
    <div class="details__item">
        <form>
            <mat-form-field>
                <mat-label>{{FieldName}}</mat-label>
                <input name="" [ngModelOptions]="{standalone: true}" matInput [(ngModel)]="{{structname}}.{{FieldName}}">
            </mat-form-field>
        </form>
    </div>`,

	NgDetailHtmlBool: `
    <!-- -->
    <div class="details__item">
        <form>
            <mat-checkbox [formControl]="{{FieldName}}FormControl">{{FieldName}}</mat-checkbox>
        </form>
    </div>`,

	NgDetailHtmlTimeDuration: `
    <!-- -->
    <div class="details__item">
        <mat-grid-list cols="3" rowHeight="4:1">
            <mat-grid-tile>
                <form>
                    <mat-form-field class="details_hours_width">
                        <mat-label>{{FieldName}} Hours</mat-label>
                        <input type="number" [ngModelOptions]="{standalone: true}" matInput
                            [(ngModel)]="{{FieldName}}_Hours">
                    </mat-form-field>
                </form>
            </mat-grid-tile>
            <mat-grid-tile>
                <form>
                    <mat-form-field class="details_minutes_width">
                        <mat-label>{{FieldName}} Minutes</mat-label>
                        <input type="number" [ngModelOptions]="{standalone: true}" matInput
                            [(ngModel)]="{{FieldName}}_Minutes">
                    </mat-form-field>
                </form>
            </mat-grid-tile>
            <mat-grid-tile>
                <form>
                    <mat-form-field class="details_seconds_width">
                        <mat-label>{{FieldName}} Seconds</mat-label>
                        <input type="number" [ngModelOptions]="{standalone: true}" matInput
                            [(ngModel)]="{{FieldName}}_Seconds">
                    </mat-form-field>
                </form>
            </mat-grid-tile>
        </mat-grid-list>
    </div>`,

	NgDetailPointerToStructHtmlFormField: `
    <div class="details__item">
        <mat-form-field>
            <mat-label>{{FieldName}}</mat-label>
            <mat-select [(ngModel)]="{{structname}}.{{FieldName}}">
                <mat-option>None</mat-option>
                <mat-option *ngFor="let {{assocStructName}} of frontRepo.{{AssocStructName}}s_array" [value]="{{assocStructName}}">
                    {{{{assocStructName}}.Name}}
                </mat-option>
            </mat-select>
        </mat-form-field>
    </div>`,

	NgDetailSliceOfPointerToStructHtml: `
    <mat-grid-list cols="2" rowHeight="6:1">
        <mat-grid-tile>
            <div style="width: 100vw; text-align: right;">
                <button mat-raised-button
                    (click)="openReverseSelection('{{AssocStructName}}', '{{Structname}}_{{FieldName}}DBID')">{{FieldName}}</button>
            </div>
        </mat-grid-tile>
        <mat-grid-tile>
            <div style="width: 100vw; text-align: left;">
                <button mat-raised-button (click)="openDragAndDropOrdering('{{AssocStructName}}', '{{Structname}}_{{FieldName}}DBID')">
                    <mat-icon>
                        shuffle
                    </mat-icon>
                </button>
            </div>
        </mat-grid-tile>
    </mat-grid-list>`,

	NgDetailSliceOfPointerToStructReverseHtml: `
    <div class="details__item">
    <mat-form-field>
        <mat-label><- {{FieldName}}</mat-label>
        <mat-select [(ngModel)]="{{structname}}.{{AssocStructName}}_{{FieldName}}_reverse">
            <mat-option>None</mat-option>
            <mat-option *ngFor="let {{assocStructName}} of frontRepo.{{AssocStructName}}s_array" [value]="{{assocStructName}}">
                {{{{assocStructName}}.Name}}
            </mat-option>
        </mat-select>
    </mat-form-field>
</div>`,
}
