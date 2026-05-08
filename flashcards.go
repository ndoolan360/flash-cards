package flashcards

import (
	"embed"
)

//go:embed index.html
//go:embed src
//go:embed styles
//go:embed external
var FS embed.FS
