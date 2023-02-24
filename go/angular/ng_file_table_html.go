package angular

const NgTableTemplateHTML = `<section class="container mat-elevation-z8" tabindex="0">
    <mat-form-field>
        <mat-label>Filter</mat-label>
        <input matInput (keyup)="applyFilter($event)" placeholder="Ex. ium" #input>
    </mat-form-field>
    <h1 *ngIf="dialogData">{{structname}}</h1>
    <table class="table" mat-table [dataSource]="matTableDataSource" matSort>
        <!-- Checkbox Column, only used when the component is used to select items for a field that is a slice of pointers -->
        <ng-container matColumnDef="select">
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
            <th mat-header-cell *matHeaderCellDef mat-sort-header> ID. </th>
            <td mat-cell *matCellDef="let {{Structname}}"> {{{{Structname}}.ID}} </td>
        </ng-container>

        <ng-container matColumnDef="Delete" sticky>
            <th mat-header-cell *matHeaderCellDef> Delete </th>
            <td mat-cell *matCellDef="let {{structname}};  let j = index;">
                <i class="material-icons my-delete-button" [ngStyle]="{'color':'rgba(0,0,0,.50)'}" (click)="delete{{Structname}}({{structname}}.ID, {{structname}}); $event.stopPropagation()">delete</i>
            </td>
        </ng-container>

        <tr mat-header-row *matHeaderRowDef="displayedColumns;"></tr>

        <tr mat-row *matRowDef="
        let row; 
        columns: displayedColumns;
        " (click)="setEditorRouterOutlet( row.ID ) " class="row-link">
        </tr>

        <!-- Row shown when there is no matching data. -->
        <tr class="mat-row" *matNoDataRow>
            <td class="mat-cell" colspan="4">No data matching the filter "{{input.value}}"</td>
        </tr>
    </table>
    <mat-paginator [pageSizeOptions]="[50, 100, 500, 1000]" showFirstLastButtons></mat-paginator>
    <button class="table__save" color="primary" *ngIf="dialogData" mat-raised-button (click)="save()">
        Save
    </button>
</section>`

// insertion points in the main template
type NgTableHtmlInsertionPoint int

const (
	NgTableHtmlInsertionColumn NgTableHtmlInsertionPoint = iota
	NgTableHtmlInsertionsNb
)

type NgTableHTMLSubTemplate int

const (
	NgTableHTMLBasicField NgTableHTMLSubTemplate = iota
	NgTableHTMLEnumIntField
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
        <ng-container matColumnDef="{{FieldName}}"{{sticky}}>
            <th mat-header-cell *matHeaderCellDef mat-sort-header> {{FieldName}} </th>
            <td mat-cell *matCellDef="let {{Structname}}">
                {{{{Structname}}.{{FieldName}}}}
            </td>
        </ng-container>`,

	NgTableHTMLEnumIntField: `
        <!-- -->
        <ng-container matColumnDef="{{FieldName}}">
            <th mat-header-cell *matHeaderCellDef mat-sort-header> {{FieldName}} </th>
            <td mat-cell *matCellDef="let {{Structname}}">
                {{{{Structname}}.{{FieldName}}_string}}
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
