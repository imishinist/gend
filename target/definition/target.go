package definition

type Target map[string][]string

func New() Target {
	return make(map[string][]string)
}

func (t Target) Add(key string, value string) {
	if _, ok := t[key]; !ok {
		t[key] = []string{value}
		return
	}
	t[key] = append(t[key], value)
}
