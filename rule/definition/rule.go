package definition

type Rule struct {
	Key       string    `json:"key" yaml:"key"`
	Value     Value     `json:"value" yaml:"value"`
	Length    Length    `json:"length" yaml:"length"`
	Generator Generator `json:"generator" yaml:"generator"`
}
