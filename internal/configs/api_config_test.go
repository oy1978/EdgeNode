package configs_test

import (
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/oy1978/EdgeNode/internal/configs"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestLoadAPIConfig(t *testing.T) {
	config, err := configs.LoadAPIConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", config)

	configData, err := yaml.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(configData))
}
