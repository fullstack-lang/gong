package models

const NgSidebarTemplateHTML = `<mat-tree [dataSource]="dataSource" [treeControl]="treeControl">

    <!-- This is the tree node template for leaf nodes -->
    <mat-tree-node class="node-link" *matTreeNodeDef="let node" matTreeNodePadding>
        <!-- use a disabled button to provide padding for tree leaf -->
        <button mat-icon-button disabled></button>
        <span (click)="setTableRouterOutletFromTree( node.name + 's', node.type, node.structName, node.id)">{{node.name}} &nbsp;</span>
        <mat-icon class="node-link-icon" *ngIf='node.level==0' (click)="setEditorRouterOutlet( node.name + '-adder' )">
            add_circle_outline
        </mat-icon>
        <mat-icon class="node-link-icon" *ngIf='node.level==2 && node.type=="ONE__ZERO_MANY_ASSOCIATION"' (click)="setEditorSpecialRouterOutlet( node )" >
            add_circle_outline
        </mat-icon>
    </mat-tree-node>

    <!-- This is the tree node template for expandable nodes -->
    <mat-tree-node class="node-link" *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
        <button mat-icon-button matTreeNodeToggle [attr.aria-label]="'Toggle ' + node.name">
            <mat-icon class="mat-icon-rtl-mirror">
                {{treeControl.isExpanded(node) ?
                'expand_more' :
                'chevron_right'
                }}
            </mat-icon>
        </button>
        <span (click)="setTableRouterOutletFromTree( node.name + 's', node.type, node.structName, node.id)">{{node.name}}  &nbsp;</span>
        <mat-icon class="node-link-icon" *ngIf='node.level==0' (click)="setEditorRouterOutlet( node.name + '-adder' )">
            add_circle_outline
        </mat-icon>
        <mat-icon class="node-link-icon" *ngIf='node.level==2 && node.type=="ONE__ZERO_MANY_ASSOCIATION"' (click)="setEditorSpecialRouterOutlet( node )" >
            add_circle_outline
        </mat-icon>
    </mat-tree-node>
</mat-tree>
<h3>&nbsp;&nbsp;Nb commits:&nbsp; &nbsp;&nbsp;{{commitNb}}</h3>`

// insertion points
type NgSidebarHtmlInsertionPoint int

const (
	NgSidebarHtmlStruct NgSidebarHtmlInsertionPoint = iota
	NgSidebarHtmlNbInsertionPoints
)

//
// Sub Templates
//
type NgSidebarHtmlSubTemplate int

const (
	NgSidebarHtmlField NgSidebarHtmlSubTemplate = iota
)

var NgSidebarHtmlSubTemplateCode map[NgSidebarHtmlSubTemplate]string = map[NgSidebarHtmlSubTemplate]string{
	NgSidebarHtmlField: `
    <mat-list-item class="row-link">
        <span (click)="setTableRouterOutlet( '{{structname}}s' )">{{Structname}}s &nbsp;</span>
        <mat-icon (click)="setEditorRouterOutlet( '{{structname}}-adder' )">
            add_circle_outline
        </mat-icon>
    </mat-list-item>,`,
}
