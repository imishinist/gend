package definition

type Config struct {
	Rules     []Rule     `json:"rules" yaml:"rules"`
	Generator *Generator `json:"generator,omitempty" yaml:"generator"`
}
