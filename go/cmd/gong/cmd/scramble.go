package cmd

import (
	"log"

	"github.com/fullstack-lang/gong/go/golang"
	"github.com/spf13/cobra"
)

var (
	scrambleOutputPath string
	scrambleInPlace    bool
	scrambleMode       string
	scrambleSalt       string
)

var scrambleCmd = &cobra.Command{
	Use:   "scramble [flags] <stage.go>",
	Short: "Scrambles all string fields in a stage file with irreversible cryptography",
	Long: `Scrambles all string fields of staged model instances in a Gong stage file.

The scrambling is cryptographically irreversible, using HMAC-SHA256 with an ephemeral
256-bit random salt from crypto/rand. Because the salt is never saved, anyone possessing
the open-source Gong codebase cannot unscramble the string fields or run dictionary attacks.

Modes:
  - words (default): Preserves word boundaries, character length, and casing using
    pronounceable pseudo-words, keeping delimiters (hyphens, spaces) intact. This ensures
    diagram layouts, box widths, and DSM compound shape names remain consistent.
  - tokens: Replaces each unique string with an opaque random token (e.g. Scrambled_a9f1b4...).

Examples:
  gong scramble data/stage.go
  gong scramble data/stage.go -o data/stage-scrambled.go
  gong scramble data/stage.go -i
  gong scramble data/stage.go --mode tokens
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.SetPrefix("gong scramble: ")
		log.SetFlags(0)

		if scrambleInPlace && scrambleOutputPath != "" {
			log.Fatalf("cannot use both --in-place (-i) and --output (-o)")
		}

		mode := golang.ScrambleMode(scrambleMode)
		if mode != golang.ScrambleModeWords && mode != golang.ScrambleModeTokens {
			log.Fatalf("invalid mode %q: must be 'words' or 'tokens'", scrambleMode)
		}

		opts := golang.ScrambleOptions{
			Mode:       mode,
			Salt:       scrambleSalt,
			InPlace:    scrambleInPlace,
			OutputPath: scrambleOutputPath,
		}

		outPath, err := golang.ScrambleStageFile(args[0], opts)
		if err != nil {
			log.Fatalf("Error scrambling stage file: %v", err)
		}

		log.Printf("Successfully scrambled string fields in %s -> %s", args[0], outPath)
	},
}

func init() {
	scrambleCmd.Flags().StringVarP(&scrambleOutputPath, "output", "o", "", "output file for the scrambled stage (default: <input>-scrambled.go)")
	scrambleCmd.Flags().BoolVarP(&scrambleInPlace, "in-place", "i", false, "overwrite input file in place")
	scrambleCmd.Flags().StringVarP(&scrambleMode, "mode", "m", "words", "scrambling mode: 'words' (preserves words/delimiters/lengths) or 'tokens'")
	scrambleCmd.Flags().StringVarP(&scrambleSalt, "salt", "s", "", "optional custom salt for reproducible scrambling")

	rootCmd.AddCommand(scrambleCmd)
}
