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
	testApp.Load(filename)
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

	testApp := App{"testWorkspace", nil, nil}
	testApp.Save(filename)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"defaultWorkspace":"testWorkspace"}`
	if string(contents) != expected {
		t.Fatalf("Failed to save config properly\nExpected: %v\n  Actual: %v", expected, string(contents))
	}
}
