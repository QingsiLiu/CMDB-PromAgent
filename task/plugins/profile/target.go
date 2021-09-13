package profile

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"promagent/utils"
)

type TargetConfig struct {
	Targets []string `yaml:"targets"`
}

func NewTargetConfig(targets ...string) *TargetConfig {
	return &TargetConfig{targets}
}

func writeTarget(path string, targets []*Target) error {
	utils.MKPdir(path)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	addrs := make([]string, len(targets))
	for i, target := range targets {
		fmt.Println("----------", target.Addr)
		addrs[i] = target.Addr
	}

	encoder := yaml.NewEncoder(file)
	config := []*TargetConfig{NewTargetConfig(addrs...)}
	return encoder.Encode(&config)
}
