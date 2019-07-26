package test

import (
	"os"
	"testing"
)

func TestWorkingDirectory(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log("cwd: " + wd)
}
