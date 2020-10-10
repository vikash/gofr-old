package gofr

import (
	"testing"

	"github.com/vikash/gofr/pkg/gofr/testUtil"
)

func TestLogger_Log(t *testing.T) {

	logStatement := "hello log!"
	expectedLog := "hello log!\n"

	f := func() {
		logger := newLogger()
		logger.Log(logStatement)
	}

	output := testUtil.StdoutOutputForFunc(f)

	if output != expectedLog {
		t.Errorf("Stdout mismatch. Expected: %s Got: %s", expectedLog, output)
	}
}

func TestLogger_Logf(t *testing.T) {

	logStatement := "hello log!"
	expectedLog := "hello log!"

	f := func() {
		logger := newLogger()
		logger.Logf("%s", logStatement)
	}

	output := testUtil.StdoutOutputForFunc(f)

	if output != expectedLog {
		t.Errorf("Stdout mismatch. Expected: %s Got: %s", expectedLog, output)
	}
}

func TestLogger_Error(t *testing.T) {

	logStatement := "hello error!"
	expectedLog := "hello error!\n"

	f := func() {
		logger := newLogger()
		logger.Error(logStatement)
	}

	output := testUtil.StderrOutputForFunc(f)

	if output != expectedLog {
		t.Errorf("Stdout mismatch. Expected: %s Got: %s", expectedLog, output)
	}
}

func TestLogger_Errorf(t *testing.T) {

	logStatement := "hello error!"
	expectedLog := "hello error!"

	f := func() {
		logger := newLogger()
		logger.Errorf("%s", logStatement)
	}

	output := testUtil.StderrOutputForFunc(f)

	if output != expectedLog {
		t.Errorf("Stdout mismatch. Expected: %s Got: %s", expectedLog, output)
	}
}
