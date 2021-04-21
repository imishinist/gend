package definition

import "github.com/imishinist/gend/rule/definition"

type Rule struct {
	Key   string           `json:"key" yaml:"key"`
	Value definition.Value `json:"value" yaml:"value"`
}

func FromDefinitionRule(rule definition.Rule) Rule {
	return Rule{
		Key:   rule.Key,
		Value: rule.Value,
	}
}
