// Copyright 1999-2020. Plesk International GmbH.

package cmd

import (
	"git.plesk.ru/projects/SBX/repos/pleskapp/actions"
	"git.plesk.ru/projects/SBX/repos/pleskapp/config"
	"git.plesk.ru/projects/SBX/repos/pleskapp/locales"
	"git.plesk.ru/projects/SBX/repos/pleskapp/utils"
	"github.com/spf13/cobra"
)

var reauthCmd = &cobra.Command{
	Use:   locales.L.Get("server.reauth.cmd"),
	Short: locales.L.Get("server.reauth.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := config.GetServer(args[0])
		if err != nil {
			return err
		}

		cmd.SilenceUsage = true
		return utils.Log.PrintSuccessOrError("server.reauth.success", nil, actions.ServerReauth(*server))
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	ServersCmd.AddCommand(reauthCmd)
}
