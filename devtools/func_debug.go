//go:build debug
// +build debug

package devtools

import (
	"embed"
	"io/fs"
)

//go:embed front_end
var Devtools embed.FS

func FS() (fs.FS, error) {
	return fs.Sub(Devtools, "front_end")
}
