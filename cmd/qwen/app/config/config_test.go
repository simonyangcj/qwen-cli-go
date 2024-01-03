package config

import (
	"encoding/json"
	"qwen-cli/pkg/util"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("qwen-cli-go config test", func() {
	pConfigPath := "/tmp/p-config.yaml"
	aConfigPath := "/tmp/a-config.yaml"
	var pConfig *ParameterConfig
	var aConfig *ApiConfig
	BeforeEach(func() {
		pConfig = DefaultParameterConfig()
		aConfig = &ApiConfig{
			ApiURL: "https://api.qwen.ai",
			ApiKey: "test-key",
		}
		result, err := json.Marshal(pConfig)
		Expect(err).To(BeNil())
		err = util.WriteFileWithNosec(pConfigPath, result)
		Expect(err).To(BeNil())
		result, err = json.Marshal(aConfig)
		Expect(err).To(BeNil())
		err = util.WriteFileWithNosec(aConfigPath, result)
		Expect(err).To(BeNil())
	})

	Describe("read parameter config from file", func() {
		It("Should be no error", func() {
			result, err := ReadParameterConfigFile(pConfigPath)
			Expect(err).To(BeNil())
			Expect(result.EnableSearch).To(Equal(pConfig.EnableSearch))
			Expect(result.IncrementalOutput).To(Equal(pConfig.IncrementalOutput))
			Expect(result.MaxTokens).To(Equal(pConfig.MaxTokens))
			Expect(result.RepetitionPenalty).To(Equal(pConfig.RepetitionPenalty))
			Expect(result.ResultFormat).To(Equal(pConfig.ResultFormat))
			Expect(result.Seed).To(Equal(pConfig.Seed))
		})
	})

	Describe("read api config from file", func() {
		It("Should be no error", func() {
			result, err := ReadApiConfigFile(aConfigPath)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(aConfig))
		})
	})
})
