package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/fullstack-lang/gong/go/angular"
	"github.com/fullstack-lang/gong/go/flutter"
	"github.com/fullstack-lang/gong/go/golang"
	"github.com/fullstack-lang/gong/go/golang/diagrams"
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/vscode"
	"github.com/spf13/cobra"
)

var (
	skipSwagger       bool
	skipNg            bool
	skipFlutter       bool
	skipCoder         bool
	skipSerialize     bool
	skipNpmWorkspaces bool
	skipStager        bool
	level1            bool
	addr              string
	run               bool
	skipGoModCommands bool
	skipNpmInstall    bool
	compileForDebug   bool
	dbLite            bool
	debouncedMarshall bool
)

var generateCmd = &cobra.Command{
	Use:     "generate [path to models package]",
	Short:   "Generates the code from the models package",
	Aliases: []string{"g"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.SetPrefix("gongc: ")
		log.SetFlags(0)

		if level1 {
			skipNg = true
		}

		var pkgPath string
		if len(args) > 0 {
			pkgPath = args[0]
		} else {
			pkgPath = "."
		}

		// remove gong generated files
		golang.RemoveGeneratedGongFilesButDocs(pkgPath)

		// initiate model package
		modelStage := gong_models.NewStage("")
		modelPkg, _ := gong_models.LoadSource(modelStage, pkgPath)

		// check wether the package name follows gong naming convention
		if strings.ContainsAny(modelPkg.Name, "-") {
			log.Panicln(modelPkg.Name + " is not OK for a gong package name because it contains a - (dash) " +
				"and it cannot be used for naming a typescript class (in the generated front end lib)")
		}

		//
		// compute paths to generation directory
		//
		{

			directory, err := filepath.Abs(filepath.Join(pkgPath, ".."))
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
			modelPkg.DbOrmPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "orm/dbgorm")
			os.RemoveAll(modelPkg.DbOrmPkgGenPath)
			modelPkg.DbLiteOrmPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "orm/dblite")
			os.RemoveAll(modelPkg.DbLiteOrmPkgGenPath)
			modelPkg.DbPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "db")
			os.RemoveAll(modelPkg.DbPkgGenPath)
			modelPkg.ControllersPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "controllers")
			os.RemoveAll(modelPkg.ControllersPkgGenPath)
			modelPkg.FullstackPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "fullstack")
			os.RemoveAll(modelPkg.FullstackPkgGenPath)
			modelPkg.StackPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "stack")
			os.RemoveAll(modelPkg.StackPkgGenPath)
			modelPkg.Level1StackPkgGenPath = filepath.Join(modelPkg.PathToGoSubDirectory, "level1stack")
			os.RemoveAll(modelPkg.Level1StackPkgGenPath)
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
						filepath.Join(pkgPath,
							fmt.Sprintf("../../%s/projects/%s/src/lib", modelPkg.NgWorkspaceName, modelPkg.Name)))
				modelPkg.NgDataLibrarySourceCodeDirectory = directory
				if err != nil {
					log.Panic("Problem with frontend target path " + err.Error())
				}
			}

			{
				directory, err :=
					filepath.Abs(
						filepath.Join(pkgPath,
							fmt.Sprintf("../../%s/projects/%sspecific/src/lib", modelPkg.NgWorkspaceName, modelPkg.Name)))
				modelPkg.NgSpecificLibrarySourceCodeDirectory = directory
				if err != nil {
					log.Panic("Problem with frontend target path " + err.Error())
				}
			}

			{
				directory, err :=
					filepath.Abs(
						filepath.Join(pkgPath,
							fmt.Sprintf("../../%s/projects/%sdatamodel/src/lib", modelPkg.NgWorkspaceName, modelPkg.Name)))
				modelPkg.MaterialLibDatamodelTargetPath = directory
				if err != nil {
					log.Panic("Problem with frontend target path " + err.Error())
				}
			}

			if !skipNg {
				// log.Println("Removing all content of " + modelPkg.NgDataLibrarySourceCodeDirectory)
				gong_models.RemoveContents(modelPkg.NgDataLibrarySourceCodeDirectory)
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
				cmd.Dir, _ = filepath.Abs(filepath.Join(pkgPath, "../.."))

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

		if !skipNpmWorkspaces && !skipNg && !skipNpmInstall {
			// we need to use npm package (because of angular 17/esbuild)
			// check wether a package.json is present, otherwise generate it
			packageJsonFilePath := filepath.Join(pkgPath, "../../package.json")
			_, errd := os.Stat(packageJsonFilePath)
			if os.IsNotExist(errd) {
				gong_models.VerySimpleCodeGenerator(
					modelPkg,
					packageJsonFilePath,
					golang.NpmPackageJsonTemplate)
			}
		}

		if !skipNpmWorkspaces && !skipNg && !skipNpmInstall {
			// check wether a package.json is present, otherwise generate it
			packageJsonFilePath := filepath.Join(pkgPath, "../../.gitignore")
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
			diagramsDocFilePath := filepath.Join(pkgPath, "../diagrams/docs.go")
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

		// check existance of .vscode directory. If absent, generates default vscode configurations
		// that are usefull for development
		{
			vscodeDirFilePath := filepath.Join(pkgPath, "../../.vscode")

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

		golang.GeneratesGoCode(modelPkg, pkgPath, skipCoder, dbLite, skipSerialize, skipStager, level1, debouncedMarshall)

		// since go mod vendor brings angular dependencies into the vendor directory
		// the go mod vendor command has to be issued before the ng build command
		if !skipNg && !level1 {
			angular.GeneratesAngularCode(modelPkg, pkgPath, skipNpmInstall, skipGoModCommands, addr)
		}

		if !skipFlutter {
			flutter.GenFlutter(modelPkg, pkgPath)
		}

		apiYamlFilePath := fmt.Sprintf("%s/%sapi.yml", modelPkg.ControllersPkgGenPath, modelPkg.Name)
		if !skipSwagger {

			// compute source path
			sourcePath, errd2 := filepath.Abs(pkgPath)
			if errd2 != nil {
				log.Panic("Problem with source path " + errd2.Error())
			}

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

		// go mod tidy
		if !skipGoModCommands {
			start := time.Now()

			cmd := exec.Command("go", "mod", "tidy")

			cmd.Dir, _ = filepath.Abs(filepath.Join(pkgPath, fmt.Sprintf("../cmd/%s", gong_models.ComputePkgNameFromPkgPath(pkgPath))))
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

		// go build
		if true {
			start := time.Now()
			var cmd *exec.Cmd
			if !compileForDebug {
				cmd = exec.Command("go", "build")
			} else {
				// gcFlags allows for speedup of delve
				cmd = exec.Command("go", "build", "-gcflags", "-N -l")
			}
			cmd.Dir, _ = filepath.Abs(filepath.Join(pkgPath, fmt.Sprintf("../cmd/%s", gong_models.ComputePkgNameFromPkgPath(pkgPath))))
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
		if run {
			cmd := exec.Command("go", "run", "main.go")
			cmd.Dir, _ = filepath.Abs(filepath.Join(pkgPath,
				fmt.Sprintf("../../go/cmd/%s", gong_models.ComputePkgNameFromPkgPath(pkgPath))))
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
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolVar(&skipSwagger, "skipSwagger", true, "skip swagger file generation")
	generateCmd.Flags().BoolVar(&skipNg, "skipNg", false, "generates skipNg, skip ng operations")
	generateCmd.Flags().BoolVar(&skipFlutter, "skipFlutter", true, "do not generate flutter front")
	generateCmd.Flags().BoolVar(&skipCoder, "skipCoder", true, "do not generate coder file")
	generateCmd.Flags().BoolVar(&skipSerialize, "skipSerialize", false, "do not generate models/gong_serialize code for xl ouput")
	generateCmd.Flags().BoolVar(&skipNpmWorkspaces, "skipNpmWorkspaces", false, "do not generate package.json at the root for npm workspaces")
	generateCmd.Flags().BoolVar(&skipStager, "skipStager", true, "do not generate blolerplate stager.go in models and in main.go")
	generateCmd.Flags().BoolVar(&level1, "level1", false, "generates a level1 gong application (implies -skipNg)")
	generateCmd.Flags().StringVar(&addr, "addr", "localhost:8080/api", "network address addr where the angular generated service will lookup the server")
	generateCmd.Flags().BoolVar(&run, "run", false, "run 'go run main.go' after compilation")
	generateCmd.Flags().BoolVar(&skipGoModCommands, "skipGoModCommands", false, "avoid calls to go mod init, tidy and vendor")
	generateCmd.Flags().BoolVar(&skipNpmInstall, "skipNpmInstall", false, "skip the npm install command")
	generateCmd.Flags().BoolVar(&compileForDebug, "compileForDebug", false, "The go debugger can be slow to start (more than 60'). A workaround is to generate a go build with with '-N -l' options")
	generateCmd.Flags().BoolVar(&dbLite, "dbLite", true, "If true, the database in all stack instances are purely in memory. If false, it is sqlite and can be persisted to a sqlite file")
	generateCmd.Flags().BoolVar(&debouncedMarshall, "debouncedMarshall", false, "If false, each marshall call is blocking. If true, the marshalling is non blocking and debounced")
}
