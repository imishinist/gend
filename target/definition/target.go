package definition

type TargetKV struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

func NewTarget(key string) *TargetKV {
	return &TargetKV{
		Key:    key,
		Values: make([]string, 0),
	}
}

func (t *TargetKV) Add(value string) {
	t.Values = append(t.Values, value)
}
