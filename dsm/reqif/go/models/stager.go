// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fullstack-lang/gong/dsm/reqif/go/icons"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/gin-gonic/gin"

	table "github.com/fullstack-lang/gong/lib/table/go/models"
	table_stack "github.com/fullstack-lang/gong/lib/table/go/stack"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"

	ssg_level1stack "github.com/fullstack-lang/gong/lib/ssg/go/level1stack"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"

	markdown "github.com/fullstack-lang/gong/lib/markdown/go/models"
	markdown_stack "github.com/fullstack-lang/gong/lib/markdown/go/stack"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

type ModelGeneratorInterface interface {
	GenerateModels(stager *Stager)
}

// data types

type DataTypesTreeUpdaterInterface interface {
	UpdateAndCommitDataTypeTreeStage(stager *Stager)
}

// spec types

type SpecTypesTreeUpdaterInterface interface {
	UpdateAndCommitSpecTypesTreeStage(stager *Stager)
}

// instances

type SpecObjectsTreeUpdaterInterface interface {
	UpdateAndCommitSpecObjectsTreeStage(stager *Stager)
}

type SpecRelationsTreeUpdaterInterface interface {
	UpdateAndCommitSpecRelationsTreeStage(stager *Stager)
}

type SpecificationsTreeUpdaterInterface interface {
	UpdateAndCommitSpecificationsTreeStage(stager *Stager)
	UpdateAndCommitSpecificationsMarkdownStage(stager *Stager)
	UpdateAttributeDefinitionNb(
		stager *Stager,
	)
}

type Exporter interface {
	ExportReqif(stager *Stager)
	ExportRenderingConf(stager *Stager)
	ExportAnonymousReqif(stager *Stager)
}

type ObjectNamerInterface interface {
	SetNamesToElements(stage *Stage, reqif *REQ_IF)
}

