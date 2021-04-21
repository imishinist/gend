package definition

type Value struct {
	Static    string          `json:"static,omitempty" yaml:"static"`
	Allowed   []string        `json:"allowed,omitempty" yaml:"allowed"`
	Range     [2]string       `json:"range,omitempty" yaml:"range"`
	Generator *ValueGenerator `json:"generator,omitempty" yaml:"generator"`
}

type ValueGenerator struct {
	Bash string `json:"bash,omitempty" yaml:"bash"`
}
