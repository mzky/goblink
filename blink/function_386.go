//go:build 386
// +build 386

package blink

import (
	_ "embed"
)

//go:embed blink32.dll
var Node []byte
