package definition

type Length struct {
	Static     int       `json:"static" yaml:"static"`
	Occurrence []float64 `json:"occurrence" yaml:"occurrence"`
}
