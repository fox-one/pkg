package config

import (
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
)

func TestAutomaticLoadEnv(t *testing.T) {
	os.Setenv("PKG_NAME", "pkg")
	os.Setenv("PKG_author", "yiplee")
	os.Setenv("PKG_DB__BROKER_ID", "localhost")
	os.Setenv("PKG_DB__PASSWORD", "PWD")

	AutomaticLoadEnv("pkg")
	assert.Equal(t, viper.GetString("name"), "pkg")
	assert.Equal(t, viper.GetString("author"), "yiplee")
	assert.Equal(t, viper.GetString("db.broker_id"), "localhost")
	assert.Equal(t, viper.GetString("db.password"), "PWD")
}
