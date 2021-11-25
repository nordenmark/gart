package config

type ImageParameter struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	DefaultValue interface{} `json:"defaultValue"`
	Min          float32     `json:"min"`
	Max          float32     `json:"max"`
	Step         float32     `json:"step"`
}

type ImageObject struct {
	Name       string           `json:"name"`
	Parameters []ImageParameter `json:"parameters"`
}

type Configuration interface {
	Parameters() []ImageParameter
}
