package funcs

import (
	"strings"
	"text/template"
)

var Map = template.FuncMap{
	"join": strings.Join,
}
