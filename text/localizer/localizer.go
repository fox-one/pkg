package localizer

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Localizer struct {
	bundle    *i18n.Bundle
	languages []string
}

func New(defaultLanguage language.Tag, messageFiles ...string) *Localizer {
	b := i18n.NewBundle(defaultLanguage)
	b.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	b.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	for _, file := range messageFiles {
		if _, err := b.LoadMessageFile(file); err != nil {
			panic(fmt.Errorf("load message file at %s failed: %w", file, err))
		}
	}

	return NewWithBundle(b)
}

func NewWithBundle(bundle *i18n.Bundle) *Localizer {
	return &Localizer{
		bundle: bundle,
	}
}

func WithLanguage(l *Localizer, langs ...string) *Localizer {
	return &Localizer{
		bundle:    l.bundle,
		languages: append(l.languages, langs...),
	}
}

func (l *Localizer) Localize(id string, args ...interface{}) (string, error) {
	return i18n.NewLocalizer(l.bundle, l.languages...).Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: buildData(args...),
	})
}

func (l *Localizer) MustLocalize(id string, args ...interface{}) string {
	v, err := l.Localize(id, args...)
	if err != nil {
		panic(err)
	}

	return v
}

func (l *Localizer) LocalizeOr(id, fallback string, args ...interface{}) string {
	v, err := l.Localize(id, args...)
	if err != nil {
		v = fallback
	}

	return v
}

func buildData(args ...interface{}) interface{} {
	if len(args) == 1 {
		return args[0]
	}

	data := make(map[string]interface{}, len(args)/2)
	for idx := 0; idx+1 < len(args); idx += 2 {
		data[args[idx].(string)] = args[idx+1]
	}

	return data
}
