package app

import (
	"encoding/json"
	"fmt"
	"qwen-cli/cmd/qwen/app/config"
	"qwen-cli/cmd/qwen/app/option"
	"qwen-cli/pkg/api"
	"qwen-cli/pkg/conversation"
	"qwen-cli/pkg/qwen"
	"qwen-cli/pkg/util"

	"github.com/spf13/cobra"
)

func CreatePromptCommand(opts *option.GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "prompt",
		Short:   "create a prompt to play with qwen",
		Long:    "create a prompt to play with qwen",
		Example: "",
	}

	cmd.AddCommand(createPromptCommand(opts))
	return cmd
}

func createPromptCommand(opts *option.GlobalOptions) *cobra.Command {
	promptOption := &option.PromptOptions{}
	runCmd := &cobra.Command{
		Use:     "create [prompt]",
		Short:   "create a prompt",
		Long:    "create a prompt",
		Example: renderPromptExample(),
		Args:    cobra.ExactArgs(1),
		RunE:    createPromptRun(opts, promptOption),
	}
	promptOption.BindPromptFlags(runCmd.Flags())
	return runCmd
}

func renderPromptExample() string {
	helpText := `# create a prompt without any conversation context
qwen-cli prompt create "你好"
# create a prompt with conversation context
qwen-cli prompt create "你好" -c c1
# create a prompt from file with conversation context
qwen-cli prompt create -f /tmp/your-prompt -c c1`

	return helpText
}

func preCheckArgs(promptOpts *option.PromptOptions) error {
	if promptOpts == nil {
		return nil
	}

	switch promptOpts.Model {
	case "qwen-turbo":
	case "qwen-plus":
	case "qwen-max":
	case "qwen-max-1201":
	case "qwen-max-longcontext":
	default:
		return fmt.Errorf("model %s is not supported currently support qwen-turbo,qwen-plus,qwen-max,qwen-max-1201,qwen-max-longcontext", promptOpts.Model)
	}

	return nil
}

func createPromptRun(opts *option.GlobalOptions, promptOpts *option.PromptOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		err := preCheckArgs(promptOpts)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// get base api info from config file
		apiCfg, err := config.ReadApiConfigFile(opts.ApiConfigFile)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if apiCfg.ApiKey == "" {
			fmt.Println(fmt.Errorf("api key is empty"))
			return fmt.Errorf("api key is empty")
		}

		if apiCfg.ApiURL == "" {
			fmt.Println(fmt.Errorf("api url is empty"))
			return fmt.Errorf("api url is empty")
		}

		var conversationList []api.Message
		if promptOpts.Conversation != "" {
			conversationList, err = conversation.ReadFromConversationFile(opts.ConversationStorageDir, promptOpts.Conversation)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}

		var parameters *api.ApiParameter
		if promptOpts.EnableParameter {
			cPara, err := config.ReadParameterConfigFile(opts.ParameterConfigFile)
			if err != nil {
				fmt.Println(err)
				return err
			}
			tmpP := api.ApiParameter(*cPara)
			parameters = &tmpP
		}

		prompt := args[0]
		if promptOpts.FilePrompt {
			// if file prompt is set then we read prompt from file
			data, err := util.ReadTxtFile(prompt)
			if err != nil {
				fmt.Println(err)
				return err
			}
			prompt = string(data)
		}

		// always append user prompt to conversation list
		conversationList = append(conversationList, api.Message{
			Content: prompt,
			Role:    "user",
		})

		input := &api.Input{
			Conversation: conversationList,
		}

		response, err := qwen.SendToQwen(apiCfg, input, parameters, promptOpts.Model, promptOpts.EnableSSE, promptOpts.Timeout)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if promptOpts.Conversation != "" {
			// save conversation
			// fill the conversation with choices if we send down the parameter.result_format = "message"
			for _, choice := range response.Output.Choices {
				conversationList = append(conversationList, choice.Message)
			}

			// always check if response.output.Text if we didn't send parameter.result_format = "test"
			if response.Output.Text != "" {
				conversationList = append(conversationList, api.Message{
					Content: response.Output.Text,
					Role:    "system", // now system is qwen
				})
			}
			// update conversation file
			err = conversation.WriteToConversationFile(opts.ConversationStorageDir, promptOpts.Conversation, conversationList)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}

		// print response
		if promptOpts.RawOutput {
			data, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(string(data))
		} else {
			if promptOpts.EnableParameter && parameters.ResultFormat == "message" {
				// print choices
				for _, choice := range response.Output.Choices {
					fmt.Println(choice.Message.Content)
				}
			} else {
				// print text
				fmt.Println(response.Output.Text)
			}
		}

		return nil
	}
}
