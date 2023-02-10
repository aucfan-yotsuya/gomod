package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) { m.Run() }
func TestLoadConfig(t *testing.T) {
	type (
		childSection struct {
			Key1 string `toml:"key1"`
			Key2 string `toml:"key2"`
		}
		Section struct {
			Key string `toml:"key"`
		}
		Config struct {
			Section  Section        `toml:"section"`
			Sections []childSection `toml:"sections"`
		}
	)
	var c Config
	assert.NoError(t, LoadConfig("config_test.toml", &c))
	assert.Equal(t, c, Config{
		Section: Section{"value"},
		Sections: []childSection{
			{"value1", "value2"},
			{"value1", "value2"},
		},
	})
	t.Log(c)
}
