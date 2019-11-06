package localizer

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	jsoniter "github.com/json-iterator/go"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestLoadToml(t *testing.T) {
	b := i18n.NewBundle(language.English)
	b.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	f, err := b.LoadMessageFile("./testdata/active.en.toml")
	assert.Nil(t, err)
	for _, msg := range f.Messages {
		data, _ := jsoniter.MarshalToString(msg)
		t.Log(data)
	}
}

func TestLoadYaml(t *testing.T) {
	b := i18n.NewBundle(language.English)
	b.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	f, err := b.LoadMessageFile("./testdata/active.en.yaml")
	assert.Nil(t, err)
	for _, msg := range f.Messages {
		data, _ := jsoniter.MarshalToString(msg)
		t.Log(data)
	}
}
