package generator

import "github.com/imishinist/gend/rule/definition"

func Use(use definition.Use) bool {
	if use.Percent >= 1e-10 {
		return Percent(use.Percent)
	}
	return use.Static
}
