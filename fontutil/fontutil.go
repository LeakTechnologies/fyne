// Package fontutil exposes auxiliary font registration for the VideoTools font fallback system.
// It provides a thin public wrapper over the internal painter's SetAuxiliaryFont mechanism,
// which allows a supplementary font (e.g. Aboriginal Sans) to be injected into the face stack
// after the primary font so that only characters missing from the primary are rendered with it.
package fontutil

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/cache"
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

// ClearFontCache flushes the font face cache and the font metric cache.
// Call this after SetTheme or SetAuxiliaryFont to ensure the next render
// builds fresh face stacks from the current theme.
func ClearFontCache() {
	painter.ClearFontCache()
	cache.ResetThemeCaches()
}

// SetFontCacheDebugCallback registers an optional callback invoked the first time
// a face list is built for a given text style. styleName is "regular", "bold",
// "italic", "bold+italic", or "monospace". fontResourceName is the Name() of the
// font resource that becomes faces[0] in the face stack.
// Pass nil to clear. For diagnostic use only; not needed in production.
func SetFontCacheDebugCallback(fn func(styleName, fontResourceName string)) {
	painter.FontCacheDebugCallback = fn
}
