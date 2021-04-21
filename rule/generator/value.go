package generator

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/imishinist/gend/rule/definition"
)

var (
	ErrValueDefinition = errors.New("value definition invalid")
	ErrFormatInvalid   = errors.New("invalid format")
)

func Value(ctx context.Context, value definition.Value) (string, error) {
	if value.Static != "" {
		return value.Static, nil
	}

	if value.Allowed != nil && len(value.Allowed) != 0 {
		return enumValue(value), nil
	}

	if value.Range[0] != "" && value.Range[1] != "" {
		return rangeValue(value)
	}

	if value.Generator != nil {
		// TODO
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

	if strings.Contains(from, ".") {
		ffrom, err := strconv.ParseFloat(from, 64)
		if err != nil {
			return "", fmt.Errorf("failed to parse range.0 as float: %w", err)
		}
		fto, err := strconv.ParseFloat(to, 64)
		if err != nil {
			return "", fmt.Errorf("failed to parse range.1 as float: %w", err)
		}
		return genRangeFloat(ffrom, fto), nil
	}

	ifrom, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse range.0 as int: %w", err)
	}
	ito, err := strconv.ParseInt(to, 10, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse range.1 as int: %w", err)
	}
	return genRangeInt(ifrom, ito), nil
}

func genRangeFloat(from, to float64) string {
	t := rand.Float64()
	ret := t*(to-from) + from
	return fmt.Sprint(ret)
}

func genRangeInt(from, to int64) string {
	diff := to - from
	ret := from + rand.Int63n(diff)
	return strconv.FormatInt(ret, 10)
}