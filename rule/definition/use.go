package definition

type Use struct {
	Static  bool    `json:"static"`
	Percent float64 `json:"percent,omitempty" yaml:"percent"`
}
