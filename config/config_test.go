package config

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func newErrCheck(t *testing.T, expected string) {
	if _, err := New(); err != nil && strings.Index(err.Error(), expected) == -1 {
		t.Errorf("error has no expected words. expected: %v, result: %v", expected, err)
	}
}

func TestNewFailWithoutFile(t *testing.T) {
	temp, _ := ioutil.TempDir("", "moesia-test")
	defer os.Remove(temp)
	filename = path.Join(temp, "no_existent_file")
	newErrCheck(t, "failed to make initial config file:")
}

func TestNewFailToMkdir(t *testing.T) {
	temp, _ := ioutil.TempDir("", "moesia-test")
	defer os.Remove(temp)
	dirWithoutPermission := path.Join(temp, "no_permission")
	os.Mkdir(dirWithoutPermission, 0000)
	newErrCheck(t, "failed to mkdir")
}

func TestNewFailWithInvalidFile(t *testing.T) {
	temp, _ := ioutil.TempFile("", "moesia-test-")
	defer os.Remove(temp.Name())
	filename = temp.Name()
	newErrCheck(t, "failed to load config file:")
}

func TestNewFailWithoutPermission(t *testing.T) {
	temp, _ := ioutil.TempDir("", "moesia-test-")
	defer os.Remove(temp)
	fileWithoutPermission := path.Join(temp, "no-permission")
	os.Create(fileWithoutPermission)
	os.Chmod(fileWithoutPermission, 0000)
	newErrCheck(t, "failed to open")
}

func TempNewFailWithInvalidJSON(t *testing.T) {
	temp, _ := ioutil.TempFile("", "moesia-test-")
	defer os.Remove(temp.Name())
	filename = temp.Name()
	ioutil.WriteFile(filename, []byte("invalid JSON"), 0666)
	newErrCheck(t, "failed to decode config")
}
