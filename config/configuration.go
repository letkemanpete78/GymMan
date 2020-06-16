package config

// Configuration is the data structure for the config.yml file
type Configuration struct {
	Server      ServerConfiguration
	Database    DatabaseConfiguration
	APIVersions APIVersionsConfiguration
}
