package models

const NgFileAppComponentHtml = `<as-split *ngIf="view==default" direction="vertical">
    <as-split-area [size]=05>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <mat-radio-group aria-label="Select an option" [(ngModel)]="view">
            <mat-radio-button *ngFor="let view of views" [value]="view">
                {{view}}&nbsp;&nbsp;&nbsp;
            </mat-radio-button>
        </mat-radio-group>
    </as-split-area>
    <as-split-area [size]=95>
	    <app-{{pkgname}}-splitter></app-{{pkgname}}-splitter>
    </as-split-area>
</as-split>

<as-split *ngIf="view==meta" direction="vertical">
    <as-split-area [size]=05>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <mat-radio-group aria-label="Select an option" [(ngModel)]="view">
            <mat-radio-button *ngFor="let view of views" [value]="view">
                {{view}}&nbsp;&nbsp;&nbsp;
            </mat-radio-button>
        </mat-radio-group>
    </as-split-area>
    <as-split-area [size]=95>
        <app-gong-splitter></app-gong-splitter>
    </as-split-area>
</as-split>


<as-split *ngIf="view==diagrams" direction="vertical">
    <as-split-area [size]=05>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <mat-radio-group aria-label="Select an option" [(ngModel)]="view">
            <mat-radio-button *ngFor="let view of views" [value]="view">
                {{view}}&nbsp;&nbsp;&nbsp;
            </mat-radio-button>
        </mat-radio-group>
    </as-split-area>
    <as-split-area [size]=95>
        <lib-pkgelt-docs></lib-pkgelt-docs>
    </as-split-area>
</as-split>


`
