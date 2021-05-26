package testutil

import (
	"crypto/tls"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// this is executed before the tests of a package, not for each individual tests!
func GlobalTearUp() {
	// set current directory by matching apollo suffix
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "parmtracker") {
		wd = filepath.Dir(wd)
	}
	_ = os.Chdir(wd)

	// avoid checking https verification
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// update packages
}

func GlobalTearDown(code int) {
	os.Exit(code)
}
