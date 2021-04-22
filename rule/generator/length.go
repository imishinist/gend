package generator

import (
	"errors"
	"math/rand"

	"github.com/imishinist/gend/rule/definition"
)

var (
	ErrLengthDefinition = errors.New("length definition invalid")
)

func Length(length definition.Length) (int, error) {
	if length.Static != 0 {
		return length.Static, nil
	}

	if length.Occurrence != nil && len(length.Occurrence) != 0 {
		return lengthOccurrence(length), nil
	}

	// default => static: 1
	return 1, nil
}

func lengthOccurrence(length definition.Length) int {
	total := 0.0
	for _, o := range length.Occurrence {
		total += o
	}

	// [0, 3, 3, 6]
	// sum = 6
	// points:
	//   [0, 0, 0.25, 0.5, 1.0]
	points := make([]float64, 0, len(length.Occurrence))
	points = append(points, 0.0)
	prog := 0.0
	for _, o := range length.Occurrence {
		prog += o
		points = append(points, prog/total)
	}

	t := rand.Float64()
	for i := 1; i < len(points); i++ {
		if points[i-1] <= t && t < points[i] {
			return i - 1
		}
	}
	return 0
}
