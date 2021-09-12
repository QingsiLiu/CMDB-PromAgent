package config

type AgentConfig struct {
	UUID         string        `mapstructure:"_"`
	Addr         string        `mapstructure:"addr"`
	ServerConfig *ServerConfig `mapstructure:"server"`
	LogConfig    *LogConfig    `mapstructure:"log"`
}

type ServerConfig struct {
	Addr  string `mapstructure:"addr"`
	Token string `mapstructure:"token"`
}

type LogConfig struct {
	Filename   string `mapstructure:"filename"`
	Maxsize    int    `mapstructure:"maxsize"`
	Maxbackups int    `mapstructure:"maxbackups"`
	Compress   bool   `mapstructure:"compress"`
}
