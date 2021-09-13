package profile

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"promagent/utils"
)

type FileSdConfig struct {
	Files []string `yaml:"files"`
}

func NewFileSdConfig(files ...string) *FileSdConfig {
	return &FileSdConfig{files}
}

type ScrapeConfig struct {
	JobName       string          `yaml:"job_name"`
	StaticConfigs interface{}     `yaml:"static_configs,omitempty"`
	FileSdConfigs []*FileSdConfig `yaml:"file_sd_configs"`
}

type PrometheusConfig struct {
	Global        interface{}     `yaml:"global"`
	Alerting      interface{}     `yaml:"alerting"`
	RuleFiles     []string        `yaml:"rule_files"`
	ScrapeConfigs []*ScrapeConfig `yaml:"scrape_configs"`
}

func NewScrapeConfig(job string) *ScrapeConfig {
	paths := []string{
		fmt.Sprintf("sd/file/%s/*.yaml", job),
		fmt.Sprintf("sd/file/%s/*.json", job),
	}

	return &ScrapeConfig{
		JobName:       job,
		FileSdConfigs: []*FileSdConfig{NewFileSdConfig(paths...)},
	}
}

func writePrometheus(tpl, path string, jobs []*Job) error {
	utils.MKPdir(path)
	file, err := os.Open(tpl)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var config PrometheusConfig
	if err := decoder.Decode(&config); err != nil {
		return err
	}

	for _, job := range jobs {
		scrape := NewScrapeConfig(job.Name)
		config.ScrapeConfigs = append(config.ScrapeConfigs, scrape)
		configPath := filepath.Join(filepath.Dir(path), fmt.Sprintf("sd/file/%s/%s.yaml", job.Name, job.Name))
		writeTarget(configPath, job.Targets)
	}

	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()

	encoder := yaml.NewEncoder(output)
	if err := encoder.Encode(config); err != nil {
		return err
	}

	return nil
}
