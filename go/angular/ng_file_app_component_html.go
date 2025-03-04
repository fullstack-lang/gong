package angular

const NgFileAppComponentHtml = `<!-- button bar that is present in all tabs -->
<ng-template #radioToolbar>
    <mat-radio-group aria-label="Select an option" [(ngModel)]="view">
      @for (view of views; track view) {
        <mat-radio-button [value]="view">
          {{view}}&nbsp;&nbsp;&nbsp;
        </mat-radio-button>
      }
    </mat-radio-group>
</ng-template>

<as-split unit="pixel" *ngIf="view=={{pkgname}}" direction="vertical">
    <as-split-area [size]=40>
        <ng-container *ngTemplateOutlet="radioToolbar"></ng-container>
    </as-split-area>
    <as-split-area [size]="400">
        <lib-{{pkgname}}-specific GONG__StackPath="{{pkgname}}">
        </lib-{{pkgname}}-specific>
    </as-split-area>
    <as-split-area [size]="$any('*')">

        <as-split direction="vertical">
            <as-split-area [size]=50>
                <as-split direction="horizontal">
                    <as-split-area [size]="20">
                        <lib-tree-specific [GONG__StackPath]="StackName+TableExtraPathEnum.StackNamePostFixForTableForMainTree" name="gong"></lib-tree-specific>
                    </as-split-area>
                    <as-split-area [size]="50">
                        <div [ngStyle]="scrollStyle">
                            <lib-table-specific TableName="Table" [DataStack]="StackName+TableExtraPathEnum.StackNamePostFixForTableForMainTable"></lib-table-specific>
                        </div>
                    </as-split-area>
                    <as-split-area [size]="30">
                        <lib-form-specific FormName="Form" [DataStack]="StackName+TableExtraPathEnum.StackNamePostFixForTableForMainForm"></lib-form-specific>
                    </as-split-area>
                </as-split>
            </as-split-area>
            <as-split-area [size]=50>
                <lib-doc-specific [GONG__StackPath]="StackType"></lib-doc-specific>
            </as-split-area>
        </as-split>

    </as-split-area>
</as-split>

<as-split unit="pixel" *ngIf="view==probe" direction="vertical">
    <as-split-area [size]=40>
        <ng-container *ngTemplateOutlet="radioToolbar"></ng-container>
    </as-split-area>
    <as-split-area [size]="$any('*')">
    
        <as-split direction="vertical">
            <as-split-area [size]=50>
                <as-split direction="horizontal">
                    <as-split-area [size]="20">
                        <lib-tree-specific [GONG__StackPath]="StackName+TableExtraPathEnum.StackNamePostFixForTableForMainTree" name="gong"></lib-tree-specific>
                    </as-split-area>
                    <as-split-area [size]="50">
                        <div [ngStyle]="scrollStyle">
                            <lib-table-specific TableName="Table" [DataStack]="StackName+TableExtraPathEnum.StackNamePostFixForTableForMainTable"></lib-table-specific>
                        </div>
                    </as-split-area>
                    <as-split-area [size]="30">
                        <lib-form-specific FormName="Form" [DataStack]="StackName+TableExtraPathEnum.StackNamePostFixForTableForMainForm"></lib-form-specific>
                    </as-split-area>
                </as-split>
            </as-split-area>
            <as-split-area [size]=50>
                <lib-doc-specific [GONG__StackPath]="StackType"></lib-doc-specific>
            </as-split-area>
        </as-split>

    </as-split-area>

</as-split>
`
