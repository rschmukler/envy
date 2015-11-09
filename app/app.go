package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const DEFAULT_CONFIG_PATH = "~/.envyconfig.json"

type val struct {
	Name       string
	Workspaces []string
	Value      string
}

type App struct {
	DefaultWorkspace string   `json:"defaultWorkspace"`
	Workspaces       []string `json:"workspaces,omitempty"`
	Vals             []val    `json:"vals,omitempty"`
}

func NewApp(path string) App {
	app := App{}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		app.Save(path)
	} else {
		app.Load(path)
	}
	return app
}

func (a *App) Load(path string) error {
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

func (a *App) Save(path string) error {
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
