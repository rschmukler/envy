package app

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	f, err := ioutil.TempFile("", "envy-load-test")
	if err != nil {
		t.Fatal(err)
	}
	filename := f.Name()
	defer f.Close()
	defer func() { os.Remove(filename) }()

	data := `{"defaultWorkspace": "home"}`

	if err := ioutil.WriteFile(filename, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	testApp := App{}
	testApp.path = filename
	testApp.Load()
	if testApp.DefaultWorkspace != "home" {
		t.Fatalf("Failed to load config properly")
	}
}

func TestSave(t *testing.T) {
	t.Parallel()

	f, err := ioutil.TempFile("", "envy-save-test")
	if err != nil {
		t.Fatal(err)
	}
	filename := f.Name()
	f.Close()
	defer func() { os.Remove(filename) }()

	testApp := App{"testWorkspace", nil, nil, filename}
	testApp.Save()

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"defaultWorkspace":"testWorkspace"}`
	if string(contents) != expected {
		t.Fatalf("Failed to save config properly\nExpected: %v\n  Actual: %v", expected, string(contents))
	}
}

func TestHasWorkspace(t *testing.T) {
	t.Parallel()
	app := App{"", []string{"test"}, nil, ""}
	if present, index := app.hasWorkspace("test"); !present {
		if index != 0 {
			t.Fatalf("hasWorkspace returned incorrect index")
		}
		t.Fatalf("hasWorkspace returned a false negative")
	}
	if present, _ := app.hasWorkspace("no"); present {
		t.Fatal("hasWorkspace returned a false positive")
	}
}

func TestAddWorkspaceBasic(t *testing.T) {
	t.Parallel()
	app := App{}
	app.AddWorkspace("home")

	if app.Workspaces == nil {
		t.Fatalf("AddWorkspace did not initialize Workspaces")
	}

	if app.Workspaces[0] != "home" {
		t.Fatalf("AddWorkspace added incorrect value")
	}
}

func TestAddWorkspaceDuplicateError(t *testing.T) {
	t.Parallel()
	app := App{}
	app.AddWorkspace("home")
	duplicateError := app.AddWorkspace("home")
	if len(app.Workspaces) != 1 {
		t.Fatalf("AddWorkspace did not deduplicate itself")
	}
	if duplicateError == nil {
		t.Fatalf("AddWorkspace did not throw error on duplicate entry")
	}
}

func TestAddWorkspaceNameValidation(t *testing.T) {
	t.Parallel()
	app := App{}
	invalidNameErr := app.AddWorkspace("invalid workspace name")

	if invalidNameErr == nil {
		t.Fatalf("AddWorkspace did not validate workspace name")
	}

	if app.Workspaces != nil {
		t.Fatalf("AddWorkspace did not prevent invalid entry")
	}
}

func TestRemoveWorkspace(t *testing.T) {
	t.Parallel()
	app := App{"", []string{"test", "other"}, nil, ""}
	app.RemoveWorkspace("test")

	if len(app.Workspaces) == 2 {
		t.Fatalf("RemoveWorkspace did not remove the workspace")
	} else if app.Workspaces[0] != "other" {
		t.Fatalf("RemoveWorkspace removed the wrong workspace")
	}
}

func TestRemoveNonExistant(t *testing.T) {
	t.Parallel()
	app := App{"", []string{"test", "other"}, nil, ""}
	removeNonExistantError := app.RemoveWorkspace("non-existant")

	if removeNonExistantError == nil {
		t.Fatalf("RemoveWorkspace did not throw error on non-existant workspace removal")
	}
}
