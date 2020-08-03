// Copyright 1999-2020. Plesk International GmbH.

package actions

import (
	"os"
	"path/filepath"
	"strings"

	"git.plesk.ru/projects/SBX/repos/pleskapp/api/factory"
	"git.plesk.ru/projects/SBX/repos/pleskapp/locales"
	"git.plesk.ru/projects/SBX/repos/pleskapp/types"
	"git.plesk.ru/projects/SBX/repos/pleskapp/upload"
	"git.plesk.ru/projects/SBX/repos/pleskapp/utils"
)

func getPrereq(host types.Server, domain types.Domain) (*types.FtpUser, *string, *string, error) {
	var fullpath string
	var docroot string
	// TODO: Ideally, there should be no need to do this
	{
		api := factory.GetDomainManagement(host.GetServerAuth())
		i, err := api.GetDomain(domain.Name)
		if err != nil {
			return nil, nil, nil, err
		}

		fullpath = i.WWWRoot
		parts := strings.Split(i.WWWRoot, domain.Name)
		docroot = parts[len(parts)-1]
	}

	ftp := FindCachedFtpUser(domain)
	if ftp == nil {
		var err error
		ftp, err = FtpUserCreate(host, domain, nil)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return ftp, &docroot, &fullpath, nil
}

func UploadFileToRoot(host types.Server, domain types.Domain, ovw bool, file string) (string, error) {
	ftp, docroot, path, err := getPrereq(host, domain)
	if err != nil {
		return "", err
	}

	connection, err := upload.Connect(*ftp, host.Host, "/")
	if err != nil {
		return "", err
	}

	cr, f := utils.GetClientRootName(file)
	return strings.Split(*path, *docroot)[0], connection.UploadFile(cr, "/", f, true, host.Info.IsWindows)
}

func UploadFile(host types.Server, domain types.Domain, ovw bool, file string) error {
	ftp, docroot, _, err := getPrereq(host, domain)
	if err != nil {
		return err
	}

	connection, err := upload.Connect(*ftp, host.Host, *docroot)
	if err != nil {
		return err
	}

	cr, f := utils.GetClientRootName(file)
	return connection.UploadFile(cr, *docroot, f, ovw, host.Info.IsWindows)
}

func UploadDirectory(host types.Server, domain types.Domain, ovw bool, dry bool, dir string, root *string) error {
	ftp, docroot, _, err := getPrereq(host, domain)
	if err != nil {
		return err
	}

	connection, err := upload.Connect(*ftp, host.Host, *docroot)
	if err != nil {
		return err
	}

	serverPath := root
	if serverPath == nil {
		serverPath = docroot
	}

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		pathPart := strings.TrimPrefix(path, dir)
		if pathPart == "." {
			return nil
		}

		if dry {
			utils.Log.Print(locales.L.Get("upload.dry.run.upload", dir+"/"+pathPart, *serverPath+"/"+pathPart))
			return nil
		}

		return connection.UploadFile(dir+"/", *serverPath+"/", pathPart, ovw, host.Info.IsWindows)
	})
}
