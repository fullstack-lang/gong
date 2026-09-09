package golang

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

// ScrambleMode defines the strategy used to scramble strings.
type ScrambleMode string

const (
	// ScrambleModeWords preserves word boundaries, casing, character length,
	// and delimiters (spaces, hyphens, punctuation). Ideal for diagram layouts.
	ScrambleModeWords ScrambleMode = "words"

	// ScrambleModeTokens replaces each unique string with an opaque random token.
	ScrambleModeTokens ScrambleMode = "tokens"
)

// ScrambleOptions configures the scrambling process.
type ScrambleOptions struct {
	Mode       ScrambleMode
	Salt       string // optional custom salt; if empty, 256-bit crypto/rand is used
	InPlace    bool
	OutputPath string
}

// ScrambleStageFile reads a stage Go file, scrambles all string fields, and writes the output.
func ScrambleStageFile(inputPath string, opts ScrambleOptions) (string, error) {
	if inputPath == "" {
		return "", errors.New("no input stage file provided")
	}

	src, err := os.ReadFile(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", inputPath, err)
	}

	scrambledBytes, err := ScrambleStageSource(src, opts)
	if err != nil {
		return "", fmt.Errorf("failed to scramble %s: %w", inputPath, err)
	}

	outPath := opts.OutputPath
	if opts.InPlace {
		outPath = inputPath
	} else if outPath == "" {
		dir := filepath.Dir(inputPath)
		base := filepath.Base(inputPath)
		ext := filepath.Ext(base)
		stem := strings.TrimSuffix(base, ext)
		outPath = filepath.Join(dir, stem+"-scrambled"+ext)
	}

	dir := filepath.Dir(outPath)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return "", fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	if err := os.WriteFile(outPath, scrambledBytes, 0o644); err != nil {
		return "", fmt.Errorf("failed to write scrambled stage file to %s: %w", outPath, err)
	}

	return outPath, nil
}

// ScrambleStageSource scrambles all string fields in a Go source byte slice representing a stage file.
func ScrambleStageSource(src []byte, opts ScrambleOptions) ([]byte, error) {
	fset := token.NewFileSet()
	inFile, err := parser.ParseFile(fset, "stage.go", src, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("unable to parse stage file: %w", err)
	}

	stagingFuncs := findStagingFunctions(inFile)
	if len(stagingFuncs) == 0 {
		return nil, errors.New("no staging function found in stage file")
	}

	scrambler, err := newStringScrambler(opts.Mode, opts.Salt)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize scrambler: %w", err)
	}

	for _, fn := range stagingFuncs {
		scrambleStagingFunction(fn, scrambler)
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, inFile); err != nil {
		return nil, fmt.Errorf("error formatting scrambled AST: %w", err)
	}

	return buf.Bytes(), nil
}

func findStagingFunctions(inFile *ast.File) []*ast.FuncDecl {
	var fns []*ast.FuncDecl
	for _, decl := range inFile.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if isStagingFunction(fn) {
				fns = append(fns, fn)
			}
		}
	}
	return fns
}

