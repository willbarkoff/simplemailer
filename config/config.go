package config

import "github.com/BurntSushi/toml"

// Config represents a configuration
type Config struct {
	FromAddress  string
	FromDisplay  string
	SendDomain   string
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

//Configure reads a configuration file and stores the values in the Config variable.
func Configure(file string) Config {
	config := Config{}
	_, err := toml.DecodeFile(file, &config)
	if err != nil {
		panic(err)
	}
	return config
}
