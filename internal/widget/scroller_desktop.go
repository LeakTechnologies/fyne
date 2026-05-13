//go:build !ci && !no_glfw && !android && !ios && !mobile

package widget

import "fyne.io/fyne/v2"

// DragEnd will stop scrolling on desktop
func (s *Scroll) DragEnd() {
}

// Dragged will scroll on any drag - bar or otherwise - for desktop
func (s *Scroll) Dragged(e *fyne.DragEvent) {
	if s.updateOffset(e.Dragged.DX, e.Dragged.DY) {
		s.refreshWithoutOffsetUpdate()
	}
}
