package option

import (
	"path"

	"github.com/spf13/pflag"
)

type GlobalOptions struct {
	ApiConfigFile          string
	ParameterConfigFile    string
	ConversationStorageDir string
}

type PromptOptions struct {
	Model           string
	Conversation    string
	RawOutput       bool
	FilePrompt      bool
	EnableSSE       bool
	EnableParameter bool
	Timeout         int
}

func (option *GlobalOptions) BindFlags(fs *pflag.FlagSet, homeDir string) {
	fs.StringVarP(&option.ParameterConfigFile, "parameter-config", "p", path.Join(homeDir, ".qwen-cli", "parameter-config.json"), "parameter config file location")
	fs.StringVarP(&option.ApiConfigFile, "api-config", "a", path.Join(homeDir, ".qwen-cli", "api-config.json"), "api config file location")
	fs.StringVarP(&option.ConversationStorageDir, "conversation-storage-dir", "s", path.Join(homeDir, ".qwen-cli", "conversation"), "conversation storage directory")
}

func (option *PromptOptions) BindPromptFlags(fs *pflag.FlagSet) {
	// by default we use qwen-max model also we support qwen-turbo qwen-plus qwen-max qwen-max-1201,qwen-max-longcontext if you care enough to pay
	// see https://dashscope.console.aliyun.com/billing for detail
	fs.StringVarP(&option.Model, "model", "m", "qwen-max", "model to use")
	fs.StringVarP(&option.Conversation, "conversation", "c", "", "conversation to use if not set then no conversation context will be sent to qwen api")
	fs.BoolVarP(&option.RawOutput, "raw-output", "r", false, "raw output json from qwen api")
	fs.BoolVarP(&option.FilePrompt, "file-prompt", "f", false, "prompt from file you can put you prompt in a file then is will be read from cli")
	fs.BoolVarP(&option.EnableSSE, "enable-sse", "e", false, "enable http SSE for prompt")
	fs.BoolVarP(&option.EnableParameter, "enable-parameters", "t", false, "enable parameters for prompt")
	// for much complicated conversation you may need more time to wait for response
	fs.IntVarP(&option.Timeout, "timeout", "w", 30, "how many seconds to wait for response from qwen api")
}

func NewDefaultOption() *GlobalOptions {
	return &GlobalOptions{}
}
