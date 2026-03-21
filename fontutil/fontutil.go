// Package fontutil exposes auxiliary font registration for the VideoTools font fallback system.
// It provides a thin public wrapper over the internal painter's SetAuxiliaryFont mechanism,
// which allows a supplementary font (e.g. Aboriginal Sans) to be injected into the face stack
// after the primary font so that only characters missing from the primary are rendered with it.
package fontutil

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/painter"
)

// SetAuxiliaryFont registers a supplementary font added to every face stack immediately after
// the primary theme font. Characters supported by the primary font (e.g. Latin in IBM Plex Mono)
// continue to render with it; characters not found there (e.g. UCAS syllabics) fall through to
// the auxiliary. Pass nil to clear.
//
// Call fyne.CurrentApp().Settings().SetTheme(…) after changing this to flush cached face lists.
func SetAuxiliaryFont(r fyne.Resource) {
	painter.SetAuxiliaryFont(r)
}
