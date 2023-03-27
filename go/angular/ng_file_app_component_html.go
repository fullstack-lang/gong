package angular

const NgFileAppComponentHtml = `<!-- button bar that is present in all tabs -->
<ng-template #radioToolbar>
    <mat-radio-group aria-label="Select an option" [(ngModel)]="view">
        <mat-radio-button *ngFor="let view of views" [value]="view">
            {{view}}&nbsp;&nbsp;&nbsp;
        </mat-radio-button>
    </mat-radio-group>
</ng-template>

<as-split unit="pixel" *ngIf="view==default" direction="vertical">
    <as-split-area [size]=40>
        <ng-container *ngTemplateOutlet="radioToolbar"></ng-container>
    </as-split-area>
    <as-split-area [size]="$any('*')">
        <app-{{pkgname}}-splitter GONG__StackPath=""></app-{{pkgname}}-splitter>
    </as-split-area>
</as-split>


<as-split unit="pixel" *ngIf="view==model" direction="vertical">
    <as-split-area [size]=40>
        <ng-container *ngTemplateOutlet="radioToolbar"></ng-container>
    </as-split-area>
    <as-split-area [size]="$any('*')">
        <lib-panel [GONG__StackPath]="GONG__StackPath"></lib-panel>
    </as-split-area>
</as-split>
`
