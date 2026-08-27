package golang

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Regex to identify Gong instance identifiers: __<StructName>__<Digits>_<optional_suffix>
var gongIdentRegex = regexp.MustCompile(`^__([A-Za-z0-9_]+)__(\d+)_(.*)$`)

// ParsedStageFile holds the parsed information of a stage Go file.
type ParsedStageFile struct {
	FilePath    string
	PackageName string
	Imports     []*ImportEntry
	FuncDecl    *ast.FuncDecl
	HasCommit   bool
	FileSet     *token.FileSet
	AstFile     *ast.File
}

// ImportEntry represents an import declaration.
type ImportEntry struct {
	Alias string // alias if any, e.g. "split"
	Path  string // quoted import path, e.g. "\"github.com/...\""
}

// MergeStageFiles merges multiple Gong stage Go source files into a single stage file.
// It recomputes instance IDs to prevent collisions across any data model.
func MergeStageFiles(inputPaths []string, outputPath string, pkgOverride string) error {
	if len(inputPaths) == 0 {
		return errors.New("no input stage files provided")
	}

	mergedSource, err := MergeStageFilesToString(inputPaths, pkgOverride)
	if err != nil {
		return err
	}

	if outputPath == "" {
		outputPath = "merged.go"
	}

	// Ensure directory exists
	dir := filepath.Dir(outputPath)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	if err := os.WriteFile(outputPath, []byte(mergedSource), 0o644); err != nil {
		return fmt.Errorf("failed to write merged stage file to %s: %w", outputPath, err)
	}

	return nil
}

// MergeStageFilesToString merges multiple Gong stage Go source files into a formatted Go source string.
func MergeStageFilesToString(inputPaths []string, pkgOverride string) (string, error) {
	if len(inputPaths) == 0 {
		return "", errors.New("no input stage files provided")
	}

	parsedFiles := make([]*ParsedStageFile, 0, len(inputPaths))
	for _, path := range inputPaths {
		parsed, err := parseStageFile(path)
		if err != nil {
			return "", fmt.Errorf("error parsing %s: %w", path, err)
		}
		parsedFiles = append(parsedFiles, parsed)
	}

	return mergeParsedStageFiles(parsedFiles, pkgOverride)
}

// MergeStageBytesToString merges multiple in-memory stage source code blobs.
func MergeStageBytesToString(sources [][]byte, pkgOverride string) (string, error) {
	if len(sources) == 0 {
		return "", errors.New("no input sources provided")
	}

	parsedFiles := make([]*ParsedStageFile, 0, len(sources))
	for i, src := range sources {
		fset := token.NewFileSet()
		inFile, err := parser.ParseFile(fset, fmt.Sprintf("input_%d.go", i), src, parser.ParseComments)
		if err != nil {
			return "", fmt.Errorf("error parsing input %d: %w", i, err)
		}
		parsed, err := extractParsedStageFile(inFile, fset, fmt.Sprintf("input_%d.go", i))
		if err != nil {
			return "", fmt.Errorf("error in input %d: %w", i, err)
		}
		parsedFiles = append(parsedFiles, parsed)
	}

	return mergeParsedStageFiles(parsedFiles, pkgOverride)
}

func parseStageFile(path string) (*ParsedStageFile, error) {
	fset := token.NewFileSet()
	inFile, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file: %w", err)
	}

	return extractParsedStageFile(inFile, fset, path)
}

func extractParsedStageFile(inFile *ast.File, fset *token.FileSet, path string) (*ParsedStageFile, error) {
	parsed := &ParsedStageFile{
		FilePath:    path,
		PackageName: inFile.Name.Name,
		FileSet:     fset,
		AstFile:     inFile,
	}

	// Extract imports
	for _, imp := range inFile.Imports {
		entry := &ImportEntry{
			Path: imp.Path.Value,
		}
		if imp.Name != nil {
			entry.Alias = imp.Name.Name
		}
		parsed.Imports = append(parsed.Imports, entry)
	}

	// Find staging function: func _(stage *...Stage) or any func with stage param
	for _, decl := range inFile.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if isStagingFunction(fn) {
				parsed.FuncDecl = fn
				break
			}
		}
	}

	if parsed.FuncDecl == nil {
		return nil, fmt.Errorf("file %s does not contain a stage initialization function", path)
	}

	// Check if file contains stage.Commit()
	ast.Inspect(parsed.FuncDecl.Body, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
				if sel.Sel.Name == "Commit" {
					if id, ok := sel.X.(*ast.Ident); ok && id.Name == "stage" {
						parsed.HasCommit = true
						return false
					}
				}
			}
		}
		return true
	})

	return parsed, nil
}