type Stager struct {
	stage      *Stage
	splitStage *split.Stage

	welcomeTabButtonStage   *button.Stage
	renderingTabButtonStage *button.Stage
	anonymousButtonStage    *button.Stage

	summaryTableStage *table.Stage

	dataTypeTreeStage       *tree.Stage
	specTypesTreeStage      *tree.Stage
	specObjectsTreeStage    *tree.Stage
	specRelationsTreeStage  *tree.Stage
	specificationsTreeStage *tree.Stage

	// rendering

	markdownStage *markdown.Stage

	rootReqif             *REQ_IF
	PathToReqifFile       string
	pathToRenderingConf   string
	pathToOutputReqifFile string

	// those interface allows for processing outside the models
	// package. This is for compilation time sake
	// indeed, since the models is quite complex, the compilation
	// of the models package can last an unusual long time
	// for a go compilation. Therefore, processing is delegated to other
	// package where modification will be compiled in a much faster time
	modelGenerator ModelGeneratorInterface

	dataTypesTreeUpdater DataTypesTreeUpdaterInterface
	specTypesTreeUpdater SpecTypesTreeUpdaterInterface
	specObjectsUX        SpecObjectsTreeUpdaterInterface
	specRelationsUX      SpecRelationsTreeUpdaterInterface
	specificationsUX     SpecificationsTreeUpdaterInterface
	reqifExporter        Exporter
	objectNamer          ObjectNamerInterface

	// maps for navigating the ReqIF data
	Map_id_DATATYPE_DEFINITION_XHTML       map[string]*DATATYPE_DEFINITION_XHTML
	Map_id_DATATYPE_DEFINITION_STRING      map[string]*DATATYPE_DEFINITION_STRING
	Map_id_DATATYPE_DEFINITION_BOOLEAN     map[string]*DATATYPE_DEFINITION_BOOLEAN
	Map_id_DATATYPE_DEFINITION_INTEGER     map[string]*DATATYPE_DEFINITION_INTEGER
	Map_id_DATATYPE_DEFINITION_REAL        map[string]*DATATYPE_DEFINITION_REAL
	Map_id_DATATYPE_DEFINITION_DATE        map[string]*DATATYPE_DEFINITION_DATE
	Map_id_DATATYPE_DEFINITION_ENUMERATION map[string]*DATATYPE_DEFINITION_ENUMERATION

	Map_id_SPEC_OBJECT_TYPE   map[string]*SPEC_OBJECT_TYPE
	Map_id_SPECIFICATION_TYPE map[string]*SPECIFICATION_TYPE
	Map_id_SPEC_OBJECT        map[string]*SPEC_OBJECT

	Map_id_ATTRIBUTE_DEFINITION_XHTML       map[string]*ATTRIBUTE_DEFINITION_XHTML
	Map_id_ATTRIBUTE_DEFINITION_STRING      map[string]*ATTRIBUTE_DEFINITION_STRING
	Map_id_ATTRIBUTE_DEFINITION_BOOLEAN     map[string]*ATTRIBUTE_DEFINITION_BOOLEAN
	Map_id_ATTRIBUTE_DEFINITION_INTEGER     map[string]*ATTRIBUTE_DEFINITION_INTEGER
	Map_id_ATTRIBUTE_DEFINITION_DATE        map[string]*ATTRIBUTE_DEFINITION_DATE
	Map_id_ATTRIBUTE_DEFINITION_REAL        map[string]*ATTRIBUTE_DEFINITION_REAL
	Map_id_ATTRIBUTE_DEFINITION_ENUMERATION map[string]*ATTRIBUTE_DEFINITION_ENUMERATION

	Map_SPEC_OBJECT_TYPE_Spec_nbInstance   map[*SPEC_OBJECT_TYPE]int
	Map_SPECIFICATION_TYPE_Spec_nbInstance map[*SPECIFICATION_TYPE]int
	Map_SPEC_RELATION_TYPE_Spec_nbInstance map[*SPEC_RELATION_TYPE]int

	Map_ATTRIBUTE_DEFINITION_XHTML_Spec_nbInstance       map[*ATTRIBUTE_DEFINITION_XHTML]int
	Map_ATTRIBUTE_DEFINITION_STRING_Spec_nbInstance      map[*ATTRIBUTE_DEFINITION_STRING]int
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_Spec_nbInstance     map[*ATTRIBUTE_DEFINITION_BOOLEAN]int
	Map_ATTRIBUTE_DEFINITION_INTEGER_Spec_nbInstance     map[*ATTRIBUTE_DEFINITION_INTEGER]int
	Map_ATTRIBUTE_DEFINITION_DATE_Spec_nbInstance        map[*ATTRIBUTE_DEFINITION_DATE]int
	Map_ATTRIBUTE_DEFINITION_REAL_Spec_nbInstance        map[*ATTRIBUTE_DEFINITION_REAL]int
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_Spec_nbInstance map[*ATTRIBUTE_DEFINITION_ENUMERATION]int

	Map_ENUM_VALUE_Spec_nbInstance map[*ENUM_VALUE]int

	Map_id_ENUM_VALUE map[string]*ENUM_VALUE

	Map_id_SPEC_RELATION_TYPE map[string]*SPEC_RELATION_TYPE

	// attributes for generating static web site
	ssgStage *ssg.Stage

	// for generating the static site, one needs to
	// (1) extract the default files of SSG into a directory
	pathToExtractedGongSsgDefaultStaticFiles string
	pathToExtractedGongSsgDefaultLayoutFiles string

	// (2) copy those files to a directory
	// where images generated by one will be added
	pathToFilesUsedForSsgGeneration string

	// (3) the will serve for the generation of markdown files
	rootPathToImageInputs                   string
	pathToFilesForSsgGeneratedMarkdownFiles string

	// (4) that will be used for the path to the ssg site
	rootPathToStaticOutputs string

	loadReqifStage         *load.Stage // load the reqif file
	loadRenderingConfStage *load.Stage // load the rendering conf file

	// allow for navigation from spec object to their relations
	Map_SPEC_OBJECT_relations_sources map[*SPEC_OBJECT][]*SPEC_RELATION
	Map_SPEC_OBJECT_relations_targets map[*SPEC_OBJECT][]*SPEC_RELATION
}

func (stager *Stager) GetStage() (stage *Stage) {
	stage = stager.stage
	return
}

// Tree for Data Types

func (stager *Stager) GetDataTypeTreeStage() (dataTypeTreeStage *tree.Stage) {
	dataTypeTreeStage = stager.dataTypeTreeStage
	return
}

// Tree for spec types

func (stager *Stager) GetSpecTypesTreeStage() (s *tree.Stage) {
	s = stager.specTypesTreeStage
	return
}

// Tree for spec objects

func (stager *Stager) GetSpecObjectsTreeStage() (s *tree.Stage) {
	s = stager.specObjectsTreeStage
	return
}

