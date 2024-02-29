package structs


type Config struct {
	Backend struct {
		Port int `yaml:"port"`
	}`yaml:"backend"`
	Elastic struct {
		Host string `yaml:"host"`
		Port int `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
	}`yaml:"elastic"`
}
