package gui

import (
	"fyne.io/fyne/v2"

	"github.com/afa7789/skene/internal/localization"
)

func (g *GUI) MainMenu() *fyne.MainMenu {
	// Create language menu dynamically based on available locales
	availableLanguages := localization.GetAvailableLanguages()
	languageDisplayKeys := localization.GetLanguageDisplayKeyMap()
	languageItems := make([]*fyne.MenuItem, 0, len(availableLanguages))

	for _, lang := range availableLanguages {
		lang := lang // capture loop variable
		var displayName string

		// Try to get the translation key from the map, fallback to language code
		if translationKey, exists := languageDisplayKeys[lang]; exists {
			displayName = localization.T(translationKey)
		} else {
			displayName = lang // fallback to language code
		}

		languageItems = append(languageItems, fyne.NewMenuItem(displayName, func() {
			localization.SetLanguage(lang)
			g.UpdateLanguage()
		}))
	}

	languageMenu := fyne.NewMenu(localization.T("menu_language"), languageItems...)

	// Top menu
	fileMenu := fyne.NewMenu(localization.T("menu_file"),
		fyne.NewMenuItem(localization.T("menu_exit"), func() {
			g.app.Quit()
		}),
	)

	mainMenu := fyne.NewMainMenu(fileMenu, languageMenu)

	return mainMenu
}
