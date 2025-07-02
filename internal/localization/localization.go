// Package localization provides internationalization support for the Skene application.
//
// This package uses embedded locale files for the following reasons:
//
// 1. **Simplicity**: No need to manage external files or complex deployment
// 2. **Reliability**: Locale files are always available and cannot be accidentally deleted
// 3. **Security**: Translations cannot be tampered with by external users
// 4. **Performance**: Faster access to locale data (no file system calls)
// 5. **Distribution**: Single binary deployment without external dependencies
// 6. **Version Control**: Locale files are versioned together with the code
//
// The embedded approach ensures that translations are always consistent with the
// application version and eliminates potential runtime errors from missing files.
//
// Usage:
//
//	import "github.com/afa7789/skene/internal/localization"
//
//	// Automatically uses embedded locales from "./locales/*.json"
//	text := localization.T("welcome_message")
//
//	// Change language
//	localization.SetLanguage("es")
//	text := localization.T("welcome_message")
package localization

import (
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.json
var localeFS embed.FS

var (
	bundle           *i18n.Bundle
	localizer        *i18n.Localizer
	currentLanguage  string
	availableLocales []string
)

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load all available locales dynamically
	availableLocales = loadAvailableLocales()

	for _, locale := range availableLocales {
		filename := fmt.Sprintf("locales/%s.json", locale)
		file, err := localeFS.ReadFile(filename)
		if err != nil {
			panic(fmt.Sprintf("failed to load %s: %v", filename, err))
		}
		_, err = bundle.ParseMessageFileBytes(file, filepath.Base(filename))
		if err != nil {
			panic(fmt.Sprintf("failed to parse %s: %v", filename, err))
		}
	}

	// Set default to English if available, otherwise first available
	defaultLang := "en"
	if !contains(availableLocales, defaultLang) && len(availableLocales) > 0 {
		defaultLang = availableLocales[0]
	}
	SetLanguage(defaultLang)
}

func loadAvailableLocales() []string {
	entries, err := localeFS.ReadDir("locales")
	if err != nil {
		panic(fmt.Sprintf("failed to read locales directory: %v", err))
	}

	var locales []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			locale := strings.TrimSuffix(entry.Name(), ".json")
			locales = append(locales, locale)
		}
	}
	return locales
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// SetLanguage changes the current language
func SetLanguage(lang string) {
	currentLanguage = lang
	localizer = i18n.NewLocalizer(bundle, lang)
}

// GetAvailableLanguages returns all available languages found in locales directory
func GetAvailableLanguages() []string {
	return availableLocales
}

// GetCurrentLanguage returns the current language
func GetCurrentLanguage() string {
	return currentLanguage
}

// GetLanguageDisplayKeyMap returns a map of language codes to their translation keys
// by reading the display_key from each locale file
func GetLanguageDisplayKeyMap() map[string]string {
	displayMap := make(map[string]string)

	for _, locale := range availableLocales {
		filename := fmt.Sprintf("locales/%s.json", locale)
		file, err := localeFS.ReadFile(filename)
		if err != nil {
			// Fallback to language code if file can't be read
			displayMap[locale] = locale
			continue
		}

		var localeData map[string]interface{}
		if err := json.Unmarshal(file, &localeData); err != nil {
			// Fallback to language code if JSON can't be parsed
			displayMap[locale] = locale
			continue
		}

		// Try to get display_key, fallback to language code
		if displayKey, exists := localeData["display_key"]; exists {
			if displayStr, ok := displayKey.(string); ok {
				displayMap[locale] = displayStr
			} else {
				displayMap[locale] = locale
			}
		} else {
			displayMap[locale] = locale
		}
	}

	return displayMap
}

// T is a shorthand function for translating text
func T(messageID string, data ...interface{}) string {
	if localizer == nil {
		return messageID
	}

	config := &i18n.LocalizeConfig{
		MessageID: messageID,
	}

	// If data is provided, assume it's template data
	if len(data) > 0 {
		if templateData, ok := data[0].(map[string]interface{}); ok {
			config.TemplateData = templateData
		}
	}

	result, err := localizer.Localize(config)
	if err != nil {
		return messageID
	}

	return result
}
