package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/fullstack-lang/gong/go/golang"
	"github.com/fullstack-lang/gong/go/golang/controllers"
	"github.com/fullstack-lang/gong/go/golang/diagrams"
	"github.com/fullstack-lang/gong/go/golang/fullstack"
	"github.com/fullstack-lang/gong/go/golang/models"
	"github.com/fullstack-lang/gong/go/golang/orm"
	"github.com/fullstack-lang/gong/go/golang/probe"
	"github.com/fullstack-lang/gong/go/golang/stack"
	"github.com/fullstack-lang/gong/go/golang/static"

	"github.com/fullstack-lang/gong/go/vscode"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

const COMPUTED_FROM_PKG_PATH string = "computed from pkgPath (path to package for analysis)"

var (
	pkgPath = flag.String("pkgPath", ".", "path to the models package."+
		"For instance, gongc go/models")
	skipSwagger       = flag.Bool("skipSwagger", true, "skip swagger file generation")
	skipNg            = flag.Bool("skipNg", false, "generates skipNg, skip ng operations")
	skipFlutter       = flag.Bool("skipFlutter", true, "do not generate flutter front")
	skipCoder         = flag.Bool("skipCoder", true, "do not generate coder file")
	skipSerialize     = flag.Bool("skipSerialize", false, "do not generate models/gong_serialize code for xl ouput")
	skipNpmWorkspaces = flag.Bool("skipNpmWorkspaces", false, "do not generate package.json at the root for npm workspaces")

	clean = flag.Bool("clean", false, "let gongc remove files & dir that are generated. The program then exits.")

	addr = flag.String("addr", "localhost:8080/api",
		"network address addr where the angular generated service will lookup the server")
	run               = flag.Bool("run", false, "run 'go run main.go' after compilation")
	skipGoModCommands = flag.Bool("skipGoModCommands", false, "avoid calls to go mod init, tidy and vendor")

	skipNpmInstall = flag.Bool("skipNpmInstall", false, "skip the npm install command")

	compileForDebug = flag.Bool("compileForDebug", false, "The go debugger can be slow to start (more than 60')."+
		"A workaround is to generate a go build with with '-N -l' options")
)