func isStagingFunction(fn *ast.FuncDecl) bool {
	if fn.Type == nil || fn.Type.Params == nil {
		return false
	}
	for _, field := range fn.Type.Params.List {
		// Look for type *models.Stage or *Stage
		if starExpr, ok := field.Type.(*ast.StarExpr); ok {
			switch t := starExpr.X.(type) {
			case *ast.SelectorExpr:
				if t.Sel.Name == "Stage" {
					return true
				}
			case *ast.Ident:
				if t.Name == "Stage" {
					return true
				}
			}
		}
	}
	return false
}

func mergeParsedStageFiles(parsedFiles []*ParsedStageFile, pkgOverride string) (string, error) {
	if len(parsedFiles) == 0 {
		return "", errors.New("no files to merge")
	}

	pkgName := parsedFiles[0].PackageName
	if pkgOverride != "" {
		pkgName = pkgOverride
	}

	// Deduplicate imports
	importsMap := make(map[string]*ImportEntry)
	for _, pf := range parsedFiles {
		for _, imp := range pf.Imports {
			key := imp.Alias + ":" + imp.Path
			if _, exists := importsMap[key]; !exists {
				importsMap[key] = imp
			}
		}
	}

	// Make sure "time" and "slices" are included if needed
	var importLines []string
	modelsImport := ""
	var otherImports []string

	for _, imp := range importsMap {
		line := ""
		if imp.Alias != "" {
			line = fmt.Sprintf("\t%s %s", imp.Alias, imp.Path)
		} else {
			line = fmt.Sprintf("\t%s", imp.Path)
		}

		if strings.Contains(imp.Path, "/models\"") {
			modelsImport = line
		} else {
			otherImports = append(otherImports, line)
		}
	}

	sort.Strings(otherImports)
	if modelsImport != "" {
		importLines = append(otherImports, "\n"+modelsImport)
	} else {
		importLines = otherImports
	}

	// Track next available index per struct name across all files
	nextAvailableIndex := make(map[string]int)

	const (
		CategoryDeclaration  = 0
		CategoryValueInit    = 1
		CategoryPointerSetup = 2
		CategoryOther        = 3
	)

	var allDeclarations []string
	var allValueInits []string
	var allPointerSetups []string
	var allSequentialStmts []string

	anyHasCommit := false
	for _, pf := range parsedFiles {
		if pf.HasCommit {
			anyHasCommit = true
		}
	}

	for fileIdx, pf := range parsedFiles {
		// Step 1: Scan all identifiers declared/used in this file
		fileDeclaredInstances := make(map[string]int) // structName -> max orig index in this file
		ast.Inspect(pf.FuncDecl.Body, func(n ast.Node) bool {
			if ident, ok := n.(*ast.Ident); ok {
				if match := gongIdentRegex.FindStringSubmatch(ident.Name); len(match) >= 3 {
					structName := match[1]
					idx, err := strconv.Atoi(match[2])
					if err == nil {
						if currentMax, ok := fileDeclaredInstances[structName]; !ok || idx > currentMax {
							fileDeclaredInstances[structName] = idx
						}
					}
				}
			}
			return true
		})

		// Step 2: Build identifier renaming map for this file
		renameMap := make(map[string]string)
		fileMaxNewIndex := make(map[string]int)

		ast.Inspect(pf.FuncDecl.Body, func(n ast.Node) bool {
			if ident, ok := n.(*ast.Ident); ok {
				if match := gongIdentRegex.FindStringSubmatch(ident.Name); len(match) >= 4 {
					structName := match[1]
					origIdx, err := strconv.Atoi(match[2])
					suffix := match[3]
					if err == nil {
						offset := nextAvailableIndex[structName]
						newIdx := origIdx + offset
						newIdent := fmt.Sprintf("__%s__%08d_%s", structName, newIdx, suffix)
						renameMap[ident.Name] = newIdent
						if currentMax, ok := fileMaxNewIndex[structName]; !ok || newIdx > currentMax {
							fileMaxNewIndex[structName] = newIdx
						}
					}
				}
			}
			return true
		})

		// Step 3: Rewrite identifiers in the AST of this file
		ast.Inspect(pf.FuncDecl.Body, func(n ast.Node) bool {
			if ident, ok := n.(*ast.Ident); ok {
				if newIdent, exists := renameMap[ident.Name]; exists {
					ident.Name = newIdent
				}
			}
			return true
		})

		// Update nextAvailableIndex for subsequent files
		for structName, maxNewIdx := range fileMaxNewIndex {
			nextAvailableIndex[structName] = maxNewIdx + 1
		}

		// Step 4: Process and categorize statements in this function
		for _, stmt := range pf.FuncDecl.Body.List {
			var buf bytes.Buffer
			if err := printer.Fprint(&buf, pf.FileSet, stmt); err != nil {
				return "", fmt.Errorf("error printing statement from file %d: %w", fileIdx, err)
			}
			stmtCode := buf.String()

			cat := classifyStatement(stmt)
			switch cat {
			case CategoryDeclaration:
				allDeclarations = append(allDeclarations, stmtCode)
			case CategoryValueInit:
				allValueInits = append(allValueInits, stmtCode)
			case CategoryPointerSetup:
				allPointerSetups = append(allPointerSetups, stmtCode)
			case CategoryOther:
				allPointerSetups = append(allPointerSetups, stmtCode)
			}

			allSequentialStmts = append(allSequentialStmts, stmtCode)
		}
	}

	// Generate source code
	var sb strings.Builder
	sb.WriteString("// file generated by gong\n")
	sb.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	sb.WriteString("import (\n")
	sb.WriteString("\t\"slices\"\n")
	sb.WriteString("\t\"time\"\n")
	for _, imp := range importLines {
		if !strings.Contains(imp, "\"slices\"") && !strings.Contains(imp, "\"time\"") {
			sb.WriteString(imp + "\n")
		}
	}
	sb.WriteString("\t// injection point for ident package import declaration\n")
	sb.WriteString(")\n\n")
	sb.WriteString("// generated in order to avoid error in the package import\n")
	sb.WriteString("// if there are no elements in the stage to marshall\n")
	sb.WriteString("var (\n")
	sb.WriteString("\t_ time.Time\n")
	sb.WriteString("\t_ = slices.Index[[]int, int]\n")
	sb.WriteString(")\n\n")
	sb.WriteString("// function will stage objects\n")
	sb.WriteString("func _(stage *models.Stage) {\n\n")

	if anyHasCommit {
		for _, stmtCode := range allSequentialStmts {
			sb.WriteString("\t" + stmtCode + "\n")
		}
	} else {
		sb.WriteString("\t// insertion point for declaration of instances to stage\n\n")
		for _, decl := range allDeclarations {
			sb.WriteString("\t" + decl + "\n")
		}
		sb.WriteString("\n\t// insertion point for initialization of values\n\n")
		for _, valInit := range allValueInits {
			sb.WriteString("\t" + valInit + "\n")
		}
		sb.WriteString("\n\t// insertion point for setup of pointers\n")
		for _, ptrSetup := range allPointerSetups {
			sb.WriteString("\t" + ptrSetup + "\n")
		}
	}

	sb.WriteString("}\n")

	// Format final source code using gofmt rules
	formatted, err := format.Source([]byte(sb.String()))
	if err != nil {
		// If formatting failed, return raw string along with formatting error for diagnostics
		return sb.String(), fmt.Errorf("error formatting generated source: %w", err)
	}

	return string(formatted), nil
}