// Tree for spec relations

func (stager *Stager) GetSpecRelationsTreeStage() (s *tree.Stage) {
	s = stager.specRelationsTreeStage
	return
}

// Tree for specification

func (stager *Stager) GetSpecificationsTreeStage() (s *tree.Stage) {
	s = stager.specificationsTreeStage
	return
}

func (stager *Stager) GetMarkdownStage() (s *markdown.Stage) {
	s = stager.markdownStage
	return
}

func (stager *Stager) GetSsgStage() (s *ssg.Stage) {
	s = stager.ssgStage
	return
}

func (stager *Stager) GetLoadStage() (s *load.Stage) {
	s = stager.loadReqifStage
	return
}

func (stager *Stager) GetPathToOutputReqifFile() string {
	return stager.pathToOutputReqifFile
}

func (stager *Stager) GetPathToReqifFile() string {
	return stager.PathToReqifFile
}

// OTHERS

func (stager *Stager) GetRootREQIF() (rootReqif *REQ_IF) {
	rootReqif = stager.rootReqif
	return
}

func (stager *Stager) SetModelGenerator(modelGenerator ModelGeneratorInterface) {
	stager.modelGenerator = modelGenerator
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	pathToReqifFile string,
	pathToRenderingConf string,
	pathToOutputReqifFile string,

	dataTypesTreeUpdater DataTypesTreeUpdaterInterface,
	specTypesTreeUpdater SpecTypesTreeUpdaterInterface,

	specObjectsTreeUpdater SpecObjectsTreeUpdaterInterface,
	specRelationsTreeUpdater SpecRelationsTreeUpdaterInterface,
	specificationsTreeUpdater SpecificationsTreeUpdaterInterface,

	exporter Exporter,

	objectNamer ObjectNamerInterface) (
	stager *Stager,
) {

	stager = new(Stager)

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.splitStage = splitStage

	stager.stage = stage
	stager.splitStage = splitStage
	stager.PathToReqifFile = pathToReqifFile
	stager.pathToRenderingConf = pathToRenderingConf
	stager.pathToOutputReqifFile = pathToOutputReqifFile

	stager.dataTypesTreeUpdater = dataTypesTreeUpdater
	stager.specTypesTreeUpdater = specTypesTreeUpdater
	stager.specObjectsUX = specObjectsTreeUpdater
	stager.specRelationsUX = specRelationsTreeUpdater
	stager.specificationsUX = specificationsTreeUpdater

	stager.reqifExporter = exporter

	stager.objectNamer = objectNamer

	stager.ssgStage = ssg_level1stack.NewLevel1Stack(stage.GetName(), "", "", true, true).Stage
	stager.loadReqifStage = load_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage
	stager.loadRenderingConfStage = load_stack.NewStack(r, stage.GetName()+"-rendering", "", "", "", true, true).Stage
	stager.markdownStage = markdown_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage

	stager.summaryTableStage = table_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage
	stager.dataTypeTreeStage = tree_stack.NewStack(r, stage.GetName()+"-data types", "", "", "", true, true).Stage
	stager.specTypesTreeStage = tree_stack.NewStack(r, stage.GetName()+"-spec types", "", "", "", true, true).Stage

	stager.specObjectsTreeStage = tree_stack.NewStack(r, stage.GetName()+"-spec objects", "", "", "", true, true).Stage
	stager.specRelationsTreeStage = tree_stack.NewStack(r, stage.GetName()+"-spec relations", "", "", "", true, true).Stage
	stager.specificationsTreeStage = tree_stack.NewStack(r, stage.GetName()+"-specifications", "", "", "", true, true).Stage
	stager.welcomeTabButtonStage = button_stack.NewStack(r, stage.GetName()+"-Specification Tab", "", "", "", true, true).Stage
	stager.renderingTabButtonStage = button_stack.NewStack(r, stage.GetName()+"-Render Tab", "", "", "", true, true).Stage
	stager.anonymousButtonStage = button_stack.NewStack(r, stage.GetName()+"-Anonymous", "", "", "", true, true).Stage

	// add a title to the web application
	(&split.Title{Name: "GONG ReqIF"}).Stage(stager.splitStage)
	(&split.FavIcon{Name: "Test",
		SVG: icons.ReqifIconSVG,
	}).Stage(stager.splitStage)
	logoOnTheLeft := (&split.LogoOnTheLeft{
		Name:   "Gong Reqif Icons",
		Width:  108,
		Height: 127,
		SVG:    icons.ReqifIconSVG,
	}).Stage(stager.splitStage)
	logoOnTheLeft.Width /= 2.0
	logoOnTheLeft.Height /= 2.0

	// StageBranch will stage on the the first argument
	// all instances related to the second argument
	split.StageBranch(stager.splitStage, &split.View{
		Name: "Welcome to reqif",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "room for tree",
				ShowNameInHeader: false,
				Size:             100,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name: "Summary + upload + button",
							Size: 20,
							AsSplit: (&split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "Summary table",
										ShowNameInHeader: false,
										Size:             50,
										Table: &split.Table{
											StackName: stager.summaryTableStage.GetName(),
										},
									},
									{
										Name: "Upload Reqif File",
										Size: 25,
										AsSplit: (&split.AsSplit{
											Direction: split.Horizontal,
											AsSplitAreas: []*split.AsSplitArea{
												load.NewStager(r, stager.loadReqifStage, stager.splitStage).GetAsSplitArea()},
										}),
									},
									{
										Name:             "Buttons",
										ShowNameInHeader: false,
										Size:             25,
										Button: &split.Button{
											StackName: stager.welcomeTabButtonStage.GetName(),
										},
									},
								},
							}),
						},

						{
							Size: 40,
							Tree: &split.Tree{
								StackName: stager.dataTypeTreeStage.GetName(),
							},
						},
						{
							Size: 40,
							Tree: &split.Tree{
								StackName: stager.specTypesTreeStage.GetName(),
							},
						},
					},
				}),
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "REQIF Data",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "room for tree",
				ShowNameInHeader: false,
				Size:             100,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 33,
							Tree: &split.Tree{
								StackName: stager.specObjectsTreeStage.GetName(),
							},
						},
						{
							Size: 33,
							Tree: &split.Tree{
								StackName: stager.specRelationsTreeStage.GetName(),
							},
						},
						{
							Size: 34,
							Tree: &split.Tree{
								StackName: stager.specificationsTreeStage.GetName(),
							},
						},
						// {
						// 	Name:             "To be completed",
						// 	ShowNameInHeader: false,

						// 	Size: 15,
						// },
					},
				}),
			},
		},
	})

	asSplitAreaRenderingConfStage := load.NewStager(r, stager.loadRenderingConfStage, stager.splitStage).GetAsSplitArea()
	asSplitAreaRenderingConfStage.Size = 20

	split.StageBranch(stager.splitStage, &split.View{
		Name: "REQIF Render",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "room for tree",
				ShowNameInHeader: false,
				Size:             100,
				AsSplit: &split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 19.1,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Size: 70,
										Tree: &split.Tree{
											StackName: stager.specificationsTreeStage.GetName(),
										},
									},
									{
										Name: "Upload Reqif File",
										Size: 20,
										AsSplit: (&split.AsSplit{
											Direction: split.Horizontal,
											AsSplitAreas: []*split.AsSplitArea{
												load.NewStager(r, stager.loadReqifStage, stager.splitStage).GetAsSplitArea()},
										}),
									},
									{
										Size: 10,
										Button: &split.Button{
											StackName: stager.anonymousButtonStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size: 19.1,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Size: 70,
										Tree: &split.Tree{
											StackName: stager.specTypesTreeStage.GetName(),
										},
									},
									asSplitAreaRenderingConfStage,
									{
										Size: 10,
										Button: &split.Button{
											StackName: stager.renderingTabButtonStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size: 61.8,
							Markdown: &split.Markdown{
								StackName: stager.markdownStage.GetName(),
							},
						},
						// {
						// 	Name:             "To be completed",
						// 	ShowNameInHeader: false,

						// 	Size: 15,
						// },
					},
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) REQIF Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}),
			}),
		},
		IsSecondaryView: true,
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) summary table probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.summaryTableStage.GetProbeSplitStageName(),
				}),
			}),
		},
		IsSecondaryView: true,
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) data type tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.dataTypeTreeStage.GetProbeSplitStageName(),
				}),
			}),
		},
		IsSecondaryView: true,
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) spec type tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.specTypesTreeStage.GetProbeSplitStageName(),
				}),
			}),
		},
		IsSecondaryView: true,
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) markdownStage",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.markdownStage.GetProbeSplitStageName(),
				}),
			}),
		},
		IsSecondaryView: true,
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) load Reqif Stage",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.loadReqifStage.GetProbeSplitStageName(),
				}),
			}),
		},
		IsSecondaryView: true,
	})
	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) load Rendering Conf Stage",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.loadRenderingConfStage.GetProbeSplitStageName(),
				}),
			}),
		},
		IsSecondaryView: true,
	})

	stager.splitStage.Commit()

	if pathToReqifFile != "" {
		// Open the XML file
		reqifData, err := os.ReadFile(pathToReqifFile) // os.ReadFile is simpler
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", pathToReqifFile, err)
		}

		if strings.HasSuffix(pathToReqifFile, ".reqifz") {
			reqifData, svgImages, jpgImages, pngImages, err := extractReqifFromZip(reqifData)
			if err != nil {
				fmt.Printf("Error extracting reqif from zip %s: %v\n", pathToReqifFile, err)
			}
			stager.processReqifData(reqifData, svgImages, jpgImages, pngImages, pathToReqifFile)

			for _, svgImage := range svgImages {
				_ = svgImage
			}
		} else {
			stager.processReqifData(reqifData, []*EmbeddedSvgImage{}, []*EmbeddedJpgImage{}, []*EmbeddedPngImage{}, pathToReqifFile)
		}
	}

	if pathToRenderingConf != "" {
		// file
		renderingConf, err := os.ReadFile(pathToRenderingConf) // os.ReadFile is simpler
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", pathToRenderingConf, err)
		}

		stageForRenderingConf := NewStage("renderingConf")
		ParseAstFromBytes(stageForRenderingConf, renderingConf)

		stager.processRenderingConf(stageForRenderingConf)
	}

	stager.UpdateAndCommitLoadReqifStage()
	stager.updateAndCommitLoadRenderingConfStage()
	stager.UpdateAndCommitWelcomeTabButtonStage()
	stager.UpdateAndCommitRenderingTabButtonStage()
	stager.UpdateAndCommitAnonymousButtonStage()

	// hook the stage on a kill command
	stage.OnAfterKillCreateCallback = &OnAfterKillCreateCallback{}

	return
}