func scrambleStagingFunction(fn *ast.FuncDecl, scrambler *stringScrambler) {
	if fn.Body == nil {
		return
	}

	for _, stmt := range fn.Body.List {
		assignStmt, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}

		// Case 1: Declaration with Stage call:
		// __Struct__00000000_ := (&models.Struct{Name: "..."}).Stage(stage)
		if assignStmt.Tok == token.DEFINE {
			for _, rhs := range assignStmt.Rhs {
				scrambleCompositeLiteralsInExpr(rhs, scrambler)
			}
			continue
		}

		// Case 2: Direct assignment:
		// __Struct__00000000_.Field = "..."
		// or append: __Struct__00000000_.Field = append(__Struct__00000000_.Field, "...")
		if assignStmt.Tok == token.ASSIGN {
			if len(assignStmt.Lhs) > 0 {
				if sel, ok := assignStmt.Lhs[0].(*ast.SelectorExpr); ok {
					if id, ok := sel.X.(*ast.Ident); ok && gongIdentRegex.MatchString(id.Name) {
						if len(assignStmt.Rhs) > 0 {
							rhs := assignStmt.Rhs[0]
							// Check if it is a time.Parse call
							if call, ok := rhs.(*ast.CallExpr); ok {
								if isTimeParseCall(call) {
									// Skip time parsing expressions
									continue
								}
								// Check for append call
								if isAppendCall(call) {
									for i := 1; i < len(call.Args); i++ {
										if strVal, ok := extractStringVal(call.Args[i]); ok && strVal != "" {
											scrambled := scrambler.Scramble(strVal)
											call.Args[i] = makeStringLit(call.Args[i].Pos(), scrambled)
										}
									}
									continue
								}
							}

							// Regular string assignment
							if strVal, ok := extractStringVal(rhs); ok && strVal != "" {
								scrambled := scrambler.Scramble(strVal)
								assignStmt.Rhs[0] = makeStringLit(rhs.Pos(), scrambled)
							}
						}
					}
				}
			}
		}
	}
}

func scrambleCompositeLiteralsInExpr(expr ast.Expr, scrambler *stringScrambler) {
	ast.Inspect(expr, func(n ast.Node) bool {
		compLit, ok := n.(*ast.CompositeLit)
		if !ok {
			return true
		}

		for _, elt := range compLit.Elts {
			kv, ok := elt.(*ast.KeyValueExpr)
			if !ok {
				continue
			}

			if strVal, ok := extractStringVal(kv.Value); ok && strVal != "" {
				scrambled := scrambler.Scramble(strVal)
				kv.Value = makeStringLit(kv.Value.Pos(), scrambled)
			}
		}

		return true
	})
}

func isTimeParseCall(call *ast.CallExpr) bool {
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if xIdent, ok := sel.X.(*ast.Ident); ok {
			if xIdent.Name == "time" && sel.Sel.Name == "Parse" {
				return true
			}
		}
	}
	return false
}

func isAppendCall(call *ast.CallExpr) bool {
	if id, ok := call.Fun.(*ast.Ident); ok && id.Name == "append" {
		return true
	}
	return false
}

func extractStringVal(expr ast.Expr) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			val, err := strconv.Unquote(e.Value)
			if err != nil {
				val = strings.Trim(e.Value, "`\"")
			}
			return val, true
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			s1, ok1 := extractStringVal(e.X)
			s2, ok2 := extractStringVal(e.Y)
			if ok1 && ok2 {
				return s1 + s2, true
			}
		}
	}
	return "", false
}

func makeStringLit(pos token.Pos, strVal string) *ast.BasicLit {
	return &ast.BasicLit{
		ValuePos: pos,
		Kind:     token.STRING,
		Value:    ToRawStringLiteral(strVal),
	}
}

// ToRawStringLiteral formats a string into safe Go source code wrapped in backticks.
func ToRawStringLiteral(s string) string {
	escaped := strings.ReplaceAll(s, "`", "` + \"`\" + `")
	result := "`" + escaped + "`"
	result = strings.ReplaceAll(result, "`` + ", "")
	result = strings.ReplaceAll(result, " + ``", "")
	return result
}

// stringScrambler manages bijective string obfuscation with cryptographic salt.
type stringScrambler struct {
	mode            ScrambleMode
	salt            []byte
	stringCache     map[string]string // complete string -> scrambled string
	wordToScrambled map[string]string // word -> scrambled word
	scrambledToWord map[string]string // scrambled word -> original word (collision check)
}

