package angular

const NgFileAppComponentHtml = `<!-- button bar that is present in all tabs -->
<ng-template #radioToolbar>
    <mat-radio-group aria-label="Select an option" [(ngModel)]="view">
        <mat-radio-button *ngFor="let view of views" [value]="view">
            {{view}}&nbsp;&nbsp;&nbsp;
        </mat-radio-button>
    </mat-radio-group>
</ng-template>

<as-split unit="pixel" *ngIf="view=={{pkgname}}" direction="vertical">
    <as-split-area [size]=40>
        <ng-container *ngTemplateOutlet="radioToolbar"></ng-container>
    </as-split-area>
    <as-split-area [size]="$any('*')">
        <lib-{{pkgname}}specific GONG__StackPath="{{pkgname}}">

        </lib-{{pkgname}}specific>
    </as-split-area>
</as-split>
`
