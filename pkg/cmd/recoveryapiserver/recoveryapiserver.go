package recoveryapiserver

import (
	"github.com/spf13/cobra"
)

func NewRecoveryAPIServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "recovery-apiserver",
	}

	cmd.AddCommand(NewCreateCommand())
	cmd.AddCommand(NewDestroyCommand())

	return cmd
}