func newStringScrambler(mode ScrambleMode, customSalt string) (*stringScrambler, error) {
	var salt []byte
	if customSalt != "" {
		salt = []byte(customSalt)
	} else {
		// Generate 256-bit cryptographically secure random salt
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, fmt.Errorf("failed to generate random salt: %w", err)
		}
	}

	if mode == "" {
		mode = ScrambleModeWords
	}

	return &stringScrambler{
		mode:            mode,
		salt:            salt,
		stringCache:     make(map[string]string),
		wordToScrambled: make(map[string]string),
		scrambledToWord: make(map[string]string),
	}, nil
}

// Scramble transforms an input string into a scrambled version.
func (s *stringScrambler) Scramble(input string) string {
	if input == "" {
		return ""
	}

	if cached, exists := s.stringCache[input]; exists {
		return cached
	}

	var result string
	switch s.mode {
	case ScrambleModeTokens:
		result = s.scrambleAsToken(input)
	case ScrambleModeWords:
		fallthrough
	default:
		result = s.scrambleWords(input)
	}

	s.stringCache[input] = result
	return result
}

func (s *stringScrambler) scrambleAsToken(input string) string {
	h := hmac.New(sha256.New, s.salt)
	h.Write([]byte(input))
	sum := h.Sum(nil)

	// Short token derived from HMAC (e.g. 12 hex chars)
	tokenHex := hex.EncodeToString(sum[:6])
	return "Scrambled_" + tokenHex
}

func (s *stringScrambler) scrambleWords(input string) string {
	var sb strings.Builder
	var currentWord strings.Builder

	flushWord := func() {
		if currentWord.Len() > 0 {
			w := currentWord.String()
			sb.WriteString(s.getOrGenerateScrambledWord(w))
			currentWord.Reset()
		}
	}

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			currentWord.WriteRune(r)
		} else {
			flushWord()
			sb.WriteRune(r)
		}
	}
	flushWord()

	return sb.String()
}

func (s *stringScrambler) getOrGenerateScrambledWord(word string) string {
	if cached, exists := s.wordToScrambled[word]; exists {
		return cached
	}

	// Consonants and vowels for readable pseudo-words
	const consonants = "bcdfghjklmnpqrstvwxyz"
	const vowels = "aeiou"

	counter := 0
	for {
		h := hmac.New(sha256.New, s.salt)
		h.Write([]byte(word))
		if counter > 0 {
			h.Write([]byte(fmt.Sprintf("_%d", counter)))
		}
		hash := h.Sum(nil)

		runes := []rune(word)
		outRunes := make([]rune, len(runes))

		// Check if word is all uppercase
		isAllUpper := true
		for _, r := range runes {
			if unicode.IsLetter(r) && !unicode.IsUpper(r) {
				isAllUpper = false
				break
			}
		}

		// Alternate starting consonant or vowel based on first byte of hash
		startWithVowel := (hash[0] % 2) == 1

		for i, r := range runes {
			if unicode.IsDigit(r) {
				digitVal := hash[(i+1)%len(hash)] % 10
				outRunes[i] = rune('0' + digitVal)
			} else if unicode.IsLetter(r) {
				var charRune rune
				var pickVowel bool
				if startWithVowel {
					pickVowel = (i % 2) == 0
				} else {
					pickVowel = (i % 2) == 1
				}

				byteVal := int(hash[(i*3+2)%len(hash)])
				if pickVowel {
					charRune = rune(vowels[byteVal%len(vowels)])
				} else {
					charRune = rune(consonants[byteVal%len(consonants)])
				}

				if isAllUpper || unicode.IsUpper(r) {
					charRune = unicode.ToUpper(charRune)
				} else {
					charRune = unicode.ToLower(charRune)
				}
				outRunes[i] = charRune
			} else {
				outRunes[i] = r
			}
		}

		scrambled := string(outRunes)
		// Check for collision with a different original word
		if orig, exists := s.scrambledToWord[scrambled]; exists && orig != word {
			counter++
			continue
		}

		s.wordToScrambled[word] = scrambled
		s.scrambledToWord[scrambled] = word
		return scrambled
	}
}
