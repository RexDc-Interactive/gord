//+build !windows

package tviewutil

import (
	"github.com/cainy-a/gord/tview"
)

// Escape delegates to tview escape, optionally doing additional escaping.
func Escape(text string) string {
	return tview.Escape(text)
}
