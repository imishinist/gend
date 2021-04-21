package definition

type Generator struct {
	StaticJoin string   `json:"static_join" yaml:"static_join"`
	Command    string   `json:"command" yaml:"command"`
	Templates  string   `json:"templates" yaml:"templates"`
}
