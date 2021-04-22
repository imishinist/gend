package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/imishinist/gend/rule/definition"
	targetdef "github.com/imishinist/gend/target/definition"
)

type IGenerator interface {
	Generate(ctx context.Context, env map[string]interface{}) (string, error)
	io.Closer
}

func Generator(ctx context.Context, gtx *Context, target *targetdef.TargetKV, rule definition.Rule) (string, error) {
	if rule.Use != nil && !Use(*rule.Use) {
		return "", errors.New("don't use")
	}
	if rule.Generator != nil {
		res, err := gtx.Generator.Generate(ctx, rule.Key, map[string]interface{}{
			"key":    target.Key,
			"values": target.Values,
		})
		if err != nil {
			return "", err
		}
		return res, nil
	}

	// default: json encoder
	tmp := map[string][]string{
		target.Key: target.Values,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(tmp); err != nil {
		return "", err
	}
	return buf.String(), nil
}
