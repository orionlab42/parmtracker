package settings_test

import (
	"github.com/annakallo/parmtracker/testutil"
	"testing"
)

func TestMain(m *testing.M) {
	testutil.GlobalTearUp()
	code := m.Run()
	testutil.GlobalTearDown(code)
}
