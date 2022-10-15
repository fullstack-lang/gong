package angular

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "embed"

	"github.com/fullstack-lang/gong/go/models"
)

//go:embed ng_file_sidebar.css
var NgFileSidebarCssTmpl string

const NgSidebarTemplateTS = `import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code{{` + string(rune(NgSidebarTsInsertionPerStructImports)) + `}}

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-{{pkgname}}-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbFromBackService: CommitNbFromBackService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration{{` + string(rune(NgSidebarTsInsertionPerStructServiceDeclaration)) + `}}
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger{{` + string(rune(NgSidebarTsInsertionPerStructObservableForRefresh)) + `}}
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction{{` + string(rune(NgSidebarTsInsertionPerStruct)) + `}}

      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    });

    // fetch the number of commits
    this.commitNbFromBackService.getCommitNbFromBack().subscribe(
      commitNbFromBack => {
        this.commitNbFromBack = commitNbFromBack
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        {{PkgPathRootWithoutSlashes}}_table: ["{{PkgPathRootWithoutSlashes}}-" + path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          {{PkgPathRootWithoutSlashes}}_table: ["{{PkgPathRootWithoutSlashes}}-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          {{PkgPathRootWithoutSlashes}}_presentation: ["{{PkgPathRootWithoutSlashes}}-" + structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        {{PkgPathRootWithoutSlashes}}_editor: ["{{PkgPathRootWithoutSlashes}}-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        {{PkgPathRootWithoutSlashes}}_editor: ["{{PkgPathRootWithoutSlashes}}-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}
`

// insertion points in the main template
type NgSidebarTsInsertionPoint int

const (
	NgSidebarTsInsertionPerStruct NgSidebarTsInsertionPoint = iota
	NgSidebarTsInsertionPerStructImports
	NgSidebarTsInsertionPerStructServiceDeclaration
	NgSidebarTsInsertionPerStructObservableForRefresh
	NgSidebarTsInsertionsNb
)

// Sub Templates
type NgSidebarTsSubTemplate int

const (
	NgSidebarTsPerStructNode NgSidebarTsSubTemplate = iota
	NgSidebarTsPerStructNodeImports
	NgSidebarTsPerStructNodeServiceDeclaration
	NgSidebarTsPerStructNodeObservableForRefresh
)

var NgSidebarTsSubTemplateCode map[NgSidebarTsSubTemplate]string = map[NgSidebarTsSubTemplate]string{
	NgSidebarTsPerStructNode: `
      /**
      * fill up the {{Structname}} part of the mat tree
      */
      let {{structname}}GongNodeStruct: GongNode = {
        name: "{{Structname}}",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "{{Structname}}",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push({{structname}}GongNodeStruct)

      this.frontRepo.{{Structname}}s_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.{{Structname}}s_array.forEach(
        {{structname}}DB => {
          let {{structname}}GongNodeInstance: GongNode = {
            name: {{structname}}DB.Name,
            type: GongNodeType.INSTANCE,
            id: {{structname}}DB.ID,
            uniqueIdPerStack: get{{Structname}}UniqueID({{structname}}DB.ID),
            structName: "{{Structname}}",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          {{structname}}GongNodeStruct.children!.push({{structname}}GongNodeInstance)

          // insertion point for per field code{{` + string(rune(NgSidebarTsSubInsertionPerField)) + `}}
        }
      )
`,

	NgSidebarTsPerStructNodeImports: `
import { {{Structname}}Service } from '../{{structname}}.service'
import { get{{Structname}}UniqueID } from '../front-repo.service'`,

	NgSidebarTsPerStructNodeServiceDeclaration: `
    private {{structname}}Service: {{Structname}}Service,`,

	NgSidebarTsPerStructNodeObservableForRefresh: `
    // observable for changes in structs
    this.{{structname}}Service.{{Structname}}ServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )`,
}

// insertion points in the main template
type NgSidebarTsSubInsertionPoint int

const (
	NgSidebarTsSubInsertionPerField NgSidebarTsSubInsertionPoint = iota
	NgSidebarTsSubInsertionsNb
)

// Sub Templates
type NgSidebarTsStructSubTemplate int

const (
	NgSidebarTsPerStructPointerToStructFieldTemplateNode NgSidebarTsStructSubTemplate = iota
	NgSidebarTsPerStructSliceOfPointerToStructFieldTemplateNode
)

