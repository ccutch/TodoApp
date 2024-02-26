package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func storeToken(token string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to locate home directory")
	}
	appData := filepath.Join(homeDir, ".local", "share", "todos")
	err = os.MkdirAll(appData, 0700)
	if err != nil {
		return errors.Wrap(err, "failed to create app directory")
	}
	path := filepath.Join(appData, "jwt.token")
	err = os.WriteFile(path, []byte(token), 0600)
	return errors.Wrap(err, "failed to write token to file")
}

func readToken() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("failed to locate home directory: " + err.Error())
	}
	path := filepath.Join(homeDir, ".local", "share", "todos", "jwt.token")
	b, err := ioutil.ReadFile(path)
	return string(b), errors.Wrap(err, "failed to read token file")
}