func classifyStatement(stmt ast.Stmt) int {
	const (
		CategoryDeclaration  = 0
		CategoryValueInit    = 1
		CategoryPointerSetup = 2
		CategoryOther        = 3
	)

	assignStmt, ok := stmt.(*ast.AssignStmt)
	if !ok {
		return CategoryOther
	}

	if assignStmt.Tok == token.DEFINE {
		return CategoryDeclaration
	}

	if assignStmt.Tok == token.ASSIGN {
		// Check RHS to see if it's a pointer or slice association assignment
		if len(assignStmt.Rhs) > 0 {
			rhs := assignStmt.Rhs[0]
			if isPointerOrSliceExpr(rhs) {
				return CategoryPointerSetup
			}
		}
		return CategoryValueInit
	}

	return CategoryOther
}

func isPointerOrSliceExpr(expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.Ident:
		if e.Name == "nil" {
			return true
		}
		if gongIdentRegex.MatchString(e.Name) {
			return true
		}
	case *ast.CallExpr:
		// Check for append(..., __...__) or slices.Insert(...) or slices.Delete(...)
		if sel, ok := e.Fun.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok && id.Name == "slices" {
				return true
			}
		}
		if id, ok := e.Fun.(*ast.Ident); ok && id.Name == "append" {
			return true
		}
	}

	// Deep check if any child node contains an instance identifier
	hasInstanceIdent := false
	ast.Inspect(expr, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok {
			if gongIdentRegex.MatchString(id.Name) {
				hasInstanceIdent = true
				return false
			}
		}
		return true
	})

	return hasInstanceIdent
}
