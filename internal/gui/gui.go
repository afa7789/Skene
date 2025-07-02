package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/afa7789/skene/internal/counter"
	"github.com/afa7789/skene/internal/localization"
)

// built in icon.go by the fyne bundle command
// var resourceIconSvg fyne.Resource
var appID = "com.afa7789.skene"

type GUI struct {
	app            fyne.App
	window         fyne.Window
	counterService *counter.CounterService
}

func NewGUI() *GUI {
	myApp := app.NewWithID(appID)
	myWindow := myApp.NewWindow(localization.T("app_title"))

	// Cria os serviços
	counterService := counter.NewCounterService()

	return &GUI{
		app:            myApp,
		window:         myWindow,
		counterService: counterService,
	}
}

func (g *GUI) updateContent() {
	// Cria counter widget com injeção de dependência
	counterWidget := NewCounterWidget(g.counterService)

	// Conteúdo principal
	content := container.NewVBox(
		widget.NewLabel(localization.T("hello_skene")),
		counterWidget.GetContainer(),
	)

	g.window.SetContent(content)
}

func (g *GUI) UpdateLanguage() {
	// Atualiza título da janela
	g.window.SetTitle(localization.T("app_title"))

	// Recria conteúdo para atualizar textos (isso cria um novo widget com textos atualizados)
	g.updateContent()
}

func (g *GUI) Serve() {
	// g.app.SetIcon(resourceIconSvg) // TODO: Fix icon
	g.window.Resize(fyne.NewSize(400, 300))

	// Configura menu
	g.window.SetMainMenu(g.MainMenu())

	// Configura conteúdo inicial
	g.updateContent()

	g.window.ShowAndRun()
}
