package models

const NgTableTemplateHTML = `<div>
    <mat-form-field>
        <mat-label>Filter</mat-label>
        <input matInput (keyup)="applyFilter($event)" placeholder="Ex. ium" #input>
    </mat-form-field>
    <h1 *ngIf="dialogData">{{structname}}</h1>
    <table class="table" mat-table [dataSource]="matTableDataSource" matSort>
        <!-- Checkbox Column -->
        <ng-container matColumnDef="select" sticky>
            <th mat-header-cell *matHeaderCellDef>
                <mat-checkbox (change)="$event ? masterToggle() : null" [checked]="selection.hasValue() && isAllSelected()" [indeterminate]="selection.hasValue() && !isAllSelected()">
                </mat-checkbox>
            </th>
            <td mat-cell *matCellDef="let row">
                <mat-checkbox (click)="$event.stopPropagation()" (change)="$event ? selection.toggle(row) : null" [checked]="selection.isSelected(row)">
                </mat-checkbox>
            </td>
        </ng-container>

        <!-- insertion point for fields specific code -->{{` + string(rune(NgTableHtmlInsertionColumn)) + `}}

        <ng-container matColumnDef="ID" sticky>
            <th mat-header-cell *matHeaderCellDef> ID. </th>
            <td mat-cell *matCellDef="let {{Structname}}"> {{{{Structname}}.ID}} </td>
        </ng-container>

        <ng-container matColumnDef="Edit">
            <th mat-header-cell *matHeaderCellDef> Edit </th>
            <td mat-cell *matCellDef="let {{structname}};  let j = index;">
                <i class="material-icons" [ngStyle]="{'color':'rgba(0,0,0,.50)'}" (click)="setEditorRouterOutlet({{structname}}.ID)">edit</i>
            </td>
        </ng-container>

        <ng-container matColumnDef="Delete">
            <th mat-header-cell *matHeaderCellDef> Delete </th>
            <td mat-cell *matCellDef="let {{structname}};  let j = index;">
                <i class="material-icons" [ngStyle]="{'color':'rgba(0,0,0,.50)'}" (click)="delete{{Structname}}({{structname}}.ID, {{structname}})">delete</i>
            </td>
        </ng-container>

        <tr mat-header-row *matHeaderRowDef="displayedColumns; sticky: true"></tr>

        <tr mat-row *matRowDef="
        let row; 
        columns: displayedColumns;
        " (click)="setPresentationRouterOutlet( row.ID ) " class="row-link">
        </tr>

        <!-- Row shown when there is no matching data. -->
        <tr class="mat-row" *matNoDataRow>
            <td class="mat-cell" colspan="4">No data matching the filter "{{input.value}}"</td>
        </tr>
    </table>
    <mat-paginator [pageSizeOptions]="[10, 20, 50, 100, 500, 1000]" showFirstLastButtons></mat-paginator>
</div>
<button class="table__save" color="primary" *ngIf="dialogData" mat-raised-button (click)="save()">
    Save
</button>`

// insertion points in the main template
type NgTableHtmlInsertionPoint int

const (
	NgTableHtmlInsertionColumn NgTableHtmlInsertionPoint = iota
	NgTableHtmlInsertionsNb
)

type NgTableHTMLSubTemplate int

const (
	NgTableHTMLBasicField NgTableHTMLSubTemplate = iota
	NgTableHTMLTimeField
	NgTableHTMLBasicFloat64Field
	NgTableHTMLBasicFieldTimeDuration
	NgTableHTMLBool
	NgTablePointerToStructHTMLFormField
	NgTablePointerToSliceOfGongStructHTMLFormField
)

var NgTableHTMLSubTemplateCode map[NgTableHTMLSubTemplate]string = map[NgTableHTMLSubTemplate]string{

	NgTableHTMLBasicField: `
        <!-- -->
        <ng-container matColumnDef="{{FieldName}}">
            <th mat-header-cell *matHeaderCellDef mat-sort-header> {{FieldName}} </th>
            <td mat-cell *matCellDef="let {{Structname}}">
                {{{{Structname}}.{{FieldName}}}}
            </td>
        </ng-container>`,

	NgTableHTMLTimeField: `
        <!-- -->
        <ng-container matColumnDef="{{FieldName}}">
            <th mat-header-cell *matHeaderCellDef mat-sort-header> {{FieldName}} </th>
            <td mat-cell *matCellDef="let {{Structname}}">
                {{{{Structname}}.{{FieldName}}}}
            </td>
        </ng-container>`,

	NgTableHTMLBasicFieldTimeDuration: `
        <!-- -->
        <ng-container matColumnDef="{{FieldName}}">
            <th mat-header-cell *matHeaderCellDef mat-sort-header> {{FieldName}} </th>
            <td mat-cell *matCellDef="let {{Structname}}">
                {{{{Structname}}.{{FieldName}}_string}}
            </td>
        </ng-container>`,

	NgTableHTMLBasicFloat64Field: `
        <!-- -->
        <ng-container matColumnDef="{{FieldName}}">
            <th mat-header-cell *matHeaderCellDef mat-sort-header> {{FieldName}} </th>
            <td mat-cell *matCellDef="let {{Structname}}">
                {{{{Structname}}.{{FieldName}}.toPrecision(5)}}
            </td>
        </ng-container>`,

	NgTableHTMLBool: `
        <!-- -->
        <ng-container matColumnDef="{{FieldName}}">
            <th mat-header-cell *matHeaderCellDef mat-sort-header> {{FieldName}} </th>
            <td mat-cell *matCellDef="let {{Structname}}">
                <mat-checkbox [checked]="{{Structname}}.{{FieldName}}" disabled=true></mat-checkbox>
            </td>
        </ng-container>`,

	NgTablePointerToStructHTMLFormField: `
        <!-- -->
        <ng-container matColumnDef="{{FieldName}}">
            <th mat-header-cell *matHeaderCellDef mat-sort-header> {{FieldName}} </th>
            <td mat-cell *matCellDef="let {{Structname}}">
                {{ {{Structname}}.{{FieldName}} ? {{Structname}}.{{FieldName}}.Name : ''}}
            </td>
        </ng-container>`,

	NgTablePointerToSliceOfGongStructHTMLFormField: `
        <!-- -->
        <ng-container matColumnDef="{{AssocStructName}}_{{FieldName}}">
            <th mat-header-cell *matHeaderCellDef mat-sort-header> <-- ({{AssocStructName}}) {{FieldName}} </th>
            <td mat-cell *matCellDef="let {{structname}}">
                {{frontRepo.{{AssocStructName}}s.get({{structname}}.{{AssocStructName}}_{{FieldName}}DBID.Int64)?.Name}}
            </td>
        </ng-container>`,
}
