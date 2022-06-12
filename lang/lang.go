package lang

import (
	"fmt"
	"github.com/RestartFU/gophig"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"golang.org/x/text/language"
)

// translationData contains the properties and translations of a language.
type translationData struct {
	Properties struct {
		Name  string `json:"name"`
		Image string `json:"image"`
	} `json:"properties"`
	Translations map[string]string `json:"translations"`
}

// translations stores a mapping between the language and the translation data.
var translations = make(map[language.Tag]translationData)

// Register registers a translation file and adds the decoded data to the translations list.
func Register(lang language.Tag) {
	var data translationData
	if err := gophig.GetConfComplex(fmt.Sprintf("assets/translations/%v.json", lang.String()), gophig.JSONMarshaler{}, &data); err != nil {
		panic(err)
	}
	translations[lang] = data
}

// Properties returns the name and image of a language.
func Properties(lang language.Tag) (string, string, bool) {
	data, ok := translations[lang]
	return data.Properties.Name, data.Properties.Image, ok
}

// Translatef returns the translated version of a string.
func Translatef(lang language.Tag, key string, a ...interface{}) string {
	return text.Colourf(Translate(lang, key), a...)
}

// Translate returns the translated version of a string.
func Translate(lang language.Tag, key string) string {
	t, ok := translations[lang]
	if !ok {
		t = translations[language.English]
	}
	return t.Translations[key]
}