func main() {

	log.SetPrefix("gongc: ")
	log.SetFlags(0)
	flag.Parse()
	if len(flag.Args()) > 1 {
		log.Fatal("surplus arguments")
	}
	if len(flag.Args()) == 1 {
		*pkgPath = flag.Arg(0)
	}

	// remove gong generated files
	golang.RemoveGeneratedGongFilesButDocs(*pkgPath)

	// initiate model package
	modelStage := gong_models.NewStage("")
	modelPkg, _ := gong_models.LoadSource(modelStage, *pkgPath)

	// check wether the package name follows gong naming convention
	if strings.ContainsAny(modelPkg.Name, "-") {
		log.Panicln(modelPkg.Name + " is not OK for a gong package name because it contains a - (dash) " +
			"and it cannot be used for naming a typescript class (in the generated front end lib)")
	}

	//
	// compute paths to generation directory
	//
	{

		directory, err := filepath.Abs(filepath.Join(*pkgPath, ".."))
		if err != nil {
			log.Fatal("Problem with backend target path " + err.Error())
		}

		// check existance of path
		fileInfo, err := os.Stat(directory)
		if os.IsNotExist(err) {
			log.Panicf("Folder %s does not exist.", directory)
		}
		if fileInfo.Mode().Perm()&(1<<(uint(7))) == 0 {
			log.Panicf("Folder %s is not writtable", directory)
		}
		modelPkg.PathToGoSubDirectory = directory
		// log.Println("backend target path " + modelPkg.PathToGoSubDirectory)

		modelPkg.OrmPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "orm")
		os.RemoveAll(modelPkg.OrmPkgGenPath)
		modelPkg.ControllersPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "controllers")
		os.RemoveAll(modelPkg.ControllersPkgGenPath)
		modelPkg.FullstackPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "fullstack")
		os.RemoveAll(modelPkg.FullstackPkgGenPath)
		modelPkg.StackPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "stack")
		os.RemoveAll(modelPkg.StackPkgGenPath)
		modelPkg.StaticPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "static")
		os.RemoveAll(modelPkg.StaticPkgGenPath)
		modelPkg.ProbePkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "probe")
		os.RemoveAll(modelPkg.ProbePkgGenPath)
		{
			// compute the name of the ng workspace
			// Convert to lowercase
			ngWorkspaceName := strings.ToLower(modelPkg.PkgPath)

			// remove trailing "/ng/models"
			ngWorkspaceName = strings.ReplaceAll(ngWorkspaceName, "/go/models", "")

			// Replace disallowed characters with hyphens
			re := regexp.MustCompile(`[^a-z0-9\.-]+`)
			ngWorkspaceName = re.ReplaceAllString(ngWorkspaceName, "-")
			// Trim leading and trailing hyphens, handle leading dots
			ngWorkspaceName = strings.Trim(ngWorkspaceName, "-")
			ngWorkspaceName = "ng-" + ngWorkspaceName // Prefix with 'ng' if starts with a dot
			log.Println(ngWorkspaceName)
			modelPkg.NgWorkspaceName = ngWorkspaceName

		}
		{
			directory, err :=
				filepath.Abs(
					filepath.Join(*pkgPath,
						fmt.Sprintf("../../%s/projects/%s/src/lib", modelPkg.NgWorkspaceName, modelPkg.Name)))
			modelPkg.NgDataLibrarySourceCodeDirectory = directory
			if err != nil {
				log.Panic("Problem with frontend target path " + err.Error())
			}
		}

		{
			directory, err :=
				filepath.Abs(
					filepath.Join(*pkgPath,
						fmt.Sprintf("../../%s/projects/%sspecific/src/lib", modelPkg.NgWorkspaceName, modelPkg.Name)))
			modelPkg.NgSpecificLibrarySourceCodeDirectory = directory
			if err != nil {
				log.Panic("Problem with frontend target path " + err.Error())
			}
		}

		{
			directory, err :=
				filepath.Abs(
					filepath.Join(*pkgPath,
						fmt.Sprintf("../../%s/projects/%sdatamodel/src/lib", modelPkg.NgWorkspaceName, modelPkg.Name)))
			modelPkg.MaterialLibDatamodelTargetPath = directory
			if err != nil {
				log.Panic("Problem with frontend target path " + err.Error())
			}
		}

		if !*skipNg {
			// log.Println("Removing all content of " + modelPkg.NgDataLibrarySourceCodeDirectory)
			gong_models.RemoveContents(modelPkg.NgDataLibrarySourceCodeDirectory)
		}

		// idealy, one would like to regenerate the datamodel library at each gongc but
		// there is no "ng" command to remove an existing library from a workspace
		// if !*skipNg {
		// 	datamodelLibDir := filepath.Join(modelPkg.MaterialLibDatamodelTargetPath, "../..")
		// 	log.Println("Removing all content of " + datamodelLibDir)
		// 	gong_models.RemoveContents(datamodelLibDir)
		// 	err = os.RemoveAll(datamodelLibDir)
		// 	if err != nil {
		// 		log.Println("Unable de remove", datamodelLibDir)
		// 	}
		// }

		if *clean {
			os.Exit(0)
		}

	}

	// check wether one is in a git managed directory. If absent, use "git init"
	{
		gitStatusCommand := exec.Command("git", "status")

		var isGitActive bool

		// Execute the command
		if err := gitStatusCommand.Run(); err != nil {
			isGitActive = false
			// log.Println("not a git directory")
		} else {
			isGitActive = true
			// log.Println("a git directory")
		}

		if !isGitActive {
			// log.Printf("git is not active, hence gong is generating a git directory with git init command")

			// git init
			cmd := exec.Command("git", "init")
			cmd.Dir, _ = filepath.Abs(filepath.Join(*pkgPath, "../.."))

			// https://stackoverflow.com/questions/48253268/print-the-stdout-from-exec-command-in-real-time-in-go
			var stdBuffer bytes.Buffer
			mw := io.MultiWriter(os.Stdout, &stdBuffer)

			cmd.Stdout = mw
			cmd.Stderr = mw

			log.Println(cmd.String())
			log.Println(stdBuffer.String())

			// Execute the command
			if err := cmd.Run(); err != nil {
				log.Panic(err)
			}

			f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}

			defer f.Close()

			if _, err := f.WriteString("go/cmd/" + modelPkg.Name + "/" + modelPkg.Name + "\n"); err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
			if _, err := f.WriteString("__debug_bin*" + "\n"); err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
			if _, err := f.WriteString("*.exe" + "\n"); err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
			if _, err := f.WriteString(".DS_Store" + "\n"); err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
			if _, err := f.WriteString("node_modules" + "\n"); err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
		}
	}

	if !*skipNpmWorkspaces && !*skipNg && !*skipNpmInstall {
		// we need to use npm package (because of angular 17/esbuild)
		// check wether a package.json is present, otherwise generate it
		packageJsonFilePath := filepath.Join(*pkgPath, "../../package.json")
		_, errd := os.Stat(packageJsonFilePath)
		if os.IsNotExist(errd) {
			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				packageJsonFilePath,
				golang.NpmPackageJsonTemplate)
		}
	}

	if !*skipNpmWorkspaces && !*skipNg && !*skipNpmInstall {
		// we need to use npm package (because of angular 17/esbuild)
		// check wether a package.json is present, otherwise generate it
		packageJsonFilePath := filepath.Join(*pkgPath, "../../.gitignore")
		_, errd := os.Stat(packageJsonFilePath)
		if os.IsNotExist(errd) {
			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				packageJsonFilePath,
				golang.GitIgnoreTempl)
		}
	}

	// generate diagrams/docs.go if absent
	{
		diagramsDocFilePath := filepath.Join(*pkgPath, "../diagrams/docs.go")
		_, errd := os.Stat(diagramsDocFilePath)
		if os.IsNotExist(errd) {
			log.Printf("../diagrams/docs.go does not exist, gongc creates a default one")

			diagramsDocFileDirPath := filepath.Dir(diagramsDocFilePath)
			diagramsDocFileDirAbsPath, _ := filepath.Abs(diagramsDocFileDirPath)

			errd := os.MkdirAll(diagramsDocFileDirAbsPath, os.ModePerm)
			if os.IsNotExist(errd) {
				log.Println("creating directory : " + diagramsDocFileDirAbsPath)
			}
			if os.IsExist(errd) {
				log.Println("directory " + diagramsDocFileDirAbsPath + " allready exists")
			}
			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				diagramsDocFilePath,
				diagrams.DiagramsDocFile)
		}
	}

	// generate main.go if absent
	{
		// check existance of "main.go" file and generate a default "main.go" if absent
		mainFilePath := filepath.Join(*pkgPath, fmt.Sprintf("../cmd/%s/main.go", gong_models.ComputePkgNameFromPkgPath(*pkgPath)))

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

			// sometimes on windows, directory creation is not completed before creation of file/directory (this
			// leads to non reproductible "access denied")
			time.Sleep(1000 * time.Millisecond)
			golang.CodeGeneratorPackageMain(
				modelPkg,
				modelPkg.Name,
				modelPkg.PkgPath,
				mainFilePath,
				*skipNg)
		}
	}

	// check existance of .vscode directory. If absent, generates default vscode configurations
	// that are usefull for development
	{
		vscodeDirFilePath := filepath.Join(*pkgPath, "../../.vscode")

		_, errd := os.Stat(vscodeDirFilePath)
		if os.IsNotExist(errd) {
			log.Printf(".vscode directory does not exist, gongc creates a default .vscode directory and the debug and launch configs")

			errd := os.MkdirAll(vscodeDirFilePath, os.ModePerm)
			if os.IsNotExist(errd) {
				log.Println("creating directory : " + vscodeDirFilePath)
			}
			if os.IsExist(errd) {
				log.Fatal("directory " + vscodeDirFilePath + " should not allready exists")
			}

			// sometimes on windows, directory creation is not completed before creation of file/directory (this
			// leads to non reproductible "access denied")
			time.Sleep(1000 * time.Millisecond)
			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				filepath.Join(vscodeDirFilePath, "launch.json"),
				vscode.VsCodeLaunchConfig)

			gong_models.VerySimpleCodeGenerator(
				modelPkg,
				filepath.Join(vscodeDirFilePath, "tasks.json"),
				vscode.VsCodeTasksConfig)
		}

	}

	// generate directory for orm package
	errd := os.MkdirAll(modelPkg.OrmPkgGenPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + modelPkg.OrmPkgGenPath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + modelPkg.OrmPkgGenPath + " allready exists")

		// supppress all files in it
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

	// compute source path
	sourcePath, errd2 := filepath.Abs(*pkgPath)
	if errd2 != nil {
		log.Panic("Problem with source path " + errd2.Error())
	}

	caserEnglish := cases.Title(language.English)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../../go/embed.go"),
		golang.EmebedGoDirTemplate)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(*pkgPath, "../fullstack/new_stack_instance.go"),
		fullstack.FullstackNewStackInstanceTemplate,
		fullstack.ModelGongNewStackInstanceStructSubTemplateCode)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../stack/stack.go"),
		stack.StackInstanceTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../static/serve_static_files.go"),
		static.ServeStaticFilesTemplate)

	models.CodeGeneratorModelGong(
		modelPkg,
		modelPkg.Name,
		*pkgPath)

	models.CodeGeneratorModelGongEnum(
		modelPkg,
		modelPkg.Name,
		*pkgPath)

	models.CodeGeneratorModelGongMarshall(
		modelPkg,
		modelPkg.Name,
		*pkgPath)

	models.CodeGeneratorModelGongGraph(
		modelPkg,
		modelPkg.Name,
		*pkgPath)

	models.CodeGeneratorModelGongSlice(
		modelPkg,
		modelPkg.Name,
		*pkgPath,
		modelPkg.PkgPath,
	)

	if !*skipCoder {
		models.CodeGeneratorModelGongCoder(
			modelPkg,
			modelPkg.Name,
			*pkgPath)
	}

	models.GongAstGenerator(modelPkg, *pkgPath)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../orm/back_repo.go"),
		orm.BackRepoTemplateCode, orm.BackRepoSubTemplate)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../orm/get_instance_db_from_instance.go"),
		orm.GetInstanceDBFromInstanceTemplateCode, orm.GetInstanceDBFromInstanceSubTemplate)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../orm/back_repo_data.go"),
		orm.BackRepoDataTemplateCode, orm.BackRepoDataSubTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../orm/int_slice.go"),
		orm.IntSliceTemplateCode)

	// for the replacement of the of the first bar in the Gongstruct Type def
	orm.ReplaceInFile("../orm/get_instance_db_from_instance.go", "	 | ", "	")

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../controllers/register_controllers.go"),
		controllers.ControllersRegisterTemplate, controllers.ControllersRegistrationsSubTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../controllers/controller.go"),
		controllers.ControllerTemplate)

	gong_models.SimpleCodeGeneratorForGongStructWithNameField(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../models/gong_callbacks.go"),
		models.ModelGongCallbacksFileTemplate, models.ModelGongCallbacksStructSubTemplateCode)

	gong_models.CodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(*pkgPath, "../models/gong_orchestrator.go"),
		models.ModelGongOrchestratorFileTemplate,
		models.ModelGongOrchestratorStructSubTemplateCode,
		map[string]string{}, map[string]string{},
		true,
		true)

	if !*skipSerialize {
		gong_models.SimpleCodeGeneratorForGongStructWithNameField(
			modelPkg,
			modelPkg.Name,
			modelPkg.PkgPath,
			filepath.Join(*pkgPath, "../models/gong_serialize.go"),
			models.ModelGongSerializeFileTemplate, models.ModelGongSerializeStructSubTemplateCode)
	}

	models.CodeGeneratorModelGongWop(modelPkg, modelPkg.Name, *pkgPath)

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
		filepath.Join(*pkgPath, "../docs.go"),
		golang.RootFileDocsTemplate,
		map[string]string{})

	probe.CodeGeneratorFillUpForm(
		modelPkg,
		modelPkg.Name,
		*pkgPath,
		modelPkg.PkgPath,
	)

	orm.CodeGeneratorGetReverseFieldOwnerName(
		modelPkg,
		modelPkg.Name,
		*pkgPath,
		modelPkg.PkgPath,
	)

	probe.CodeGeneratorModelFormCallback(
		modelPkg,
		modelPkg.Name,
		*pkgPath,
		modelPkg.PkgPath,
	)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(*pkgPath, "../probe/tree_node_impl_gongstruct.go"),
		probe.TreeNodeImplGongstructFileTemplate, probe.TreeNodeImplGongstructSubTemplateCode)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../probe/button_impl_gongstruct.go"),
		probe.ButtonImplGongstructFileTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../probe/button_impl_refresh.go"),
		probe.ButtonImplRefreshFileTemplate)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(*pkgPath, "../probe/cell_delete_icon_impl.go"),
		probe.CellDeleteIconImplTemplate, probe.CellDeleteIconImplSubTemplateCode)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(*pkgPath, "../probe/fill_up_form_from_gongstruct.go"),
		probe.FillUpFormFromGongstructTemplate, probe.FillUpFormFromGongstructSubTemplateCode)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(*pkgPath, "../probe/fill_up_form_from_gongstruct_name.go"),
		probe.FillUpFormFromGongstructNameTemplate, probe.FillUpFormFromGongstructNameSubTemplateCode)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(*pkgPath, "../probe/fill_up_table.go"),
		probe.FillUpTableTemplate, probe.FillUpTableSubTemplateCode)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../probe/probe.go"),
		probe.ProbeTemplate)

	gong_models.SimpleCodeGenerator(
		modelPkg,
		caserEnglish.String(modelPkg.Name),
		modelPkg.PkgPath, filepath.Join(*pkgPath, "../probe/fill_up_tree.go"),
		probe.FillUpTree, probe.FillUpTreeSubTemplateCode)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../probe/basic_field_to_form.go"),
		probe.BasicFieldtoFormTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../probe/association_to_form.go"),
		probe.AssociationFieldToFormTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../probe/association_slice_to_form.go"),
		probe.AssociationSliceToFormTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../probe/enum_type_to_form.go"),
		probe.EnumTypeStringToForm)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../probe/assoc_sorting_button_impl.go"),
		probe.NewOnSortingEditonTemplate)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		filepath.Join(*pkgPath, "../probe/form_div_field.go"),
		probe.FormDivToFieldTemplate)

	// since go mod vendor brings angular dependencies into the vendor directory
	// the go mod vendor command has to be issued before the ng build command
	if !*skipNg {
		genAngular(modelPkg, *skipNpmInstall, *skipGoModCommands)
	}

	if !*skipFlutter {
		genFlutter(modelPkg)
	}

	apiYamlFilePath := fmt.Sprintf("%s/%sapi.yml", modelPkg.ControllersPkgGenPath, modelPkg.Name)
	if !*skipSwagger {

		// generate open api specification with swagger
		cmd := exec.Command("swagger",
			"generate", "spec",
			"-w", filepath.Dir(sourcePath), // swagger is interested in the "docs.go" in the package
			// by convention, this file is located on the parent dir
			"-o", apiYamlFilePath)
		log.Printf("Running swagger command and waiting for it to finish...\n")

		// https://stackoverflow.com/questions/48253268/print-the-stdout-from-exec-command-in-real-time-in-go
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		if err := cmd.Run(); err != nil {
			log.Printf("problem with swagger : %s", err.Error())
		}
	}

	// if *skipNg {
	// 	return
	// }

	// go build
	if true {
		start := time.Now()
		var cmd *exec.Cmd
		if !*compileForDebug {
			cmd = exec.Command("go", "build")
		} else {
			// gcFlags allows for speedup of delve
			cmd = exec.Command("go", "build", "-gcflags", "-N -l")
		}
		cmd.Dir, _ = filepath.Abs(filepath.Join(*pkgPath, fmt.Sprintf("../cmd/%s", gong_models.ComputePkgNameFromPkgPath(*pkgPath))))
		log.Printf("Running %s command in directory %s and waiting for it to finish...\n", cmd.Args, cmd.Dir)

		// https://stackoverflow.com/questions/48253268/print-the-stdout-from-exec-command-in-real-time-in-go
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("go build over and took %s", time.Since(start))
	}

	// run application
	if *run {
		cmd := exec.Command("go", "run", "main.go")
		cmd.Dir, _ = filepath.Abs(filepath.Join(*pkgPath,
			fmt.Sprintf("../../go/cmd/%s", gong_models.ComputePkgNameFromPkgPath(*pkgPath))))
		log.Printf("Running %s command in directory %s and waiting for it to finish...\n", cmd.Args, cmd.Dir)

		// https://stackoverflow.com/questions/48253268/print-the-stdout-from-exec-command-in-real-time-in-go
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
	}
}
