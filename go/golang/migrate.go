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
	"strings"
)

// MigrateStageFile converts a single-stage Gong data file to a multiple-stage StageSet data file.
func MigrateStageFile(inputPath string, outputPath string, pkgOverride string, inPlace bool) (string, error) {
	if inputPath == "" {
		return "", errors.New("no input stage file provided")
	}

	src, err := os.ReadFile(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to read input file %s: %w", inputPath, err)
	}

	migratedSource, err := MigrateStageSource(src, pkgOverride, inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to migrate %s: %w", inputPath, err)
	}

	if inPlace {
		outputPath = inputPath
	} else if outputPath == "" {
		dir := filepath.Dir(inputPath)
		base := filepath.Base(inputPath)
		if base == "stage.go" {
			outputPath = filepath.Join(dir, "stageset.go")
		} else if strings.HasSuffix(base, ".go") {
			outputPath = filepath.Join(dir, strings.TrimSuffix(base, ".go")+"_stageset.go")
		} else {
			outputPath = inputPath + "_stageset.go"
		}
	}

	// Ensure destination directory exists
	dir := filepath.Dir(outputPath)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return "", fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	if err := os.WriteFile(outputPath, []byte(migratedSource), 0o644); err != nil {
		return "", fmt.Errorf("failed to write migrated file to %s: %w", outputPath, err)
	}

	return outputPath, nil
}

