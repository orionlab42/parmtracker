package settings_test

import (
	"github.com/orionlab42/parmtracker/testutil"
	"testing"
)

func TestMain(m *testing.M) {
	testutil.GlobalTearUp()
	code := m.Run()
	testutil.GlobalTearDown(code)
}
