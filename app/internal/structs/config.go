package structs

// Config - used to store all configurations needed by the backend.
type Config struct {
	Backend BackendConfig `yaml:"backend"`
	Mongo   MongoConfig   `yaml:"mongodb"`
	Elastic ElasticConfig `yaml:"elastic"`
}

// BackendConfig - all backend-related configurations.
type BackendConfig struct {
	Port int `yaml:"port"`
}

// MongoConfig - all mongodb-related configurations.
type MongoConfig struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// ElasticConfig - all elasticsearch-related configurations.
type ElasticConfig struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
