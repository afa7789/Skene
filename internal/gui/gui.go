package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/afa7789/skene/internal/localization"
)

// built in icon.go by the fyne bundle command
// var resourceIconSvg fyne.Resource
var appID = "com.afa7789.skene"

type GUI struct {
	app    fyne.App
	window fyne.Window
}

func NewGUI() *GUI {
	myApp := app.NewWithID(appID)
	myWindow := myApp.NewWindow(localization.T("app_title"))

	return &GUI{
		app:    myApp,
		window: myWindow,
	}
}

func (g *GUI) Serve() {
	g.app.SetIcon(resourceIconSvg)
	g.updateContent()
	// importing menu
	mainMenu := g.MainMenu()
	g.window.SetMainMenu(mainMenu)

	g.window.Resize(fyne.NewSize(400, 300))
	g.window.ShowAndRun()
}

// updateContent creates and sets the main content
func (g *GUI) updateContent() {
	label := widget.NewLabel(localization.T("hello_skene"))
	content := container.NewVBox(label)
	g.window.SetContent(content)
}

// UpdateLanguage updates all UI elements with new language
func (g *GUI) UpdateLanguage() {
	g.window.SetTitle(localization.T("app_title"))

	// Recreate content with new language
	g.updateContent()

	// Recreate menu with new language
	mainMenu := g.MainMenu()
	g.window.SetMainMenu(mainMenu)
}
