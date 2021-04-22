package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"text/template"

	"github.com/imishinist/gend/funcs"
	"github.com/imishinist/gend/rule/definition"
	targetdef "github.com/imishinist/gend/target/definition"
)

func Generator(ctx context.Context, target *targetdef.TargetKV, rule definition.Rule) (string, error) {
	if rule.Use != nil && !Use(*rule.Use) {
		return "", errors.New("don't use")
	}
	if rule.Generator != nil {
		if rule.Generator.Bash != "" {
			cmd, clean := runAsBash(ctx, rule.Generator.Bash, map[string]string{
				"key":    target.Key,
				"values": strings.Join(target.Values, ","),
			})
			defer clean()
			output, err := cmd.Output()
			if err != nil {
				return "", ErrRunCommand
			}
			return string(output), nil
		}
		if rule.Generator.Templates != "" {
			t, err := template.New(rule.Key).Funcs(funcs.Map).Parse(rule.Generator.Templates)
			if err != nil {
				return "", ErrTemplateInvalid
			}
			buf := new(bytes.Buffer)
			if err := t.Execute(buf, map[string]interface{}{
				"key":    target.Key,
				"values": target.Values,
			}); err != nil {
				return "", ErrTemplateInvalid
			}
			return buf.String(), nil
		}
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
