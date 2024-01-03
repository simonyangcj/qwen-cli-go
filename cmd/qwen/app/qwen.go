package app

import (
	"fmt"

	"qwen-cli/cmd/qwen/app/option"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

func NewCommand(version string) *cobra.Command {
	opts := &option.GlobalOptions{}
	cmd := &cobra.Command{
		Use:     "qwen-cli",
		Long:    "qwen-cli is command line tool to have fun with qwen model",
		Example: figure.NewColorFigure("Fun From Maas Team", "", "green", true).String(),
	}

	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Print version and exit",
		Long:    "version subcommand will print version and exit",
		Example: "qwen-cli version",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("version:", version)
		},
	}

	opts.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(versionCmd, CreateConversationCommand(opts.ConversationStorageDir), CreatePromptCommand(opts))

	return cmd
}
