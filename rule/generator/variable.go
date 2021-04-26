package generator

type Variable struct {
	Name  string
	Value string
}

type Variables []Variable

func (v *Variables) ForTemplate() map[string]interface{} {
	ret := make(map[string]interface{})
	for _, vr := range *v {
		ret[vr.Name] = vr.Value
	}
	return ret
}
