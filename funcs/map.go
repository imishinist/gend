package funcs

import (
	"strings"
	"text/template"
)

var Map = template.FuncMap{
	"join":     strings.Join,
	"trimjoin": trimJoin,
	"trim":     strings.TrimSpace,
}

func trimJoin(elems []string, sep string) string {
	trimed := make([]string, 0, len(elems))
	for _, e := range elems {
		trimed = append(trimed, strings.TrimSpace(e))
	}
	return strings.Join(trimed, sep)
}
