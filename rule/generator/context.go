package generator

import (
	"context"
	"fmt"
	"io"

	"github.com/imishinist/gend/rule/definition"
)

type Context struct {
	VGenerator cache
	Generator  cache
}

func (c *Context) Close() error {
	c.VGenerator.Close()
	c.Generator.Close()
	return nil
}

func Build(conf definition.Config) (*Context, error) {
	vg := make(cache)
	g := make(cache)

	for _, rule := range conf.Rules {
		if rule.Value.Generator != nil {
			gen, err := buildValueGenerator(rule.Key, rule.Value.Generator)
			if err != nil {
				return nil, err
			}
			vg.Register(rule.Key, gen)
		}

		if rule.Generator == nil {
			return nil, fmt.Errorf("generator: %w", ErrGenerator)
		}
		if rule.Generator.Bash != "" {
			bash, err := NewBash(rule.Generator.Bash)
			if err != nil {
				return nil, err
			}
			g.Register(rule.Key, bash)
		} else if rule.Generator.Templates != "" {
			t, err := NewTemplates(rule.Key, rule.Generator.Templates)
			if err != nil {
				return nil, err
			}
			g.Register(rule.Key, t)
		} else {
			return nil, fmt.Errorf("empty error: %w", ErrGenerator)
		}
	}
	return &Context{
		VGenerator: vg,
		Generator:  g,
	}, nil
}

type cache map[string]IGenerator

func (c *cache) Register(key string, value IGenerator) {
	(*c)[key] = value
}

func (c *cache) Generate(ctx context.Context, key string, env map[string]interface{}, out io.Writer) error {
	if gen, ok := (*c)[key]; ok {
		return gen.Generate(ctx, env, out)
	}
	panic("not registered")
}

func (c *cache) Close() error {
	for _, v := range *c {
		v.Close()
	}
	return nil
}