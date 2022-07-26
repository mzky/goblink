//go:build amd64
// +build amd64

package blink

import (
	_ "embed"
)

//go:embed blink64.dll
var Node []byte
