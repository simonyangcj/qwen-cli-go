package app

import (
	"fmt"
	"path"
	"qwen-cli/pkg/util"

	"github.com/spf13/cobra"
)

func CreateConversationCommand(conversationDir string) *cobra.Command {
	runCmd := &cobra.Command{
		Use:     "conversation",
		Short:   "manager conversation",
		Long:    "manager conversation",
		Example: "",
	}

	runCmd.AddCommand(createConversationCommand(conversationDir))
	runCmd.AddCommand(deleteConversationCommand(conversationDir))
	runCmd.AddCommand(listConversationCommand(conversationDir))
	runCmd.AddCommand(showConversationCommand(conversationDir))
	return runCmd
}

func createConversationCommand(conversationDir string) *cobra.Command {
	createCMD := &cobra.Command{
		Use:     "create [conversation name]",
		Short:   "create a conversation",
		Long:    "create a conversation",
		Example: "qwen-cli conversation create c1",
		Args:    cobra.ExactArgs(1),
		RunE:    createConversationRun(conversationDir),
	}
	return createCMD
}

func createConversationRun(conversationDir string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if conversationDir == "" {
			return fmt.Errorf("conversationDir is empty")
		}
		targetDir := path.Join(conversationDir, args[0])
		result := util.PathExist(targetDir)
		errMsg := ""
		if result {
			errMsg = fmt.Sprintf("failed to create conversation %s due to already exists", args[0])
			return fmt.Errorf(errMsg)
		}

		err := util.CreateDirIfNotExists(targetDir)
		if err != nil {
			errMsg = fmt.Sprintf("failed to create conversation %s due to error: %s", args[0], err)
			return fmt.Errorf(errMsg)
		}

		err = util.WriteFileWithNosec(path.Join(targetDir, "conversation.json"), []byte("{}"))
		if err != nil {
			errMsg = fmt.Sprintf("failed to write conversation.json %s due to error: %s", args[0], err)
			return fmt.Errorf(errMsg)
		}

		listConversationRun(conversationDir)(nil, nil)

		return nil
	}
}

func deleteConversationCommand(conversationDir string) *cobra.Command {
	deleteCMD := &cobra.Command{
		Use:     "delete",
		Short:   "delete a conversation",
		Long:    "delete a conversation",
		Example: "qwen-cli conversation delete c1",
		Args:    cobra.MaximumNArgs(1),
		RunE:    deleteConversationRun(conversationDir),
	}
	return deleteCMD
}

func deleteConversationRun(conversationDir string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if conversationDir == "" {
			return fmt.Errorf("conversationDir is empty")
		}
		targetDir := path.Join(conversationDir, args[0])

		err := util.DeleteDirs(targetDir)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("failed to delete conversation %s due to error: %s", args[0], err))
		}

		return nil
	}
}

func listConversationCommand(conversationDir string) *cobra.Command {
	listCMD := &cobra.Command{
		Use:     "list",
		Short:   "list all conversations",
		Long:    "list all conversations",
		Example: "qwen-cli conversation list",
		Args:    cobra.MaximumNArgs(0),
		RunE:    listConversationRun(conversationDir),
	}
	return listCMD
}

func listConversationRun(conversationDir string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if conversationDir == "" {
			return fmt.Errorf("conversationDir is empty")
		}
		result, err := util.ListDirs(conversationDir)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("failed to list conversations due to error: %s", err))
		}

		if len(result) == 0 {
			fmt.Println("no conversation exists")
			return nil
		}

		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	}
}

func showConversationCommand(conversationDir string) *cobra.Command {
	showCMD := &cobra.Command{
		Use:     "show",
		Short:   "show a conversation",
		Long:    "show a conversation",
		Example: "qwen-cli conversation show c1",
		Args:    cobra.MaximumNArgs(1),
		RunE:    showConversationRun(conversationDir),
	}
	return showCMD
}

func showConversationRun(conversationDir string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if conversationDir == "" {
			return fmt.Errorf("conversationDir is empty")
		}
		targetDir := path.Join(conversationDir, args[0])

		result := util.PathExist(targetDir)
		if !result {
			return fmt.Errorf(fmt.Sprintf("failed to show conversation %s due to not exists", args[0]))
		}

		data, err := util.ReadTxtFile(path.Join(targetDir, "conversation.json"))
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("failed to show conversation %s due to error: %s", args[0], err))
		}
		fmt.Println(string(data))
		return nil
	}
}
