package modules

import (
	"bytes"
	"errors"
	"os"
	"strings"

	"github.com/lmorg/murex/builtins/core/httpclient"
	"github.com/lmorg/murex/config/profile"
	profilepaths "github.com/lmorg/murex/config/profile/paths"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/cd"
	"github.com/lmorg/readline/v4"
)

func packageDirExists(pack string) error {
	_, err := os.Stat(pack)
	if os.IsNotExist(err) {
		return nil
	}

	return errors.New("a file or directory already exists with that package name")
}

func importModules(p *lang.Process) error {
	modulePath := profilepaths.ModulePath()
	path, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	if path == modulePath+profile.PackagesFile {
		return errors.New("you cannot import the same file as the master packages.json file")
	}

	if utils.IsURL(path) {
		resp, err := httpclient.Request(p.Context, "GET", path, nil, p.Config, true)
		if err != nil {
			return err
		}

		f, err := utils.NewTempFile(resp.Body, "_package.json")
		if err != nil {
			return err
		}

		path = f.FileName

		defer f.Close()
	}

	importDb, err := readPackagesFile(path)
	if err != nil {
		return err
	}

	db, err := readPackagesFile(modulePath + profile.PackagesFile)
	if err != nil {
		return err
	}

	for i := range importDb {
		err = cd.Chdir(p, modulePath)
		if err != nil {
			write(p, "{RED}%s{RESET}", err.Error())
			continue
		}

		p.Stdout.Writeln(bytes.Repeat([]byte{'-'}, readline.GetTermWidth()))
		write(p, "Importing '%s'....", importDb[i].Package)
		err = packageDirExists(importDb[i].Package)
		if err != nil {
			write(p, "{RED}%s{RESET}", err.Error())
			continue
		}

		importDb[i].Package, _, err = getPackage(p, importDb[i].URI)
		if err != nil {
			write(p, "{RED}%s{RESET}", err.Error())
			continue
		}

		db = append(db, importDb[i])

		_, err = profile.LoadPackage(modulePath+importDb[i].Package, true, true)
		if err != nil {
			write(p, "{RED}%s{RESET}", err.Error())
		}
		_, err = profile.LoadPackage(modulePath+importDb[i].Package, true, false)
		if err != nil {
			write(p, "{RED}%s{RESET}", err.Error())
		}
	}

	var message string

	err = writePackagesFile(&db)
	if err != nil {
		message += err.Error() + utils.NewLineString
	}

	if message != "" {
		return errors.New(strings.TrimSpace(message))
	}

	return nil
}