// MigrateStageSource converts single-stage Gong source code bytes to StageSet source code string.
func MigrateStageSource(src []byte, pkgOverride string, filePath string) (string, error) {
	fset := token.NewFileSet()
	inFile, err := parser.ParseFile(fset, filePath, src, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("unable to parse file: %w", err)
	}

	// 1. Check if the file is already a StageSet file
	for _, decl := range inFile.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if isStageSetFunction(fn) {
				// Already migrated, format and return
				var buf bytes.Buffer
				if err := format.Node(&buf, fset, inFile); err == nil {
					return buf.String(), nil
				}
				return string(src), nil
			}
		}
	}

	// 2. Find the staging function
	var stagingFunc *ast.FuncDecl
	for _, decl := range inFile.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if isSingleStageFunction(fn) {
				stagingFunc = fn
				break
			}
		}
	}
	if stagingFunc == nil {
		return "", fmt.Errorf("file %s does not contain a single-stage initialization function", filePath)
	}

	// 3. Package name
	pkgName := inFile.Name.Name
	if pkgOverride != "" {
		pkgName = pkgOverride
	}

	// 4. Identify models import path
	modelsImportPath, otherImports := extractModelsImport(inFile, stagingFunc, filePath)
	if modelsImportPath == "" {
		return "", fmt.Errorf("unable to determine models package path for %s", filePath)
	}

	// 5. Scan all identifiers declared in DEFINE statements: __<Struct>__<Digits>_<Name>
	// Build map from old identifier to new StageSet identifier: __stage_0__<Struct>__<08d>_
	renameMap := make(map[string]string)
	structCounts := make(map[string]int)

	for _, stmt := range stagingFunc.Body.List {
		if assignStmt, ok := stmt.(*ast.AssignStmt); ok && assignStmt.Tok == token.DEFINE {
			if len(assignStmt.Lhs) > 0 {
				if ident, ok := assignStmt.Lhs[0].(*ast.Ident); ok {
					structName := extractStructNameFromIdent(ident.Name)
					if structName == "" {
						// Fallback: inspect the RHS CompositeLit type
						structName = extractStructNameFromRHS(assignStmt.Rhs[0])
					}
					if structName != "" {
						idx := structCounts[structName]
						structCounts[structName]++
						newIdent := fmt.Sprintf("__stage_0__%s__%08d_", structName, idx)
						renameMap[ident.Name] = newIdent
					}
				}
			}
		}
	}

	// 6. Rewrite AST nodes in staging function body:
	// - Identifiers in renameMap
	// - (&models.Struct{...}).Stage(stage) -> (&__stage_0__.Struct{...}).Stage(stageSet.Stage)
	// - stage.Commit() -> stageSet.Commit()
	for _, stmt := range stagingFunc.Body.List {
		if assignStmt, ok := stmt.(*ast.AssignStmt); ok && assignStmt.Tok == token.DEFINE {
			// Rewrite LHS identifier
			if len(assignStmt.Lhs) > 0 {
				if ident, ok := assignStmt.Lhs[0].(*ast.Ident); ok {
					if newIdent, exists := renameMap[ident.Name]; exists {
						ident.Name = newIdent
					}
				}
			}
			// Rewrite RHS:
			// Change (&models.Foo{...}).Stage(stage)
			// to (&__stage_0__.Foo{...}).Stage(stageSet.Stage)
			rewriteStageCall(assignStmt.Rhs[0])
		}

		// Rewrite all other occurrences of renamed identifiers (in LHS, RHS, appends, etc.)
		ast.Inspect(stmt, func(n ast.Node) bool {
			if ident, ok := n.(*ast.Ident); ok {
				if newIdent, exists := renameMap[ident.Name]; exists {
					ident.Name = newIdent
				}
			}
			// Rewrite stage.Commit() -> stageSet.Commit()
			if call, ok := n.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if id, ok := sel.X.(*ast.Ident); ok && id.Name == "stage" {
						id.Name = "stageSet"
					}
				}
			}
			return true
		})
	}

	// 7. Categorize statements into the 3 phases
	const (
		Phase1Declarations = 1
		Phase2ValueInits   = 2
		Phase3Pointers     = 3
	)

	var phase1Stmts []string
	var phase2Stmts []string
	var phase3Stmts []string
	var otherStmts []string

	for _, stmt := range stagingFunc.Body.List {
		var buf bytes.Buffer
		if err := printer.Fprint(&buf, fset, stmt); err != nil {
			return "", fmt.Errorf("error printing statement: %w", err)
		}
		stmtCode := buf.String()

		switch classifyMigrationStatement(stmt) {
		case Phase1Declarations:
			phase1Stmts = append(phase1Stmts, stmtCode)
		case Phase2ValueInits:
			phase2Stmts = append(phase2Stmts, stmtCode)
		case Phase3Pointers:
			phase3Stmts = append(phase3Stmts, stmtCode)
		default:
			otherStmts = append(otherStmts, stmtCode)
		}
	}

	// 8. Build migrated source code string
	var sb strings.Builder
	sb.WriteString("// file generated by gong\n")
	sb.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	sb.WriteString("import (\n")
	sb.WriteString("\t\"slices\"\n")
	sb.WriteString("\t\"time\"\n\n")
	sb.WriteString(fmt.Sprintf("\t__stage_0__ \"%s\"\n", modelsImportPath))
	for _, imp := range otherImports {
		if imp != "\"slices\"" && imp != "\"time\"" {
			sb.WriteString(fmt.Sprintf("\t%s\n", imp))
		}
	}
	sb.WriteString(")\n\n")

	sb.WriteString("var (\n")
	sb.WriteString("\t_ time.Time\n")
	sb.WriteString("\t_ = slices.Index[[]int, int]\n\n")
	sb.WriteString("\t_ *__stage_0__.Stage\n")
	sb.WriteString(")\n\n")

	sb.WriteString("// function will stage objects across all coordinated stages\n")
	sb.WriteString("func _(stageSet *__stage_0__.StageSet) {\n\n")

	sb.WriteString("\t// ------------------------------------------------------------------------\n")
	sb.WriteString("\t// Phase 1: Declarations (in topological order: leaves first)\n")
	sb.WriteString("\t// ------------------------------------------------------------------------\n")
	for _, s := range phase1Stmts {
		sb.WriteString(fmt.Sprintf("\t%s\n", s))
	}

	sb.WriteString("\n\t// ------------------------------------------------------------------------\n")
	sb.WriteString("\t// Phase 2: Value Initializations\n")
	sb.WriteString("\t// ------------------------------------------------------------------------\n")
	for _, s := range phase2Stmts {
		sb.WriteString(fmt.Sprintf("\t%s\n", s))
	}

	sb.WriteString("\n\t// ------------------------------------------------------------------------\n")
	sb.WriteString("\t// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)\n")
	sb.WriteString("\t// ------------------------------------------------------------------------\n")
	for _, s := range phase3Stmts {
		sb.WriteString(fmt.Sprintf("\t%s\n", s))
	}

	if len(otherStmts) > 0 {
		sb.WriteString("\n")
		for _, s := range otherStmts {
			sb.WriteString(fmt.Sprintf("\t%s\n", s))
		}
	}

	sb.WriteString("}\n")

	// Format final Go source
	formattedBytes, err := format.Source([]byte(sb.String()))
	if err != nil {
		// Return unformatted source if format fails so error can be investigated
		return sb.String(), fmt.Errorf("error formatting generated source: %w", err)
	}

	return string(formattedBytes), nil
}

