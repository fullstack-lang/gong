package models

const NgPresentationTemplateHTML = `
<div class="presentation__header" (click)="setEditorRouterOutlet({{structname}}.ID)">
    <h1>{{Structname}}'s fields <span class="material-icons">edit</span></h1>
</div>
<table *ngIf="{{structname}}" class="presentation__table mat-table">
    <tr class="mat-header-row">
        <th class="mat-header-cell">Name</th>
        <th class="mat-header-cell">Value</th>
    </tr>
    <!-- insertion point for fields specific code -->{{` + string(rune(NgPresentationHtmlField)) + `}}
</table>
<!-- hack to force loading of mat table style an error is generated but it is a know angular issue https://github.com/angular/components/issues/10941 -->
<table mat-table [dataSource]="dataSource" class="mat-elevation-z8">
    <!-- Position Column -->
    <tr mat-header-row *matHeaderRowDef="displayedColumns"></tr>
    <tr mat-row *matRowDef="let row; columns: displayedColumns;"></tr>
</table>`

//
// Sub Templates
//
type NgPresentationHtmlSubTemplateId int

const (
	NgPresentationHtmlField NgPresentationHtmlSubTemplateId = iota
	NgPresentationHtmlEnumString
	NgPresentationHtmlEnumInt
	NgPresentationHtmlBasicField
	NgPresentationHtmlTimeField
	NgPresentationHtmlBasicFieldTimeDuration
	NgPresentationHtmlBool
	NgPresentationPointerToStructHtmlFormField
)

var NgPresentationEnumHtmlSubTemplateCode map[NgPresentationHtmlSubTemplateId]string = map[NgPresentationHtmlSubTemplateId]string{
	NgPresentationHtmlEnumString: `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{structname}}.{{FieldName}}}}</td>
    </tr>`,

	NgPresentationHtmlEnumInt: `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{FieldName}}_Value}}</td>
    </tr>`,

	NgPresentationHtmlBasicField: `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{structname}}.{{FieldName}}}}</td>
    </tr>`,

	NgPresentationHtmlTimeField: `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{structname}}.{{FieldName}}}}</td>
    </tr>`,

	NgPresentationHtmlBasicFieldTimeDuration: `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{FieldName}}_Hours}}H {{{{FieldName}}_Minutes}}M {{{{FieldName}}_Seconds}}S</td>
    </tr>`,

	NgPresentationHtmlBool: `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">
            <mat-checkbox [checked]="{{structname}}.{{FieldName}}" disabled=true>
            </mat-checkbox>
        </td>
    </tr>`,

	NgPresentationPointerToStructHtmlFormField: `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{structname}}.{{FieldName}} ? {{structname}}.{{FieldName}}.Name : ""}}</td>
    </tr>`,
}
