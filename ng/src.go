package ng

import "embed"

// Src is the export of the angular source code. This allows for go vendoring to import ng sources
// and align go source and ng source of the gong stack
//go:embed src
var Src embed.FS