func isSingleStageFunction(fn *ast.FuncDecl) bool {
	if fn.Type == nil || fn.Type.Params == nil {
		return false
	}
	for _, field := range fn.Type.Params.List {
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

func isStageSetFunction(fn *ast.FuncDecl) bool {
	if fn.Type == nil || fn.Type.Params == nil {
		return false
	}
	for _, field := range fn.Type.Params.List {
		if starExpr, ok := field.Type.(*ast.StarExpr); ok {
			switch t := starExpr.X.(type) {
			case *ast.SelectorExpr:
				if t.Sel.Name == "StageSet" {
					return true
				}
			case *ast.Ident:
				if t.Name == "StageSet" {
					return true
				}
			}
		}
	}
	return false
}

func extractModelsImport(inFile *ast.File, fn *ast.FuncDecl, filePath string) (string, []string) {
	var modelsPath string
	var otherImports []string

	// Check function param selector name if present, e.g. `stage *models.Stage`
	paramPkgAlias := ""
	if fn.Type != nil && fn.Type.Params != nil {
		for _, field := range fn.Type.Params.List {
			if starExpr, ok := field.Type.(*ast.StarExpr); ok {
				if sel, ok := starExpr.X.(*ast.SelectorExpr); ok {
					if id, ok := sel.X.(*ast.Ident); ok {
						paramPkgAlias = id.Name
					}
				}
			}
		}
	}

	for _, imp := range inFile.Imports {
		rawPath := strings.Trim(imp.Path.Value, "\"")
		alias := ""
		if imp.Name != nil {
			alias = imp.Name.Name
		}

		isModels := false
		if alias != "" && alias == paramPkgAlias {
			isModels = true
		} else if strings.HasSuffix(rawPath, "/models") {
			isModels = true
		} else if paramPkgAlias == "models" && strings.Contains(rawPath, "models") {
			isModels = true
		}

		if isModels && modelsPath == "" {
			modelsPath = rawPath
		} else {
			if alias != "" {
				otherImports = append(otherImports, fmt.Sprintf("%s \"%s\"", alias, rawPath))
			} else {
				otherImports = append(otherImports, fmt.Sprintf("\"%s\"", rawPath))
			}
		}
	}

	// Fallback if not found directly in imports
	if modelsPath == "" && filePath != "" {
		modelsPath = inferModelsPathFromDir(filePath)
	}

	return modelsPath, otherImports
}

func inferModelsPathFromDir(filePath string) string {
	abs, err := filepath.Abs(filePath)
	if err != nil {
		return ""
	}
	curr := filepath.Dir(abs)
	for curr != "/" && curr != "." {
		goModPath := filepath.Join(curr, "go.mod")
		if data, err := os.ReadFile(goModPath); err == nil {
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "module ") {
					modName := strings.TrimSpace(strings.TrimPrefix(line, "module"))
					rel, _ := filepath.Rel(curr, filepath.Dir(abs))
					// Look for /go/models
					if strings.Contains(rel, "go") {
						return modName + "/go/models"
					}
					return modName + "/models"
				}
			}
		}
		curr = filepath.Dir(curr)
	}
	return ""
}

func extractStructNameFromIdent(identName string) string {
	// Pattern: __<StructName>__<Digits>_<optional>
	match := gongIdentRegex.FindStringSubmatch(identName)
	if len(match) >= 2 {
		return match[1]
	}
	return ""
}

func extractStructNameFromRHS(expr ast.Expr) string {
	var structName string
	ast.Inspect(expr, func(n ast.Node) bool {
		if compLit, ok := n.(*ast.CompositeLit); ok {
			switch t := compLit.Type.(type) {
			case *ast.SelectorExpr:
				structName = t.Sel.Name
				return false
			case *ast.Ident:
				structName = t.Name
				return false
			}
		}
		return true
	})
	return structName
}

func rewriteStageCall(rhs ast.Expr) {
	// Look for call to .Stage(stage)
	if call, ok := rhs.(*ast.CallExpr); ok {
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "Stage" {
			// Change argument from `stage` to `stageSet.Stage`
			call.Args = []ast.Expr{
				&ast.SelectorExpr{
					X:   ast.NewIdent("stageSet"),
					Sel: ast.NewIdent("Stage"),
				},
			}
			// Check receiver: (&models.A{...})
			if paren, ok := sel.X.(*ast.ParenExpr); ok {
				rewriteCompositeLitType(paren.X)
			} else {
				rewriteCompositeLitType(sel.X)
			}
		}
	}
}

func rewriteCompositeLitType(expr ast.Expr) {
	if unary, ok := expr.(*ast.UnaryExpr); ok && unary.Op == token.AND {
		expr = unary.X
	}
	if compLit, ok := expr.(*ast.CompositeLit); ok {
		switch t := compLit.Type.(type) {
		case *ast.SelectorExpr:
			// Change package from `models` to `__stage_0__`
			t.X = ast.NewIdent("__stage_0__")
		case *ast.Ident:
			// Change from `A` to `__stage_0__.A`
			compLit.Type = &ast.SelectorExpr{
				X:   ast.NewIdent("__stage_0__"),
				Sel: ast.NewIdent(t.Name),
			}
		}
	}
}

func classifyMigrationStatement(stmt ast.Stmt) int {
	const (
		Phase1Declarations = 1
		Phase2ValueInits   = 2
		Phase3Pointers     = 3
		PhaseOther         = 4
	)

	assignStmt, ok := stmt.(*ast.AssignStmt)
	if !ok {
		return PhaseOther
	}

	if assignStmt.Tok == token.DEFINE {
		return Phase1Declarations
	}

	if assignStmt.Tok == token.ASSIGN {
		if len(assignStmt.Rhs) > 0 {
			rhs := assignStmt.Rhs[0]
			if isPointerOrSliceExpr(rhs) {
				return Phase3Pointers
			}
		}
		return Phase2ValueInits
	}

	return PhaseOther
}
