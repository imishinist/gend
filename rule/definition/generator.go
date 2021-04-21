package definition

type Generator struct {
	StaticJoin string `json:"static_join,omitempty" yaml:"static_join"`
	Command    string `json:"command,omitempty" yaml:"command"`
	Templates  string `json:"templates,omitempty" yaml:"templates"`
}
