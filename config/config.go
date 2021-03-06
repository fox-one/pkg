package config

import (
	"bytes"
	"encoding/base64"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

// -X 'github.com/fox-one/pkg/config.DATA=$(DATA)'
var DATA string

func Load(configFile string, typ string, v interface{}) error {
	viper.SetConfigType(typ)
	viper.AutomaticEnv()

	if configFile != "" {
		f, err := os.Open(configFile)
		if err != nil {
			return err
		}

		defer f.Close()

		if err := viper.ReadConfig(f); err != nil {
			return err
		}
	}

	data, err := jsoniter.Marshal(viper.AllSettings())
	if err != nil {
		return err
	}

	if err := jsoniter.Unmarshal(data, v); err != nil {
		return err
	}

	if DATA != "" {
		data, err := base64.StdEncoding.DecodeString(DATA)
		if err != nil {
			return err
		}

		embed := viper.New()
		embed.SetConfigType("yaml")

		if err := embed.ReadConfig(bytes.NewBuffer(data)); err != nil {
			return err
		}

		embedData, err := jsoniter.Marshal(embed.AllSettings())
		if err != nil {
			return err
		}

		if err := jsoniter.Unmarshal(embedData, v); err != nil {
			return err
		}
	}

	return nil
}

func LoadYaml(configFile string, v interface{}) error {
	return Load(configFile, "yaml", v)
}

func LoadJson(configFile string, v interface{}) error {
	return Load(configFile, "json", v)
}

type H map[string]interface{}

func SetDefaults(defaults H) {
	for k, v := range defaults {
		SetDefault(k, v)
	}
}

func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

// load os environments to viper defaults
// with prefix "fox"
// FOX_DB__NAME -> db.name
// FOX_DB__BROKER_ID -> db.broker_id
func AutomaticLoadEnv(prefix string) {
	prefix = strings.ToLower(prefix) + "_"
	for _, env := range os.Environ() {
		items := strings.Split(env, "=")
		k, v := strings.ToLower(items[0]), items[1]
		if v == "" || !strings.HasPrefix(k, prefix) {
			continue
		}

		k = strings.TrimPrefix(k, prefix)
		k = strings.ReplaceAll(k, "__", ".")
		SetDefault(k, v)
	}
}