var NgSidebarTsSubPerFieldTemplateCode map[NgSidebarTsStructSubTemplate]string = map[NgSidebarTsStructSubTemplate]string{
	NgSidebarTsPerStructPointerToStructFieldTemplateNode: `
          /**
          * let append a node for the association {{Fieldname}}
          */
          let {{Fieldname}}GongNodeAssociation: GongNode = {
            name: "({{AssociatedStructname}}) {{Fieldname}}",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: {{structname}}DB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "{{Structname}}",
            associationField: "{{Fieldname}}",
            associatedStructName: "{{AssociatedStructname}}",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          {{structname}}GongNodeInstance.children!.push({{Fieldname}}GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation {{Fieldname}}
            */
          if ({{structname}}DB.{{Fieldname}} != undefined) {
            let {{structname}}GongNodeInstance_{{Fieldname}}: GongNode = {
              name: {{structname}}DB.{{Fieldname}}.Name,
              type: GongNodeType.INSTANCE,
              id: {{structname}}DB.{{Fieldname}}.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * get{{Structname}}UniqueID({{structname}}DB.ID)
                + 5 * get{{AssociatedStructname}}UniqueID({{structname}}DB.{{Fieldname}}.ID),
              structName: "{{AssociatedStructname}}",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            {{Fieldname}}GongNodeAssociation.children.push({{structname}}GongNodeInstance_{{Fieldname}})
          }
`,
	NgSidebarTsPerStructSliceOfPointerToStructFieldTemplateNode: `
          /**
          * let append a node for the slide of pointer {{Fieldname}}
          */
          let {{Fieldname}}GongNodeAssociation: GongNode = {
            name: "({{AssociatedStructname}}) {{Fieldname}}",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: {{structname}}DB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "{{Structname}}",
            associationField: "{{Fieldname}}",
            associatedStructName: "{{AssociatedStructname}}",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          {{structname}}GongNodeInstance.children.push({{Fieldname}}GongNodeAssociation)

          {{structname}}DB.{{Fieldname}}?.forEach({{associatedStructname}}DB => {
            let {{associatedStructname}}Node: GongNode = {
              name: {{associatedStructname}}DB.Name,
              type: GongNodeType.INSTANCE,
              id: {{associatedStructname}}DB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * get{{Structname}}UniqueID({{structname}}DB.ID)
                + 11 * get{{AssociatedStructname}}UniqueID({{associatedStructname}}DB.ID),
              structName: "{{AssociatedStructname}}",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            {{Fieldname}}GongNodeAssociation.children.push({{associatedStructname}}Node)
          })
`,
}

