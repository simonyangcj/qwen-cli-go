package config

import (
	"encoding/json"
	"qwen-cli/pkg/util"
)

type ParameterConfig struct {
	ResultFormat      string   `json:"result_format,omitempty"`
	Seed              int      `json:"seed,omitempty"`
	MaxTokens         int      `json:"max_tokens,omitempty"`
	TopP              float32  `json:"top_p,omitempty"`
	TopK              int      `json:"top_k,omitempty"`
	RepetitionPenalty float32  `json:"repetition_penalty,omitempty"`
	Temperature       float32  `json:"temperature,omitempty"`
	StopTokenId       []int    `json:"stop,omitempty"`
	StopToken         []string `json:"stop,omitempty"`
	EnableSearch      bool     `json:"enable_search,omitempty"`
	IncrementalOutput bool     `json:"incremental_output,omitempty"`
}

type ApiConfig struct {
	ApiURL string `json:"api_url"`
	ApiKey string `json:"api_key"`
}

func DefaultParameterConfig() *ParameterConfig {
	return &ParameterConfig{
		ResultFormat:      "text",
		Seed:              65535,
		MaxTokens:         1500,
		TopP:              0.8,
		TopK:              50,
		RepetitionPenalty: 1.1,
		Temperature:       1.0,
		StopTokenId:       []int{},
		StopToken:         []string{},
		EnableSearch:      false,
		IncrementalOutput: false,
	}
}

func WriteParameterConfigFile(pathName string, data []byte) error {
	// #nosec G306, Expect WriteFile permissions to be 0600 or less
	return util.WriteFileWithNosec(pathName, data)
}

func WriteApiConfigFile(pathName string, data []byte) error {
	// #nosec G306, Expect WriteFile permissions to be 0600 or less
	return util.WriteFileWithNosec(pathName, data)
}

func ReadParameterConfigFile(pathName string) (*ParameterConfig, error) {
	result, err := util.ReadTxtFile(pathName)
	if err != nil {
		return nil, err
	}
	pConfig := &ParameterConfig{}
	err = json.Unmarshal(result, pConfig)
	if err != nil {
		return nil, err
	}
	return pConfig, nil
}

func ReadApiConfigFile(pathName string) (*ApiConfig, error) {
	result, err := util.ReadTxtFile(pathName)
	if err != nil {
		return nil, err
	}
	aConfig := &ApiConfig{}
	err = json.Unmarshal(result, aConfig)
	if err != nil {
		return nil, err
	}
	return aConfig, nil
}
