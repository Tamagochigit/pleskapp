// Copyright 1999-2020. Plesk International GmbH. All rights reserved.

package cmd

import (
	"git.plesk.ru/~abashurov/pleskapp/actions"
	"git.plesk.ru/~abashurov/pleskapp/config"
	"git.plesk.ru/~abashurov/pleskapp/locales"
	"git.plesk.ru/~abashurov/pleskapp/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   locales.L.Get("database.delete.cmd"),
	Short: locales.L.Get("database.delete.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := config.GetServer(args[0])
		if err != nil {
			return err
		}

		cmd.SilenceUsage = true
		return utils.Log.PrintSuccessOrError("database.delete.success", nil, actions.DatabaseDelete(*server, args[1]))
	},
	Args: cobra.ExactArgs(2),
}

func init() {
	DatabasesCmd.AddCommand(deleteCmd)
}
