package config

import (
	"bytes"
	"encoding/base64"
	"os"

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

	if DATA != "" {
		data, err := base64.StdEncoding.DecodeString(DATA)
		if err != nil {
			return err
		}

		if err := viper.MergeConfig(bytes.NewBuffer(data)); err != nil {
			return err
		}
	}

	data, err := jsoniter.Marshal(viper.AllSettings())
	if err != nil {
		return err
	}

	return jsoniter.Unmarshal(data, v)
}

func LoadYaml(configFile string, v interface{}) error {
	return Load(configFile, "yaml", v)
}

func LoadJson(configFile string, v interface{}) error {
	return Load(configFile, "json", v)
}
