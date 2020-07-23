// Copyright 1999-2020. Plesk International GmbH. All rights reserved.

package factory

import (
	"git.plesk.ru/projects/SBX/repos/pleskapp/api"
	"git.plesk.ru/projects/SBX/repos/pleskapp/api/json"
)

func GetDomainManagement(a api.Auth) api.DomainManagement {
	return json.NewDomains(a)
}

func GetFTPUserManagement(a api.Auth) api.FTPManagement {
	return json.NewFTP(a)
}

func GetDatabaseManagement(a api.Auth) api.DatabaseManagement {
	return json.NewDatabases(a)
}

func GetDatabaseUserManagement(a api.Auth) api.DatabaseUserManagement {
	return json.NewDatabases(a)
}

func GetAuthentication(a api.AuthClient) api.Authentication {
	return json.NewAuth(a)
}

func GetServerInfo(a api.Auth) api.Server {
	return json.NewInfo(a)
}
