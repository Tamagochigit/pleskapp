// Copyright 1999-2020. Plesk International GmbH.

package actions

import (
	"github.com/plesk/pleskapp/plesk/api/factory"
	"github.com/plesk/pleskapp/plesk/config"
	"github.com/plesk/pleskapp/plesk/types"
	"github.com/plesk/pleskapp/plesk/utils"
)

func FindCachedFtpUser(domain types.Domain) *types.FtpUser {
	if len(domain.FTPUsers) != 0 {
		return &domain.FTPUsers[0]
	}
	return nil
}

func FtpUserCreate(host types.Server, domain types.Domain, user *types.FtpUser) (*types.FtpUser, error) {
	if user == nil {
		user = &types.FtpUser{
			Login:    utils.GenUsername(16),
			Password: utils.GenPassword(32),
		}
	}

	api := factory.GetFTPUserManagement(host.GetServerAuth())
	_, err := api.CreateFtpUser(domain.Name, *user)
	if err != nil {
		return nil, err
	}

	domain.FTPUsers = append(domain.FTPUsers, *user)
	config.SetDomain(host, domain)

	return user, nil
}
