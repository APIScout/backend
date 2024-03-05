package structs

type Config struct {
	Backend Backend `yaml:"backend"`
	Mongo   Mongo   `yaml:"mongodb"`
	Elastic Elastic `yaml:"elastic"`
}

type Backend struct {
	Port int `yaml:"port"`
}

type Mongo struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Elastic struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
