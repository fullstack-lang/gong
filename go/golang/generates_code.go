package golang

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/fullstack-lang/gong/go/golang/cmd"
	"github.com/fullstack-lang/gong/go/golang/controllers"
	"github.com/fullstack-lang/gong/go/golang/db"
	"github.com/fullstack-lang/gong/go/golang/fullstack"
	"github.com/fullstack-lang/gong/go/golang/models"
	"github.com/fullstack-lang/gong/go/golang/orm"
	"github.com/fullstack-lang/gong/go/golang/orm/dbgorm"
	"github.com/fullstack-lang/gong/go/golang/probe"
	"github.com/fullstack-lang/gong/go/golang/stack"
	"github.com/fullstack-lang/gong/go/golang/static"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

func GeneratesGoCode(modelPkg *gong_models.ModelPkg,
	pkgPath string, skipCoder bool, dbLite bool, skipSerialize bool, skipStager bool) {

	// generate main.go if absent
	{
		// check existance of "main.go" file and generate a default "main.go" if absent
		mainFilePath := filepath.Join(pkgPath, fmt.Sprintf("../cmd/%s/main.go", gong_models.ComputePkgNameFromPkgPath(pkgPath)))

		_, errd := os.Stat(mainFilePath)
		if os.IsNotExist(errd) {
			log.Printf("maing.go does not exist, gongc creates a default main.go")

			mainFileDirPath := filepath.Dir(mainFilePath)
			mainFileDirAbsPath, _ := filepath.Abs(mainFileDirPath)

			errd := os.MkdirAll(mainFileDirAbsPath, os.ModePerm)
			if os.IsNotExist(errd) {
				log.Println("creating directory : " + mainFileDirAbsPath)
			}
			if os.IsExist(errd) {
				log.Println("directory " + mainFileDirAbsPath + " allready exists")
			}

			// creates a "data" directory as well to store the stage.go files
			dataDirPath := filepath.Join(mainFileDirAbsPath, "data")
			errd = os.MkdirAll(dataDirPath, os.ModePerm)
			if os.IsNotExist(errd) {
				log.Println("creating directory : " + dataDirPath)
			}
			if os.IsExist(errd) {
				log.Println("directory " + dataDirPath + " allready exists")
			}

			// sometimes on windows, directory creation is not completed before creation of file/directory (this
			// leads to non reproductible "access denied")
			time.Sleep(1000 * time.Millisecond)
			cmd.CodeGeneratorPackageMain(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				mainFilePath,
				skipStager)
		}
	}

	// stager.go
	{
		// check existance of "main.go" file and generate a default "main.go" if absent
		coderFilePath := filepath.Join(pkgPath, "stager.go")

		_, errd := os.Stat(coderFilePath)
		if os.IsNotExist(errd) && !skipStager {
			log.Printf("stager.go does not exist, gongc creates a default stager.go")
			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				coderFilePath,
				models.StagerFileTemplate)
		}

	}

	errd := os.MkdirAll(modelPkg.OrmPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.OrmPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.OrmPkgGenPath + " allready exists")
	}

	// generate directory for orm package
	errd = os.MkdirAll(modelPkg.DbOrmPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.DbOrmPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.DbOrmPkgGenPath + " allready exists")
	}

	// generate directory for orm package
	errd = os.MkdirAll(modelPkg.DbLiteOrmPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.DbLiteOrmPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.DbLiteOrmPkgGenPath + " allready exists")
	}

	// generate directory for orm package
	errd = os.MkdirAll(modelPkg.DbPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.DbPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.DbPkgGenPath + " allready exists")
	}

	// generate directory for controllers package
	errd = os.MkdirAll(modelPkg.ControllersPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.ControllersPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.ControllersPkgGenPath + " allready exists")
	}

	// generate directory for fullstack package
	errd = os.MkdirAll(modelPkg.FullstackPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.FullstackPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.FullstackPkgGenPath + " allready exists")
	}

	errd = os.MkdirAll(modelPkg.StackPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.StackPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.StackPkgGenPath + " allready exists")
	}

	// generate directory for static package
	errd = os.MkdirAll(modelPkg.StaticPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.StaticPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.StaticPkgGenPath + " allready exists")
	}

	// generate directory for Data package
	errd = os.MkdirAll(modelPkg.ProbePkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.ProbePkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.ProbePkgGenPath + " allready exists")
	}

	caserEnglish := cases.Title(language.English)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../../go/embed.go"),
		EmebedGoDirTemplate)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(pkgPath, "../fullstack/new_stack_instance.go"),
		fullstack.FullstackNewStackInstanceTemplate,
		fullstack.ModelGongNewStackInstanceStructSubTemplateCode)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../stack/stack.go"),
		stack.StackInstanceTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../static/serve_static_files.go"),
		static.ServeStaticFilesTemplate)

	models.CodeGeneratorModelGong(
		modelPkg,
		modelPkg.Name,
		pkgPath)

	models.CodeGeneratorModelGongEnum(
		modelPkg,
		modelPkg.Name,
		pkgPath)

	models.CodeGeneratorModelGongMarshall(
		modelPkg,
		modelPkg.Name,
		pkgPath)

	models.CodeGeneratorModelGongGraph(
		modelPkg,
		modelPkg.Name,
		pkgPath)

	models.CodeGeneratorModelGongSlice(
		modelPkg,
		modelPkg.Name,
		pkgPath,
		modelPkg.PkgPath,
	)

	if !skipCoder {
		models.CodeGeneratorModelGongCoder(
			modelPkg,
			modelPkg.Name,
			pkgPath)
	}

	models.GongAstGenerator(modelPkg, pkgPath)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(pkgPath, "../orm/back_repo.go"),
		orm.BackRepoTemplateCode, orm.BackRepoSubTemplate)

	// back repo is either with gorm + sqlite or with lite
	if dbLite {
		orm.RemoveTargetedLines(filepath.Join(pkgPath, "../orm/back_repo.go"), orm.Lite)
	} else {
		orm.RemoveTargetedLines(filepath.Join(pkgPath, "../orm/back_repo.go"), orm.Gorm)
	}

	gong_models.SimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(pkgPath, "../orm/get_instance_db_from_instance.go"),
		orm.GetInstanceDBFromInstanceTemplateCode, orm.GetInstanceDBFromInstanceSubTemplate)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(pkgPath, "../orm/back_repo_data.go"),
		orm.BackRepoDataTemplateCode, orm.BackRepoDataSubTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../db/db_interface.go"),
		db.DbInterfaceTmpl)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../orm/dbgorm/db.go"),
		dbgorm.DbTmpl)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(pkgPath, "../orm/db.go"),
		orm.DbTmpl, orm.DBliteSubTemplates)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../orm/int_slice.go"),
		orm.IntSliceTemplateCode)

	// for the replacement of the of the first bar in the Gongstruct Type def
	orm.ReplaceInFile("../orm/get_instance_db_from_instance.go", "	 | ", "	")

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(pkgPath, "../controllers/register_controllers.go"),
		controllers.ControllersRegisterTemplate, controllers.ControllersRegistrationsSubTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../controllers/controller.go"),
		controllers.ControllerTemplate)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(pkgPath, "../models/gong_callbacks.go"),
		models.ModelGongCallbacksFileTemplate, models.ModelGongCallbacksStructSubTemplateCode)

	gong_models.CodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(pkgPath, "../models/gong_orchestrator.go"),
		models.ModelGongOrchestratorFileTemplate,
		models.ModelGongOrchestratorStructSubTemplateCode,
		map[string]string{}, map[string]string{},
		true,
		true)

	if !skipSerialize {
		gong_models.SimpleCodeGeneratorForGongStructWithNameField(
			modelPkg,
			modelPkg.Name,
			modelPkg.PkgPath,
			filepath.Join(pkgPath, "../models/gong_serialize.go"),
			models.ModelGongSerializeFileTemplate, models.ModelGongSerializeStructSubTemplateCode)
	}

	models.CodeGeneratorModelGongWop(modelPkg, modelPkg.Name, pkgPath)

	orm.MultiCodeGeneratorBackRepo(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		modelPkg.OrmPkgGenPath)

	controllers.MultiCodeGeneratorControllers(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		modelPkg.ControllersPkgGenPath)

	gong_models.SimpleCodeGenerator(modelPkg,
		modelPkg.Name, modelPkg.PkgPath,
		filepath.Join(pkgPath, "../docs.go"),
		RootFileDocsTemplate,
		map[string]string{})

	probe.CodeGeneratorFillUpForm(
		modelPkg,
		modelPkg.Name,
		pkgPath,
		modelPkg.PkgPath,
	)

	models.CodeGeneratorGongReverse(
		modelPkg,
		modelPkg.Name,
		pkgPath,
		modelPkg.PkgPath,
	)

	probe.CodeGeneratorModelFormCallback(
		modelPkg,
		modelPkg.Name,
		pkgPath,
		modelPkg.PkgPath,
	)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(pkgPath, "../probe/tree_node_impl_gongstruct.go"),
		probe.TreeNodeImplGongstructFileTemplate, probe.TreeNodeImplGongstructSubTemplateCode)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../probe/button_impl_gongstruct.go"),
		probe.ButtonImplGongstructFileTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../probe/button_impl_refresh.go"),
		probe.ButtonImplRefreshFileTemplate)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(pkgPath, "../probe/cell_delete_icon_impl.go"),
		probe.CellDeleteIconImplTemplate, probe.CellDeleteIconImplSubTemplateCode)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(pkgPath, "../probe/fill_up_form_from_gongstruct.go"),
		probe.FillUpFormFromGongstructTemplate, probe.FillUpFormFromGongstructSubTemplateCode)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(pkgPath, "../probe/fill_up_form_from_gongstruct_name.go"),
		probe.FillUpFormFromGongstructNameTemplate, probe.FillUpFormFromGongstructNameSubTemplateCode)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(pkgPath, "../probe/fill_up_table.go"),
		probe.FillUpTableTemplate, probe.FillUpTableSubTemplateCode)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../probe/probe.go"),
		probe.ProbeTemplate)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(pkgPath, "../probe/fill_up_tree.go"),
		probe.FillUpTree, probe.FillUpTreeSubTemplateCode)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../probe/basic_field_to_form.go"),
		probe.BasicFieldtoFormTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../probe/association_to_form.go"),
		probe.AssociationFieldToFormTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../probe/association_slice_to_form.go"),
		probe.AssociationSliceToFormTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../probe/enum_type_to_form.go"),
		probe.EnumTypeStringToForm)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../probe/assoc_sorting_button_impl.go"),
		probe.NewOnSortingEditonTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(pkgPath, "../probe/form_div_field.go"),
		probe.FormDivToFieldTemplate)
}
