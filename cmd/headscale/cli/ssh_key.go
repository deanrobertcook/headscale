package cli

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sshKeysCmd)
	sshKeysCmd.AddCommand(saveSSHKeyCmd)
}

var sshKeysCmd = &cobra.Command{
	Use:     "sshkeys",
	Short:   "Handle SSH keys for users of headscale",
	Aliases: []string{"sshkey", "ssh"},
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var saveSSHKeyCmd = &cobra.Command{
	Use:     "save",
	Short:   "Saves a provided public SSH key for a given user",
	Long:    `TODO`,
	Aliases: []string{"s"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errMissingParameter
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")

		log.Trace().
			Msg("Preparing to save SSH key")

		userName := args[0]
		sshKey := args[1]

		SuccessOutput(userName, sshKey, output)
	},
}
