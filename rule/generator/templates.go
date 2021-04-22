package generator

import (
	"bytes"
	"context"
	"text/template"

	"github.com/imishinist/gend/funcs"
)

// compile time implementation check
var _ IGenerator = (*Templates)(nil)

type Templates struct {
	inner *template.Template
}

func NewTemplates(key string, templatestr string) (*Templates, error) {
	t, err := template.New(key).Funcs(funcs.Map).Parse(templatestr)
	if err != nil {
		return nil, err
	}
	return &Templates{
		inner: t,
	}, nil
}

func (t *Templates) Generate(ctx context.Context, env map[string]interface{}) (string, error) {
	buf := new(bytes.Buffer)
	if err := t.inner.Execute(buf, env); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (t *Templates) Close() error {
	return nil
}
