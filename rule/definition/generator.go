package definition

type Generator struct {
	Bash      string `json:"bash,omitempty" yaml:"bash"`
	Templates string `json:"templates,omitempty" yaml:"templates"`
}