// MultiCodeGeneratorNgSidebar parses mdlPkg and generates the code for the
// Sidebar components
func CodeGeneratorNgSidebar(
	mdlPkg *models.ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string) {

	// create the component directory
	dirPath := filepath.Join(matTargetPath, "sidebar")
	errd := os.MkdirAll(dirPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + dirPath)
	}

	// generate the css file
	models.VerySimpleCodeGenerator(mdlPkg,
		pkgName,
		pkgGoPath,
		filepath.Join(dirPath, "sidebar.component.css"),
		NgFileSidebarCssTmpl,
	)

	// have alphabetical order generation
	structList := []*models.GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	// generate the typescript file
	codeHTML := NgSidebarTemplateHTML

	HTMLinsertions := make(map[NgSidebarHtmlInsertionPoint]string)
	for insertion := NgSidebarHtmlInsertionPoint(0); insertion < NgSidebarHtmlNbInsertionPoints; insertion++ {
		HTMLinsertions[insertion] = ""
	}

	// generate the typescript file
	codeTS := NgSidebarTemplateTS

	TSinsertions := make(map[NgSidebarTsInsertionPoint]string)
	for insertion := NgSidebarTsInsertionPoint(0); insertion < NgSidebarTsInsertionsNb; insertion++ {
		TSinsertions[insertion] = ""
	}

	for _, _struct := range structList {

		HTMLinsertions[NgSidebarHtmlStruct] += models.Replace2(
			NgSidebarHtmlSubTemplateCode[NgSidebarHtmlField],
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))

		// only generate a STRUCT node if there is a "Name" field in the gong struct
		if _struct.HasNameField() {

			TSinsertions[NgSidebarTsInsertionPerStructImports] += models.Replace2(
				NgSidebarTsSubTemplateCode[NgSidebarTsPerStructNodeImports],
				"{{Structname}}", _struct.Name,
				"{{structname}}", strings.ToLower(_struct.Name))

			TSinsertions[NgSidebarTsInsertionPerStructServiceDeclaration] += models.Replace2(
				NgSidebarTsSubTemplateCode[NgSidebarTsPerStructNodeServiceDeclaration],
				"{{Structname}}", _struct.Name,
				"{{structname}}", strings.ToLower(_struct.Name))

			TSinsertions[NgSidebarTsInsertionPerStructObservableForRefresh] += models.Replace2(
				NgSidebarTsSubTemplateCode[NgSidebarTsPerStructNodeObservableForRefresh],
				"{{Structname}}", _struct.Name,
				"{{structname}}", strings.ToLower(_struct.Name))

			TSinsertions[NgSidebarTsInsertionPerStruct] += models.Replace2(
				NgSidebarTsSubTemplateCode[NgSidebarTsPerStructNode],
				"{{Structname}}", _struct.Name,
				"{{structname}}", strings.ToLower(_struct.Name))

			// prepare the insertions per field
			TSSubInsertions := make(map[NgSidebarTsSubInsertionPoint]string)
			for insertion := NgSidebarTsSubInsertionPoint(0); insertion < NgSidebarTsSubInsertionsNb; insertion++ {
				TSSubInsertions[insertion] = ""
			}

			// Generate association nodes
			for _, field := range _struct.Fields {
				switch field := field.(type) {
				case *models.PointerToGongStructField:

					TSSubInsertions[NgSidebarTsSubInsertionPerField] += models.Replace4(
						NgSidebarTsSubPerFieldTemplateCode[NgSidebarTsPerStructPointerToStructFieldTemplateNode],
						"{{Structname}}", _struct.Name,
						"{{structname}}", strings.ToLower(_struct.Name),
						"{{Fieldname}}", field.Name,
						"{{AssociatedStructname}}", field.GongStruct.Name)

				case *models.SliceOfPointerToGongStructField:

					TSSubInsertions[NgSidebarTsSubInsertionPerField] += models.Replace5(
						NgSidebarTsSubPerFieldTemplateCode[NgSidebarTsPerStructSliceOfPointerToStructFieldTemplateNode],
						"{{Structname}}", _struct.Name,
						"{{structname}}", strings.ToLower(_struct.Name),
						"{{Fieldname}}", field.Name,
						"{{associatedStructname}}", strings.ToLower(field.GongStruct.Name),
						"{{AssociatedStructname}}", field.GongStruct.Name)
				}
			}
			// substitutes {{<<insertion points>>}} stuff with generated code
			for subInsertion := NgSidebarTsSubInsertionPoint(0); subInsertion < NgSidebarTsSubInsertionsNb; subInsertion++ {
				toReplace := "{{" + string(rune(subInsertion)) + "}}"
				TSinsertions[NgSidebarTsInsertionPerStruct] =
					strings.ReplaceAll(TSinsertions[NgSidebarTsInsertionPerStruct], toReplace, TSSubInsertions[subInsertion])
			}

		}
	}

	// substitutes {{<<insertion points>>}} stuff with generated code
	for insertion := NgSidebarHtmlInsertionPoint(0); insertion < NgSidebarHtmlNbInsertionPoints; insertion++ {
		toReplace := "{{" + string(rune(insertion)) + "}}"
		codeHTML = strings.ReplaceAll(codeHTML, toReplace, HTMLinsertions[insertion])
	}

	for insertion := NgSidebarTsInsertionPoint(0); insertion < NgSidebarTsInsertionsNb; insertion++ {
		toReplace := "{{" + string(rune(insertion)) + "}}"
		codeTS = strings.ReplaceAll(codeTS, toReplace, TSinsertions[insertion])
	}

	codeHTML = models.Replace4(codeHTML,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", strings.Title(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""))

	{
		file, err := os.Create(filepath.Join(dirPath, "sidebar.component.html"))
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		fmt.Fprint(file, codeHTML)
	}

	pkgPathRootWithoutSlashes := strings.ReplaceAll(pkgGoPath, "/models", "")
	pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, "/", "_")
	pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, "-", "_")
	pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, ".", "_")

	codeTS = models.Replace5(codeTS,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", strings.Title(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
		"{{PkgPathRootWithoutSlashes}}", pkgPathRootWithoutSlashes)

	{
		file, err := os.Create(filepath.Join(dirPath, "sidebar.component.ts"))
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		fmt.Fprint(file, codeTS)
	}
}
