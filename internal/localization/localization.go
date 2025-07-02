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
