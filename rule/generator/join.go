package generator

import (
	"context"
	"io"
	"strings"
)

type Join struct {
	key string
	sep string
}

func NewJoin(key, sep string) *Join {
	return &Join{
		key: key,
		sep: sep,
	}
}

func (j *Join) Generate(ctx context.Context, env map[string]interface{}, out io.Writer) error {
	if _, ok := env[j.key]; !ok {
		return nil
	}
	v := env[j.key]
	switch v := v.(type) {
	case []string:
		if _, err := io.WriteString(out, strings.Join(v, j.sep)); err != nil {
			return err
		}
	}

	return nil
}

func (j *Join) Close() error {
	return nil
}
