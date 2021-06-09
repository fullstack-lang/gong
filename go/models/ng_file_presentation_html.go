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
    <!-- insertion point for fields specific code -->{{` + string(NgPresentationHtmlField) + `}}
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
type NgPresentationHtmlSubTemplate string

const (
	NgPresentationHtmlField                    NgPresentationHtmlSubTemplate = "NgPresentationHtmlField"
	NgPresentationHtmlEnum                     NgPresentationHtmlSubTemplate = "NgPresentationHtmlEnum"
	NgPresentationHtmlBasicField               NgPresentationHtmlSubTemplate = "NgPresentationHtmlBasicField"
	NgPresentationHtmlTimeField                NgPresentationHtmlSubTemplate = "NgPresentationHtmlTimeField"
	NgPresentationHtmlBasicFieldTimeDuration   NgPresentationHtmlSubTemplate = "NgPresentationHtmlBasicFieldTimeDuration"
	NgPresentationHtmlBool                     NgPresentationHtmlSubTemplate = "NgPresentationHtmlBool"
	NgPresentationPointerToStructHtmlFormField NgPresentationHtmlSubTemplate = "NgPresentationPointerToStructHtmlFormField"
	// NgPresentationSliceOfPointerToStructHtml   NgPresentationHtmlSubTemplate = "NgPresentationSliceOfPointerToStructHtml"
)

var NgPresentationEnumHtmlSubTemplateCode map[string]string = map[string]string{
	string(NgPresentationHtmlEnum): `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{structname}}.{{FieldName}}}}</td>
    </tr>`,
}

var NgPresentationBasicFieldHtmlSubTemplateCode map[string]string = map[string]string{
	string(NgPresentationHtmlBasicField): `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{structname}}.{{FieldName}}}}</td>
    </tr>`,
}

var NgPresentationTimeFieldHtmlSubTemplateCode map[string]string = map[string]string{
	string(NgPresentationHtmlTimeField): `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{structname}}.{{FieldName}}}}</td>
    </tr>`,
}

var NgPresentationBasicFieldTimeDurationHtmlSubTemplateCode map[string]string = map[string]string{
	string(NgPresentationHtmlBasicFieldTimeDuration): `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{FieldName}}_Hours}}H {{{{FieldName}}_Minutes}}M {{{{FieldName}}_Seconds}}S</td>
    </tr>`,
}

var NgPresentationBoolHtmlSubTemplateCode map[string]string = map[string]string{
	string(NgPresentationHtmlBool): `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">
            <mat-checkbox [checked]="{{structname}}.{{FieldName}}" disabled=true>
            </mat-checkbox>
        </td>
    </tr>`,
}

var NgPresentationPointerToStructHtmlSubTemplateCode map[string]string = map[string]string{
	string(NgPresentationPointerToStructHtmlFormField): `
    <tr class="mat-row">
        <td class="mat-cell">{{FieldName}}</td>
        <td class="mat-cell">{{{{structname}}.{{FieldName}} ? {{structname}}.{{FieldName}}.Name : ""}}</td>
    </tr>`,
}

// var NgPresentationSliceOfPointerToStructHtmlSubTemplateCode map[string]string = map[string]string{
//     string(NgPresentationSliceOfPointerToStructHtml): `
//     <div class="mat-row">
//         <button mat-raised-button (click)="openReverseSelection('{{AssocStructName}}', '{{Structname}}_{{FieldName}}DBID')">{{FieldName}}</button>
//     </div>`,
// }
