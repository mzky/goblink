package goblink

import (
	"github.com/mzky/goblink/win32"
)

func NewWebBrowser(url, title, devtoolsPath string, max, mb, ib bool, width, height int) error {
	return win32.StartBlinkMain(url, title, devtoolsPath, max, mb, ib, width, height)
}
