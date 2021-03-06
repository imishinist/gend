package generator

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/imishinist/gend/rule/definition"
)

var (
	ErrValueDefinition   = errors.New("value definition invalid")
	ErrFormatInvalid     = errors.New("invalid format")
	ErrCommandDefinition = errors.New("command definition invalid")

	ErrRunCommand      = errors.New("run command failure")
	ErrTemplateInvalid = errors.New("template is invalid")

	ErrGenerator = errors.New("generator is invalid")
)

func Value(ctx context.Context, gtx *Context, rule definition.Rule) (string, error) {
	value := rule.Value
	if value.Static != "" {
		return value.Static, nil
	}

	if value.Allowed != nil && len(value.Allowed) != 0 {
		return enumValue(value), nil
	}

	if value.Range != nil && len(value.Range) >= 2 {
		return rangeValue(value)
	}

	if value.Generator != nil {
		buf := new(bytes.Buffer)
		err := gtx.VGenerator.Generate(ctx, rule.Key, nil, buf)
		if err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	return "", ErrValueDefinition
}

func enumValue(value definition.Value) string {
	n := rand.Intn(len(value.Allowed))
	return value.Allowed[n]
}

func rangeValue(value definition.Value) (string, error) {
	from := value.Range[0]
	to := value.Range[1]
	step := "1"
	if len(value.Range) == 3 {
		step = value.Range[2]
	}

	if strings.Contains(from, ".") {
		ffrom, err := strconv.ParseFloat(from, 64)
		if err != nil {
			return "", fmt.Errorf("failed to parse range.0 as float: %w", err)
		}
		fto, err := strconv.ParseFloat(to, 64)
		if err != nil {
			return "", fmt.Errorf("failed to parse range.1 as float: %w", err)
		}
		fstep, err := strconv.ParseFloat(step, 64)
		if err != nil {
			return "", fmt.Errorf("failed to parse range.1 as float: %w", err)
		}
		return genRangeFloat(ffrom, fto, fstep), nil
	}

	ifrom, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse range.0 as int: %w", err)
	}
	ito, err := strconv.ParseInt(to, 10, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse range.1 as int: %w", err)
	}
	istep, err := strconv.ParseInt(step, 10, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse range.1 as int: %w", err)
	}
	return genRangeInt(ifrom, ito, istep), nil
}

func genRangeFloat(from, to, step float64) string {
	t := rand.Float64()
	ret := t*(to-from) + from
	mod := math.Mod(ret, step)
	return fmt.Sprint(ret - mod)
}

func genRangeInt(from, to, step int64) string {
	diff := to - from
	ret := from + rand.Int63n(diff)
	mod := ret % step
	return strconv.FormatInt(ret-mod, 10)
}

func buildValueGenerator(key string, gen *definition.ValueGenerator) (IGenerator, error) {
	if gen.Bash != "" {
		bash, err := NewBash(gen.Bash)
		if err != nil {
			return nil, err
		}
		return bash, nil
	}

	return nil, ErrCommandDefinition
}