func (stager *Stager) processReqifData(reqifData []byte, svgImages []*EmbeddedSvgImage, jpgImages []*EmbeddedJpgImage, pngImages []*EmbeddedPngImage, pathToReqifFile string) {

	stager.stage.Reset()
	stager.PathToReqifFile = pathToReqifFile

	// Unmarshal the XML into the Reqif struct
	var req_if REQ_IF
	err := xml.Unmarshal(reqifData, &req_if)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	StageBranch(stager.stage, &req_if)

	for _, svgImage := range svgImages {
		svgImage.Stage(stager.stage)
	}
	for _, jpgImage := range jpgImages {
		jpgImage.Stage(stager.stage)
	}
	for _, pngImage := range pngImages {
		pngImage.Stage(stager.stage)
	}

	// fetch the root REQ_IF element and exit otherwise
	for reqif := range *GetGongstructInstancesSet[REQ_IF](stager.stage) {
		stager.rootReqif = reqif
	}

	if stager.rootReqif == nil {
		log.Fatal("No REQ_IF element found in file", pathToReqifFile)
	}

	stager.objectNamer.SetNamesToElements(stager.stage, &req_if)

	stager.enforceModelSemantic()

	stager.stage.Commit()

	stager.updateAndCommitSummaryTableStage()

	stager.dataTypesTreeUpdater.UpdateAndCommitDataTypeTreeStage(stager)
	stager.specTypesTreeUpdater.UpdateAndCommitSpecTypesTreeStage(stager)

	stager.specObjectsUX.UpdateAndCommitSpecObjectsTreeStage(stager)
	stager.specRelationsUX.UpdateAndCommitSpecRelationsTreeStage(stager)
	stager.specificationsUX.UpdateAndCommitSpecificationsTreeStage(stager)
	stager.specificationsUX.UpdateAndCommitSpecificationsMarkdownStage(stager)

	// stager.UpdateAndCommitButtonStage()

}

func (stager *Stager) GetSpecificationsTreeUpdater() (specificationsTreeUpdater SpecificationsTreeUpdaterInterface) {
	return stager.specificationsUX
}

func (stager *Stager) GetSpecTypesTreeUpdater() (specTypesTreeUpdater SpecTypesTreeUpdaterInterface) {
	return stager.specTypesTreeUpdater
}
