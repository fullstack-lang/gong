package models

const NgDetailTemplateHTML = `<form *ngIf="{{structname}}" class="details">
    <h2 class="details__title">{{structname}}</h2>
    <!-- insertion point for fields specific code -->{{` + string(rune(NgDetailHtmlInsertionPerStructFields)) + `}}
    <div class="details__save">
        <button mat-raised-button color="primary" (click)="save()">
			Save {{structname}}
		</button>
    </div>
</form>`

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
	NgDetailHtmlBasicStringField
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
    <mat-form-field appearance="fill" class="detail-full-width">
        <mat-label>{{FieldName}}</mat-label>
        <mat-select [(ngModel)]="{{structname}}.{{FieldName}}" [ngModelOptions]="{standalone: true}">
            <mat-option *ngFor="let enum of {{EnumName}}List" [value]="enum.value">
                {{enum.viewValue}}
            </mat-option>
        </mat-select>
    </mat-form-field>
`,

	NgDetailHtmlBasicField: `
    <!-- -->
    <mat-form-field class="detail-full-width">
        <mat-label>{{FieldName}}</mat-label>
        <input {{TypeInput}} matInput [(ngModel)]="{{structname}}.{{FieldName}}" [ngModelOptions]="{standalone: true}">
    </mat-form-field>
`,

	NgDetailHtmlBasicStringField: `
    <mat-grid-list *ngIf='!isATextArea("{{FieldName}}")' cols="5" rowHeight="2:1">
        <mat-grid-tile [colspan]="4">
            <mat-form-field mat-form-field class="detail-full-width">
                <mat-label>{{FieldName}}</mat-label>
                <input name="" [ngModelOptions]="{standalone: true}" matInput [(ngModel)]="{{structname}}.{{FieldName}}">
            </mat-form-field>
        </mat-grid-tile>
        <mat-grid-tile>
            <button mat-raised-button (click)="toggleTextArea('{{FieldName}}')">
                <mat-icon>
                    expand_more
                </mat-icon>
            </button>
        </mat-grid-tile>
    </mat-grid-list>
    <mat-form-field *ngIf='isATextArea("{{FieldName}}")' mat-form-field class="detail-full-width">
        <mat-label>{{FieldName}}</mat-label>
        <textarea name="" [ngModelOptions]="{standalone: true}" matInput [(ngModel)]="{{structname}}.{{FieldName}}"></textarea>
    </mat-form-field>
`,
	NgDetailHtmlTimeField: `
    <!-- -->
    <mat-form-field class="detail-full-width">
        <mat-label>{{FieldName}}</mat-label>
        <input name="" [ngModelOptions]="{standalone: true}" matInput [(ngModel)]="{{structname}}.{{FieldName}}">
    </mat-form-field>
`,

	NgDetailHtmlBool: `
    <!-- -->
    <mat-checkbox [formControl]="{{FieldName}}FormControl">{{FieldName}}</mat-checkbox>
`,

	NgDetailHtmlTimeDuration: `
    <!-- -->
    <mat-grid-list cols="3" rowHeight="4:1">
        <mat-grid-tile>
            <mat-form-field class="details_hours_width">
                <mat-label>{{FieldName}} Hours</mat-label>
                <input type="number" [ngModelOptions]="{standalone: true}" matInput
                    [(ngModel)]="{{FieldName}}_Hours">
            </mat-form-field>
        </mat-grid-tile>
        <mat-grid-tile>
            <mat-form-field class="details_minutes_width">
                <mat-label>{{FieldName}} Minutes</mat-label>
                <input type="number" [ngModelOptions]="{standalone: true}" matInput
                    [(ngModel)]="{{FieldName}}_Minutes">
            </mat-form-field>
        </mat-grid-tile>
        <mat-grid-tile>
            <mat-form-field class="details_seconds_width">
                <mat-label>{{FieldName}} Seconds</mat-label>
                <input type="number" [ngModelOptions]="{standalone: true}" matInput
                    [(ngModel)]="{{FieldName}}_Seconds">
            </mat-form-field>
        </mat-grid-tile>
    </mat-grid-list>
`,

	NgDetailPointerToStructHtmlFormField: `
    <mat-form-field class="detail-full-width">
        <mat-label>{{FieldName}}</mat-label>
        <mat-select [compareWith]="compareObjects" [(value)]="{{structname}}.{{FieldName}}" (selectionChange)="fillUpNameIfEmpty($event)">
            <mat-option>None</mat-option>
            <mat-option *ngFor="let {{assocStructName}} of frontRepo.{{AssocStructName}}s_array" [value]="{{assocStructName}}">
                {{{{assocStructName}}.Name}}
            </mat-option>
        </mat-select>
    </mat-form-field>
`,

	NgDetailSliceOfPointerToStructHtml: `
    <mat-grid-list cols="5" rowHeight="2:1">
        <mat-grid-tile [colspan]="4">
            <button mat-raised-button
                (click)="openReverseSelection('{{AssocStructName}}', '{{Structname}}_{{FieldName}}DBID')">{{FieldName}}</button>
        </mat-grid-tile>
        <mat-grid-tile>
            <button mat-raised-button (click)="openDragAndDropOrdering('{{AssocStructName}}', '{{Structname}}_{{FieldName}}DBID')">
                <mat-icon>
                    shuffle
                </mat-icon>
            </button>
        </mat-grid-tile>
    </mat-grid-list>`,

	NgDetailSliceOfPointerToStructReverseHtml: `
    <mat-form-field class="detail-full-width">
        <mat-label>({{AssocStructName}}) {{FieldName}}</mat-label>
        <mat-select [compareWith]="compareObjects" [(ngModel)]="{{structname}}.{{AssocStructName}}_{{FieldName}}_reverse" [ngModelOptions]="{standalone: true}">
            <mat-option>None</mat-option>
            <mat-option *ngFor="let {{assocStructName}} of frontRepo.{{AssocStructName}}s_array" [value]="{{assocStructName}}">
                {{{{assocStructName}}.Name}}
            </mat-option>
        </mat-select>
    </mat-form-field>`,
}
