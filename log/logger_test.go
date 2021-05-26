package log_test

import (
	"github.com/annakallo/parmtracker/log"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// Test coverage for Logger
func TestNewLogger(t *testing.T) {
	// check stdout
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	l := log.GetInstance()
	l.SetLevel(log.LevelInfo)
	l.Debug("test1", "message")
	l.Info("test2", "message")

	// capture the reading
	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	expected := "INFO  [test2] message"
	if !strings.Contains(string(out), expected) {
		t.Errorf("Printed %v should contain %v", string(out), expected)
	}
}

// Get the current level
func TestLogger_GetInitialLevel(t *testing.T) {
	initialLevel := log.GetInstance().GetLevel()
	log.GetInstance().SetLevel(log.LevelTrace)
	level := log.GetInstance().GetLevel()
	assert.Equal(t, log.LevelTrace, level)
	log.GetInstance().SetLevel(initialLevel)
}
