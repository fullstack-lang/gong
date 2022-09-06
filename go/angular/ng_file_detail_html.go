package angular

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
	NgDetailHtmlInsertionPerStructFieldsManyMany
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
	NgDetailSliceOfPointerToStructManyManyHtml
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
        <mat-grid-tile [colspan]="3">
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
    <mat-form-field>
        <mat-label>{{FieldName}}</mat-label>
        <input matInput [ngModelOptions]="{standalone: true}" [ngxMatDatetimePicker]="picker{{FieldName}}" [(ngModel)]="{{structname}}.{{FieldName}}">
        <mat-datepicker-toggle matSuffix [for]="$any(picker{{FieldName}})"></mat-datepicker-toggle>
        <ngx-mat-datetime-picker [showSeconds]="true" #picker{{FieldName}}></ngx-mat-datetime-picker>
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
                <input type="number" [ngModelOptions]="{standalone: true}" matInput [(ngModel)]="{{FieldName}}_Hours">
            </mat-form-field>
        </mat-grid-tile>
        <mat-grid-tile>
            <mat-form-field class="details_minutes_width">
                <mat-label>{{FieldName}} Minutes</mat-label>
                <input type="number" [ngModelOptions]="{standalone: true}" matInput [(ngModel)]="{{FieldName}}_Minutes">
            </mat-form-field>
        </mat-grid-tile>
        <mat-grid-tile>
            <mat-form-field class="details_seconds_width">
                <mat-label>{{FieldName}} Seconds</mat-label>
                <input type="number" [ngModelOptions]="{standalone: true}" matInput [(ngModel)]="{{FieldName}}_Seconds">
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
        <mat-grid-tile [colspan]="3">
            <button mat-raised-button (click)="openReverseSelection('{{AssocStructName}}', '{{Structname}}_{{FieldName}}DBID', 'ONE_MANY_ASSOCIATION_MODE', '', '', '')">{{FieldName}}</button>
        </mat-grid-tile>
        <!-- insertion point for the button of the MANY_MANY association{{` + string(rune(NgDetailHtmlInsertionPerStructFieldsManyMany)) + `}}-->
        <mat-grid-tile>
            <button mat-raised-button (click)="openDragAndDropOrdering('{{AssocStructName}}', '{{Structname}}_{{FieldName}}DBID')">
                <mat-icon>
                    shuffle
                </mat-icon>
            </button>
        </mat-grid-tile>
    </mat-grid-list>`,

	NgDetailSliceOfPointerToStructManyManyHtml: `-->
        <mat-grid-tile>
            <button mat-raised-button (click)="openReverseSelection('{{AssocStructName}}', '{{Structname}}_{{FieldName}}DBID', 
            'MANY_MANY_ASSOCIATION_MODE', '{{FieldName}}', '{{IntermediateStructField}}', '{{NextAssociatedStruct}}')">
                <mat-icon>
                    list
                </mat-icon>
            </button>
        </mat-grid-tile> <!-- end of insertion `,

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
