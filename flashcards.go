package flashcards

import (
	"embed"
	"path"

	sitetools "go.doolan.dev/site-tools"
)

//go:embed index.html
//go:embed src
//go:embed styles
//go:embed external
var FS embed.FS

func AddToBuild(build *sitetools.Build, mountPath string) error {
	var local sitetools.Build
	if err := local.FromDir(FS, "."); err != nil {
		return err
	}
	for _, a := range local.Assets {
		a.Path = path.Join("/", mountPath, a.Path)
	}
	build.Assets = append(build.Assets, local.Assets...)
	return nil
}
