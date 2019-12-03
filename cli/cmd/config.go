package cmd

import (
	"encoding/json"
)

type Config struct {
	ByFile        bool    `json:"byFile"`
	MinSupport    float64 `json:"minSupport"`
	MinConfidence float64 `json:"minConfidence"`
	MinLift       float64 `json:"minLift"`
	MaxLength     int     `json:"maxLength"`
}

func NewConfig(body []byte) *Config {
	config := &Config{}
	json.Unmarshal(body, config)
	return config
}
