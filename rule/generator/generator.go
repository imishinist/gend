package generator

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/imishinist/gend/rule/definition"
	targetdef "github.com/imishinist/gend/target/definition"
)

type IGenerator interface {
	Generate(ctx context.Context, env map[string]interface{}, out io.Writer) error
	io.Closer
}

func KVGenerator(ctx context.Context, gtx *Context, target *targetdef.TargetKV, rule definition.Rule, out io.Writer) error {
	if rule.Use != nil && !Use(*rule.Use) {
		return errors.New("don't use")
	}
	if rule.Generator != nil {
		if err := gtx.KVGenerator.Generate(ctx, rule.Key, map[string]interface{}{
			"key":    target.Key,
			"values": target.Values,
		}, out); err != nil {
			return err
		}
		return nil
	}

	// default: json encoder
	tmp := map[string][]string{
		target.Key: target.Values,
	}
	if err := json.NewEncoder(out).Encode(tmp); err != nil {
		return err
	}
	return nil
}

func Generator(ctx context.Context, gtx *Context, items []string, out io.Writer) error {
	if err := gtx.Generator.Generate(ctx, "main", map[string]interface{}{
		"items": items,
	}, out); err != nil {
		return err
	}
	return nil
}
