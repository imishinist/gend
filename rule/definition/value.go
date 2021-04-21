package definition

type Value struct {
	Static    string         `json:"static" yaml:"static"`
	Allowed   []string       `json:"allowed" yaml:"allowed"`
	Range     [2]string      `json:"range" yaml:"range"`
	Generator ValueGenerator `json:"generator" yaml:"generator"`
}

type ValueGenerator struct {
	Comand string `json:"command" yaml:"command"`
}
