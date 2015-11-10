package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const DEFAULT_CONFIG_PATH = "~/.envyconfig.json"

type val struct {
	Name      string
	Workspace string
	Value     string
}

type App struct {
	DefaultWorkspace string   `json:"defaultWorkspace"`
	Workspaces       []string `json:"workspaces,omitempty"`
	Vals             []val    `json:"vals,omitempty"`
	path             string
}

func NewApp(path string) (App, error) {
	app := App{}
	app.path = path

	var err error

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = app.Save()
	} else {
		err = app.Load()
	}
	return app, err
}

func (a *App) Load() error {
	path := a.path
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, &a)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Save() error {
	path := a.path
	contents, err := json.Marshal(&a)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(contents), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) AddWorkspace(name string) error {
	if err := validateWorkspaceName(name); err != nil {
		return err
	} else if err = a.validateWorkspaceDoesNotExist(name); err != nil {
		return err
	}
	if a.Workspaces == nil {
		a.Workspaces = []string{name}
	} else {
		a.Workspaces = append(a.Workspaces, name)
	}
	return nil
}

func (a *App) RemoveWorkspace(name string) error {
	index, err := a.validateWorkspaceExists(name)

	if err != nil {
		return err
	}

	a.Workspaces = append(a.Workspaces[:index], a.Workspaces[index+1:]...)
	return nil
}

func (a *App) hasWorkspace(name string) (bool, int) {
	if a.Workspaces == nil || len(a.Workspaces) == 0 {
		return false, -1
	}
	for i, workspace := range a.Workspaces {
		if workspace == name {
			return true, i
		}
	}
	return false, -1
}

func validateWorkspaceName(name string) error {
	if len(name) == 0 {
		errMsg := fmt.Sprintf("Missing name argument for workspace")
		return errors.New(errMsg)
	}
	valid, _ := regexp.Match("^\\w+$", []byte(name))
	if !valid {
		errMsg := fmt.Sprintf("'%s' is not a valid workspace name")
		return errors.New(errMsg)
	} else {
		return nil
	}
}

func (a *App) validateWorkspaceExists(name string) (int, error) {
	present, index := a.hasWorkspace(name)

	if !present {
		errMsg := fmt.Sprintf("Workspace '%s' does not exist", name)
		return -1, errors.New(errMsg)
	}
	return index, nil
}

func (a *App) validateWorkspaceDoesNotExist(name string) error {
	if present, _ := a.hasWorkspace(name); present {
		errMsg := fmt.Sprintf("Workspace '%s' already exists", name)
		return errors.New(errMsg)
	}
	return nil
}
