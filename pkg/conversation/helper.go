package conversation

import (
	"encoding/json"
	"path"
	"qwen-cli/pkg/api"
	"qwen-cli/pkg/util"
)

func WriteToConversationFile(conversationDir string, name string, conversation []api.Message) error {
	targetPath := path.Join(conversationDir, name, "conversation.json")
	data, err := json.Marshal(conversation)
	if err != nil {
		return err
	}
	return util.WriteFileWithNosec(targetPath, data)
}

func ReadFromConversationFile(conversationDir string, name string) ([]api.Message, error) {
	targetPath := path.Join(conversationDir, name, "conversation.json")
	result, err := util.ReadTxtFile(targetPath)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return []api.Message{}, nil
	}
	conversation := []api.Message{}
	err = json.Unmarshal(result, &conversation)
	if err != nil {
		return nil, err
	}
	return conversation, nil
}
