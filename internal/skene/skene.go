// client package for skene
package skene

import "github.com/afa7789/skene/internal/gui"

// Run starts the main client (GUI for now)
func Run() {
	gui := gui.NewGUI()
	gui.Serve()
}
