package definition

type Length struct {
	Static     int       `json:"static,omitempty" yaml:"static"`
	Occurrence []float64 `json:"occurrence,omitempty" yaml:"occurrence"`
}
